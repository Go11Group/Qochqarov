package handler

import (
	postgres "my_mod/storage/postgres"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Userss      *postgres.UserRepo
	Courses     *postgres.CousresRepo
	Lesson      *postgres.LessonRepo
	Emrollments *postgres.EnrollmentsRepo
	Additional  *postgres.AdditionalRepo
}

func NewHandler(handler *Handler) *gin.Engine {
	//user uchun API
	router := gin.Default()
	router.GET("/user", handler.UserGets)
	router.POST("/user", handler.UserPost)
	router.PUT("/user", handler.UserPut)
	router.DELETE("/user", handler.UserDelete)
	router.GET("/userall", handler.GetAllUsers)

	// course bn ishlash API
	router.GET("/course", handler.CoursGets)
	router.POST("/course", handler.CoursPost)
	router.PUT("/course", handler.CoursPut)
	router.DELETE("/course/:id", handler.CoursDelete)
	router.GET("/courseall", handler.GetAllCourse)

	//lesson ucun API
	router.GET("/lesson", handler.LessonGets)
	router.POST("/lesson", handler.LessonPost)
	router.PUT("/lesson", handler.LessonPut)
	router.DELETE("/lesson/:id", handler.LessonDelete)
	router.GET("/lessonall", handler.GetAllLesson)

	//enrollments uchun API
	router.GET("/enrol", handler.EmrolGets)
	router.POST("/enrol", handler.EmrolPost)
	router.PUT("/enrol", handler.EmrolPut)
	router.DELETE("/enrol/:id", handler.EmrolDelete)
	router.GET("/enrolall", handler.GetAllEnrolments)

	//qoshimcha apilar

	router.GET("/CoursesbyUsers/:id", handler.GetCoursesbyUsersApi)
	router.GET("/lessonsbyCourse/:id", handler.GetLessonsbyCourseApi)
	router.GET("/enrolledUsersbyCourse/:id", handler.GetEnrolledUsersbyCourseApi)
	router.GET("/searchUser/", handler.SearchUsersApi)
	router.GET("/getMost/", handler.GetMostPopularCourses)

	router.Run(":8090")
	return router
}
