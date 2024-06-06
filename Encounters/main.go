package main

import (
	"context"
	"encgo/config"
	"encgo/handler"
	"encgo/model"
	user_experience "encgo/proto/user-experience"
	"encgo/repo"
	"encgo/service"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	cfg "github.com/banja001/SOA/Encounters/config"
	saga "github.com/banja001/SOA/saga/messaging"
	"github.com/banja001/SOA/saga/messaging/nats"
	"github.com/dgrijalva/jwt-go"
	"github.com/tamararankovic/microservices_demo/ordering_service/startup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnectionString() string {
	connectionString, isPresent := os.LookupEnv("DATABASE_URL")
	if isPresent {
		return connectionString
	} else {
		return "host=localhost user=postgres password=super dbname=encountersdb port=5432 sslmode=disable"
	}

}
func initDB() *gorm.DB {
	connectionStr := GetConnectionString()
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionStr,
	}), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}
	err = database.AutoMigrate(&model.UserExperience{})
	if err != nil {
		log.Fatal("Error while running migration for user experience")
	}
	err = database.AutoMigrate(&model.Challenge{})
	if err != nil {
		log.Fatal("Error while running migration for challenges")
	}
	err = database.AutoMigrate(&model.ChallengeExecution{})
	if err != nil {
		log.Fatal("Error while running migration for challenges")
	}
	return database
}

type Server struct {
	config *config.Config
}

// func startServer(database *gorm.DB) {
// 	router := mux.NewRouter().StrictSlash(true)

// 	initUserExpirience(router, database)
// 	initChallenges(router, database)
// 	initChallengeExecution(router, database)
// 	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
// 	println("Server starting")
// 	log.Fatal(http.ListenAndServe(":8090", router))
// }

// func initUserExpirience(router *mux.Router, database *gorm.DB) {
// 	repo := &repo.UserExperienceRepository{DatabaseConnection: database}
// 	service := &service.UserExperienceService{UserExperienceRepo: repo}
// 	handler := &handler.UserExperienceHandler{UserExperienceService: service}

// 	router.HandleFunc("/userxp/{userId}", handler.GetByUserId).Methods("GET")
// 	router.HandleFunc("/userxp/add/{id}/{xp}", handler.AddXP).Methods("PUT")
// 	router.HandleFunc("/userxp/create", handler.Create).Methods("POST")
// 	router.HandleFunc("/userxp/delete/{id}", handler.Delete).Methods("DELETE")
// 	router.HandleFunc("/userxp/update", handler.Update).Methods("PUT")
// }

// func initChallenges(router *mux.Router, database *gorm.DB) {
// 	repo := &repo.ChallengeRepository{DatabaseConnection: database}
// 	service := &service.ChallengeService{ChallengeRepository: repo}
// 	handler := &handler.ChallengeHandler{ChallengeService: service}

// 	router.HandleFunc("/challenge", handler.GetAll).Methods("GET")
// 	router.HandleFunc("/challenge/{id}", handler.Delete).Methods("DELETE")
// 	router.HandleFunc("/challenge", handler.Update).Methods("PUT")
// 	router.HandleFunc("/challenge", handler.Create).Methods("POST")
// }

// func initChallengeExecution(router *mux.Router, database *gorm.DB) {
// 	repo := &repo.ChallengeExecutionRepository{DatabaseConnection: database}
// 	service := &service.ChallengeExecutionService{ChallengeExecutionRepository: repo}
// 	handler := &handler.ChallengeExecutionHandler{ChallengeExecutionService: service}
// 	router.HandleFunc("/challenge-execution/{id}", handler.Delete).Methods("DELETE")
// }

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

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	//startServer(database)
	repo := &repo.UserExperienceRepository{DatabaseConnection: database}
	service := &service.UserExperienceService{UserExperienceRepo: repo}
	handler := &handler.UserExperienceHandler{UserExperienceService: service}
	config := cfg.NewConfig()
	server := startup.NewServer(config)

	commandSubscriber := server.initSubscriber(server.config.CreateOrderCommandSubject, QueueGroup)
	replyPublisher := server.initPublisher(server.config.CreateOrderReplySubject)
	server.initCreateOrderHandler(productService, replyPublisher, commandSubscriber)

	productHandler := server.initProductHandler(productService)

	//cfg := config.GetConfig()

	listener, err := net.Listen("tcp", ":8090")

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
	user_experience.RegisterUserExperienceServiceServer(grpcServer, handler)

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

func NewServer(config *config.Config) *Server {
	return &Server{
		config: config,
	}
}

const (
	QueueGroup = "encounter_service"
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

func (server *Server) initProductService(store model.UserExperience) *service.UserExperienceService {
	return service.NewProductService(store)
}

func (server *Server) initCreateOrderHandler(service *service.UserExperienceService, publisher saga.Publisher, subscriber saga.Subscriber) {
	_, err := handler.NewAddxpCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}
