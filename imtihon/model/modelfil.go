package model

type UpdateLesson struct {
	LessonId *string 
	Title     *string
	CourseId *string
	Content   *string
}

type UpdateEnrol struct {
	EnrolId        *string
	UserId         *string
	CourseId       *string
	EnrollmentDate *string
}

type UpdateCourse struct {
	CourseId   *string
	Title       *string
	Description *string
}

type UpdateUser struct {
	UserId  *string
	Name     *string
	Email    *string
	Birthday *string
	Password *string
}
