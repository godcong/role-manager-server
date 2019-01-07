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

	verV0 := "v0"
	eng.Use(AccessControlAllow)
	g0 := eng.Group(verV0)
	eng.POST("init", func(ctx *gin.Context) {
		genesis := model.NewGenesis()
		if !roleIsExist(genesis) {
			genesis.Create()
		}
		admin := model.NewAdmin()
		if !roleIsExist(admin) {
			admin.Create()
		}
		org := model.NewOrg()
		if !roleIsExist(org) {
			org.Create()
		}
		monitor := model.NewMonitor()
		if !roleIsExist(monitor) {
			monitor.Create()
		}
		user := model.NewGod()
		if !roleIsExist(user) {
			user.Create()
		}
		success(ctx, "success")
	})
	//登录
	g0.POST("login", LoginPOST(verV0))
	//组织注册
	g0.POST("register", RegisterPOST(verV0))

	g0.GET("genesis", GenesisGET(verV0))

	v0 := g0.Group("")
	v0.Use(LoginCheck(verV0))
	v0.POST("addAdmin", AddAdminPOST(verV0))
	v0.POST("addOrg", AddOrgPOST(verV0))
	v0.POST("addUser", AddUserPOST(verV0))
	v0.POST("add", AddPOST(verV0))
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

		role, err := model.RoleBySlug(slug)
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
