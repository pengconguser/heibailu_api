package main

import (
	"./controllers"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	router.GET("/user", controllers.GetAllUser)
	router.GET("/user/:id", controllers.GetUserById)
	router.POST("/user", controllers.CreateUser)

	router.Run() //..监听8080端口
}
