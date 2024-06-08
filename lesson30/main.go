package main

import (
	"fmt"
	"may_mod/crud"
	"may_mod/model"
)

func main() {
	db, err := crud.ConnectDb()
	if err != nil {
		panic(err)
	}
	var user *model.Users
	crud.CreateS(*db, model.Users{})

	user = crud.Read(*db, model.Users{})
	fmt.Println(user)
	crud.Update(*db, model.Users{})

	crud.Delete(*db, model.Users{})

	crud.CreateS(*db, model.Users{})

	praduct := crud.P_Read(*db, model.Product{})
	fmt.Println(praduct)
	crud.P_Update(*db, model.Product{})

	crud.P_Delete(*db, model.Product{})
}
