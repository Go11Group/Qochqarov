package model

//lessonni filter qilsih ucun ishlatilgan struct
type UpdateLesson struct {
	LessonId *string 
	Title     *string
	CourseId *string
	Content   *string
}

//enrolmentsni filter qilsih ucun ishlatilgan struct
type UpdateEnrol struct {
	EnrolId        *string
	UserId         *string
	CourseId       *string
	EnrollmentDate *string
}

//coursesi filter qilsih ucun ishlatilgan struct

type UpdateCourse struct {
	CourseId   *string
	Title       *string
	Description *string
}


//lessonni filter qilsih ucun ishlatilgan struct

type UpdateUser struct {
	UserId  *string
	Name     *string
	Email    *string
	Birthday *string
	Password *string
}
