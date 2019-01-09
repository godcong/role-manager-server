package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"log"
)

// DashboardRoleDelete ...
func DashboardRoleDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		failed(ctx, "can't delete role now")
		return
		success(ctx, "")
	}
}

// DashboardRoleUpdate ...
func DashboardRoleUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		failed(ctx, "can't change role now")
		return
		success(ctx, "")
	}
}

// DashboardRoleAdd ...
func DashboardRoleAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		failed(ctx, "can't add new role now")
		return
		success(ctx, "")
	}
}

// DashboardRoleList ...
func DashboardRoleList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		r := model.NewRole()
		roles, err := r.ALL()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, roles)
	}
}

// DashboardRoleShow ...
func DashboardRoleShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		r := model.NewRole()
		r.ID = model.ID(id)
		err := r.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		permissions, err := r.Permissions()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		success(ctx, permissions)
	}
}

// DashboardPermissionDelete ...
func DashboardPermissionDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		p := model.NewPermission()
		p.ID = model.ID(id)
		err := p.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		err = p.Delete()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, p)
	}
}

// DashboardPermissionUpdate ...
func DashboardPermissionUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		p := model.NewPermission()
		p.ID = model.ID(id)
		err := p.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		name := ctx.PostForm("name")
		if name != "" {
			p.Name = name
			p.Description = name
		}
		des := ctx.PostForm("description")
		if des != "" {
			p.Description = des
		}
		slug := ctx.PostForm("slug")
		if slug != "" {
			p.Slug = slug
		}
		err = p.Update()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		success(ctx, p)
	}
}

// DashboardPermissionAdd ...
func DashboardPermissionAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := model.NewPermission()
		slug := ctx.PostForm("slug")
		name := ctx.PostForm("name")
		des := ctx.PostForm("description")
		if des == "" {
			des = name
		}
		p.Slug = slug
		p.Name = name
		p.Description = des
		err := p.CreateIfNotExist()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, p)
	}
}

// DashboardPermissionList ...
func DashboardPermissionList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := model.NewPermission()
		permissions, err := p.ALL()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, permissions)
	}
}

// DashboardUserDelete ...
/**
* @api {delete} /v0/dashboard/user/:id 删除用户
* @apiName DashboardUserDelete
* @apiGroup DashboardUser
* @apiVersion  0.0.1
*
* @apiUse Success
* @apiSuccess (detail) {json} id 文件名ID
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "ID": "5c3596d716fbec777db5a645",
*		        "CreatedAt": "2019-01-09T14:38:15.191+08:00",
*		        "UpdatedAt": "2019-01-09T14:38:15.191+08:00",
*		        "DeletedAt": null,
*		        "Version": 1,
*		        "Name": "genesis",
*		        "Username": "",
*		        "Email": "",
*		        "Mobile": "",
*		        "IDCardFacade": "",
*		        "IDCardObverse": "",
*		        "OrganizationID": "000000000000000000000000",
*		        "Password": "DBD978CCDBBE8B6DE77F6B37B5DF9B5B62A7E892A501C3B53EAA16B0838BD5ED",
*		        "Certificate": "",
*		        "PrivateKey": "",
*		        "Token": ""
*		    },
*		    "message": "success"
*		}
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/user
 */
func DashboardUserDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser()
		user.ID = model.ID(id)
		log.Println(user)
		err := user.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		err = user.Delete()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, user)
	}
}

// DashboardUserUpdate ...
func DashboardUserUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := updateUser(ctx)
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, user)
	}
}

// DashboardUserList ...
func DashboardUserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.NewUser()
		users, err := user.ALL()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, users)
	}
}

// DashboardUserAdd ...
func DashboardUserAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := addUser(ctx)
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, user)
	}
}

// DashboardUserShow 查看用户信息
func DashboardUserShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewUser()
		user.ID = model.ID(id)
		err := user.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		p, _ := user.Permissions()
		r, _ := user.Role()
		o, _ := user.Organization()
		success(ctx, gin.H{
			"user":         user,
			"role":         r,
			"permissions":  p,
			"organization": o,
		})
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
