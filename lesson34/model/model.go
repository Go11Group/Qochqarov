package model

type Product struct {
	Id string
	Name       string
	Price int
	Year int
}

type User struct{
	Id string
	Name string
	Age int
	Year int
	Product_id string
}
