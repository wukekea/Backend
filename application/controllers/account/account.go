package account

import "github.com/gin-gonic/gin"

// Register godoc
// @Summary Register An Account.
// @Description Using name and password to register an account. Sex, hobbies, age and so on is optional.
// @Accept json
// @Produce json
// @Param RegisterParam body account.RegisterParam true "required param."
// @Success 200 {object} base.BasedResponse
// @Header 200 {string} Token "qwerty"
// @Router /account/register [PUT]
func Register(context *gin.Context) {
	register(context)
}

// 登录账号
func Login(context *gin.Context) {
	login(context)
}

// 测试recover中间件
func MyPanic(context *gin.Context) {
	myPanic(context)
}