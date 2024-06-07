package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type person struct{
	id int
	name string
}

var hashmap = map[string]string{"1": "ali", "2": "vali", "3": "salim"}

func main() {
	http.HandleFunc("/map/1", hender)
	http.HandleFunc("/sever/1", server)
	http.ListenAndServe(":8060", nil)
}

func hender(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(hashmap[strings.TrimPrefix(r.URL.Path, "/map/")]))
}

func server(w http.ResponseWriter,r *http.Request)  {
	var p person
	err:=json.NewDecoder(r.Body).Decode(p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p)
	w.Write([]byte("salom"))

}
