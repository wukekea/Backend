package main

import (
	"Backend/application/route"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"syscall"
)

func main() {
	// 默认使用Logger and Recovery中间件
	//ginObject := gin.Default()


	// 不使用中间件
	ginObject := gin.New()

	gin.ForceConsoleColor()
	// 修改日志的默认位置
	file, err := os.OpenFile("gin.log", syscall.O_RDWR, 0666)
	if err != nil {
		log.Println("fail to open gin.log.")
		return
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	ginObject.Use(gin.Logger())
	route.LoadRouteConfig(ginObject)
	_ = ginObject.Run(":8080")
}