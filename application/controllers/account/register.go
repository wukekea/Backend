package account

import (
	"Backend/modules/account"
	"Backend/modules/core/base"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
)

func register(context *gin.Context) {
	// 获取参数
	//name := context.PostForm("name")
	var param account.RegisterParam
	err := context.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		log.Println("fail to get reqParam of register.")
		base.ReplyFailWithMsg(context, base.ParamErr, err.Error())
		return
	}
	_, err = account.Register(param)
	if err != nil {
		base.ReplyFailWithMsg(context, base.ServerErr, err.Error())
		log.Println("fail to register an account.")
		return
	}
	base.ReplyOkWithoutData(context)
}