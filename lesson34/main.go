package main

import (
	"fmt"
	"my_mod/handler"
	"my_mod/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	product := postgres.NewProRepo(db)

	server := handler.NewHandler(handler.Handler{Product: product})

	err = server.ListenAndServe()
	fmt.Println("salom")
	if err != nil {
		panic(err)
	}
}
