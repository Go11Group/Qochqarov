package main

import (
	"encoding/json"
	"fmt"
	"os"

	"my_mod/crud"
	"my_mod/model"
)

func main() {
	user := []model.User{}
	f, err := os.OpenFile("user.json", os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	json.NewDecoder(f).Decode(&user)
	fmt.Println(user)

	db, err := crud.ConnectDB()
	if err != nil {
		panic(err)
	}
	pr:=crud.NewPerson(db)

	for _, v := range user {
		err := pr.Create(v)
		if err != nil {
			panic(err)
		}

	}
}
