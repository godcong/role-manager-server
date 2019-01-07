package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Router ...
func Router(eng *gin.Engine) {

	verV0 := "v0"
	eng.Use(AccessControlAllow)
	g0 := eng.Group(verV0)
	g0.GET("inited", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Code": 0})
	})
	//登录
	g0.POST("login", LoginPOST(verV0))
	//组织注册
	g0.POST("register", RegisterPOST(verV0))

	g0.GET("genesis", GenesisGet(verV0))

	v0 := g0.Group("")
	v0.Use(LoginCheck(verV0))

	v0.POST("addUser", AddUserPOST(verV0))

}
