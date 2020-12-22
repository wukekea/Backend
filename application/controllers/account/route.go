package account

import (
	"Backend/application/middleware"
	"github.com/gin-gonic/gin"
)

func InitWith(ginObject *gin.Engine) {
	group := ginObject.Group("/account/")

	group.Use(middleware.TheMiddlewarefirst)
	group.Use(middleware.TheMiddlewareSecond)

	group.POST("/login", Login)                   // 登录账号
	group.PUT("/register", Register)              // 注册账号（名字唯一）
	group.GET("/panic", MyPanic)                  // 测试panic，recovery函数
	group.DELETE("/logoff", Logoff)               // 注销账号
	group.POST("/modify_info", ModifyAccountInfo) // 通过用户id修改账号信息
}
