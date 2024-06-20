package handler

import (
	"fmt"
	"my_mod/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCoursesbyUsersApi(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	user, err := h.Additional.GetCoursesbyUser(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"errors"})
	}
	c.JSON(200, user)

}

func (h *Handler) GetLessonsbyCourseApi(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	user, err := h.Additional.GetLessonsbyCourse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"errors"})
	}
	c.JSON(200, user)

}

func (h *Handler) GetEnrolledUsersbyCourseApi(c *gin.Context) {
	id := c.Param("id")
	user, err := h.Additional.GetEnrolledUsersbyCourse(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"errors"})
	}
	c.JSON(200, user)

}

func (h *Handler) SearchUsersApi(c *gin.Context) {
	name := c.Query("name")
	email := c.Query("email")
	oneAge1 := c.Query("OneAge")
	twoAge1 := c.Query("TwoAge")
	var oneAge2 int
	var twoAge2 int
	if oneAge1 != "" {
		k, err := strconv.Atoi(oneAge1)
		oneAge2 = k
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"errors"})
		}
	}

	if twoAge1 != "" {
		n, err := strconv.Atoi(twoAge1)
		twoAge2 = n
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error":"errors"})
		}
	}
	user := model.SearchUser{
		Name:   name,
		Email:  email,
		OneAge: oneAge2,
		TwoAge: twoAge2,
	}
	fmt.Println(user)
	users, err := h.Additional.SearchUsers(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"errors"})
	}

	c.JSON(200, users)

}

func (h *Handler) GetMostPopularCourses(c *gin.Context) {
	starDate := c.Query("StartDate")
	endDate := c.Query("EndDate")
	most := model.MostPopularCourses{
		StartDate: starDate,
		EndDate:   endDate,
	}

	courses, err := h.Additional.GetMostPopularCourses(most)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"errors"})
	}
	c.JSON(200, courses)

}
