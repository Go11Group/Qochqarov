package modul

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type Person struct {
	Id   int
	Name string
	Year int
	Age  int
}

func Modul(w http.ResponseWriter, r *http.Request) {
	f, err := os.OpenFile("file.json",os.O_RDONLY,0666)
	if err != nil {
		panic(err)
	}
	switch r.Method {
	case http.MethodPost:
		byt, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}
		f.Write(byt)
	case http.MethodGet:
		person := Person{
			Id:   1,
			Name: "salim",
			Year: 2022,
			Age:  23,
		}
		byt2, err := json.Marshal(person)
		if err != nil {
			panic(err)
		}
		w.Write(byt2)
	}
}
