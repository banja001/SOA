package main

import (
	"database-example/config"
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

	"github.com/banja001/SOA/saga/messaging/nats"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	saga "github.com/tamararankovic/microservices_demo/common/saga/messaging"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
)

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

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "add_xp_service"
)

func (server *Server) initPublisher(subject string) saga.Publisher {
	publisher, err := nats.NewNATSPublisher(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func (server *Server) initSubscriber(subject, queueGroup string) saga.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		server.config.NatsHost, server.config.NatsPort,
		server.config.NatsUser, server.config.NatsPass, subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func (server *Server) initCreateOrderOrchestrator(publisher saga.Publisher, subscriber saga.Subscriber) *service.GiveXPOrchestrator {
	orchestrator, err := service.NewCreateOrderOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return orchestrator
}

func (server *Server) initOrderService(orchestrator *service.GiveXPOrchestrator) *service.SessionService {
	return service.NewSessionService(orchestrator)
}

func (server *Server) initCreateOrderHandler(service *service.SessionService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handler.NewAddXPCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func (server *Server) initOrderHandler(service *service.SessionService) *handler.SessionHandler {
	return handler.NewSessionHandler(service)
}
