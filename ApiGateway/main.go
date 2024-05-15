package main

import (
	stakeholder_service "api-gateway/proto/stakeholder-service"
	user_experience_service "api-gateway/proto/user-experience-service"
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

	userExperienceConn, err := grpc.DialContext(
		context.Background(),
		os.Getenv("ENCOUNTERS_SERVICE_ADDRESS"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		log.Fatalln("Failed to dial user experience server:", err)
	}

	gwmux := runtime.NewServeMux()

	client := stakeholder_service.NewStakeholderServiceClient(conn)
	err = stakeholder_service.RegisterStakeholderServiceHandlerClient(
		context.Background(),
		gwmux,
		client,
	)

	client2 := user_experience_service.NewUserExperienceServiceClient(userExperienceConn)
	err = user_experience_service.RegisterUserExperienceServiceHandlerClient(
		context.Background(),
		gwmux,
		client2,
	)

	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
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
