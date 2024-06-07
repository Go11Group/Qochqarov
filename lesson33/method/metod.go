package method

import (
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte("get func"))
	if err != nil {
		panic(err)
	}
}


func Post(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("post func"))
	if err != nil {
		panic(err)
	}
}
func Put(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("put func"))
	if err != nil {
		panic(err)
	}
}

func Delete(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("delete func"))
	if err != nil {
		panic(err)
	}
}
