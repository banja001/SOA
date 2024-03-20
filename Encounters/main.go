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

func startServer(database *gorm.DB) {
	router := mux.NewRouter().StrictSlash(true)

	initUserExpirience(router, database)
	initChallenges(router, database)
	initChallengeExecution(router, database)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))
	println("Server starting")
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initUserExpirience(router *mux.Router, database *gorm.DB) {
	repo := &repo.UserExperienceRepository{DatabaseConnection: database}
	service := &service.UserExperienceService{UserExperienceRepo: repo}
	handler := &handler.UserExperienceHandler{UserExperienceService: service}

	router.HandleFunc("/userxp/{userId}", handler.GetByUserId).Methods("GET")
	router.HandleFunc("/userxp/add/{id}/{xp}", handler.AddXP).Methods("PUT")
	router.HandleFunc("/userxp/create", handler.Create).Methods("POST")
}

func initChallenges(router *mux.Router, database *gorm.DB) {
	repo := &repo.ChallengeRepository{DatabaseConnection: database}
	service := &service.ChallengeService{ChallengeRepository: repo}
	handler := &handler.ChallengeHandler{ChallengeService: service}

	router.HandleFunc("/challenge", handler.GetAll).Methods("GET")
	router.HandleFunc("/challenge/{id}", handler.Delete).Methods("DELETE")
	router.HandleFunc("/challenge", handler.Update).Methods("PUT")
	router.HandleFunc("/challenge", handler.Create).Methods("POST")
}

func initChallengeExecution(router *mux.Router, database *gorm.DB) {
	repo := &repo.ChallengeExecutionRepository{DatabaseConnection: database}
	service := &service.ChallengeExecutionService{ChallengeExecutionRepository: repo}
	handler := &handler.ChallengeExecutionHandler{ChallengeExecutionService: service}
	router.HandleFunc("/challenge-execution/{id}", handler.Delete).Methods("DELETE")
}

func main() {
	database := initDB()
	if database == nil {
		print("FAILED TO CONNECT TO DB")
		return
	}

	startServer(database)
}
