package postgres

import (
	"database/sql"
	"fmt"
	pakage "my_mod/Pakage"
	"my_mod/model"
)

type AdditionalRepo struct {
	Db *sql.DB
}

func NewAdditionalRepo(db *sql.DB) *AdditionalRepo {
	return &AdditionalRepo{Db: db}
}

func (h *AdditionalRepo) GetCoursesbyUser(id string) ([]model.GetCoursesbyUsers, error) {
	rows, err := h.Db.Query("select u.id, c.id,c.title,c.Description from users as u join enrollments as e on u.id=e.user_id join courses as c on c.id=e.course_id where u.id=$1", id)
	if err != nil {
		panic(err)
	}

	var p []model.GetCoursesbyUsers
	var get model.GetCoursesbyUsers
	for rows.Next() {
		err := rows.Scan(&get.Id, &get.Course.Id, &get.Course.Title, &get.Course.Description)
		if err != nil {
			return nil, err
		}
		p = append(p, get)
	}
	return p, nil
}

func (h *AdditionalRepo) GetLessonsbyCourse(id string) ([]model.GetLessonsbyCourses, error) {
	rows, err := h.Db.Query("select c.id, l.id,l.title,l.content from courses as c join lessons as l on c.id=l.course_id where c.id=$1", id)
	if err != nil {
		panic(err)
	}

	var p []model.GetLessonsbyCourses
	var get model.GetLessonsbyCourses
	for rows.Next() {
		err := rows.Scan(&get.Id, &get.Lesson.Id, &get.Lesson.Title, &get.Lesson.Content)
		if err != nil {
			return nil, err
		}
		p = append(p, get)
	}
	return p, nil
}

func (h *AdditionalRepo) GetEnrolledUsersbyCourse(id string) ([]model.GetEnrolledUsersbyCourses, error) {
	rows, err := h.Db.Query("select c.id, u.id,u.name,u.email,u.birthday,u.password from courses as c join enrollments as e on c.id=e.course_id join users as u on u.id=e.user_id where c.id=$1 and e.delete_at=0", id)
	if err != nil {
		panic(err)
	}

	var p []model.GetEnrolledUsersbyCourses
	var get model.GetEnrolledUsersbyCourses
	for rows.Next() {
		err := rows.Scan(&get.Id, &get.User.Id, &get.User.Name, &get.User.Email, &get.User.Birthday, &get.User.Password)
		if err != nil {
			return nil, err
		}
		p = append(p, get)

	}
	return p, nil
}

func (a *AdditionalRepo) SearchUsers(f model.SearchUser) ([]model.Result, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		filter string
	)

	query := `SELECT id, name,email, birthday,password
	FROM users WHERE delete_at=0 `

	if len(f.Name) > 0 {
		params["name"] = f.Name
		filter += " and name = :name "
	}

	if len(f.Email) > 0 && f.OneAge > 0 {
		params["email"] = f.Email
		filter += " and email = :email "
	}
	fmt.Println(query)

	if f.OneAge > 0 && f.TwoAge > 0 {
		params["age1"] = f.OneAge
		params["age2"] = f.TwoAge
		filter += " and EXTRACT(YEAR FROM age(birthday)) BETWEEN :age1 AND :age2; "
	}

	query = query + filter

	query, arr = pakage.ReplaceQueryParams(query, params)
	fmt.Println(query, arr)
	rows, err := a.Db.Query(query, arr...)
	if err != nil {
		return nil, err
	}

	var users []model.Users
	for rows.Next() {
		var user model.Users
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Birthday, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	result := model.Result{
		Results: users,
	}
	var results []model.Result
	results = append(results, result)
	return results, nil
}

func (a *AdditionalRepo) GetMostPopularCourses(f model.MostPopularCourses) (*model.MostPopularCoursesstruct, error) {
	rows, err := a.Db.Query(`
		SELECT 
			c.id,
			c.title,
			COUNT(c.id) AS enrollment_count
		FROM 
			courses AS c 
		INNER JOIN
			enrollments AS e ON c.id = e.course_id
		WHERE 
			e.enrollment_date BETWEEN $1 and $2
		GROUP BY
			c.id, c.title
		HAVING
			COUNT(c.id) = (
				SELECT 
					MAX(enroll_count) 
				FROM (
					SELECT 
						COUNT(e.course_id) AS enroll_count
					FROM 
						courses AS c 
					INNER JOIN
						enrollments AS e ON c.id = e.course_id
					GROUP BY
						c.id
				) AS counts
			)
`, f.StartDate, f.EndDate)

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
	var course []model.PopularCourses
	var courses model.PopularCourses
	for rows.Next() {
		err := rows.Scan(&courses.CourseId, &courses.CourseTitle, &courses.EnrolmentsCount)
		if err != nil {
			return nil, err
		}

		course = append(course, courses)
	}
	data := model.MostPopularCoursesstruct{
		MostPopularCourse: model.MostPopularCourses{
			EndDate:   f.EndDate,
			StartDate: f.StartDate,
		},
		CopularCourses: course,
	}
	return &data, nil

}
