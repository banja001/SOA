package handler

import (
	"Followers/model"
	"Followers/service"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type KeyProduct struct{}

type FollowerHandler struct {
	logger  *log.Logger
	service *service.FollowerService
}

func NewFollowerHandler(service *service.FollowerService, logger *log.Logger) *FollowerHandler {
	return &FollowerHandler{
		service: service,
		logger:  logger,
	}
}

func (f *FollowerHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		f.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}

func (fh *FollowerHandler) GetAllFollowers(rw http.ResponseWriter, h *http.Request) {
	persons, err := fh.service.GetAllPersonsNodes()
	if err != nil {
		fh.logger.Println("Database exception:", err)
		http.Error(rw, "Database exception", http.StatusInternalServerError)
		return
	}
	err = persons.ToJSON(rw)
	if err != nil {
		fh.logger.Println("Error converting to JSON:", err)
		http.Error(rw, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
}
func (fh *FollowerHandler) IsFollowed(rw http.ResponseWriter, h *http.Request) {
	id := mux.Vars(h)["id"]
	num, _ := strconv.Atoi(id)
	id2 := mux.Vars(h)["uid"]
	num2, _ := strconv.Atoi(id2)
	isFollowed, err := fh.service.IsFollowed(num, num2)

	if err != nil {
		fh.logger.Println("Database exception:", err)
		http.Error(rw, "Database exception", http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(isFollowed)
	if err != nil {
		fh.logger.Println("Error converting to JSON:", err)
		http.Error(rw, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonData)
}

func (fh *FollowerHandler) GetAllRecomended(rw http.ResponseWriter, h *http.Request) {
	id := mux.Vars(h)["id"]
	num, _ := strconv.Atoi(id)
	id2 := mux.Vars(h)["uid"]
	num2, _ := strconv.Atoi(id2)
	persons, err := fh.service.GetAllRecomended(num, num2)

	if err != nil {
		fh.logger.Println("Database exception:", err)
		http.Error(rw, "Database exception", http.StatusInternalServerError)
		return
	}

	err = persons.ToJSON(rw)
	if err != nil {
		fh.logger.Println("Error converting to JSON:", err)
		http.Error(rw, "Error converting to JSON", http.StatusInternalServerError)
		return
	}
}

func (f *FollowerHandler) CreateFollower(rw http.ResponseWriter, r *http.Request) {
	updatedFollower := &model.Follower{}
	err := json.NewDecoder(r.Body).Decode(updatedFollower)
	if err != nil {
		f.logger.Println("Error decoding request body:", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err = f.service.RewriteFollower(updatedFollower)
	if err != nil {
		f.logger.Print("Error updating follower:", err)
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
