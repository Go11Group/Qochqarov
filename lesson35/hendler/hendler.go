package handler

import (
	"my_mod/storage/postgres"
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
	Userss         *postgres.UserRepo
	Problems       *postgres.ProblemRepo
	Solvedproblems *postgres.SolvedRepo
}

func NewHandler(handler Handler) *http.Server {
	//mux := http.NewServeMux()
	m := mux.NewRouter()

	m.HandleFunc("/api/user", handler.User).Methods(http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut)
	m.HandleFunc("/api/problem", handler.Problem).Methods(http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut)
	m.HandleFunc("/api/solved", handler.Solved_problems).Methods(http.MethodGet, http.MethodDelete, http.MethodPost, http.MethodPut)
	// m.HandleFunc("/api/solved", handler.GetBayId).Methods(http.MethodPost)

	return &http.Server{
		Addr:    ":8090",
		Handler: m,
	}
}
