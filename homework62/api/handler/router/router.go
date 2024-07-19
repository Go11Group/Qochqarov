package router

import (
	"my_mod/api/handler"

	"github.com/gin-gonic/gin"
)

func RouterApi(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	user := router.Group("api/user")
	{
		user.POST("/", h.CresteUserHandler)
		user.GET("/user:id", h.GetByIdHandler)
		user.PUT("/user:id", h.UpdateUserHandler)
		user.DELETE("/user:id", h.DeleteUserHandler)
	}

	return router
}
