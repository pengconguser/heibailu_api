package controllers

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"crypto/md5"

	"../helper"
	"../models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type UsernameLogin struct {
	Name     string `json:"name" binding:"required,min=4,max=20"`
	Password string `json:"password" binding:"required,min=4,max=20"`
}

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
	//获取密钥配置
	auth_conf, err := helper.ReadJsonFile(helper.Object_path() + "/config/auth.json")

	if err != nil {
		log.Println("密钥配置文件错误!", err)
	}

	var user models.User

	var login UsernameLogin
	if err := context.ShouldBindWith(&login, binding.JSON); err != nil {
		SendErrorResponse("输入格式不符合要求", context)
		fmt.Println(err)
		return
	}

	if err := models.DB.Where("name=?", login.Name).First(&user).Error; err != nil {
		SendErrorResponse("请求失败!,用户名不存在", context)
		return
	}

	//生成一个md5对象
	md5_obj := md5.New()
	io.WriteString(md5_obj, login.Password)        //将str写入到w中
	md5str2 := fmt.Sprintf("%x", md5_obj.Sum(nil)) //w.Sum(nil)将w的hash转成[]byte格式

	fmt.Println(md5str2)

	if md5str2 != user.Password {
		SendErrorResponse("登陆失败!密码错误", context)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": user.ID,
	})

	tokenString, err := token.SignedString([]byte(auth_conf["SecretKey"]))

	if err != nil {
		SendErrorResponse("内部错误", context)
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":      "登陆成功!",
		"access_token": tokenString,
	})
}
