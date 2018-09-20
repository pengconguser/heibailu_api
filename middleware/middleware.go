package middleware

import (
	"github.com/gin-gonic/gin"
)

func Check_token(context *gin.Context) {
	//验证头部token是否存在
	Header := context.Request.Header

	if Header.Get("Authorization") == "" {

	}

	context.Next()
}
