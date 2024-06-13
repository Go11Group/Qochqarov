package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	pakage "my_mod/Pakage"
	"my_mod/model"

	"github.com/google/uuid"
)

type CousresRepo struct {
	Db *sql.DB
}

func NewCourseRepo(db *sql.DB) *CousresRepo {
	return &CousresRepo{Db: db}
}

func (p *CousresRepo) CoursCreate(course model.Courses) error {

	_, err := p.Db.Exec("insert into course(id,title,Description,create_at,update_at) values($1,$2,$3,$4,$5)",
		uuid.NewString, course.Title, course.Description, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (pa *CousresRepo) CoursRead(course model.Courses) ([]model.Courses, error) {
	rows, err := pa.Db.Query("select Id,title, description from courses")
	if err != nil {
		return nil, err
	}

	var p []model.Courses
	for rows.Next() {
		err = rows.Scan(&course.Id, &course.Title, &course.Description)
		if err != nil {
			return nil, err
		}
		p = append(p, course)
	}
	return p, nil
}

func (u *CousresRepo) CoursUpdate(courseUpdateFilter model.UpdateCourse) error {
	var params []string
	var args []interface{}
	query := `
	SELECT id
	FROM courses
	WHERE delete_at IS NULL AND id = $1
	`

	if err := u.Db.QueryRow(query, courseUpdateFilter.CourseId).Err(); err != nil {
		return fmt.Errorf("courses by this id not found: %v", err)
	}

	query = `
	UPDATE courses SET 
	`

	if courseUpdateFilter.Title != nil {
		params = append(params, fmt.Sprintf(" title = $%d ", len(args)+1))
		args = append(args, *courseUpdateFilter.Title)
	}

	if courseUpdateFilter.Description != nil {
		params = append(params, fmt.Sprintf("description = $%d", len(args)+1))
		args = append(args, *courseUpdateFilter.Description)
	}

	params = append(params, fmt.Sprintf("update_at = $%d", len(args)+1))
	args = append(args, time.Now())

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, courseUpdateFilter.CourseId)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE id = $%d AND delete_at IS NULL", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err := u.Db.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}

	fmt.Println(query)
	return nil
}
func (u CousresRepo) CoursDelete(id string) error {

	_rows, err := u.Db.Exec(`update courses set
	delete_at = date_part('epoch', current_timestamp)::INT
   where id = $1 and delete_at = 0`, id)
	if err != nil {
		return err
	}
	rowsaff,err:=_rows.RowsAffected()
	if err != nil {
		return nil
	}

	if rowsaff==0 {
		return err
		
	}

	return nil
}

func (u *CousresRepo) GetAllCourse(f model.CourseaGetAll) ([]model.Courses, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)
	fmt.Println("salom")

	query := `SELECT id,title,Description
	FROM courses WHERE true `

	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += " and title = :title "
	}

	if len(f.Description) > 0 {
		params["Description"] = f.Description
		filter += " and Description = :Description "
	}

	if f.Offset > 0 {
		params["offset"] = f.Offset
		filter += " OFFSET :offset"
	}

	if f.Limit > 0 {
		params["limit"] = f.Limit
		filter += " LIMIT :limit"
	}
	query = query + filter

	query, arr = pakage.ReplaceQueryParams(query, params)
	fmt.Println(query, arr)
	rows, err := u.Db.Query(query, arr...)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	var courses []model.Courses
	for rows.Next() {
		var course model.Courses
		err := rows.Scan(&course.Id, &course.Title, &course.Description)

		if err != nil {
			return nil, err
		}

		courses = append(courses, course)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
