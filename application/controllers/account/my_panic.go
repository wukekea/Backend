package account

import (
	"github.com/gin-gonic/gin"
)

func myPanic(context *gin.Context) {
	panic("test a panic")
}