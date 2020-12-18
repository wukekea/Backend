package account

import (
	"Backend/application/middleware"
	"github.com/gin-gonic/gin"
)



func InitWith(ginObject *gin.Engine) {
	group := ginObject.Group("/account/")

	group.Use(middleware.TheMiddlewarefirst)
	group.Use(middleware.TheMiddlewareSecond)

	group.POST("/login", Login)
	group.PUT("/register", Register)
	group.GET("/panic", MyPanic)



}