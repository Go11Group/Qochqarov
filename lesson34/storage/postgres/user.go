package postgres

import (
	"database/sql"
	"fmt"

	"my_mod/model"
	"github.com/google/uuid"
)

func U_Create(db *sql.DB, user model.User) error {

	user = model.User{
		Id:         uuid.NewString(),
		Name:       "matem",
		Age:        34,
		Year:       324000,
		Product_id: "7d81a963-bf68-4ac0-a1ed-0168d602909b",
	}

	_, err := db.Exec("insert into users(id,name,age,year,product_id) values($1,$2,$3)", user.Id, user.Name, user.Age, user.Year, user.Product_id)
	if err != nil {
		return err
	}

	return nil
}

func U_Read(db *sql.DB, user model.User) error {
	rows, err := db.Query("select * from users;")
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&user.Id, &user.Name, &user.Year, &user.Product_id)
		if err != nil {
			return err
		}
		fmt.Println(user)
	}

	return nil

}

func U_Update(db *sql.DB, user model.User) error {
	new_name := "karim"

	_, err := db.Exec("update users set name=$1 where id='771aaa8e-be1a-4522-a4c0-757e1ee76905' ", new_name)
	if err != nil {
		return err
	}

	return nil
}

func U_Delete(db *sql.DB) error {
	id := "a8e27c9c-d6cb-4236-9413-205d8294dc06"

	_, err := db.Exec("delete from users where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
