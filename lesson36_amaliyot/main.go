package main

import (
	handler "my_mod/handler"
	"my_mod/storege"
	"net/http"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	user := postgres.NewProRepo(db)
	var server *http.Server

	server = handler.NewHandler(handler.Handler{Userss: user})

	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
