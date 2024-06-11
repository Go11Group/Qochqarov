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

	// server := handler.NewHandler(&hand)
	// server.ListenAndServe()

}
