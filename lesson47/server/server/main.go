package main

import (
	"log"
	bussSer "my_mod/generated/bus"
	weatherSer "my_mod/generated/wheather"

	"my_mod/service/transport"
	postgres "my_mod/storege"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Connect to the PostgreSQL database
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Initialize repositories for bus and weather
	busRepo := postgres.NewBussRepo(db)
	weatherRepo := postgres.NewWeatherRepo(db)

	// Create a TCP listener on port 50051
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	defer listen.Close()

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the bus and weather services with the gRPC server
	bussSer.RegisterTransportServiceServer(s, &transport.Servers{Buss: busRepo})
	weatherSer.RegisterWitherServerServer(s, &transport.Server{With: weatherRepo})

	// Start the gRPC server
	log.Println("Starting gRPC server on port 50051...")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
