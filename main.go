package main

import (
	"Backend/application/route"
	"github.com/gin-gonic/gin"
)

func main() {
	// 默认使用Logger and Recovery中间件
	//ginObject := gin.Default()


	// 不使用中间件
	ginObject := gin.New()
	ginObject.Use(gin.Logger())
	route.LoadRouteConfig(ginObject)
	_ = ginObject.Run(":8080")
}