package handler

import (
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
	json.NewEncoder(writer).Encode(&userExperience)
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
