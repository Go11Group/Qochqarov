package model

import "time"

type UserGetAll struct {
	Name          string
	Birthday      *time.Time
	Limit, Offset int
}

type CourseaGetAll struct {
	Title         string
	Description   string
	Limit, Offset int
}

type LessonGetAll struct {
	Title         string
	CourseId      string
	Content       string
	Limit, Offset int
}

type EnrolGetAll struct {
	UserId         string
	CourseId       string
	EnrollmentDate *time.Time
	Limit, Offset  int
}
