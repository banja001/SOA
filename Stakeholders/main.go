package main

import (
	"log"
	"net"
	"os"
	"os/signal"
	"stakeholders/config"
	"stakeholders/model"
	"stakeholders/proto/stakeholders"
	"stakeholders/repo"
	"stakeholders/service"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnectionString() string {
	connectionString, isPresent := os.LookupEnv("DATABASE_URL")
	if isPresent {
		return connectionString
	} else {
		return "host=localhost user=postgres password=super dbname=usersdb port=5432 sslmode=disable"
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
	err = database.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Error while running migration for user experience")
	}
	return database
}

// func startServer(database *gorm.DB) {
// 	router := mux.NewRouter().StrictSlash(true)

// 	initUsers(router, database)
// 	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
// 	println("Server starting")
// 	log.Fatal(http.ListenAndServe(":8093", router))
// }

/*unc initUsers(router *mux.Router, database *gorm.DB) {
	repo := &repo.UserRepository{DatabaseConnection: database}
	service := &service.AuthenticationService{UserRepository: repo}
	handler := &handler.AuthenticationHandler{AuthenticationService: service}

	router.HandleFunc("/users/login", handler.Login).Methods("POST")
}*/

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	//startServer(database)
	repo := &repo.UserRepository{DatabaseConnection: database}
	service := &service.AuthenticationService{UserRepository: repo}

	cfg := config.GetConfig()

	listener, err := net.Listen("tcp", cfg.Address)
	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	// Bootstrap gRPC server.
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	// Bootstrap gRPC service server and respond to request.
	//authenticationService := service.AuthenticationService{}
	stakeholders.RegisterAuthenticationServiceServer(grpcServer, service)

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
