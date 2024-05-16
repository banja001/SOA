package main

import (
	stakeholder_service "api-gateway/proto/stakeholder-service"
	tour_service "api-gateway/proto/tour-service"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	//cfg := config.GetConfig()

	conn, err := grpc.DialContext(
		context.Background(),
		os.Getenv("STAKEHOLDERS_SERVICE_ADDRESS"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()

	client := stakeholder_service.NewStakeholderServiceClient(conn)
	err = stakeholder_service.RegisterStakeholderServiceHandlerClient(
		context.Background(),
		gwmux,
		client,
	)
	if err != nil {
		log.Fatalln("Failed to register gateway stakeholders:", err)
	}


	// client tours
	conn_tours, err := grpc.DialContext(
		context.Background(),
		os.Getenv("STAKEHOLDERS_SERVICE_ADDRESS"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux_tours := runtime.NewServeMux()

	client_tours := tour_service.NewTourServiceClient(conn_tours)
	err = tour_service.RegisterTourServiceHandlerClient(
		context.Background(),
		gwmux_tours,
		client_tours,
	)

	if err != nil {
		log.Fatalln("Failed to register gateway tours:", err)
	}

	gwServer := &http.Server{
		Addr:    os.Getenv("GATEWAY_ADDRESS"),
		Handler: gwmux,
	}

	log.Println("Serving gRPC-Gateway on http://0.0.0.0:9000")

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err = gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}

}
