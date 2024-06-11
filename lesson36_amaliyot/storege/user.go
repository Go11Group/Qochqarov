package postgres

import (
	"database/sql"

	"my_mod/model"
)

type UserRepo struct {
	Db *sql.DB
}

func NewProRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (p *UserRepo) U_Create(user model.Users) error {

	_, err := p.Db.Exec("insert into users(id,first_name,last_name,age,email) values($1,$2,$3,$4,$5)", user.Id, user.First_name, user.Last_name, user.Age, user.Email)
	if err != nil {
		return err
	}

	return nil
}

func (pa *UserRepo) U_Read(user model.Users,id int) ([]model.Users, error) {
	rows, err := pa.Db.Query("select * from users where id=$1;",id)
	if err != nil {
		return nil, err
	}

	var p []model.Users
	for rows.Next() {
		err = rows.Scan(&user.Id, &user.First_name, &user.Last_name, &user.Age, &user.Email)
		if err != nil {
			return nil, err
		}
		p = append(p, user)
	}
	return p, nil
}

func (p *UserRepo) U_Update(age int, First_name string) error {

	_, err := p.Db.Exec("update users set age=$1 where first_name=$2 ", age, First_name)
	if err != nil {
		return err
	}

	return nil
}

func (p UserRepo) U_Delete(id int) error {

	_, err := p.Db.Exec("delete from users where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
