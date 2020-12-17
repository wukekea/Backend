package base

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Code int8 // 错误码
const (
	ReturnSuccess Code = iota + 1
	ServerErr
	ParamErr
)
const (
	MongodbUri        string = "mongodb://localhost:27017"
	DatabaseBlog      string = "blog"
	CollectionAccount string = "account"
)

type BasedResponse struct {
	Ret  Code         `json:"ret" binding:"required"` // 为1：正常返回；2：服务错误；3：参数错误
	Data *interface{} `json:"data,omitempty"`
	Msg  *string      `json:"msg,omitempty"`
}

// 将错误信息返回给客户端
func ReplyFailWithMsg(context *gin.Context, code Code, msg string) {
	resp := &BasedResponse{
		Ret:  code,
		Data: nil,
		Msg:  &msg,
	}
	context.PureJSON(http.StatusOK, resp)
}

// 将空数据返回给客户端
func ReplyOkWithoutData(context *gin.Context) {
	resp := &BasedResponse{
		Ret:  ReturnSuccess,
		Data: nil,
		Msg:  nil,
	}
	context.PureJSON(http.StatusOK, resp)
}