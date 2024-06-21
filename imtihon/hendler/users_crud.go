package handler

import (
	"fmt"
	"my_mod/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// bu api dagi usersning barchasini royharini chiqarib beradi

func (h *Handler) UserGets(c *gin.Context) {
	var users model.Users
	user, err := h.Userss.UserRead(users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " USER MALUMOTLARINI CHIQARISHFA HATOLIK"})
	}

	
	for _, v := range user {
		c.JSON(http.StatusOK, v)
	}
}

//bu api yangi users databasesga qoshib beradi

func (h *Handler) UserPost(c *gin.Context) {
	var users model.Users
	fmt.Println(c.BindJSON(&users))
	_, err := h.Userss.UserCreate(users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " USERGA MALUMOT QOSHISHDA HATOLIK"})
	}
	c.JSON(200, users)

}

//bu api usersi databazedagi malumotni id boyicha update qiladi

func (h *Handler) UserPut(c *gin.Context) {
	var users model.UpdateUser
	c.BindJSON(&users)
	err := h.Userss.UserUpdate(users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " USETRNI  MALUMOTLARINI YANGILASHDA HATOLIK"})
	}

}

// bu api id boyicha databasesdagi userslarni ochiradi

func (h *Handler) UserDelete(c *gin.Context) {
	id := c.Query("id")
	fmt.Println(id)
	err := h.Userss.UserDelete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " USERNI MALUMOTLARINI OCJHIRISHDA HATOLIK"})
	}
}

//bu api da databasesdagi malumotni qaysi birlari boyicha filter qilish kk bolsa osha malunmotlarni qaytaradi

func (h *Handler) GetAllUsers(c *gin.Context) {
	var users model.UserGetAll
	c.BindJSON(&users)
	user, err := h.Userss.GetAllUsers(users)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "  USERNI FILTERLAB CHIQARISHDA HATOLIK"})
	}
	c.JSON(200, user)

}
