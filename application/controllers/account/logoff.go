package account

import (
	"Backend/modules/account"
	"Backend/modules/core/base"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
)

func logoff(context *gin.Context) {
	var param account.LogoffParam
	err := context.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		log.Println("fail to get reqParam.")
		base.ReplyFailWithMsg(context, base.ParamErr, err.Error())
		return
	}

	err = account.Logoff(&param)
	if err != nil {
		log.Println("fail to logoff an account.")
		base.ReplyFailWithMsg(context, base.ServerErr, err.Error())
		return
	}
	base.ReplyOkWithoutData(context)
}