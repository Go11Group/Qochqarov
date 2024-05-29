package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "go"
	password = "1918"
)

type User struct {
	id     string
	Name   string
	age    int
	C_name string
}

func main() {
	conn := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		host, port, user, dbname, password)
	db, err := sql.Open("postgres", conn)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	user := User{}
	err = db.QueryRow("select s.id, s.name,s.age,c.name from curse as c join student as s on s.id=c.student_id where s.age>$1 offset 3;", 25).
		Scan(&user.id, &user.Name, &user.age, &user.C_name)

	if err != nil {
		fmt.Println("hi")
		panic(err)
	}
	fmt.Println("queryrow", user)
	rows, err := db.Query("select s.id, s.name,s.age,c.name from curse as c join student as s on s.id=c.student_id where s.age>$1;", 25)

	if err != nil {
		panic(err)
	}

	for rows.Next() {

		err := rows.Scan(&user.id, &user.Name, &user.age, &user.C_name)
		if err != nil {
			panic(err)
		}
		fmt.Println(user)
	}

}
