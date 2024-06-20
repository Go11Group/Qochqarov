package main

import "net/http"

func main() {
	http.Get("localhost:8080")

	client := http.Client{}
	r := http.Request{Method: "UPDATE"}
	client.Do(&r)
}