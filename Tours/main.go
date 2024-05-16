package main

import (
	"database-example/handler"
	"database-example/proto/tours"
	"database-example/repo"
	"database-example/service"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"context"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
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


	// gateway
	repo := &repo.TourKeypointRepository{DatabaseConnection: client}
	service := &service.TourKeypointService{TourKeypointRepo: repo}
	handler := &handler.TourKeypointHandler{TourKeypointService: service}

	listener, err := net.Listen("tcp", ":8093")

	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	tours.RegisterTourServiceServer(grpcServer, handler)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()

	//startServer(client)
}
