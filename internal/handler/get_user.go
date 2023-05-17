package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hoangpqse01968/hexagonal_architecture_sample/internal/application/service"
)

type GetUserHandler struct {
	userService service.UserService
}

func NewGetUserHandler(
	userService service.UserService,
) *GetUserHandler {
	return &GetUserHandler{
		userService: userService,
	}
}

func (h *GetUserHandler) HandleGetUser(c *gin.Context) {
	req := struct {
		Id int64 `json:"id"`
	}{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.GetUser(req.Id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
