package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"log"
	"strings"
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

// PermissionCheck ...
func PermissionCheck(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := strings.Replace(ctx.Request.URL.Path, "/", ".", -1)
		p := model.NewPermission()
		p.Slug = path
		err := p.Find()
		if err != nil {
			failed(ctx, err.Error())
			ctx.Abort()
			return
		}

		roles, err := p.Roles()
		user := User(ctx)
		user.Role()

		log.Println("path:", ctx.Request.URL.Path)
		log.Println(ctx.Request.URL.RawPath)
		log.Println(ctx.Request.URL.RawQuery)
		log.Println(ctx.Request.URL.Query().Encode())
	}
}
