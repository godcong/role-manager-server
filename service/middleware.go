package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"log"
)

// LoginCheck ...
func LoginCheck(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token == "" {
			failed(ctx, "token is null")
			ctx.Abort()
			return
		}
		t, err := FromToken(token)
		if err != nil {
			failed(ctx, err.Error())
			ctx.Abort()
			return
		}

		user := model.NewUser()
		log.Println(t.OID)
		user.ID = model.ID(t.OID)
		err = user.Find()
		if err != nil {
			failed(ctx, err.Error())
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
