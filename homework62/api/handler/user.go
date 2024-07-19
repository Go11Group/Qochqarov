package handler

import (
	pb "my_mod/generated/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CresteUserHandler(ctx *gin.Context) {
	var req pb.CresteUserRequest

	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "erro",
		})
		return
	}

	_, err := h.User.CresteUser(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, "ishladi")

}

func (h *Handler) GetByIdHandler(ctx *gin.Context) {
	id := ctx.Param("user_id")

	resp, err := h.User.GetByIdUser(ctx, &pb.GetByIdUserRequest{UserId: id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return

	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteUserHandler(ctx *gin.Context) {
	id := ctx.Param("user_id")

	resp, err := h.User.DeleteUser(ctx, &pb.DeleteUserRequest{UserId: id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return

	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateUserHandler(ctx *gin.Context) {
	id := ctx.Param("user_id")

	resp, err := h.User.UpdateUser(ctx, &pb.UpdateUserrequest{UserId: id})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return

	}
	ctx.JSON(http.StatusOK, resp)
}
