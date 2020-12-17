package account

import "github.com/gin-gonic/gin"

func InitWith(ginObject *gin.Engine) {
	group := ginObject.Group("/account/")
	group.POST("/login", Login)
	group.PUT("/register", Register)
}