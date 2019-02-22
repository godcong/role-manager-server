package service

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
		log "github.com/sirupsen/logrus"
	"strings"
)

// MaxMultipartMemory ...
const MaxMultipartMemory = 8 << 20

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
		user := User(ctx)
		role, err := user.Role()
		if err == nil {
			//超级管理员拥有所有权限
			if role.Slug == model.SlugGenesis {
				ctx.Next()
				return
			}
		}

		if user.Block {
			nop(ctx, "this account has been blocked")
			return
		}

		p := model.NewPermission()
		//logger := Logger(ctx)
		p.Slug = role.Slug
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

// VisitLog ...
func VisitLog(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		l := model.NewLog()
		l.Permission = handleFuncName(ctx)
		l.Method = ctx.Request.Method
		l.URL = ctx.Request.URL.String()
		detail, err := GetPostFormString(ctx)
		if err != nil {
			l.Err = err.Error()
		}
		l.Detail = detail
		user, err := decodeUser(ctx)
		if err != nil {
			l.Err = err.Error()
		}
		l.UserID = user.ID

		l.VisitIP = ctx.Request.Header.Get("REMOTE-HOST")
		err = l.Create()
		ctx.Set("logger", l)
		log.Println(	log "github.com/sirupsen/logrus", *l, err)
	}
}

// GetPostFormString ...
func GetPostFormString(ctx *gin.Context) (string, error) {
	req := ctx.Request
	req.ParseForm()
	req.ParseMultipartForm(MaxMultipartMemory)
	if len(req.PostForm) > 0 {
		bytes, err := json.Marshal(req.PostForm)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}
	if req.MultipartForm != nil && req.MultipartForm.File != nil {
		bytes, err := json.Marshal(req.PostForm)
		if err != nil {
			return "", err
		}
		return string(bytes), nil
	}

	return "", nil
}

func decodeUser(ctx *gin.Context) (*model.User, error) {
	user := User(ctx)
	if user != nil {
		return user, nil
	}

	token := ctx.Request.Header.Get("token")
	if token == "" {
		return &model.User{}, errors.New("token is null")
	}
	t, err := FromToken(token)
	if err != nil {
		return &model.User{}, err
	}

	user = model.NewUser()
	log.Println(t.OID)
	user.ID = model.ID(t.OID)
	err = user.Find()
	if err != nil {
		return &model.User{}, err
	}
	return user, nil
}
