package account

import "github.com/gin-gonic/gin"

func Hello(context *gin.Context) {
	hello(context)
}

// 注册账号
func Register(context *gin.Context) {
	register(context)
}