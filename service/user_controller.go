package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
)

// UserReport ...
/**
* @api {post} /v0/report 用户登录
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
			failed(ctx, err.Error())
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
