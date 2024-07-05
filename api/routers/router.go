package routers

import (
	"github.com/bridge71/helloStrings/api/handlers"
	"github.com/bridge71/helloStrings/api/middleware"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, userHandler *handlers.UserHandler, saleHandler *handlers.SaleHandler, postHandler *handlers.PostHandler) {
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.POST("/_login", userHandler.Login)

	router.POST("/user/register", userHandler.UserCreate)
	authorized := router.Group("/")
	authorized.Use(middleware.Authentication())
	{

		userR := authorized.Group("/user")
		{
			userR.POST("/id", userHandler.UserReadId)
			userR.POST("/ip/create", userHandler.IPCreate)
			userR.POST("/ip/read", userHandler.IPRead)
		}
		router.GET("/test", userHandler.Test)

		saleR := authorized.Group("/sale")
		{
			saleR.POST("/book/create", saleHandler.BookCreate)
			saleR.POST("/book/by", saleHandler.BookReadBy)
			saleR.POST("/book/fetch", saleHandler.BookFetch)
			saleR.POST("/book/update", saleHandler.BookUpdateStatus)
		}

		postR := authorized.Group("/post")
		{
			postR.POST("/create", postHandler.PostCreate)
			postR.GET("/fetch", postHandler.PostFetch)
			postR.POST("/read/title", postHandler.PostReadTitle)
			postR.POST("/read/nickname", postHandler.PostReadNickname)
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
