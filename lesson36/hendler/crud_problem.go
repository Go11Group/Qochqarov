package handler

import (
	"my_mod/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) P_Gets(c *gin.Context) {
	var problem model.Problem
	problems, err := h.Problems.P_Read(problem)
	if err != nil {
		panic(err)
	}

	for _, v := range problems {
		c.JSON(http.StatusOK, v)
	}
}

func (h *Handler) P_Post(c *gin.Context) {
	var problem model.Problem
	c.BindJSON(&problem)
	c.JSON(200, problem)
	h.Problems.P_Create(problem)

}

func (h *Handler) P_Put(c *gin.Context) {
	var problem model.Problem
	c.BindJSON(&problem)
	c.JSON(200, problem)
	h.Userss.U_Update(problem.Id, problem.Name)

}

func (h *Handler) P_Delete(c *gin.Context) {
	var problem model.Solved_problems
	c.BindJSON(&problem)
	c.JSON(200, problem)
	h.Userss.U_Delete(problem.Id)

}
