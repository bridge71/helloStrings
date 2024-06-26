package handlers

import (
	"github.com/bridge71/helloStrings/api/services"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) CheckUser(c *gin.Context) {
	code, message := h.UserService.CreateUser(c)
	c.JSON(code, message)
}

func (h *UserHandler) AuthUser(c *gin.Context) {
	code, message := h.UserService.AuthUser(c)
	c.JSON(code, message)
}
