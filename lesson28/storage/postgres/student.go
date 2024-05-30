package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/at_lesson/lesson28/model"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type StudentRepo struct {
	Db *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{Db: db}
}

func (u *StudentRepo) GetAllStudents() ([]model.User, error) {
	rows, err := u.Db.Query(`select s.id, s.name, s.age from student as s
					left join curse as c on s.id = c.student_id`)
	if err != nil {
		return nil, err
	}

	var users []model.User
	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *StudentRepo) GetByID() (model.User, error) {
	var user model.User

	err := u.Db.QueryRow(`select s.id, s.name, age, c.name from student s
					left join course c on s.id= c.student_id`).
		Scan(&user.ID, &user.Name, &user.Age)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func Create(db *sql.DB, user model.User) error {

	user = model.User{
		ID:   uuid.NewString(),
		Name: "karim",
		Age:  34,
	}

	_, err := db.Exec("insert into student(id,name,age) values($1,$2,$3)", user.ID, user.Name, user.Age)
	if err != nil {
		return err
	}

	return nil
}

func Read(db *sql.DB, user model.User) error {
	rows, err := db.Query("select * from student;")
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return err
		}
		fmt.Println(user)
	}

	return nil

}

func Update(db *sql.DB, user model.User) error {
	new_age := 26

	_, err := db.Exec("update student set name='salim' where age=$1 ", new_age)
	if err != nil {
		return err
	}

	return nil
}

func Delete(db *sql.DB) error {
	id := "1322a310-1363-4115-b2d3-94848c867fbd"
	_, err := db.Exec("delete from student where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
