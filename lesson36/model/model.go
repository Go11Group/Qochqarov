package model

type Problem struct {
	Id   int
	Name string
	Text string
}

type User struct {
	Id         int
	First_name string
	Last_name  string
	Age        int
	Email      string
}

type Solved_problems struct {
	Id      int
	Name    string
	Degre   string
	User_id int
}
