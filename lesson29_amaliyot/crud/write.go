package crud

import (
	"database/sql"
	"my_mod/model"
	_ "github.com/lib/pq"
)

type PersonRepo struct {
	Db *sql.DB
}

func NewPerson(db *sql.DB) *PersonRepo {
	return &PersonRepo{Db: db}
}

func (s *PersonRepo)Create( user model.User) error {
	_,err:=s.Db.Exec("insert into person(id,first_name,last_name,email,gender,age) values($1,$2,$3,$4,$5,$6)",user.ID,user.Firs_name,user.Last_name,user.Email,user.Gender,user.Age)
	if err != nil {
		return err
	}

	return nil
}