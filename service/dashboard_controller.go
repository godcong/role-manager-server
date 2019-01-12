package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
)

// DashboardRoleDelete ...
/**
* @api {delete} /v0/dashboard/role/:id 删除角色(暂不支持)
* @apiName DashboardRoleDelete
* @apiGroup DashboardRole
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/role
 */
func DashboardRoleDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		failed(ctx, "can't delete role now")
		return
		success(ctx, "")
	}
}

// DashboardRoleUpdate ...
/**
* @api {post} /v0/dashboard/role/:id 更新角色(暂不支持)
* @apiName DashboardRoleUpdate
* @apiGroup DashboardRole
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} name				名称
* @apiParam  {string} slug				角色
* @apiParam  {string} [description]		说明
* @apiParam  {string} [level]			等级
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/role
 */
func DashboardRoleUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		failed(ctx, "can't change role now")
		return
		success(ctx, "")
	}
}

// DashboardRoleAdd ...
/**
* @api {post} /v0/dashboard/role 添加角色(暂不支持)
* @apiName DashboardRoleAdd
* @apiGroup DashboardRole
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} name				名称
* @apiParam  {string} slug				角色
* @apiParam  {string} [description]		说明
* @apiParam  {string} [level]			等级
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/role
 */
func DashboardRoleAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		failed(ctx, "can't add new role now")
		return
		//TODO
		success(ctx, "")
	}
}

// DashboardRoleList ...
/**
* @api {get} /v0/dashboard/role 角色列表
* @apiName DashboardRoleList
* @apiGroup DashboardRole
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": [
*		        {
*		            "ID": "5c3596d3e0b207fb17d6ddf9",
*		            "CreatedAt": "2019-01-09T14:38:11.564+08:00",
*		            "UpdatedAt": "2019-01-09T14:38:11.564+08:00",
*		            "DeletedAt": null,
*		            "Version": 1,
*		            "Name": "超级管理员",
*		            "Slug": "genesis",
*		            "Description": "超级管理员",
*		            "Level": 0
*		        },
*		        {
*		            "ID": "5c3596d3e0b207fb17d6ddfa",
*		            "CreatedAt": "2019-01-09T14:38:11.597+08:00",
*		            "UpdatedAt": "2019-01-09T14:38:11.597+08:00",
*		            "DeletedAt": null,
*		            "Version": 1,
*		            "Name": "节点管理员",
*		            "Slug": "admin",
*		            "Description": "节点管理员",
*		            "Level": 0
*		        },
*		        {
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
*		        {
*		            "ID": "5c3596d3e0b207fb17d6ddfc",
*		            "CreatedAt": "2019-01-09T14:38:11.601+08:00",
*		            "UpdatedAt": "2019-01-09T14:38:11.601+08:00",
*		            "DeletedAt": null,
*		            "Version": 1,
*		            "Name": "监督",
*		            "Slug": "monitor",
*		            "Description": "监督",
*		            "Level": 0
*		        },
*		        {
*		            "ID": "5c3596d3e0b207fb17d6ddfd",
*		            "CreatedAt": "2019-01-09T14:38:11.603+08:00",
*		            "UpdatedAt": "2019-01-09T14:38:11.603+08:00",
*		            "DeletedAt": null,
*		            "Version": 1,
*		            "Name": "普通用户",
*		            "Slug": "user",
*		            "Description": "普通用户",
*		            "Level": 0
*		        }
*		    ],
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/role
 */
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
/**
* @api {get} /v0/dashboard/role/:id/show 角色权限
* @apiName DashboardRoleShow
* @apiGroup DashboardRole
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "ID": "5c35a4481afae2f7afac1a2c",
*		        "CreatedAt": "2019-01-09T15:35:36.44+08:00",
*		        "UpdatedAt": "2019-01-09T15:44:18.4474311+08:00",
*		        "DeletedAt": null,
*		        "Version": 4,
*		        "Name": "列表权限",
*		        "Slug": "DashboardPermissionList",
*		        "Description": "列表权限",
*		        "PermissionModel": ""
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/role
 */
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

// DashboardRolePermissionAdd ...
/**
* @api {post} /v0/dashboard/role/:id/permission 添加角色权限
* @apiName DashboardRolePermissionAdd
* @apiGroup DashboardRole
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} permission_id		权限ID
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "ID": "5c35a4481afae2f7afac1a2c",
*		        "CreatedAt": "2019-01-09T15:35:36.44+08:00",
*		        "UpdatedAt": "2019-01-09T15:44:18.4474311+08:00",
*		        "DeletedAt": null,
*		        "Version": 4,
*		        "Name": "列表权限",
*		        "Slug": "DashboardPermissionList",
*		        "Description": "列表权限",
*		        "PermissionModel": ""
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/role/:id/permission
 */
func DashboardRolePermissionAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		role := model.NewRole()
		role.ID = model.ID(id)
		err := role.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		pid := ctx.PostForm("permission_id")
		p := model.NewPermission()
		p.ID = model.ID(pid)
		err = p.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		users, err := role.Users()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		err = model.Transaction(func() error {
			pr := model.NewPermissionRole()
			pr.SetPermission(p)
			pr.SetRole(role)
			err = pr.CreateIfNotExist()
			if err != nil {
				return err
			}

			for _, user := range users {
				pu := model.NewPermissionUser()
				pu.SetPermission(p)
				pu.SetUser(user)
				err := pu.CreateIfNotExist()
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		success(ctx, gin.H{
			"role":       role,
			"permission": p,
		})
	}
}

// DashboardPermissionDelete ...
/**
* @api {post} /v0/dashboard/permission/:id 删除权限
* @apiName DashboardPermissionDelete
* @apiGroup DashboardPermission
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "ID": "5c35a4481afae2f7afac1a2c",
*		        "CreatedAt": "2019-01-09T15:35:36.44+08:00",
*		        "UpdatedAt": "2019-01-09T15:44:18.4474311+08:00",
*		        "DeletedAt": null,
*		        "Version": 4,
*		        "Name": "列表权限",
*		        "Slug": "DashboardPermissionList",
*		        "Description": "列表权限",
*		        "PermissionModel": ""
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/permission
 */
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
/**
* @api {post} /v0/dashboard/permission/:id 更新权限
* @apiName DashboardPermissionUpdate
* @apiGroup DashboardPermission
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} slug        	权限函数(apiName)
* @apiParam  {string} name     		权限名称
* @apiParam  {string} [description]   说明
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "ID": "5c35a4481afae2f7afac1a2c",
*		        "CreatedAt": "2019-01-09T15:35:36.44+08:00",
*		        "UpdatedAt": "2019-01-09T15:40:39.1569541+08:00",
*		        "DeletedAt": null,
*		        "Version": 2,
*		        "Name": "权限列表",
*		        "Slug": "DashboardPermissionList",
*		        "Description": "权限列表",
*		        "PermissionModel": ""
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/permission
 */
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
/**
* @api {post} /v0/dashboard/permission 添加权限
* @apiName DashboardPermissionAdd
* @apiGroup DashboardPermission
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} slug        	权限函数(apiName)
* @apiParam  {string} name     		权限名称
* @apiParam  {string} [description]   说明
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": [
*		        {
*		            "ID": "5c35a4481afae2f7afac1a2c",
*		            "CreatedAt": "2019-01-09T15:35:36.44+08:00",
*		            "UpdatedAt": "2019-01-09T15:41:08.262+08:00",
*		            "DeletedAt": null,
*		            "Version": 3,
*		            "Name": "列表权限",
*		            "Slug": "DashboardPermissionList",
*		            "Description": "列表权限",
*		            "PermissionModel": ""
*		        },
*               {
*                   "ID": "5c35a5d51afae2f7afac1a2d",
*                   "CreatedAt": "2019-01-09T15:42:13.416+08:00",
*                   "UpdatedAt": "2019-01-09T15:42:13.416+08:00",
*                   "DeletedAt": null,
*                   "Version": 1,
*                   "Name": "添加权限",
*                   "Slug": "DashboardPermissionAdd",
*                   "Description": "添加权限",
*                   "PermissionModel": ""
*               }
*		    ],
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/permission
 */
func DashboardPermissionAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := model.NewPermission()
		p.Slug = ctx.PostForm("slug")
		p.Name = ctx.PostForm("name")
		p.Description = ctx.DefaultPostForm("description", p.Name)
		err := p.CreateIfNotExist()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, p)
	}
}

