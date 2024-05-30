package postgres

import (
	"database/sql"
	"fmt"

	"github.com/Go11Group/at_lesson/lesson28/model"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(DB *sql.DB) *CourseRepo {
	return &CourseRepo{DB}
}
func (u *CourseRepo) GetAllCourse() ([]model.Course, error) {
	rows, err := u.DB.Query(`select c.id, c.name, c.student_id from student as s
					right join curse as c on s.id = c.student_id`)
	if err != nil {
	
		return nil, err
	}

	var courses []model.Course
	var course model.Course
	for rows.Next() {
		err = rows.Scan(&course.Id, &course.Name, &course.Student_id)
		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (u *CourseRepo) GetByID() (model.Course, error) {
	var course model.Course

	err := u.DB.QueryRow(`select s.id, s.name, age, c.name from student s
					left join course c on s.id = c.student_id`).
		Scan(&course.Id, &course.Name, &course.Student_id)
	if err != nil {
		return model.Course{}, err
	}

	return course, nil
}

func C_Create(db *sql.DB, Course model.Course) error {

	Course = model.Course{
		Id:         uuid.NewString(),
		Name:       "matem",
		Student_id: "5a580aec-9e2f-4e91-ba88-06be8ec3dfc4",
	}

	_, err := db.Exec("insert into curse(id,name,student_id) values($1,$2,$3)", Course.Id, Course.Name, Course.Student_id)
	if err != nil {
		return err
	}

	return nil
}

func C_Read(db *sql.DB, curse model.Course) error {
	rows, err := db.Query("select * from curse;")
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&curse.Id, &curse.Name, &curse.Student_id)
		if err != nil {
			return err
		}
		fmt.Println(curse)
	}

	return nil

}

func C_Update(db *sql.DB, curse model.Course) error {
	new_name := "tarix"

	_, err := db.Exec("update curse set name=$1 where id='771aaa8e-be1a-4522-a4c0-757e1ee76905' ", new_name)
	if err != nil {
		return err
	}

	return nil
}

func C_Delete(db *sql.DB) error {
	id := "a8e27c9c-d6cb-4236-9413-205d8294dc06"

	_, err := db.Exec("delete from student where id=$1", id)
	if err != nil {
		return err
	}

	return nil
}
