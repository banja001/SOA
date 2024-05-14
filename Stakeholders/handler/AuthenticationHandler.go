package handler

import (
	"encoding/json"
	"net/http"
	"stakeholders/dto"
	"stakeholders/service"
)

type AuthenticationHandler struct {
	AuthenticationService *service.AuthenticationService
}

func (handler *AuthenticationHandler) Login(writer http.ResponseWriter, req *http.Request) {
	var credentials dto.Credentials
	err := json.NewDecoder(req.Body).Decode(&credentials)
	if err != nil {
		println("Error while parsing json: Credentials: ", req.Body)
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	token, err := handler.AuthenticationService.Login(&credentials)
	if err != nil {
		println("Login error")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(token); err != nil {
		println("Error while encoding token to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
