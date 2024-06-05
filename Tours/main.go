package main

import (
	"database-example/handler"
	tours "database-example/proto/tours"
	"database-example/repo"
	"database-example/service"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"context"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

const serviceName = "go-tour-service"

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

func GetConnectionString() string {

	connectionString, isPresent := os.LookupEnv("MONGO_DB_URI")
	if isPresent {
		return connectionString
	} else {
		return "mongodb://localhost:27017"
	}

}

func startServer(client *mongo.Client) {
	router := mux.NewRouter().StrictSlash(true)

	//initTourKeypoints(router, client)
	initTours(router, client)
	initSessions(router, client)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8082", router))
}

// func initTourKeypoints(router *mux.Router, client *mongo.Client) {
// 	repo := &repo.TourKeypointRepository{DatabaseConnection: client}
// 	service := &service.TourKeypointService{TourKeypointRepo: repo}
// 	handler := &handler.TourKeypointHandler{TourKeypointService: service}
// 	router.HandleFunc("/tourKeypoints/{id}", handler.Get).Methods("GET")
// 	router.HandleFunc("/tourKeypoints/create", handler.Create).Methods("POST")
// 	router.HandleFunc("/tourKeypoints/update", handler.Update).Methods("PUT")
// 	router.HandleFunc("/tourKeypoints/delete/{id}", handler.Delete).Methods("DELETE")
// }

func initTours(router *mux.Router, client *mongo.Client) {
	repo := &repo.TourRepository{DatabaseConnection: client}
	service := &service.TourService{TourRepo: repo}
	handler := &handler.TourHandler{TourService: service}

	router.HandleFunc("/tours/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/tours/create", handler.Create).Methods("POST")
	router.HandleFunc("/tours/update", handler.Update).Methods("PUT")
	router.HandleFunc("/tours/author/{id}", handler.GetByAuthorId).Methods("GET")
	router.HandleFunc("/tours/publish/{id}", handler.Publish).Methods("PUT")
	router.HandleFunc("/tours/archive/{id}", handler.Archive).Methods("PUT")
}

func initSessions(router *mux.Router, client *mongo.Client) {
	repo := &repo.SessionRepository{DatabaseConnection: client}
	service := &service.SessionService{SessionRepo: repo}
	handler := &handler.SessionHandler{SessionService: service}

	router.HandleFunc("/sessions/create", handler.Create).Methods("POST")
	router.HandleFunc("/sessions/update", handler.Update).Methods("PUT")
	router.HandleFunc("/sessions/completeKeypoint/{sessionId}", handler.CompleteKeypoint).Methods("PUT")
}

func main() {
	// mongo
	connectionStr := GetConnectionString()
	fmt.Printf("Connecting to MongoDB with URI: %s\n", connectionStr)
	opts := options.Client().ApplyURI(connectionStr)
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Failed to create MongoDB client: %v", err)
	}

	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatalf("Failed to disconnect from MongoDB: %v", err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Println("Connected to MongoDB")

	repo := &repo.TourKeypointRepository{DatabaseConnection: client}
	service := &service.TourKeypointService{TourKeypointRepo: repo}
	handler := &handler.TourKeypointHandler{TourKeypointService: service}

	//cfg := config.GetConfig()

	listener, err := net.Listen("tcp", ":8082")

	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(TokenValidationInterceptor),
	)
	reflection.Register(grpcServer)
	tours.RegisterTourServiceServer(grpcServer, handler)

	println("Server starting")

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()
}

func TokenValidationInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}
	tokens := md.Get("authorization")
	if len(tokens) == 0 {
		return nil, fmt.Errorf("missing token")
	}
	tokenString := tokens[0]

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("explorer_secret_key"), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return handler(ctx, req)
}
