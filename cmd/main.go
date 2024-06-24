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
	//
	// router.POST("/register", handlers.Register)
	// router.GET("/", func(c *gin.Context) {
	// 	// Example of using the database
	// 	rows, err := db.Query("SELECT email from Users;")
	// 	if err != nil {
	// 		c.String(500, err.Error())
	// 		return
	// 	}
	// 	var message string
	// 	for rows.Next() {
	// 		if err := rows.Scan(&message); err != nil {
	// 			c.String(500, err.Error())
	// 			return
	// 		}
	// 	}
	// 	c.String(200, message)
	// })
	//
	router.Run(":8080")
}
