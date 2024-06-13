package handler

import (
	"fmt"
	"my_mod/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LessonGets(c *gin.Context) {
	var lesson model.Lessons
	user, err := h.Lesson.LessonRead(lesson)
	if err != nil {
		panic(err)
	}

	for _, v := range user {
		c.JSON(http.StatusOK, v)
	}
}

func (h *Handler) LessonPost(c *gin.Context) {
	var lesson model.Lessons
	c.BindJSON(&lesson)
	c.JSON(200, lesson)
	h.Lesson.LessonCreate(lesson)

}

func (h *Handler) LessonPut(c *gin.Context) {

	var updateLesson model.UpdateLesson
	c.BindJSON(&updateLesson)
	c.JSON(200, updateLesson)
	h.Lesson.LessonUpdate(updateLesson)
	fmt.Println(*updateLesson.Content)
}

func (h *Handler) LessonDelete(c *gin.Context) {
	id:=c.Param("id")
	
	err:=h.Lesson.LessonDeleted(id)
	if err != nil {
		panic(err)
	}
	
}

func (h *Handler) GetAllLesson(c *gin.Context) {
	var lessons model.LessonGetAll
	c.BindJSON(&lessons)
	lesson, err := h.Lesson.GetAllLesson(lessons)
	if err != nil {
		panic(err)
	}
	c.JSON(200, lesson)

}
