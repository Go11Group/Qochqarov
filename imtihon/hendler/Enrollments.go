package handler

import (
	"my_mod/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) EmrolGets(c *gin.Context) {
	var enrol model.Enrollments
	enrollment, err := h.Emrollments.EmrolRead(enrol)
	if err != nil {
		panic(err)
	}

	for _, v := range enrollment {
		c.JSON(http.StatusOK, v)
	}
}

func (h *Handler) EmrolPost(c *gin.Context) {
	var enrol model.Enrollments
	c.BindJSON(&enrol)
	c.JSON(200, enrol)
	h.Emrollments.EmrolCreate(enrol)

}

func (h *Handler) EmrolPut(c *gin.Context) {

	var updateEnrol model.UpdateEnrol
	c.BindJSON(&updateEnrol)
	c.JSON(200, updateEnrol)
	h.Emrollments.EmrolUpdate(updateEnrol)

}

func (h *Handler) EmrolDelete(c *gin.Context) {
	id:=c.Param("id")
	
	err:=h.Emrollments.EmrolDelete(id)
	if err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllEnrolments(c *gin.Context) {
	var enrol model.EnrolGetAll
	c.BindJSON(&enrol)
	enrols, err := h.Emrollments.GetAllEnrol(enrol)
	if err != nil {
		panic(err)
	}
	c.JSON(200, enrols)

}
