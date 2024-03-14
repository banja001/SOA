package handler

import (
	"database-example/model"
	"database-example/service"
	"encoding/json"
	"net/http"
)

type SessionHandler struct {
	SessionService *service.SessionService
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
