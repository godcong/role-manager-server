package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"log"
)

// UserReport ...
/**
* @api {post} /v0/report 用户举报
* @apiName Report
* @apiGroup Default
* @apiVersion  0.0.1
*
* @apiParam  {string} media_id       举报视频ID
* @apiParam  {string} exo_id         用户ID
* @apiParam  {string} types          举报类型
* @apiParam  {string} detail         举报详情
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*				...
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/report
 */
func UserReport(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		report := model.NewReport()

		eid := ctx.PostForm("exo_id")
		if eid == "" {
			failed(ctx, "null exo_id")
		}
		report.ExoID = model.ID(eid)
		mid := ctx.PostForm("media_id")
		if mid == "" {
			failed(ctx, "null media_id")
		}
		report.MediaID = model.ID(mid)
		report.Types = ctx.PostForm("types")
		report.Detail = ctx.PostForm("detail")
		report.ProcessResult = "commit"

		err := report.Create()
		if err != nil {
			log.Println(err)
			failed(ctx, err.Error())
			return
		}
		success(ctx, report)
		return
	}
}

// UserPermissionList ...
/**
* @api {get} /v0/user/permission 我的权限
* @apiName UserPermissionList
* @apiGroup User
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
*
* @apiUse Failed
* @apiSampleRequest /v0/user/permission
 */
func UserPermissionList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := User(ctx)
		permissions, err := user.Permissions()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, permissions)
	}
}

// UserRoleList ...
/**
* @api {get} /v0/user/role 我的角色
* @apiName UserRoleList
* @apiGroup User
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
*		        "ID": "5c35cc6b5ec8a925a4143001",
*		        "CreatedAt": "2019-01-09T18:26:51.051+08:00",
*		        "UpdatedAt": "2019-01-09T18:26:51.051+08:00",
*		        "DeletedAt": null,
*		        "Version": 1,
*		        "Name": "超级管理员",
*		        "Slug": "genesis",
*		        "Description": "超级管理员",
*		        "Level": 0
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/user/role
 */
func UserRoleList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := User(ctx)
		role, err := user.Role()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, role)
	}
}

// UserMedia ...
func UserMedia(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// UserList ...
func UserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
