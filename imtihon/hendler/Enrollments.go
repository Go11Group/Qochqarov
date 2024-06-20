package handler

import (
	"my_mod/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

// bu api dagi enrollmentsning barchasini royharini chiqarib beradi

func (h *Handler) EmrolGets(c *gin.Context) {
	var enrol model.Enrollments
	enrollment, err := h.Emrollments.EmrolRead(enrol)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " ENROLMENTSNI MALUMOTLARINI CHIQARISHFDA HATOLIK"})
	}

	for _, v := range enrollment {
		c.JSON(http.StatusOK, v)
	}
}

//bu api yangi enrolments databasesga qoshib beradi

func (h *Handler) EmrolPost(c *gin.Context) {
	var enrol model.Enrollments
	c.BindJSON(&enrol)
	c.JSON(200, enrol)
	err := h.Emrollments.EmrolCreate(enrol)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " ENROLMENTSGA MALUMOT QOSHISHDA HATOLIK"})
	}

}

//bu api enrolmentsni databazedagi malumotni id boyicha update qiladi

func (h *Handler) EmrolPut(c *gin.Context) {

	var updateEnrol model.UpdateEnrol
	c.BindJSON(&updateEnrol)
	err := h.Emrollments.EmrolUpdate(updateEnrol)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " ENROLENTSNI MALUMOTLATINI YANGILASHDA HATOLIK"})
	}

}

// bu api id boyicha databasesdagi enrolmentslarni ochiradi

func (h *Handler) EmrolDelete(c *gin.Context) {
	id := c.Param("id")

	err := h.Emrollments.EmrolDelete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " ENROLMENTSNI ALUMOTLARINI OCIRISHDA HATOLIK"})
	}
}

//bu api da databasesdagi malumotni qaysi birlari boyicha filter qilish kk bolsa osha malunmotlarni qaytaradi

func (h *Handler) GetAllEnrolments(c *gin.Context) {
	var enrol model.EnrolGetAll
	c.BindJSON(&enrol)
	enrols, err := h.Emrollments.GetAllEnrol(enrol)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": " ENRILMENTSNI MALUMOTLARINI FIRLTERLASHDA HATOLIK "})
	}
	c.JSON(200, enrols)

}
