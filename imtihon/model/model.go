package model

import "time"

type Users struct {
	Id         string
	Name       string
	Email      string
	Birthday   time.Time
	Password   string
	CreatedAt time.Time
	UpdateAt  time.Time
	DeleteAt  time.Time
}

type Courses struct {
	Id          string
	Title       string
	Description string
	CreatedAt  time.Time
	UpdateAt   time.Time
	DeleteAt   time.Time
}

type Lessons struct {
	Id         string
	Title      string
	CourseId  string
	Content    string
	CreatedAt time.Time
	UpdateAt  time.Time
	Deletet  time.Time
}

type Enrollments struct {
	Id              string
	UserId         string
	CourseId       string
	EnrollmentDate time.Time
	CreatedAt      time.Time
	UpdateAt       time.Time
	DeleteAt       time.Time
}
