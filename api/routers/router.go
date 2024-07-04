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
			userR.POST("/register", userHandler.CheckUser)
			userR.POST("/info", userHandler.GetInfoUser)
			userR.POST("/ip/store", userHandler.StoreIP)
		}
		router.GET("/test", userHandler.Test)

		saleR := authorized.Group("/sale")
		{
			saleR.POST("/book/submit", saleHandler.BookCreateSale)
			saleR.POST("/book/by", saleHandler.BookGetBy)
			saleR.POST("/book/all", saleHandler.BookGet)
		}

		postR := authorized.Group("/post")
		{
			postR.POST("/submit", postHandler.CreatePost)
			postR.GET("/fetch", postHandler.GetPostAll)
			postR.POST("/title", postHandler.GetPostTitle)
			postR.POST("/content", postHandler.GetPostContent)
			commentR := postR.Group("/comment")
			{
				commentR.POST("/submit", postHandler.CreateComment)
				commentR.POST("/fetch/postId", postHandler.CommentGetPostId)
				commentR.POST("/fetch/userId", postHandler.CommentGetUserId)
			}
			markR := postR.Group("/mark")
			{
				markR.POST("/add", postHandler.PostCommentsAdd)
				markR.POST("/read", postHandler.PostCommentsRead)
			}
			likesR := postR.Group("/likes")
			{
				likesR.POST("/change", postHandler.PostLikesChange)
				likesR.POST("/read", postHandler.PostLikesRead)
			}
		}
	}
}
