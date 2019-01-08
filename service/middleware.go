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

func handleFuncName(ctx *gin.Context) string {
	hn := strings.Split(ctx.HandlerName(), ".")
	size := len(hn)
	if size > 0 {
		size -= 2
	}
	return hn[size]
}

// PermissionCheck ...
func PermissionCheck(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		log.Println("method", ctx.Request.Method)
		log.Println("path:", ctx.Request.URL.Path)
		handle := handleFuncName(ctx)
		log.Println("handle:", handle)

		user := User(ctx)
		role, err := user.Role()
		if err == nil {
			//超级管理员拥有所有权限
			if role.Slug == model.SlugGenesis {
				ctx.Next()
				return
			}
		}

		p := model.NewPermission()
		p.Slug = handle
		err = p.Find()
		if err != nil {
			log.Println(err.Error())
			nop(ctx, err.Error())
			ctx.Abort()
			return
		}

		err = user.CheckPermission(p)
		if err != nil {
			log.Println(err.Error())
			nop(ctx, "this account has no permissions to visit this url")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
