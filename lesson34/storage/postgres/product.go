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

func (pa *ProductRepo) C_Read(product model.Product) ([]model.Product,error) {
	rows, err := pa.Db.Query("select * from product;")
	if err != nil {
		return nil,err
	}

	var p []model.Product
	for rows.Next() {
		 err = rows.Scan(&product.Id, &product.Name, &product.Price,&product.Year)
		if err != nil {
			return nil,err
		}
		p=append(p, product)
		}
	return p,nil
}

func (p *ProductRepo)C_Update( year int,name string) error {

	_, err := p.Db.Exec("update product set name=$1 where year=$2 ", name, year)
	if err != nil {
		return err
	}

	return nil
}

func (p ProductRepo) C_Delete(year int) error {

	_, err := p.Db.Exec("delete from product where year=$1", year)
	if err != nil {
		return err
	}

	return nil
}
