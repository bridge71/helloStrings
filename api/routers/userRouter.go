package routers

import (
	"github.com/bridge71/helloStrings/api/handlers"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	UserHandler *handlers.UserHandler
}

func NewUserRouter(handler *handlers.UserHandler) *UserRouter {
	return &UserRouter{UserHandler: handler}
}

func RegisterRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	router.POST("/register", userHandler.CheckUser)
}
