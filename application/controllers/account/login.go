package account

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HelloResult struct {
	Name  string `json:"name" binding:"required"`
	Age   int8   `json:"age" binding:"required"`
	Greet string `json:"greet" binding:"required"`
}

func hello(context *gin.Context) {
	result := &HelloResult{
		Name:  "wukeke",
		Age:   22,
		Greet: "hello, i am wkk.",
	}
	context.JSON(http.StatusOK, result)
}
