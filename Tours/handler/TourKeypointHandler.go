package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type TourKeypointHandler struct {
	TourKeypointService *service.TourKeypointService
}

func (handler *TourKeypointHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	tourKeypoint, err := handler.TourKeypointService.Find(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourKeypoint)
}


func (handler *TourKeypointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourKeypoint model.TourKeypoint
	err := json.NewDecoder(req.Body).Decode(&tourKeypoint)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourKeypointService.Create(&tourKeypoint)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourKeypointHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var tourKeypoint model.TourKeypoint
	err := json.NewDecoder(req.Body).Decode(&tourKeypoint)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.TourKeypointService.Update(&tourKeypoint)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
}

func (handler *TourKeypointHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	err := handler.TourKeypointService.Delete(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}