package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func TheMiddlewarefirst(context *gin.Context) {
	fmt.Println("middleware first.")
	// 如果将context.Next()写在这里的话会没有效果
	// context.Next()
	defer func() {
		if v := recover(); v != nil {
			context.PureJSON(http.StatusOK, bson.M{"err":v})
		}
	}()
	// context.Next()函数表示先执行后面的中间件，最后回来执行当前的中间件（该中间件执行的顺序甚至在接口调用的函数之外）
	// 利用Next()函数的这个特性，一般Recovery()的中间件会使用到该函数
	context.Next()
}

func TheMiddlewareSecond(context *gin.Context) {
	fmt.Println("middleware second.")
}