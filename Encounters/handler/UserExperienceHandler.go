package handler

import (
	"encgo/service"
	"encoding/json"
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
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(userExperience)
}
