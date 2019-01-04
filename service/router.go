package service

import "github.com/gin-gonic/gin"

// Router ...
func Router(eng *gin.Engine) {

	verV0 := "v0"
	g0 := eng.Group(verV0)
	g0.POST("login", LoginPOST(verV0))
	g0.POST("register", RegisterPOST(verV0))

	v0 := g0.Group("")
	v0.Use(LoginCheck(verV0))

	v0.POST("add", AddUserPOST(verV0))

}

// RegisterPOST ...
func RegisterPOST(ver string) gin.HandlerFunc {
	return nil
}

// AddUserPOST ...
func AddUserPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := addUser(ctx)
		if err != nil {

		}
	}
}
