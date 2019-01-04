package service

import "github.com/gin-gonic/gin"

func Router(eng *gin.Engine) {

	verV0 := "v0"
	v0 := eng.Group(verV0)
	v0.Use(LoginCheck(verV0))
	v0.POST("login", LoginPOST(verV0))
}
