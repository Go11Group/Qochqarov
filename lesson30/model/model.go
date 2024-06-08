package model

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	First_name string
	LastName   string
	Email      string
	Password   string
	Age        int
	Field      string
	Gender     string
	IsEmployee bool
}
type Product struct{
	gorm.Model
	Name string
	Year int
}
