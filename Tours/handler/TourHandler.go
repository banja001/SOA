package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type TourHandler struct {
	TourService *service.TourService
}

func (handler *TourHandler) Get(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	tour, err := handler.TourService.Find(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tour)
}

func (handler *TourHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("Error while parsing json: Create Tour: ", req.Body)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createdTour, err := handler.TourService.Create(&tour)
	if err != nil {
		println("Error while creating a new tour")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(createdTour); err != nil {
		println("Error while encoding tour to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// func (handler *TourHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
// 	tours, err := handler.TourService.GetAll()
// 	writer.Header().Set("Content-Type", "application/json")
// 	if err != nil {
// 		writer.WriteHeader(http.StatusInternalServerError)
// 		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
// 		return
// 	}

// 	writer.WriteHeader(http.StatusOK)
// 	json.NewEncoder(writer).Encode(tours)
// }

func (handler *TourHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var tour model.Tour
	err := json.NewDecoder(req.Body).Decode(&tour)
	if err != nil {
		println("Error while parsing json: Update Tour")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedTour, err := handler.TourService.Update(&tour)
	if err != nil {
		println("Error while updating")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(updatedTour); err != nil {
		println("Error while encoding tour to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourHandler) GetByAuthorId(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	tours, err := handler.TourService.GetByAuthorId(id)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(tours)
}

func (handler *TourHandler) Publish(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	var authorID int
	err := json.NewDecoder(req.Body).Decode(&authorID)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	publishedTour, err := handler.TourService.ChangeStatus(id, authorID, model.Published)
	if err != nil {
		println("Error while publishing")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(publishedTour); err != nil {
		println("Error while encoding tour to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *TourHandler) Archive(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	var authorID int
	err := json.NewDecoder(req.Body).Decode(&authorID)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	archivedTour, err := handler.TourService.ChangeStatus(id, authorID, model.Archived)
	if err != nil {
		println("Error while archiving")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(archivedTour); err != nil {
		println("Error while encoding tour to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
