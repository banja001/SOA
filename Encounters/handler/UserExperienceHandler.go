package handler

import (
	"encgo/model"
	"encgo/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type UserExperienceHandler struct {
	UserExperienceService *service.UserExperienceService
}

func (handler *UserExperienceHandler) GetByUserId(writer http.ResponseWriter, req *http.Request) {
	userId := mux.Vars(req)["userId"]
	num, _ := strconv.Atoi(userId)
	userExperience, err := handler.UserExperienceService.FindByUserId(num)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}

	u, _ := json.Marshal(userExperience)
	fmt.Println(string(u))

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(userExperience)
}

func (handler *UserExperienceHandler) AddXP(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	num, _ := strconv.Atoi(id)

	xp := mux.Vars(req)["xp"]
	xpNum, _ := strconv.Atoi(xp)

	userExperience, err := handler.UserExperienceService.AddXP(num, xpNum)
	if err != nil {
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(userExperience)
}

func (handler *UserExperienceHandler) Create(writer http.ResponseWriter, req *http.Request) {
	var userExperience model.UserExperience
	err := json.NewDecoder(req.Body).Decode(&userExperience)
	
	if err != nil {
		println("Error while parsing json: Create user experience")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	createdUserExperience, err := handler.UserExperienceService.Create(&userExperience)
	if err != nil {
		println("Error while creating a new user experience")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(createdUserExperience); err != nil {
		println("Error while encoding user experience to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (handler *UserExperienceHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	err := handler.UserExperienceService.Delete(id)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application-json")
}