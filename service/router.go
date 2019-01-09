package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"github.com/rakyll/statik/fs"
	"log"
)

// Router ...
func Router(eng *gin.Engine) {
	st, err := fs.New()
	if err != nil {
		log.Fatal(err)
	}
	eng.StaticFS("/doc", st)

	current := "v0"
	eng.Use(AccessControlAllow)
	//eng.Static("doc", "./doc")

	g0 := eng.Group(current)

	//登录
	g0.POST("login", LoginPOST(current))

	//管理员生成
	g0.GET("genesis", GenesisGET(current))
	//用户注册
	g0.POST("register", UserRegister(current))
	//组织申请
	g0.POST("apply", OrganizationApply(current))

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
	dashboard0.POST("role/:id/permission", DashboardRolePermissionAdd(current))

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
	admin0.POST("organization/:id/user", AdminOrganizationUserUpdate(current))

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

// AdminOrganizationUserUpdate ...
/**
* @api {post} /v0/admin/organization/:id/user 添加用户
* @apiName AdminOrganizationUserUpdate
* @apiGroup AdminOrganization
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token

* @apiParam  {string} user_id            用户ID
* @apiParam  {string} apply              类型:true
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "organization": {
*		            "ID": "5c35b06daad2d1c5eb7292bd",
*		            "CreatedAt": "2019-01-09T16:27:25.903+08:00",
*		            "UpdatedAt": "2019-01-09T16:47:22.266+08:00",
*		            "DeletedAt": null,
*		            "Version": 4,
*		            "IsDefault": false,
*		            "Verify": "return",
*		            "Name": "商户名称",
*		            "Code": "社会统一信用代码",
*		            "Contact": "商户联系人",
*		            "Position": "联系人职位",
*		            "Phone": "联系人手机号",
*		            "Mailbox": "联系人邮箱",
*		            "Description": ""
*		        },
*		        "role": {
*		            "ID": "5c3596d3e0b207fb17d6ddfb",
*		            "CreatedAt": "2019-01-09T14:38:11.599+08:00",
*		            "UpdatedAt": "2019-01-09T14:38:11.599+08:00",
*		            "DeletedAt": null,
*		            "Version": 1,
*		            "Name": "组织管理员",
*		            "Slug": "organization",
*		            "Description": "组织管理员",
*		            "Level": 0
*		        },
*		        "user": {
*		            "ID": "5c35b8355f262f9b85b765a2",
*		            "CreatedAt": "2019-01-09T17:00:37.669+08:00",
*		            "UpdatedAt": "2019-01-09T17:00:37.669+08:00",
*		            "DeletedAt": null,
*		            "Version": 1,
*		            "Name": "godcong",
*		            "Username": "godcong",
*		            "Email": "godcong@163.com",
*		            "Mobile": "13058750423",
*		            "IDCardFacade": "",
*		            "IDCardObverse": "",
*		            "OrganizationID": "5c35b06daad2d1c5eb7292bd",
*		            "Password": "DBD978CCDBBE8B6DE77F6B37B5DF9B5B62A7E892A501C3B53EAA16B0838BD5ED",
*		            "Certificate": "noacc",
*		            "PrivateKey": "noacc",
*		            "Token": ""
*		        }
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/admin/organization/:id/user
 */
func AdminOrganizationUserUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		uid := ctx.PostForm("user_id")
		user := model.NewUser()
		user.ID = model.ID(uid)
		err := user.Find()
		if err != nil {
			log.Println(err)
			failed(ctx, err.Error())
			return
		}
		org := model.NewOrganization()
		org.ID = model.ID(id)
		err = org.Find()
		if err != nil {
			log.Println(err)
			failed(ctx, err.Error())
			return
		}

		apply := ctx.PostForm("apply")
		if apply != "true" {
			failed(ctx, apply)
			return
		}

		role := model.NewOrgRole()
		err = role.Find()
		err = addPermission(func() (*model.User, error) {
			return user, nil
		}, role)
		if err != nil {
			log.Println(err)
			failed(ctx, err.Error())
			return
		}

		success(ctx, gin.H{
			"user":         user,
			"role":         role,
			"organization": org,
		})

	}
}

// OrgActivation ...
func OrgActivation(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
