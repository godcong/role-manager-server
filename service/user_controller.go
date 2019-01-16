package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// UserPermissionList ...
/**
* @api {get} /v0/user/permission 我的权限(UserPermissionList)
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
* @api {get} /v0/user/role 我的角色(UserRoleList)
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

// UserReportList ...
/**
* @api {get} /v0/user/report 用户举报(UserReportList)
* @apiName UserReportList
* @apiGroup UserReport
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
* @apiSampleRequest /v0/user/report
 */
func UserReportList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		report := model.NewReport()
		reports, err := report.ALL()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, reports)
	}
}

// UserReportUpdate ...
/**
* @api {post} /v0/user/report/:id 用户举报(UserReportUpdate)
* @apiName UserReportUpdate
* @apiGroup UserReport
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam  {string} types          		举报类型
* @apiParam  {string} detail         		举报详情
* @apiParam  {string} process_result        处理结果:"obtained",...
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": {
*		        "ID": "5c39984d90881789f46185ba",
*		        "CreatedAt": "2019-01-12T15:33:33.8847341+08:00",
*		        "UpdatedAt": "2019-01-12T15:33:33.8847341+08:00",
*		        "DeletedAt": null,
*		        "Version": 1,
*		        "MediaID": "5c384909078d4d5bd20177be",
*		        "ExoID": "5c384909078d4d5bd20177be",
*		        "Types": "ttttt",
*		        "Detail": "ddddddddddddd",
*		        "ProcessResult": "finished"
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/user/report/:id
 */
func UserReportUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		report := model.NewReport()
		report.ID = model.ID(id)
		err := report.Find()
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		report.Types = ctx.DefaultPostForm("types", report.Types)
		report.ProcessResult = ctx.DefaultPostForm("process_result", report.ProcessResult)
		report.Detail = ctx.DefaultPostForm("detail", report.Detail)
		err = report.Update()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		if report.ProcessResult == model.ReportResultObtained {
			media, err := report.Media()
			if err != nil {
				failed(ctx, err.Error())
				return
			}
			media.Block = true
			err = media.Update()
			if err != nil {
				failed(ctx, err.Error())
				return
			}
		}

		success(ctx, report)
	}
}

// UserReportAdd ...
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
*		        "ID": "5c39984d90881789f46185ba",
*		        "CreatedAt": "2019-01-12T15:33:33.8847341+08:00",
*		        "UpdatedAt": "2019-01-12T15:33:33.8847341+08:00",
*		        "DeletedAt": null,
*		        "Version": 1,
*		        "MediaID": "5c384909078d4d5bd20177be",
*		        "ExoID": "5c384909078d4d5bd20177be",
*		        "Types": "ttttt",
*		        "Detail": "ddddddddddddd",
*		        "ProcessResult": "commit"
*		    },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/report
 */
func UserReportAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		report := model.NewReport()

		err := ctx.BindJSON(report)
		//err := util.UnmarshalJSON(ctx.Request.Body, report)
		if err != nil {
			failed(ctx, err.Error())
			return
		}

		if report.ExoID == primitive.NilObjectID {
			failed(ctx, "null exo_id")
			return
		}

		if report.MediaID == primitive.NilObjectID {
			failed(ctx, "null media_id")
			return
		}

		report.ProcessResult = "commit"

		err = report.Create()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, report)
		return
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

// UserLoginGet ...
func UserLoginGet(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := User(ctx)
		if user != nil {
			failed(ctx, "user not found")
		}
		success(ctx, user)

	}

}
