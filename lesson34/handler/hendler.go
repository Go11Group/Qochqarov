package handler

import (
	"my_mod/storage/postgres"
	"net/http"
)


type Handler struct {
	Product *postgres.ProductRepo
}

func NewHandler(handler Handler) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("DELETE /api/product", handler.Products)

	return &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
}


