package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type SessionHandler struct {
	SessionService *service.SessionService
}

func NewSessionHandler(service *service.SessionService) *SessionHandler {
	return &SessionHandler{
		SessionService: service,
	}
}

func (handler *SessionHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var session model.Session
	err := json.NewDecoder(req.Body).Decode(&session)
	if err != nil {
		println("Error while parsing json")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createdSession, err := handler.SessionService.Create(&session)
	if err != nil {
		println("Error while creating a new session")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(createdSession); err != nil {
		println("Error while encoding session to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *SessionHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var session model.Session
	err := json.NewDecoder(req.Body).Decode(&session)
	if err != nil {
		println("Error while parsing json: Update")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedSession, err := handler.SessionService.Update(&session)
	if err != nil {
		println("Error while creating a new session")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(updatedSession); err != nil {
		println("Error while encoding session to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *SessionHandler) CompleteKeypoint(writer http.ResponseWriter, req *http.Request) {
	sessionId := mux.Vars(req)["sessionId"]
	var keypointId int
	err := json.NewDecoder(req.Body).Decode(&keypointId)
	if err != nil {
		println("Error while parsing json: CompleteKeypoint")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	updatedSession, err := handler.SessionService.CompleteKeypoint(sessionId, keypointId)
	if err != nil {
		println("Error while completing keypoint")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(updatedSession); err != nil {
		println("Error while encoding session to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
