package main

import (
	"Backend/application/route"
	_ "Backend/docs"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"syscall"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /
// @query.collection.format multi

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
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
	//ginObject.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	_ = ginObject.Run(":8080")
}