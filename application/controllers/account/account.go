package account

import "github.com/gin-gonic/gin"

// 注册账号
func Register(context *gin.Context) {
	register(context)
}

// 登录账号
func Login(context *gin.Context) {
	login(context)
}