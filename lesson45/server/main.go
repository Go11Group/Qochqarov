package main

import (
	"my_mod/library"
	"my_mod/method"
	"my_mod/postgres"
)



func main() {
	db,err:=postgres.ConnectDB()
	if err != nil {
		panic(err)
	}

	lib := postgres.NewLibraryStorage(db)

	server := method.Server{Library: lib}
	
}
