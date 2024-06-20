package handler

import (
	"fmt"
	"my_mod/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// bu api dagi coursning barchasini royharini chiqarib beradi

func (h *Handler) CoursGets(c *gin.Context) {
	var courses model.Courses
	course, err := h.Courses.CoursRead(courses)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":" COURSENI MALUMOTLARINI CHIQARISHDA HATOLIK"})
	}

	for _, v := range course {
		c.JSON(http.StatusOK, v)
	}
}

//bu api yangi course databasesga qoshib beradi
func (h *Handler) CoursPost(c *gin.Context) {
	var courses model.Courses
	c.BindJSON(&courses)
	err:=h.Courses.CoursCreate(courses)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":" COURSESGA ALUMOT QOSHISHDA HATOLIK"})
	}
	c.JSON(200, courses)

}

//bu api courseni databazedagi malumotni id boyicha update qiladi

func (h *Handler) CoursPut(c *gin.Context) {
	var courses model.UpdateCourse
	c.BindJSON(&courses)
	_,err:=h.Courses.CoursUpdate(courses)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":" COURSESNI MALUMOTLARINI YANGILASHDA HATOLIK"})
	}

}

// bu api id boyicha databasesdagi courselarni ochiradi
func (h *Handler) CoursDelete(c *gin.Context) {
	id := c.Param("id")

	fmt.Println(id)
	err := h.Courses.CoursDelete(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":" COURSESNI MALUMOTINI OCIRISHDA HATOLIK"})
	}
}

//bu api da databasesdagi malumotni qaysi birlari boyicha filter qilish kk bolsa osha malunmotlarni qaytaradi

func (h *Handler) GetAllCourse(c *gin.Context) {
	var course model.CourseaGetAll
	c.BindJSON(&course)
	courses, err := h.Courses.GetAllCourse(course)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":" COURSESNI FILTER QILISHDA HATOLIK"})
	}
	c.JSON(200, courses)

}
