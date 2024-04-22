package handler

import (
	"Followers/model"
	"Followers/service"
	"encoding/json"
	"log"
	"net/http"
)

type KeyProduct struct{}

type FollowerHandler struct {
	logger *log.Logger
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
    followers, err := fh.service.GetAllFollowers()
    if err != nil {
        fh.logger.Println("Database exception:", err)
        http.Error(rw, "Database exception", http.StatusInternalServerError)
        return
    }

    err = followers.ToJSON(rw)
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