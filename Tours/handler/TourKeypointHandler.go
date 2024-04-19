package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"fmt"
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
	fmt.Printf("sss:")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tourKeypoint)
}

func (handler *TourKeypointHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tourKeypoint model.TourKeypoint
	err := json.NewDecoder(req.Body).Decode(&tourKeypoint)
	if err != nil {
		println("Error while parsing json: Create Keypoint")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	createdTourKeypoint, err := handler.TourKeypointService.Create(&tourKeypoint)
	if err != nil {
		println("Error while creating a new tourkeypoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(createdTourKeypoint); err != nil {
		println("Error while encoding tour to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourKeypointHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var tourKeypoint model.TourKeypoint
	err := json.NewDecoder(req.Body).Decode(&tourKeypoint)
	if err != nil {
		println("Error while parsing json: Update Keypoint")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedTourKeypoint, err := handler.TourKeypointService.Update(&tourKeypoint)
	if err != nil {
		println("Error while updating a new tourkeypoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(updatedTourKeypoint); err != nil {
		println("Error while encoding tour to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
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
