package main

import (
	"Followers/data"
	"context"
	"log"
	"os"
	"time"
)

func main() {
	port := "8060"//os.Getenv("PORT")
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
	store, err := data.New(storeLogger)
	if err != nil {
		logger.Fatal(err)
	}
	defer store.CloseDriverConnection(timeoutContext)
	store.CheckConnection()
}