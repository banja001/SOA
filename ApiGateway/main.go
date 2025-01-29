package main

import (
	stakeholder_service "api-gateway/proto/stakeholder-service"
	tour_service "api-gateway/proto/tour-service"
	user_experience_service "api-gateway/proto/user-experience-service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const serviceName = "api-gateway"

var tp *trace.TracerProvider

func initTracer() (*trace.TracerProvider, error) {
	// Ukoliko je definisana JAEGER_ENDPOINT env var, intanciraj JagerTracer koji šalje trace-ove Jaeger-u,
	// u suprotnom instanciraj FileTracer koji upisuje trace-ove u json fajl
	url := os.Getenv("JAEGER_ENDPOINT")
	if len(url) > 0 {
		return initJaegerTracer(url)
	} else {
		return initFileTracer()
	}
}

func initFileTracer() (*trace.TracerProvider, error) {
	log.Println("Initializing tracing to traces.json")
	f, err := os.Create("traces.json")
	if err != nil {
		return nil, err
	}
	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(f),
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	), nil
}

func initJaegerTracer(url string) (*trace.TracerProvider, error) {
	log.Printf("Initializing tracing to Jaeger for service: %s at %s\n", serviceName, url)
	log.Printf("Initializing tracing to jaeger at %s\n", url)
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
		)),
	), nil
}

func main() {
	var err error
	tp, err = initTracer()
	if err != nil {
		log.Fatalf("failed to initialize tracer: %v", err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Fatalf("failed to shut down tracer: %v", err)
		}
	}()

	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// Setup Gin Router
	router := gin.New()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)
	router.Use(otelgin.Middleware(serviceName))
	


	// ZA SERVISE
	conn, err := grpc.DialContext(
		context.Background(),
		os.Getenv("STAKEHOLDERS_SERVICE_ADDRESS"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Fatalln("Failed to dial stakeholders server:", err)
	}

	userExperienceConn, err := grpc.DialContext(
		context.Background(),
		os.Getenv("ENCOUNTERS_SERVICE_ADDRESS"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Fatalln("Failed to dial user experience server:", err)
	}

	conn_tours, err := grpc.DialContext(
		context.Background(),
		os.Getenv("TOURS_SERVICE_ADDRESS"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
	)
	if err != nil {
		log.Fatalln("Failed to dial tours server:", err)
	}

	gwmux := runtime.NewServeMux()

	client := stakeholder_service.NewStakeholderServiceClient(conn)
	err = stakeholder_service.RegisterStakeholderServiceHandlerClient(
		context.Background(),
		gwmux,
		client,
	)
	if err != nil {
		log.Fatalln("Failed to register stakeholders gateway:", err)
	}

	client_uxp := user_experience_service.NewUserExperienceServiceClient(userExperienceConn)
	err = user_experience_service.RegisterUserExperienceServiceHandlerClient(
		context.Background(),
		gwmux,
		client_uxp,
	)
	if err != nil {
		log.Fatalln("Failed to register gateway user xp:", err)
	}

	client_tours := tour_service.NewTourServiceClient(conn_tours)
	err = tour_service.RegisterTourServiceHandlerClient(
		context.Background(),
		gwmux,
		client_tours,
	)
	if err != nil {
		log.Fatalln("Failed to register gateway tours:", err)
	}
	log.Println("Gateway address: ", os.Getenv("GATEWAY_ADDRESS"))
	gwServer := &http.Server{
		Addr:    os.Getenv("GATEWAY_ADDRESS"),
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:9000")

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}
