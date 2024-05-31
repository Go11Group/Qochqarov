package main

import (
	"fmt"
	"my_mod/crud"
	"my_mod/model"
)

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

// func main() {
// 	//dbURL :=
// 	db, _ := gorm.Open(postgres.Open("postgres://postgres:1918@localhost:5432/go?sslmode=disable"))

// 	db.AutoMigrate(&Product{})

// 	db.Create(&Product{Code: "D93", Price: 16004})
// 	db.Create(&Product{Code: "D453", Price: 5876})

// 	var product Product
// 	db.First(&product, 1)
// 	//db.First(&product, "code = ? and price = ?", "D42", 100) // find product with code D42
// 	//db.Model(&product).Where("code = ? and price = ?", "D42", 100).Update("Price", 200)

// 	// fmt.Println(product)

// 	// db.Delete(&product, 1)

// }

// type User struct {
// 	gorm.Model
// 	FirstName  string
// 	LastName   string
// 	Email      string
// 	Password   string
// 	Age        int
// 	Field      string
// 	Gender     string
// 	IsEmployee bool
// }

func main() {
	db, err := crud.ConnectDb()
	if err != nil {
		panic(err)
	}
	var user *model.Users
	crud.CreateS(*db, model.Users{})
	
	user=crud.Read(*db,model.Users{})
	fmt.Println(user)
	crud.Update(*db,model.Users{})

	crud.Delete(*db,model.Users{})
}
