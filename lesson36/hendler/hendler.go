package handler

import (
	postgres "my_mod/storage"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Userss         *postgres.UserRepo
	Problems       *postgres.ProblemRepo
	Solvedproblems *postgres.SolvedRepo
}

func NewHandler(handler *Handler) *gin.Engine {
	//user uchun
	router := gin.Default()
	router.GET("/user",handler.Gets)
	router.POST("/user",handler.Post)
	router.PUT("/user",handler.Put)
	router.DELETE("/user",handler.Delete)
	
	//problenm uchun
	router.GET("/problem",handler.P_Gets)
	router.POST("/peoblem",handler.P_Post)
	router.PUT("/problem",handler.P_Put)
	router.DELETE("/problem",handler.P_Delete)
	
	// SOLVED UCHUNrouter.GET("/user",handler.Gets)
	router.POST("/user",handler.Post)
	router.PUT("/user",handler.Put)
	router.DELETE("/user",handler.Delete)
	

	router.GET("/solved",handler.S_Gets)
	router.POST("/solved",handler.S_Post)
	router.PUT("/solved",handler.S_Put)
	router.DELETE("/solved",handler.S_Delete)
	
	router.Run(":8090")
	return router
}
