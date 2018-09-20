package controllers

//控制器中公用函数
import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendErrorResponse(msg string, args ...interface{}) {
	if len(args) == 0 {
		panic("缺少context")
	}

	var context *gin.Context

	if len(args) == 1 {
		c, ok := args[0].(*gin.Context)
		if !ok {
			panic("缺少context")
		}

		context = c
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "请求失败" + msg,
		"data":    gin.H{},
	})

	//终止请求链

	context.Abort()
}
