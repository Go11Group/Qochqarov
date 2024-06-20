package main

import (
	handler "my_mod/hendler"
	postgres "my_mod/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	user := postgres.NewProRepo(db)
	course := postgres.NewCourseRepo(db)
	lesson := postgres.NewLessonRepo(db)
	enrollments := postgres.NewEnrollmentsRepo(db)
	additonal := postgres.NewAdditionalRepo(db)

	hand := handler.Handler{Userss: user, Courses: course, Lesson: lesson, Emrollments: enrollments, Additional: additonal}

	handler.NewHandler(&hand)

}
