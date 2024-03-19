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

func (handler *ChallengeHandler) Update(writer http.ResponseWriter, req *http.Request) {
	var challenge model.Challenge
	err := json.NewDecoder(req.Body).Decode(&challenge)
	if err != nil {
		println("Error while parsing json: Update challenge")
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	updatedChallenge, err := handler.ChallengeService.Update(&challenge)
	if err != nil {
		println("Error while updating a challenge")
		writer.WriteHeader(http.StatusExpectationFailed)
		return
	}
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(writer).Encode(updatedChallenge); err != nil {
		println("Error while encoding tour to JSON")
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
}
