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
	routers.RegisterRoutes(router, userHandler)
	router.Run("127.0.0.1:7777")
}
