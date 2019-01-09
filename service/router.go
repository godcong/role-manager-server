package service

import (
	"github.com/gin-gonic/gin"
)

// Router ...
func Router(eng *gin.Engine) {

	current := "v0"
	eng.Use(AccessControlAllow)
	eng.Static("doc", "./doc")

	g0 := eng.Group(current)

	//登录
	g0.POST("login", LoginPOST(current))

	//管理员生成
	g0.GET("genesis", GenesisGET(current))
	//用户注册
	g0.POST("register", UserRegister(current))
	//组织申请
	g0.POST("apply", OrgApply(current))
	v0 := g0.Group("")
	v0.Use(LogOutput(current), LoginCheck(current), PermissionCheck(current))

	//超级管理员面板
	//账号、密码、所属组织、角色权限、邮箱、手机号码、授权证书和授权私钥
	dashboard0 := v0.Group("dashboard")

	dashboard0.GET("permission", DashboardPermissionList(current))
	dashboard0.POST("permission", DashboardPermissionAdd(current))
	dashboard0.POST("permission/:id", DashboardPermissionUpdate(current))
	dashboard0.DELETE("permission/:id", DashboardPermissionDelete(current))

	dashboard0.GET("role", DashboardRoleList(current))
	dashboard0.POST("role", DashboardRoleAdd(current))
	dashboard0.POST("role/:id", DashboardRoleUpdate(current))
	dashboard0.DELETE("role/:id", DashboardRoleDelete(current))
	dashboard0.GET("role/:id/show", DashboardRoleShow(current))
	dashboard0.POST("role/:id/add", DashboardRolePermissionAdd(current))

	dashboard0.POST("user", DashboardUserAdd(current))
	dashboard0.GET("user", DashboardUserList(current))
	dashboard0.POST("user/:id", DashboardUserUpdate(current))
	dashboard0.DELETE("user/:id", DashboardUserDelete(current))
	dashboard0.GET("user/:id/show", DashboardUserShow(current))

	//节点管理员
	admin0 := v0.Group("admin")
	admin0.POST("organization", AdminOrganizationAdd(current))
	admin0.GET("organization", AdminOrganizationList(current))
	admin0.POST("organization/:id", AdminOrganizationUpdate(current))
	admin0.DELETE("organization/:id", AdminOrganizationDelete(current))
	admin0.GET("organization/:id/show", AdminOrganizationShow(current))

	//组织管理员
	org0 := v0.Group("org")
	//org0.POST("org", OrgAdd(current))
	//org0.GET("org", OrgList(current))
	//org0.POST("org/:id", OrgUpdate(current))
	//org0.DELETE("org/:id", OrgDelete(current))

	org0.POST("active", OrgActivation(current))
	org0.POST("upload", OrgUpload(current))

	//监督
	monitor0 := v0.Group("monitor")
	monitor0.GET("list", MonitorList(current))

	user0 := v0.Group("user")
	user0.GET("play", UserPlayList(current))
	user0.GET("play/:id", UserPlay(current))
}

// UserRegister ...
func UserRegister(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := addUser(ctx)
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, user)
	}
}

// OrgActivation ...
func OrgActivation(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
