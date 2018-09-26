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
		//user
		api.GET("/user", controllers.GetAllUser)
		api.GET("/user/:id", controllers.GetUserById)

		api.POST("/user", controllers.CreateUser)

		//article

		api.GET("/article/:id", controllers.GetArticleById)
	}
	router.Run() //..监听8080端口
}
