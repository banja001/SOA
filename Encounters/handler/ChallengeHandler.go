package handler

import (
	"encgo/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ChallengeHandler struct {
	ChallengeService *service.ChallengeService
}

func (handler *ChallengeHandler) GetAll(writer http.ResponseWriter, req *http.Request) {
	challenges, err := handler.ChallengeService.GetAll()
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(writer).Encode(map[string]string{"error": err.Error()})
		return
	}
	//fmt.Println("Uspeno izvrseno")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(challenges)
	fmt.Println("Uspeno izvrsena GetAll metoda")
}

func (handler *ChallengeHandler) Delete(writer http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	num, _ := strconv.Atoi(id)
	err := handler.ChallengeService.Delete(num)
	writer.Header().Set("Content-Type", "application/json")
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
}
