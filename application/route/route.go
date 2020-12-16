package route

import (
	"Backend/application/controllers/account"
	"github.com/gin-gonic/gin"
)

func LoadRouteConfig(gin *gin.Engine) {
	account.InitWith(gin)
}