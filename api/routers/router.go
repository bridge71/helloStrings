package routers

import (
	"github.com/bridge71/helloStrings/api/handlers"
	"github.com/bridge71/helloStrings/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userHandler *handlers.UserHandler, saleHandler *handlers.SaleHandler, postHandler *handlers.PostHandler) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.POST("/login", userHandler.Login)

	authorized := router.Group("/")
	authorized.Use(middleware.Authentication())
	{

		userR := authorized.Group("/user")
		{
			userR.POST("/register", userHandler.UserCreate)
			userR.POST("/id", userHandler.UserReadId)
			userR.POST("/ip/store", userHandler.StoreIP)
		}
		router.GET("/test", userHandler.Test)

		saleR := authorized.Group("/sale")
		{
			saleR.POST("/book/create", saleHandler.BookCreate)
			saleR.POST("/book/by", saleHandler.BookReadBy)
			saleR.POST("/book/fetch", saleHandler.BookFetch)
		}

		postR := authorized.Group("/post")
		{
			postR.POST("/create", postHandler.PostCreate)
			postR.GET("/fetch", postHandler.PostFetch)
			postR.POST("/read/title", postHandler.PostReadTitle)
			postR.POST("/read/id", postHandler.ContentReadPostId)
			commentR := postR.Group("/comment")
			{
				commentR.POST("/create", postHandler.CommentCreate)
				commentR.POST("/read/id", postHandler.CommentReadPostId)
				commentR.POST("/read/userId", postHandler.CommentReadUserId)
			}
			markR := postR.Group("/mark")
			{
				markR.POST("/create", postHandler.CommentMarkCreate)
				markR.POST("/read", postHandler.CommentMarkReadUserId)
			}
			likesR := postR.Group("/likes")
			{
				likesR.POST("/change", postHandler.LikesChange)
				likesR.POST("/read", postHandler.LikesReadUserId)
			}
		}
	}
}
