package main

import (
	"Followers/handler"
	"Followers/repo"
	"Followers/service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	gorillaHandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8060"
	}

	// Initialize context
	timeoutContext, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	//Initialize the logger we are going to use, with prefix and datetime for every log
	logger := log.New(os.Stdout, "[followers-api] ", log.LstdFlags)
	storeLogger := log.New(os.Stdout, "[followers-store] ", log.LstdFlags)

	// NoSQL: Initialize Follower Repository store
	store, err := repo.New(storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.CloseDriverConnection(timeoutContext)
	store.CheckConnection()

	//Initialize the service and inject said logger
	followerService := service.NewFollowerService(store, logger)

	//Initialize the handler and inject said logger
	followerHandler := handler.NewFollowerHandler(followerService, logger)

	//Initialize the router and add a middleware for all the requests
	router := mux.NewRouter()

	router.Use(followerHandler.MiddlewareContentTypeSet)

	getAllFollowers := router.Methods(http.MethodGet).Subrouter()
	getAllFollowers.HandleFunc("/followers", followerHandler.GetAllFollowers)

	getAllFollowed := router.Methods(http.MethodGet).Subrouter()
	getAllFollowed.HandleFunc("/followers/recommended/{id}/{uid}", followerHandler.GetAllFollowed)

	putFollower := router.Methods(http.MethodPut).Subrouter()
	putFollower.HandleFunc("/followers/update", followerHandler.CreateFollower)

	cors := gorillaHandlers.CORS(gorillaHandlers.AllowedOrigins([]string{"*"}))

	//Initialize the server
	server := http.Server{
		Addr:         ":" + port,
		Handler:      cors(router),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	logger.Println("Server listening on port", port)
	//Distribute all the connections to goroutines
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()

	sigCh := make(chan os.Signal)
	signal.Notify(sigCh, os.Interrupt)
	signal.Notify(sigCh, os.Kill)

	sig := <-sigCh
	logger.Println("Received terminate, graceful shutdown", sig)

	//Try to shutdown gracefully
	if server.Shutdown(timeoutContext) != nil {
		logger.Fatal("Cannot gracefully shutdown...")
	}
	logger.Println("Server stopped")
}
