package main

import (
	handler "my_mod/hendler"
	"my_mod/storage"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	user := postgres.NewProRepo(db)
	problem := postgres.NewProblemRepo(db)
	solved := postgres.NewSolvedRepo(db)

	hand := handler.Handler{Userss: user, Problems: problem, Solvedproblems: solved}
	handler.NewHandler(&hand)

	// var k int
	// fmt.Println("qanday malumotlar bn ishlamoqchisiz 1-user,2-problem, 3-solved_problem:! ")
	// fmt.Scanf("%d", &k)
	// var server *http.Server

	// if k == 1 {
	// 	server = handler.NewHandler(handler.Handler{Userss: user})
	// } else if k == 2 {
	// 	server = handler.NewHandler(handler.Handler{Problems: problem})
	// } else if k == 3 {
	// 	server = handler.NewHandler(handler.Handler{Solvedproblems: solved})
	// }

	// err = server.ListenAndServe()
	// fmt.Println("salom")
	// if err != nil {
	// 	panic(err)
	// }
}
