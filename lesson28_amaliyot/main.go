package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"

	"github.com/google/uuid"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "go"
	password = "1918"
)

type User struct {
	id   string
	name string
	age  int
}

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbname, password)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		fmt.Println("hi")
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	user := User{
		id:   uuid.NewString(),
		name: "ali",
		age:  12,
	}
	_, err = db.Exec("insert into student(id,name,age) values($1,$2,$3)", &user.id, &user.name, &user.age)
	if err != nil {
		panic(err)
	}

	// _, err = db.Exec("update student set name ='vali' where age=12")
	// if err != nil {
	// 	panic(err)
	// }
	update(db)
	delete(db)
}


func update(db *sql.DB)  {
	
	_, err := db.Exec("update student set name ='vali' where age=24")
	if err != nil {
		panic(err)
	}
}


func delete(db *sql.DB)  {
	
	_, err := db.Exec("delete from student where age=12")
	if err != nil {
		panic(err)
	}
}
