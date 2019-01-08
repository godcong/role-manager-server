package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"github.com/mongodb/mongo-go-driver/bson"
	"log"
)

// Router ...
func Router(eng *gin.Engine) {

	current := "v0"
	eng.Use(AccessControlAllow)
	g0 := eng.Group(current)

	//登录
	g0.POST("login", LoginPOST(current))
	//组织注册

	g0.GET("genesis", GenesisGET(current))

	v0 := g0.Group("")
	v0.Use(LoginCheck(current), PermissionCheck(current))

	//超级管理员面板
	//账号、密码、所属组织、角色权限、邮箱、手机号码、授权证书和授权私钥
	dashboard0 := v0.Group("dashboard")
	dashboard0.GET("list", DashboardListGet(current))
	dashboard0.POST("add", DashboardAdd(current))
	//节点管理员
	admin0 := v0.Group("admin")
	admin0.POST("add", AdminAdd(current))

	//组织管理员
	org0 := v0.Group("org")
	org0.POST("verify", OrgVerify(current))

	//监督
	monitor0 := v0.Group("monitor")
	monitor0.GET("list", MonitorList(current))

	user0 := v0.Group("user")
	user0.POST("add", AddPOST(current))
}

// MonitorList ...
func MonitorList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// OrgVerify ...
func OrgVerify(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// AdminAdd ...
func AdminAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// DashboardAdd ...
func DashboardAdd(s string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// DashboardListGet ...
func DashboardListGet(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func roleIsExist(role *model.Role) bool {
	return model.IsExist(role, bson.M{
		"slug": role.Slug,
	})
}

// ValidateSlug ...
func ValidateSlug(my *model.User, slug string) error {
	myRole, err := my.Role()
	if err != nil {
		return err
	}

	if myRole.Slug == slug {
		return errors.New("can not add same slug")
	}

	switch myRole.Slug {
	case model.SlugGenesis:
		if slug == model.SlugAdmin || slug == model.SlugMonitor {
			return nil
		}
	case model.SlugAdmin:
		if slug == model.SlugOrg {
			return nil
		}
	}
	return errors.New("can not add slug between (" + myRole.Slug + "," + slug + ")")
}

// AddPOST ...
func AddPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		my := User(ctx)
		log.Printf("%+v", *my)
		slug := ctx.PostForm("Slug")
		err := ValidateSlug(my, slug)
		log.Println(err)
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		oid := ctx.PostForm("OID")
		user := model.NewUser()
		user.ID = model.ID(oid)
		err = user.Find()
		log.Println(*user, oid)
		log.Printf("%+v", *user)
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		role := model.NewRole()
		role.Slug = slug
		err = role.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		log.Printf("%+v", *role)
		ru := model.NewRoleUser()
		ru.SetRole(role)
		ru.SetUser(user)
		err = ru.CreateIfNotExist()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, "success")
		return
	}
}

// AddOrgPOST ...
func AddOrgPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// AddAdminPOST ...
func AddAdminPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
