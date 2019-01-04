package service

import "github.com/gin-gonic/gin"

func LoginCheck(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {

			ctx.Abort()
			return
		}
	}
}
