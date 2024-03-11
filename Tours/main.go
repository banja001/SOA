package main

import (
	"database-example/handler"
	"database-example/model"
	"database-example/repo"
	"database-example/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "host=localhost user=postgres password=super dbname=tourdb port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionStr,
	}), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}
	err = database.AutoMigrate(&model.Tour{})
	if err != nil {
		log.Fatal("Error while running migration for tour")
	}
	err = database.AutoMigrate(&model.TourKeypoint{})
	if err != nil {
		log.Fatal("Error while running migration for tour keypoints")
	}
	return database
}

func startServer(database *gorm.DB) {
	router := mux.NewRouter().StrictSlash(true)

	initTourKeypoints(router, database)
	initTours(router, database)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":80", router))
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
	router.HandleFunc("/tours", handler.Create).Methods("POST")
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	startServer(database)
}
