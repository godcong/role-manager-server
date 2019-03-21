package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	log "github.com/sirupsen/logrus"
)

// AdminOrganizationDelete ...
func AdminOrganizationDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// AdminOrganizationUpdate ...
/**
* @api {post} /v0/admin/organization/:id 更新组织(AdminOrganizationUpdate)
* @apiName AdminOrganizationUpdate
* @apiGroup AdminOrganization
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} verify           		申请状态: 通过(pass),申请中(application),打回(return),关闭(closed)
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "ID": "5c35b06daad2d1c5eb7292bd",
*		        "CreatedAt": "2019-01-09T16:27:25.903+08:00",
*		        "UpdatedAt": "2019-01-09T16:37:37.3798805+08:00",
*		        "DeletedAt": null,
*		        "Version": 3,
*		        "IsDefault": false,
*		        "Verify": "pass",
*		        "Name": "商户名称",
*		        "Code": "社会统一信用代码",
*		        "Contact": "商户联系人",
*		        "Position": "联系人职位",
*		        "Phone": "联系人手机号",
*		        "Mailbox": "联系人邮箱",
*		        "Description": ""
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/admin/organization/:id
 */
func AdminOrganizationUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		org := model.NewOrganization()
		org.ID = model.ID(id)
		err := org.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		verify := ctx.PostForm("verify")
		org.Verify = verify
		err = org.Update()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, org)
	}
}

// AdminOrganizationList ....
/**
* @api {get} /v0/admin/organization 组织列表(AdminOrganizationList)
* @apiName AdminOrganizationList
* @apiGroup AdminOrganization
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
*		            "ID": "5c35b06daad2d1c5eb7292bd",
*		            "CreatedAt": "2019-01-09T16:27:25.903+08:00",
*		            "UpdatedAt": "2019-01-09T16:37:37.379+08:00",
*		            "DeletedAt": null,
*		            "Version": 3,
*		            "IsDefault": false,
*		            "Verify": "pass",
*		            "Name": "商户名称",
*		            "Code": "社会统一信用代码",
*		            "Contact": "商户联系人",
*		            "Position": "联系人职位",
*		            "Phone": "联系人手机号",
*		            "Mailbox": "联系人邮箱",
*		            "Description": ""
*		        }
*		    ],
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/admin/organization
 */
func AdminOrganizationList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		org := model.NewOrganization()
		organizations, err := org.All()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, organizations)
	}
}

// AdminOrganizationAdd ...
func AdminOrganizationAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//request: v0/apply
		success(ctx, nil)
	}
}

// AdminOrganizationUserList ...
/**
* @api {get} /v0/admin/organization/:id/user 组织管理用户(AdminOrganizationUserList)
* @apiName AdminOrganizationUserList
* @apiGroup AdminOrganization
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
* @apiSampleRequest /v0/admin/organization/:id/user
 */
func AdminOrganizationUserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		org := model.NewOrganization()
		users, err := org.Users()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, users)
	}
}

// AdminOrganizationUserAdd ...
/**
* @api {post} /v0/admin/organization/:id/user 添加组织管理用户(AdminOrganizationUserAdd)
* @apiName AdminOrganizationUserAdd
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
func AdminOrganizationUserAdd(ver string) gin.HandlerFunc {
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
