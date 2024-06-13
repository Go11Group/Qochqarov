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
}

func NewHandler(handler *Handler) *gin.Engine {
	//user uchun
	router := gin.Default()
	router.GET("/user", handler.UserGets)
	router.POST("/user", handler.UserPost)
	router.PUT("/user", handler.UserPut)
	router.DELETE("/user", handler.UserDelete)
	router.GET("/userall",handler.GetAllUsers)

	// course bn ishlash
	router.GET("/course", handler.CoursGets)
	router.POST("/course", handler.CoursPost)
	router.PUT("/course", handler.CoursPut)
	router.DELETE("/course", handler.CoursDelete)
	router.GET("/courseall",handler.GetAllCourse)

	//lesson ucun
	router.GET("/lesson", handler.LessonGets)
	router.POST("/lesson", handler.LessonPost)
	router.PUT("/lesson", handler.LessonPut)
	router.DELETE("/lesson", handler.LessonDelete)
	router.GET("/lessonall",handler.GetAllLesson)

	//enrol uchun
	router.GET("/enrol", handler.EmrolGets)
	router.POST("/enrol", handler.EmrolPost)
	router.PUT("/enrol", handler.EmrolPut)
	router.DELETE("/enrol/:id", handler.EmrolDelete)
	router.GET("/enrolall", handler.GetAllEnrolments)


	router.Run(":8090")
	return router
}
