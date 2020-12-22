package account

import "github.com/gin-gonic/gin"

// Register godoc
// @Summary Register An Account.
// @Description Using name and password to register an account. Sex, hobbies, age and so on is optional.
// @Tags account
// @Accept json
// @Produce json
// @Param RegisterParam body account.RegisterParam true "required param."
// @Success 200 {object} base.BasedResponse
// @Header 200 {string} Token "qwerty"
// @Router /account/register [PUT]
func Register(context *gin.Context) {
	register(context)
}

// Login godoc
// @Summary Login An Account.
// @Description Using name and password to Login an account.
// @Tags account
// @Accept json
// @Produce json
// @Param LoginParam body account.LoginParam true "required param."
// @Success 200 {object} api.AccountInfo
// @Header 200 {string} Token "qwerty"
// @Router /account/login [POST]
func Login(context *gin.Context) {
	login(context)
}

// MyPanic godoc
// @Summary MyPanic.
// @Description Test panic and recovery function.
// @Tags account
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Header 200 {string} Token "qwerty"
// @Router /account/panic [GET]
func MyPanic(context *gin.Context) {
	myPanic(context)
}

// Logoff godoc
// @Summary Logoff
// @Description Logoff an account by name and password.
// @Tags account
// @Accept json
// @Produce json
// @Param LogoffParam body account.LogoffParam true "required param."
// @Success 200 {object} base.BasedResponse
// @Header 200 {string} Token "qwerty"
// @Router /account/logoff [DELETE]
func Logoff(context *gin.Context) {
	logoff(context)
}

// ModifyAccountInfo godoc
// @Summary ModifyAccountInfo
// @Description Modify an account by id.
// @Tags account
// @Accept json
// @Produce json
// @Param ModifyAccountInfoParam body account.ModifyAccountInfoParam true "required param."
// @Success 200 {object} base.BasedResponse
// @Header 200 {string} Token "qwerty"
// @Router /account/modify_info [POST]
func ModifyAccountInfo(context *gin.Context) {
	modifyAccountInfo(context)
}