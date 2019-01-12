package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
)

// ExorcistUserAdd ...
func ExorcistUserAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.NewExorcistUser()
		users, err := user.All()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, users)
	}
}

// ExorcistUserList ...
/**
* @api {get} /v0/exorcist/user 用户列表
* @apiName ExorcistUserList
* @apiGroup ExorcistUser
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
* @apiSampleRequest /v0/exorcist/user
 */
func ExorcistUserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.NewExorcistUser()
		users, err := user.All()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, users)
	}
}

// ExorcistUserUpdate ...
func ExorcistUserUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.NewExorcistUser()
		users, err := user.All()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, users)
	}
}

// ExorcistUserDelete ...
func ExorcistUserDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user := model.NewExorcistUser()
		users, err := user.All()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, users)
	}
}
