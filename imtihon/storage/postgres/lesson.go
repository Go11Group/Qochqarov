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

type LessonRepo struct {
	Db *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{Db: db}
}

func (p *LessonRepo) LessonCreate(lessons model.Lessons) error {

	_, err := p.Db.Exec("insert into LESSONS(id,course_id,title,Content,create_at,update_at) values($1,$2,$3,$4,$5,$6)",
		uuid.NewString, lessons.CourseId, lessons.Title, lessons.Content, time.Now(), time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (pa *LessonRepo) LessonRead(lessons model.Lessons) ([]model.Lessons, error) {
	rows, err := pa.Db.Query("select Id,course_id,titlr, Content from lessons;")
	if err != nil {
		return nil, err
	}

	var p []model.Lessons
	for rows.Next() {
		err = rows.Scan(&lessons.Id, &lessons.CourseId, &lessons.Title, &lessons.Content)
		if err != nil {
			return nil, err
		}
		p = append(p, lessons)
	}
	return p, nil
}

func (u *LessonRepo) LessonUpdate(lessonsUpdateFilter model.UpdateLesson) error {
	var params []string
	var args []interface{}
	fmt.Println(*lessonsUpdateFilter.Content)
	query := `
	SELECT id
	FROM lessons
	WHERE delete_at IS NULL AND id = $1
	`

	if err := u.Db.QueryRow(query, lessonsUpdateFilter.LessonId).Err(); err != nil {
		return fmt.Errorf("lessons by this id not found: %v", err)
	}

	query = `
	UPDATE lessons SET 
	`

	if lessonsUpdateFilter.Title != nil {
		params = append(params, fmt.Sprintf(" title = $%d ", len(args)+1))
		args = append(args, *lessonsUpdateFilter.Title)
	}

	if lessonsUpdateFilter.CourseId != nil {
		params = append(params, fmt.Sprintf(" course_id = $%d ", len(args)+1))
		args = append(args, *lessonsUpdateFilter.CourseId)
	}

	if lessonsUpdateFilter.Content != nil {
		params = append(params, fmt.Sprintf("content = $%d", len(args)+1))
		args = append(args, *lessonsUpdateFilter.Content)
	}

	params = append(params, fmt.Sprintf("updated_at = $%d", len(args)+1))
	args = append(args, time.Now())

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, lessonsUpdateFilter.LessonId)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE id = $%d AND delete_at IS NULL", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err := u.Db.Exec(query, args...)

	fmt.Println(query)
	if err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}

	return nil
}

func (u LessonRepo) LessonDeleted(id string) error {

	_rows, err := u.Db.Exec(`update lessons set
	delete_at = date_part('epoch', current_timestamp)::INT
   where id = $1 and delete_at = 0`, id)
	if err != nil {
		return err
	}
	rowsaff, err := _rows.RowsAffected()
	if err != nil {
		return nil
	}

	if rowsaff == 0 {
		return err

	}

	return nil
}

func (u *LessonRepo) GetAllLesson(f model.LessonGetAll) ([]model.Lessons, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)

	query := `SELECT id,title,course_Id ,content
	FROM lessons WHERE true `

	if len(f.Title) > 0 {
		params["title"] = f.Title
		filter += " and title = :title "
	}

	if len(f.CourseId) > 0 {
		params["bi"] = f.CourseId
		filter += " and CourseId = :courseId "
	}

	if len(f.Content) > 0 {
		params["content"] = f.Content
		filter += " and content = :content"
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

	var lesson []model.Lessons
	for rows.Next() {
		var lessons model.Lessons
		err := rows.Scan(&lessons.Id, &lessons.Title, &lessons.CourseId, &lessons.Content)

		if err != nil {
			return nil, err
		}

		lesson = append(lesson, lessons)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return lesson, nil
}
