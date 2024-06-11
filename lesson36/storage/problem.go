package postgres

import (
	"database/sql"

	"my_mod/model"
)

type ProblemRepo struct {
	Db *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{Db: db}
}

func (p *ProblemRepo) ReadById(id int) error {
	rows, err := p.Db.Query("select * from problems where id=$1", id)
	if err != nil {
		return err
	}
	var pr []model.Problem
	for rows.Next() {
		var problem model.Problem
		err = rows.Scan(&problem.Id, &problem.Name, &problem.Text)
		if err != nil {
			return err
		}
		pr = append(pr, problem)
	}
	return nil

}

func (p *ProblemRepo) P_Create(problem model.Problem) error {

	_, err := p.Db.Exec("insert into problems(id,name,text) values($1,$2,$3)", problem.Id, problem.Name, problem.Text)
	if err != nil {
		return err
	}

	return nil
}

func (pa *ProblemRepo) P_Read(problem model.Problem) ([]model.Problem, error) {
	rows, err := pa.Db.Query("select * from problems;")
	if err != nil {
		return nil, err
	}

	var p []model.Problem
	for rows.Next() {
		err = rows.Scan(&problem.Id, &problem.Name, &problem.Text)
		if err != nil {
			return nil, err
		}
		p = append(p, problem)
	}
	return p, nil
}

func (p *ProblemRepo) P_Update(id int, name string) error {

	_, err := p.Db.Exec("update problems set firstname=$1 where id=$2 ", name, id)
	if err != nil {
		return err
	}

	return nil
}

func (p ProblemRepo) P_Delete(id int) error {

	_, err := p.Db.Exec("delete from problems where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
