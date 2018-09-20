package controllers

import (
	"fmt"
	"net/http"

	"../models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetAllUser(context *gin.Context) {
	users := models.GetAllUser()

	context.JSON(http.StatusOK, gin.H{
		"data":    users,
		"message": "请求成功!",
	})
}

func CreateUser(context *gin.Context) {

}

func GetUserById(context *gin.Context) {
	user := models.GetIdUser(context.Param("id"))

	context.JSON(http.StatusOK,
		gin.H{
			"data":    user,
			"message": "请求成功!",
		},
	)
}

func UserLogin(context *gin.Context) {
	var user models.User

	if err := models.DB.Where("name=?", context.Param("name")).Error; err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "请求失败!,用户名密码错误",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	fmt.Println(token)
}
