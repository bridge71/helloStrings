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

func (h *UserHandler) UserCreate(c *gin.Context) {
	code, message := h.UserService.UserCreate(c)
	c.JSON(code, message)
}

func (h *UserHandler) IPCreate(c *gin.Context) {
	code, message := h.UserService.IPCreate(c)
	c.JSON(code, message)
}

func (h *UserHandler) IPRead(c *gin.Context) {
	code, message := h.UserService.IPRead(c)
	c.JSON(code, message)
}

func (h *UserHandler) UserReadId(c *gin.Context) {
	code, message := h.UserService.UserReadId(c)
	c.JSON(code, message)
}

func (h *UserHandler) Login(c *gin.Context) {
	code, message := h.UserService.Login(c)
	c.JSON(code, message)
}

func (h *UserHandler) Test(c *gin.Context) {
	code, message := h.UserService.Test(c)
	c.JSON(code, message)
}
