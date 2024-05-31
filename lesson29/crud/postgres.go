package crud

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open("postgres://postgres:1918@localhost:5432/go?sslmode=disable"))
	if err != nil {
		return nil, err
	}

	return db, nil
}
