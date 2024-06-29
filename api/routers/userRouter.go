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
	router.GET("/test", userHandler.Test)

	router.POST("/sale/book/submit", saleHandler.BookCreateSale)
	router.POST("/sale/book/name", saleHandler.BookGetName)
	router.POST("/sale/book/profession", saleHandler.BookGetProfession)
	router.POST("/sale/book/course", saleHandler.BookGetCourse)
	router.POST("/sale/book/author", saleHandler.BookGetAuthor)
	router.POST("/sale/book/all", saleHandler.BookGet)
}
