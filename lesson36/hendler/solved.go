package handler

import (
	"my_mod/model"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) S_Gets(c *gin.Context) {
	var solved model.Solved_problems
	solveds, err := h.Solvedproblems.S_Read(solved)
	if err != nil {
		panic(err)
	}

	for _, v := range solveds {
		c.JSON(http.StatusOK, v)
	}
}

func (h *Handler) S_Post(c *gin.Context) {
	var solved model.Solved_problems
	c.BindJSON(&solved)
	c.JSON(200, solved)
	h.Solvedproblems.S_Create(solved)

}

func (h *Handler) S_Put(c *gin.Context) {
	var solved model.Solved_problems
	c.BindJSON(&solved)
	c.JSON(200, solved)
	h.Userss.U_Update(solved.Id, solved.Name)

}

func (h *Handler) S_Delete(c *gin.Context) {
	var solved model.Solved_problems
	c.BindJSON(&solved)
	c.JSON(200, solved)
	h.Userss.U_Delete(solved.Id)

}
