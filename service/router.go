package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
)

// Router ...
func Router(eng *gin.Engine) {

	verV0 := "v0"
	g0 := eng.Group(verV0)
	g0.POST("login", LoginPOST(verV0))
	g0.POST("register", RegisterPOST(verV0))

	v0 := g0.Group("")
	v0.Use(LoginCheck(verV0))

	v0.POST("addUser", AddUserPOST(verV0))

}

// RegisterPOST ...
func RegisterPOST(ver string) gin.HandlerFunc {

	return func(ctx *gin.Context) {

	}
}

// AddUserPOST ...
func AddUserPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := addUser(ctx)
		if err != nil {

		}
	}
}

// User ...
func User(ctx *gin.Context) *model.User {
	if v, b := ctx.Get("user"); b {
		if v0, b := v.(*model.User); b {
			return v0
		}
	}
	return nil
}
