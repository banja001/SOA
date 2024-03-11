package main

import (
	"encgo/handler"
	"encgo/model"
	"encgo/repo"
	"encgo/service"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connectionStr := "host=localhost user=postgres password=super dbname=encountersdb port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.New(postgres.Config{
		DSN: connectionStr,
	}), &gorm.Config{})
	if err != nil {
		print(err)
		return nil
	}
	err = database.AutoMigrate(&model.UserExperience{}) 
	if (err != nil){
		log.Fatal("Error while running migration for user experience")
	}
	return database
}

func startServer(handler *handler.UserExperienceHandler) {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/userxp/{userId}", handler.GetByUserId).Methods("GET")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":80", router))
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}
	repo := &repo.UserExperienceRepository{DatabaseConnection: database}
	service := &service.UserExperienceService{UserExperienceRepo: repo}
	handler := &handler.UserExperienceHandler{UserExperienceService: service}

	startServer(handler)
}
