package main

import (
	"Followers/handler"
	"Followers/repo"
	"Followers/service"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.23.0"
)

const serviceName = "go-follower-service"

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

// func main() {
	

// 	port := os.Getenv("PORT")
// 	if len(port) == 0 {
// 		port = "8060"
// 	}

// 	// Initialize context
// 	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()

// 	//Initialize the logger we are going to use, with prefix and datetime for every log
// 	logger := log.New(os.Stdout, "[followers-api] ", log.LstdFlags)
// 	storeLogger := log.New(os.Stdout, "[followers-store] ", log.LstdFlags)

// 	// NoSQL: Initialize Follower Repository store
// 	store, err := repo.New(storeLogger)
// 	if err != nil {
// 		logger.Fatal(err)
// 	}
// 	defer store.CloseDriverConnection(timeoutContext)
// 	store.CheckConnection()

// 	//Initialize the service and inject said logger
// 	followerService := service.NewFollowerService(store, logger)

// 	//Initialize the handler and inject said logger
// 	followerHandler := handler.NewFollowerHandler(followerService, logger)

// 	//Initialize the router and add a middleware for all the requests
// 	router := mux.NewRouter()

// 	router.Use(followerHandler.MiddlewareContentTypeSet)

// 	getAllFollowers := router.Methods(http.MethodGet).Subrouter()
// 	getAllFollowers.HandleFunc("/followers", followerHandler.GetAllFollowers)

// 	getAllRecomended := router.Methods(http.MethodGet).Subrouter()
// 	getAllRecomended.HandleFunc("/followers/recommended/{id}/{uid}", followerHandler.GetAllRecomended)

// 	getAllFollowed := router.Methods(http.MethodGet).Subrouter()
// 	getAllFollowed.HandleFunc("/followers/followed/{id}/{uid}", followerHandler.IsFollowed)

// 	putFollower := router.Methods(http.MethodPut).Subrouter()
// 	putFollower.HandleFunc("/followers/update", followerHandler.CreateFollower)

// 	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

// 	//Initialize the server
// 	server := http.Server{
// 		Addr:         ":" + port,
// 		Handler:      cors(router),
// 		IdleTimeout:  120 * time.Second,
// 		ReadTimeout:  5 * time.Second,
// 		WriteTimeout: 5 * time.Second,
// 	}

// 	logger.Println("Server listening on port", port)
// 	//Distribute all the connections to goroutines
// 	go func() {
// 		err := server.ListenAndServe()
// 		if err != nil {
// 			logger.Fatal(err)
// 		}
// 	}()

// 	sigCh := make(chan os.Signal)
// 	signal.Notify(sigCh, os.Interrupt)
// 	signal.Notify(sigCh, os.Kill)

// 	sig := <-sigCh
// 	logger.Println("Received terminate, graceful shutdown", sig)

// 	//Try to shutdown gracefully
// 	if server.Shutdown(timeoutContext) != nil {
// 		logger.Fatal("Cannot gracefully shutdown...")
// 	}
// 	logger.Println("Server stopped")
// }

func main() {


	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[followers-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[followers-store] ", log.LstdFlags)

	// NoSQL: Initialize Follower Repository store
	store, err := repo.New(storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.CloseDriverConnection(timeoutContext)
	store.CheckConnection()

	// Initialize the service and inject said logger
	followerService := service.NewFollowerService(store, logger)

	// Initialize the handler and inject said logger
	followerHandler := handler.NewFollowerHandler(followerService, logger)
	// OpenTelemetry
	tp, err = initTracer()
	if err != nil {
		logger.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			logger.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	// Initialize the Gin router
	router := gin.New()
	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)
	router.Use(otelgin.Middleware(serviceName))

	// Add routes
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the followers API")
	})
	router.GET("/followers", func(c *gin.Context) {
		followerHandler.GetAllFollowers(c.Writer, c.Request)
	})
	router.GET("/followers/recommended/:id/:uid", func(c *gin.Context) {
		followerHandler.GetAllRecomended(c.Writer, c.Request)
	})
	router.GET("/followers/followed/:id/:uid", func(c *gin.Context) {
		followerHandler.IsFollowed(c.Writer, c.Request)
	})
	router.PUT("/followers/update", func(c *gin.Context) {
		followerHandler.CreateFollower(c.Writer, c.Request)
	})

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8060"
	}

	router.Run(fmt.Sprintf(":%s", port))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      router,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	logger.Println("Server listening on port", port)
	// Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	// Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
