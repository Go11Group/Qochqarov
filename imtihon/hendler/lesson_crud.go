package handler

import (
	"my_mod/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// bu api dagi lessonnning barchasini royharini chiqarib beradi

func (h *Handler) LessonGets(c *gin.Context) {
	var lesson model.Lessons
	user, err := h.Lesson.LessonRead(lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " LESSONNI MALUMOTLARINI CHIQARISHDA HATOLIK"})
	}

	for _, v := range user {
		c.JSON(http.StatusOK, v)
	}
}

// bu api yangi lesson databasesga qoshib beradi
func (h *Handler) LessonPost(c *gin.Context) {
	var lesson model.Lessons
	c.BindJSON(&lesson)

	c.JSON(200, lesson)
	err := h.Lesson.LessonCreate(lesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " LESSONSGA YANGI MALUMOT QOSHISHGA HATOLIK"})
	}

}

//bu api lessonni databazedagi malumotni id boyicha update qiladi

func (h *Handler) LessonPut(c *gin.Context) {

	var updateLesson model.UpdateLesson
	c.BindJSON(&updateLesson)
	err := h.Lesson.LessonUpdate(updateLesson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " LESSONSNI MALUMOTLARINI YANGILASHDA HATOLIK"})
	}

}

// bu api id boyicha databasesdagi lessonnnilarni ochiradi

func (h *Handler) LessonDelete(c *gin.Context) {
	id := c.Param("id")

	err := h.Lesson.LessonDeleted(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " LESSONSNI MALUMOTLARINI OCIRISHDA HATOLIK"})
	}

}

//bu api da databasesdagi malumotni qaysi birlari boyicha filter qilish kk bolsa osha malunmotlarni qaytaradi

func (h *Handler) GetAllLesson(c *gin.Context) {
	var lessons model.LessonGetAll
	c.BindJSON(&lessons)
	lesson, err := h.Lesson.GetAllLesson(lessons)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " LESSONSNI MALUMOTLARINI FILTERLASHDA HAGTOLIK"})
	}
	c.JSON(200, lesson)

}
