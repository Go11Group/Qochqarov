package main

import (
	"net/http"
	"my_mod/modul"
)

func main() {
	http.HandleFunc("GET /task",modul.Modul)
	http.ListenAndServe(":8080", nil)
}
