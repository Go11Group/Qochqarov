package model

import "time"

// users ning structi filter uchun

type UserGetAll struct {
	Name          string
	Birthday      *time.Time
	Limit, Offset int
}

// coursesning structi filter uchun

type CourseaGetAll struct {
	Title         string
	Description   string
	Limit, Offset int
}

// lessonsning structi filter uchun

type LessonGetAll struct {
	Title         string
	CourseId      string
	Content       string
	Limit, Offset int
}

// enrolmentsning structi filter uchun

type EnrolGetAll struct {
	UserId         string
	CourseId       string
	EnrollmentDate *time.Time
	Limit, Offset  int
}
