package postgres

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	pakage "my_mod/Pakage"
	"my_mod/model"
)

type EnrollmentsRepo struct {
	Db *sql.DB
}

func NewEnrollmentsRepo(db *sql.DB) *EnrollmentsRepo {
	return &EnrollmentsRepo{Db: db}
}

//bu method databasesdagi enrolments tablega yangi malumot qoshish ucun ishlatiladi

func (p *EnrollmentsRepo) EmrolCreate(enrol model.Enrollments) error {

	_, err := p.Db.Exec("insert into Enrollments(user_id,course_id,enrollment_date,created_at,update_at) values($1,$2,$3,$4,$5)",
		enrol.UserId, enrol.CourseId, enrol.EnrollmentDate, time.Now(), time.Now())
		fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}

//bu method databasesdagi enrolmentsni  barcha malumot larini oqish ucun hizmat qiladi

func (pa *EnrollmentsRepo) EmrolRead(enrol model.Enrollments) ([]model.Enrollments, error) {
	rows, err := pa.Db.Query("select Id,user_id,course_id,enrollment_date from Enrollments ;")
	if err != nil {
		return nil, err
	}
	
	var p []model.Enrollments
	for rows.Next() {
		err = rows.Scan(&enrol.Id, &enrol.UserId, &enrol.CourseId, &enrol.EnrollmentDate)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}
		p = append(p, enrol)
	}
	return p, nil
}

// bu method enrolmentsdagi malumotlarni kelgan malumotlari boyicha update qilish ucun ishlatiladi

func (u *EnrollmentsRepo) EmrolUpdate(enrolUpdateFilter model.UpdateEnrol) error {
	var params []string
	var args []interface{}
	query := `
	SELECT id
	FROM ENROLLMENTS
	WHERE delete_at IS NULL AND id = $1
	`

	if err := u.Db.QueryRow(query, *enrolUpdateFilter.EnrolId).Err(); err != nil {
		return fmt.Errorf("enrollments by this id not found: %v", err)
	}

	query = `
	UPDATE enrollments SET 
	`

	if enrolUpdateFilter.UserId != nil {
		params = append(params, fmt.Sprintf(" USER_ID = $%d ", len(args)+1))
		args = append(args, *enrolUpdateFilter.UserId)
	}

	if enrolUpdateFilter.CourseId != nil {
		params = append(params, fmt.Sprintf(" course_id = $%d ", len(args)+1))
		args = append(args, *enrolUpdateFilter.CourseId)
	}

	if enrolUpdateFilter.EnrollmentDate != nil {
		params = append(params, fmt.Sprintf("enrollment_date = $%d", len(args)+1))
		args = append(args, *enrolUpdateFilter.EnrollmentDate)
	}

	params = append(params, fmt.Sprintf("update_at = $%d", len(args)+1))
	args = append(args, time.Now())

	if len(params) == 0 {
		return fmt.Errorf("no fields to update")
	}

	args = append(args, enrolUpdateFilter.EnrolId)
	query += strings.Join(params, ", ") + fmt.Sprintf(" WHERE id = $%d AND delete_at =0", len(args))

	fmt.Println("Executing query:", query)
	fmt.Println("With arguments:", args)
	_, err := u.Db.Exec(query, args...)

	if err != nil {
		return fmt.Errorf("failed executing query: %v", err)
	}

	fmt.Println(query)
	return nil
}

// bu method enrolmentsdagi malumotni berikgan id boyicha ociradi

func (u EnrollmentsRepo) EmrolDelete(id string) error {

	_rows, err := u.Db.Exec(`update enrollments set
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

//bu method enrolments ni berilgan malumotlari boyicha filter qiuladi

func (u *EnrollmentsRepo) GetAllEnrol(f model.EnrolGetAll) ([]model.Enrollments, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)
	fmt.Println("salom")

	query := `SELECT id, user_id,course_id, enrollment_date
	FROM enrollments WHERE true `

	if len(f.UserId) > 0 {
		params["user_id"] = f.UserId
		filter += " and user_id = :user_id "
	}

	if len(f.CourseId) > 0 {
		params["course_id"] = f.CourseId
		filter += " and course_id = :birthday "
	}

	if f.EnrollmentDate != nil {
		params["enrolment_date"] = f.EnrollmentDate
		filter += " and enrolment_date = :enrolment_date "
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

	var enrols []model.Enrollments
	for rows.Next() {
		var enrol model.Enrollments
		err := rows.Scan(&enrol.Id, &enrol.UserId, &enrol.CourseId, &enrol.EnrollmentDate)

		if err != nil {
			return nil, err
		}

		enrols = append(enrols, enrol)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return enrols, nil
}
