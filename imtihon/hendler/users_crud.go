package handler

import (
	"fmt"
	"my_mod/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) UserGets(c *gin.Context) {
	var users model.Users
	user, err := h.Userss.UserRead(users)
	if err != nil {
		panic(err)
	}

	for _, v := range user {
		c.JSON(http.StatusOK, v)
	}
}

func (h *Handler) UserPost(c *gin.Context) {
	var users model.Users
	c.BindJSON(&users)
	c.JSON(200, users)
	h.Userss.UserCreate(users)

}

func (h *Handler) UserPut(c *gin.Context) {
	var users model.UpdateUser
	c.BindJSON(&users)
	c.JSON(200, users)
	h.Userss.UserUpdate(users)

}

func (h *Handler) UserDelete(c *gin.Context){
	id:=c.Param("id")
	
	err:=h.Userss.UserDelete(id)
	if err != nil {
		panic(err)
	}
	

}

func (h *Handler) GetAllUsers(c *gin.Context) {
	fmt.Println("salom")
	var users model.UserGetAll
	c.BindJSON(&users)
	user, err := h.Userss.GetAllUsers(users)
	if err != nil {
		panic(err)
	}
	c.JSON(200, user)

}
