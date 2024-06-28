package main

import (
	"log"
	"my_Ipa/api"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	router := api.Router(conn)

	err = router.Run(":8080")

	if err != nil {
		log.Fatal(err)
	}
}
