package handler

import (
	"encgo/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ChallengeExecutionHandler struct {
	ChallengeExecutionService *service.ChallengeExecutionService
}

func (handler *ChallengeExecutionHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	num, _ := strconv.Atoi(id)
	err := handler.ChallengeExecutionService.Delete(num)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}
