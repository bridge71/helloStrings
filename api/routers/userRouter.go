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

func RegisterRoutes(router *gin.Engine, userHandler *handlers.UserHandler, saleHandler *handlers.SaleHandler) {
	router.POST("/user/register", userHandler.CheckUser)
	router.POST("/user/login", userHandler.AuthUser)
	router.POST("/user/info", userHandler.GetInfoUser)
	router.GET("/test", userHandler.Test)

	router.POST("/sale/book/submit", saleHandler.BookCreateSale)
	router.POST("/sale/book/by", saleHandler.BookGetBy)
	router.POST("/sale/book/all", saleHandler.BookGet)
}
