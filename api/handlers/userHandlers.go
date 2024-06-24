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

//	func (h *UserHandler) CheckUser(ctx context.Context, user *models.User) {
//		h.UserService.CheckUser(ctx, user)
//	}
func (h *UserHandler) CheckUser(c *gin.Context) {
	h.UserService.CheckUser(c)
	// c.String(200, "sss")
}
