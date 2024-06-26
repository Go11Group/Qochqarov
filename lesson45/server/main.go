package main

import (
	"log"
	"my_mod/library"
	"my_mod/method"
	"my_mod/postgres"
	"net"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	lib := postgres.NewLibraryStorage(db)

	// server := method.Server{Library: lib}

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println(err)
	}

	defer listen.Close()

	s := grpc.NewServer()

	library.RegisterLibraryServerServer(s, &method.Server{Library: lib})

	err=s.Serve(listen)
	if err != nil {
		log.Println(err)
	}

}
