package main

import (
	"log"
	"my_mod/api/handler"
	"my_mod/api/handler/router"
	"my_mod/config"
	"my_mod/generated/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.Load()
	hand := NewConnect()

	router := router.RouterApi(hand)

	log.Fatal(router.Run(cfg.HTTP_PORT))
}

func NewConnect() *handler.Handler {
	usersConn, err := grpc.Dial("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil
	}

	usersService := users.NewUserServiceClient(usersConn)

	return &handler.Handler{
		User: usersService,
	}
}
