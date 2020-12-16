package main

import(
	"Backend/application/route"
	"github.com/gin-gonic/gin"
)

func main() {
	ginObject := gin.Default()
	route.LoadRouteConfig(ginObject)
	//engine.GET("/test", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message":"pong",
	//	})
	//})
	_ = ginObject.Run(":8080")
}