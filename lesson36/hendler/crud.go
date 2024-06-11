package handler

import (
	"my_mod/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) Gets(c *gin.Context) {
	var users model.User
	user, err := h.Userss.U_Read(users)
	if err != nil {
		panic(err)
	}

	for _, v := range user {
		c.JSON(http.StatusOK, v)
	}
}

func (h *Handler) Post(c *gin.Context) {
	var users model.User
	c.BindJSON(&users)
	c.JSON(200, users)
	h.Userss.U_Create(users)

}

func (h *Handler) Put(c *gin.Context) {
	var users model.User
	c.BindJSON(&users)
	c.JSON(200, users)
	h.Userss.U_Update(users.Age,users.First_name)

}

func (h *Handler) Delete(c *gin.Context)  {
	var users model.User
	c.BindJSON(&users)
	c.JSON(200, users)
	h.Userss.U_Delete(users.Id)

}