package main

import ("net/http"
	"my_mod/method"
)



func main()  {
	mux:=http.NewServeMux()
	
	mux.HandleFunc("GET /getadd", method.Get)
	mux.HandleFunc("GET /postadd", method.Post)
	mux.HandleFunc("Delete /deladd", method.Put)
	mux.HandleFunc("put /putadd", method.Delete)

	err:=http.ListenAndServe(":8070",mux)
	if err != nil {
		panic(err)
	}
	
}	