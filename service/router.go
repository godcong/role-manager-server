package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// Router ...
func Router(eng *gin.Engine) {
	prefix := "api"
	current := "v0"

	st, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	eng.StaticFS("/doc", st)
	eng.StaticFS("/webui", st)

	eng.NoRoute(func(ctx *gin.Context) {
		ctx.Redirect(http.StatusMovedPermanently, "/webui")
	})
	_ = prefix
	g0 := eng.Group(current)
	g0.Use(AccessControlAllow, VisitLog(current))

	//登录
	g0.POST("login", LoginPOST(current))

	//管理员生成
	g0.GET("genesis", GenesisGET(current))
	//用户注册
	g0.POST("register", UserRegister(current))
	//组织申请
	g0.POST("apply", OrganizationApply(current))
	g0.POST("report", UserReportAdd(current))
	g0.GET("play", UserPlayList(current))
	g0.GET("play/:id", UserPlay(current))
	g0.POST("censor/callback", CensorBack(current))
	g0.POST("node/callback", NodeBack(current))

	v0 := g0.Group("")
	v0.Use(LogOutput(current), LoginCheck(current), PermissionCheck(current))

	//超级管理员面板
	//账号、密码、所属组织、角色权限、邮箱、手机号码、授权证书和授权私钥
	dashboard0 := v0.Group("dashboard")
	dashboard0.GET("log", DashboardLogList(current))

	dashboard0.GET("menu", DashboardMenuList(current))
	dashboard0.POST("menu", DashboardMenuAdd(current))
	dashboard0.POST("menu/:id", DashboardMenuUpdate(current))
	dashboard0.DELETE("menu/:id", DashboardMenuDelete(current))

	dashboard0.GET("permission", DashboardPermissionList(current))
	dashboard0.POST("permission", DashboardPermissionAdd(current))
	dashboard0.POST("permission/:id", DashboardPermissionUpdate(current))
	dashboard0.DELETE("permission/:id", DashboardPermissionDelete(current))

	dashboard0.GET("role", DashboardRoleList(current))
	dashboard0.POST("role", DashboardRoleAdd(current))
	dashboard0.POST("role/:id", DashboardRoleUpdate(current))
	dashboard0.DELETE("role/:id", DashboardRoleDelete(current))
	dashboard0.GET("role/:id", DashboardRoleShow(current))
	dashboard0.POST("role/:id/permission/:pid", DashboardRolePermissionAdd(current))

	dashboard0.POST("user", DashboardUserAdd(current))
	dashboard0.GET("user", DashboardUserList(current))
	dashboard0.POST("user/:id", DashboardUserUpdate(current))
	dashboard0.DELETE("user/:id", DashboardUserDelete(current))
	dashboard0.GET("user/:id", DashboardUserShow(current))
	dashboard0.POST("user/:id/permission/:pid", DashboardUserPermissionAdd(current))
	dashboard0.POST("user/:id/role/:rid", DashboardUserRoleAdd(current))
	//节点管理员
	admin0 := v0.Group("admin")

	admin0.POST("organization", AdminOrganizationAdd(current))
	admin0.GET("organization", AdminOrganizationList(current))
	admin0.POST("organization/:id", AdminOrganizationUpdate(current))
	admin0.DELETE("organization/:id", AdminOrganizationDelete(current))
	admin0.GET("organization/:id/user", AdminOrganizationUserList(current))
	admin0.POST("organization/:id/user", AdminOrganizationUserAdd(current))

	//组织管理员
	org0 := v0.Group("org")

	org0.GET("media", OrgMediaList(current))
	org0.POST("media", OrgMediaAdd(current))
	org0.POST("media/:id", OrgMediaUpdate(current))
	org0.GET("media/:id/censor", OrgMediaCensorList(current))
	org0.POST("media/:id/censor", OrgMediaCensorUpdate(current))
	org0.GET("censor/:id", OrgCensorList(current))
	org0.POST("censor/:id", OrgCensorUpdate(current))

	org0.POST("active", OrgActivation(current))
	//org0.POST("upload", OrgUpload(current))

	//监督
	//monitor0 := v0.Group("monitor")
	//monitor0.GET("list", MonitorList(current))

	user0 := v0.Group("user")
	user0.GET("menu", UserMenuList(current))
	user0.GET("media", UserMediaList(current))
	user0.GET("permission", UserPermissionList(current))
	user0.GET("role", UserRoleList(current))
	user0.GET("report", UserReportList(current))
	user0.POST("report/:id", UserReportUpdate(current))
	user0.GET("login", UserLoginGet(current))

	exo0 := v0.Group("exorcist")
	exo0.GET("user", ExorcistUserList(current))
	exo0.POST("user/:id", ExorcistUserUpdate(current))

}

// OrgActivation ...
func OrgActivation(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// JSON ...
type JSON struct {
	Code    int                 `json:"code"`
	Message string              `json:"message"`
	Detail  []*model.ResultData `json:"detail,omitempty"`
}