// DashboardPermissionList ...
/**
* @api {get} /v0/dashboard/permission 权限列表
* @apiName DashboardPermissionList
* @apiGroup DashboardPermission
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": [
*		        {
*		            "ID": "5c35a4481afae2f7afac1a2c",
*		            "CreatedAt": "2019-01-09T15:35:36.44+08:00",
*		            "UpdatedAt": "2019-01-09T15:41:08.262+08:00",
*		            "DeletedAt": null,
*		            "Version": 3,
*		            "Name": "列表权限",
*		            "Slug": "DashboardPermissionList",
*		            "Description": "列表权限",
*		            "PermissionModel": ""
*		        }
*		    ],
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/permission
 */
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

// DashboardUserList ...
/**
* @api {get} /v0/dashboard/user 管理用户列表
* @apiName DashboardUserList
* @apiGroup DashboardUser
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": [
*		        {
*		            "ID": "5c3596d716fbec777db5a645",
*		            ...
*		        },
*		        {
*		            "ID": "5c3596d716fbec777db5a646",
*		            ...
*		        },
*		    ],
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/user
 */
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
/**
* @api {post} /v0/dashboard/user 添加管理用户
* @apiName DashboardUserAdd
* @apiGroup DashboardUser
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} name            名称
* @apiParam  {string} username        用户名
* @apiParam  {string} email           邮件
* @apiParam  {string} mobile          移动电话
* @apiParam  {string} id_card_facade  身份证(正)
* @apiParam  {string} id_card_obverse 身份证(反)
* @apiParam  {string} organization_id 组织ID
* @apiParam  {string} password        密码
* @apiParam  {string} certificate     证书
* @apiParam  {string} private_key     私钥
*
* @apiParam  {string} slug     		  用户角色:genesis,admin,organization,monitor,user
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
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
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/user
 */
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

// DashboardUserUpdate ...
/**
* @api {post} /v0/dashboard/user/:id 更新管理用户
* @apiName DashboardUserUpdate
* @apiGroup DashboardUser
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} name           	名称
* @apiParam  {string} username       	用户名
* @apiParam  {string} email          	邮件
* @apiParam  {string} mobile         	移动电话
* @apiParam  {string} id_card_facade 	身份证(正)
* @apiParam  {string} id_card_obverse	身份证(反)
* @apiParam  {string} organization_id	组织ID
* @apiParam  {string} password       	密码
* @apiParam  {string} certificate    	证书
* @apiParam  {string} private_key    	私钥
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
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

// DashboardUserDelete ...
/**
* @api {delete} /v0/dashboard/user/:id 删除管理用户
* @apiName DashboardUserDelete
* @apiGroup DashboardUser
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
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

// DashboardUserShow 查看用户信息
/**
* @api {get} /v0/dashboard/user/:id/show 管理用户信息
* @apiName DashboardUserShow
* @apiGroup DashboardUser
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "organization": {
*		            "ID": "000000000000000000000000",
*		            "CreatedAt": "2019-01-09T15:02:19.2075483+08:00",
*		            "UpdatedAt": "2019-01-09T15:02:19.2075483+08:00",
*		            "DeletedAt": null,
*		            "Version": 1,
*		            "IsDefault": false,
*		            "Verify": "",
*		            "Name": "",
*		            "Code": "",
*		            "Contact": "",
*		            "Position": "",
*		            "Phone": "",
*		            "Mailbox": "",
*		            "Description": ""
*		        },
*		        "permissions": null,
*		        "role": {
*		            "ID": "5c3596d3e0b207fb17d6ddf9",
*		            "CreatedAt": "2019-01-09T14:38:11.564+08:00",
*		            "UpdatedAt": "2019-01-09T14:38:11.564+08:00",
*		            "DeletedAt": null,
*		            "Version": 1,
*		            "Name": "超级管理员",
*		            "Slug": "genesis",
*		            "Description": "超级管理员",
*		            "Level": 0
*		        },
*		        "user": {
*		            "ID": "5c3596d716fbec777db5a645",
*		            "CreatedAt": "2019-01-09T14:38:15.191+08:00",
*		            "UpdatedAt": "2019-01-09T14:38:15.191+08:00",
*		            "DeletedAt": null,
*		            "Version": 1,
*		            "Name": "genesis",
*		            "Username": "",
*		            "Email": "",
*		            "Mobile": "",
*		            "IDCardFacade": "",
*		            "IDCardObverse": "",
*		            "OrganizationID": "000000000000000000000000",
*		            "Password": "DBD978CCDBBE8B6DE77F6B37B5DF9B5B62A7E892A501C3B53EAA16B0838BD5ED",
*		            "Certificate": "",
*		            "PrivateKey": "",
*		            "Token": ""
*		        }
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/dashboard/user
 */
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
