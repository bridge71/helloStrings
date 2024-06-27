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
	router.POST("/user/register", userHandler.CheckUser)
	router.POST("/user/login", userHandler.AuthUser)
	router.GET("/test", userHandler.Test)
}
