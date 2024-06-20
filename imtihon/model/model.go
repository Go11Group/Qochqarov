package model

import (
	"time"
)

// users ning structi
type Users struct {
	Id       string
	Name     string
	Email    string
	Birthday string

	Password string
}

// coursesning structi
type Courses struct {
	Id          string
	Title       string
	Description string
}

// lessonsning structi
type Lessons struct {
	Id       string
	Title    string
	CourseId string
	Content  string
}

// enrolments ning structi
type Enrollments struct {
	Id             string
	UserId         string
	CourseId       string
	EnrollmentDate time.Time
}
