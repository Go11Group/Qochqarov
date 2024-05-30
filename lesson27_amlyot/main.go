package main

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "postgres"
	password = "1918"
)

type User struct {
	ID     string
	Name   string
	Age    int
	Gender string
	Course string
}

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Hello")
		panic(err)
	}

	user := User{}
	err = db.QueryRow(`select s.id, s.name, age, gender, c.name from student s
    	left join cours c on c.id = s.cours_id offset 2`).
		Scan(&user.ID, &user.Name, &user.Age, &user.Gender, &user.Course)
	if err != nil {
		panic(err)
	}
	
	_ , err = db.Exec("insert into student(id, name, age) values ($1, $2, $3)",
		uuid.NewString(), "Ibrohim",6)
	if err != nil {
		panic(err)
	}
	db.Exec("update student set name=$1, age=$2  where age=34", "wdfe", 234)
	_,err=db.Exec("delete from student where age=6 ")
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

	fmt.Println("Successfully connected!")
}