package handler

import (
	"my_mod/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CoursGets(c *gin.Context) {
	var courses model.Courses
	course, err := h.Courses.CoursRead(courses)
	if err != nil {
		panic(err)
	}

	for _, v := range course {
		c.JSON(http.StatusOK, v)
	}
}

func (h *Handler) CoursPost(c *gin.Context) {
	var courses model.Courses
	c.BindJSON(&courses)
	c.JSON(200, courses)
	h.Courses.CoursCreate(courses)

}

func (h *Handler) CoursPut(c *gin.Context) {
	var courses model.UpdateCourse
	c.BindJSON(&courses)
	c.JSON(200, courses)
	h.Courses.CoursUpdate(courses)

}

func (h *Handler) CoursDelete(c *gin.Context) {
	id:=c.Param("id")
	
	err:=h.Courses.CoursDelete(id)
	if err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllCourse(c *gin.Context) {
	var course model.CourseaGetAll
	c.BindJSON(&course)
	courses, err := h.Courses.GetAllCourse(course)
	if err != nil {
		panic(err)
	}
	c.JSON(200, courses)

}
