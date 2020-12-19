package account

import (
	"Backend/modules/account"
	"Backend/modules/core/base"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"log"
	"net/http"
)

func login(context *gin.Context) {
	var param account.LoginParam
	err := context.ShouldBindBodyWith(&param, binding.JSON)
	if err != nil {
		log.Println("fail to get reqParam.")
		base.ReplyFailWithMsg(context, base.ParamErr, err.Error())
		return
	}
	result, err := account.Login(param)
	if err != nil {
		log.Println("fail to login.")
		base.ReplyFailWithMsg(context, base.ServerErr, err.Error())
		return
	}
	context.JSON(http.StatusOK, result)
}
