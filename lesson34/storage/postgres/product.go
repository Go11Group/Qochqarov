package postgres

import (
	"database/sql"
	"fmt"

	"my_mod/model"

	"github.com/google/uuid"
)

type ProductRepo struct {
	Db *sql.DB
}

func NewProRepo(db *sql.DB) *ProductRepo {
	return &ProductRepo{Db: db}
}

func (p *ProductRepo) C_Create(product model.Product) error {
	fmt.Println(product.Name)

	_, err := p.Db.Exec("insert into product(id,name,price,year) values($1,$2,$3,$4)", uuid.NewString(), product.Name, product.Price, product.Year)
	if err != nil {
		return err
	}

	return nil
}

// func C_Read(db *sql.DB, product model.Product) error {
// 	rows, err := db.Query("select * from product;")
// 	if err != nil {
// 		return err
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(&product.Id, &product.Name, &product.Price, &product.Year)
// 		if err != nil {
// 			return err
// 		}
// 		fmt.Println(product)
// 	}

// 	return nil

// }

// func C_Update(db *sql.DB, curse model.Product, id string) error {
// 	new_name := "gilos"

// 	_, err := db.Exec("update product set name=$1 where id=&2 ", new_name, id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func C_Delete(db *sql.DB, id string) error {

// 	_, err := db.Exec("delete from product where id=$1", id)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
