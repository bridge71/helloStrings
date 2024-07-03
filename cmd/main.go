package main

import (
	"github.com/bridge71/helloStrings/api/configs"
	"github.com/bridge71/helloStrings/api/handlers"
	"github.com/bridge71/helloStrings/api/repositories"
	"github.com/bridge71/helloStrings/api/routers"
	"github.com/bridge71/helloStrings/api/services"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	configs.LoadConfigs()

	userService := services.NewUserService(&repositories.UserRepository{})
	userHandler := handlers.NewUserHandler(userService)

	saleService := services.NewSaleService(&repositories.SaleRepository{})
	saleHandler := handlers.NewSaleHandler(saleService)

	postService := services.NewPostService(&repositories.PostRepository{})
	postHandler := handlers.NewPostHandler(postService)

	routers.RegisterRoutes(router, userHandler, saleHandler, postHandler)

	router.Static("/static", "../../contents")

	// Example route to handle image requests
	router.GET("/contents/:filename", func(c *gin.Context) {
		filename := c.Param("filename")
		c.File("../../contents/" + filename)
	})

	router.Run("127.0.0.1:7777")
}
