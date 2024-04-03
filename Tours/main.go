package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnectionString() string {
	connectionString, isPresent := os.LookupEnv("DATABASE_URL2")
	if isPresent {
		return connectionString
	} else {
		return "host=localhost user=postgres password=super dbname=tourdb port=5432 sslmode=disable"
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

	err = database.AutoMigrate(&model.Tour{})
	if err != nil {
		log.Fatal("Error while running migration for tours")
	}
	err = database.AutoMigrate(&model.TourKeypoint{})
	if err != nil {
		log.Fatal("Error while running migration for tour keypoints")
	}
	err = database.AutoMigrate(&model.Session{})
	if err != nil {
		log.Fatal("Error while running migration for sessions")
	}
	return database
}

func startServer(database *gorm.DB) {
	router := mux.NewRouter().StrictSlash(true)

	initTourKeypoints(router, database)
	initTours(router, database)
	initSessions(router, database)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func initTourKeypoints(router *mux.Router, database *gorm.DB) {
	repo := &repo.TourKeypointRepository{DatabaseConnection: database}
	service := &service.TourKeypointService{TourKeypointRepo: repo}
	handler := &handler.TourKeypointHandler{TourKeypointService: service}

	router.HandleFunc("/tourKeypoints/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/tourKeypoints/create", handler.Create).Methods("POST")
	router.HandleFunc("/tourKeypoints/update", handler.Update).Methods("PUT")
	router.HandleFunc("/tourKeypoints/delete/{id}", handler.Delete).Methods("DELETE")
}

func initTours(router *mux.Router, database *gorm.DB) {
	repo := &repo.TourRepository{DatabaseConnection: database}
	service := &service.TourService{TourRepo: repo}
	handler := &handler.TourHandler{TourService: service}

	router.HandleFunc("/tours/{id}", handler.Get).Methods("GET")
	router.HandleFunc("/tours/create", handler.Create).Methods("POST")
	router.HandleFunc("/tours", handler.GetAll).Methods("GET")
	router.HandleFunc("/tours/update", handler.Update).Methods("PUT")
	router.HandleFunc("/tours/author/{id}", handler.GetByAuthorId).Methods("GET")
	router.HandleFunc("/tours/publish/{id}", handler.Publish).Methods("PUT")
	router.HandleFunc("/tours/archive/{id}", handler.Archive).Methods("PUT")
}

func initSessions(router *mux.Router, database *gorm.DB) {
	repo := &repo.SessionRepository{DatabaseConnection: database}
	service := &service.SessionService{SessionRepo: repo}
	handler := &handler.SessionHandler{SessionService: service}

	router.HandleFunc("/sessions/create", handler.Create).Methods("POST")
	router.HandleFunc("/sessions/update", handler.Update).Methods("PUT")
	router.HandleFunc("/sessions/completeKeypoint/{sessionId}", handler.CompleteKeypoint).Methods("PUT")
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	startServer(database)
}
