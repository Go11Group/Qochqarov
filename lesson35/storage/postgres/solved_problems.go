package postgres

import (
	"database/sql"

	"my_mod/model"
)

type SolvedRepo struct {
	Db *sql.DB
}

func NewSolvedRepo(db *sql.DB) *SolvedRepo {
	return &SolvedRepo{Db: db}
}

func (p *SolvedRepo) S_Create(solved model.Solved_problems) error {

	_, err := p.Db.Exec("insert into solved_problems(id,name,degre,user_id) values($1,$2,$3,$4)", solved.Id, solved.Name, solved.Degre,solved.User_id)
	if err != nil {
		return err
	}

	return nil
}

func (pa *SolvedRepo) S_Read(solved model.Solved_problems) ([]model.Solved_problems, error) {
	rows, err := pa.Db.Query("select * from solved_problems;")
	if err != nil {
		return nil, err
	}

	var p []model.Solved_problems
	for rows.Next() {
		err = rows.Scan(&solved.Id,&solved.Name,&solved.Degre,&solved.User_id)
		if err != nil {
			return nil, err
		}
		p = append(p, solved)
	}
	return p, nil
}

func (p *SolvedRepo) S_Update(id int, name string) error {

	_, err := p.Db.Exec("update solved_problems set name=$1 where id=$2 ", name, id)
	if err != nil {
		return err
	}

	return nil
}

func (p *SolvedRepo) S_Delete(id int) error {

	_, err := p.Db.Exec("delete from solved_problems where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
