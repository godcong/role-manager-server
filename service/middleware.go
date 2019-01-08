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
		log.Println("path:", ctx.Request.URL.Path)
		log.Println("rp:", ctx.Request.URL.RawPath)
		log.Println("rq:", ctx.Request.URL.RawQuery)
		log.Println("q:", ctx.Request.URL.Query().Encode())
		user := User(ctx)
		role, err := user.Role()
		if err == nil {
			//超级管理员拥有所有权限
			if role.Slug == model.SlugGenesis {
				ctx.Next()
				return
			}
		}
		log.Println(err)

		path := strings.Replace(ctx.Request.URL.Path, "/", ".", -1)
		p := model.NewPermission()
		p.Slug = path
		err = p.Find()
		if err != nil {
			log.Println(err.Error())
			failed(ctx, err.Error())
			ctx.Abort()
			return
		}

		err = user.CheckPermission(p)
		if err != nil {
			log.Println(err.Error())
			failed(ctx, "this account has no permissions to visit")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
