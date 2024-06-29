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
	routers.RegisterRoutes(router, userHandler, saleHandler)
	router.Run("127.0.0.1:7777")
}
