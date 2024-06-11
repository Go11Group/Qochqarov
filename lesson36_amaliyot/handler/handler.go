package handler

import (
	"my_mod/storege"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Userss *postgres.UserRepo
}

func NewHandler(handler Handler) *http.Server {
	//mux := http.NewServeMux()
	riuter := mux.NewRouter()
	user := riuter.PathPrefix("/user").Subrouter()
	user.HandleFunc("", handler.User).Methods(http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut)

	return &http.Server{
		Addr:    ":8090",
		Handler: riuter,
	}
}
