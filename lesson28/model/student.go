package model

type User struct {
	ID     string
	Name   string
	Age    int

}

type Users struct{
	User []User
}
