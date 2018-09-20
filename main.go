package main

import (
	"./controllers"
	"./middleware"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.POST("/login", controllers.UserLogin)

	api := router.Group("/api", middleware.Check_token)
	{
		api.GET("/user", controllers.GetAllUser)
		api.GET("/user/:id", controllers.GetUserById)

		api.POST("/user", controllers.CreateUser)
	}
	router.Run() //..监听8080端口
}
