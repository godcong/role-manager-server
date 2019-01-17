package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"strconv"
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
* @api {get} /v0/exorcist/user 用户列表(ExorcistUserList)
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
*		    "code": 0,
*		    "detail": [
*		        {
*		            "ID": "5c00a23099972d23e41e45e5",
*		            "Sn": [
*		                "0"
*		            ],
*		            "IpfsID": [
*		                "0"
*		            ],
*		            "QuestList": [
*		                {
*		                    "ID": "b3f288725c5cd80f",
*		                    "Code": "dau"
*		                },
*		                {
*		                    "ID": "2dc91542e7067a1d",
*		                    "Code": "dau"
*		                },
*		                {
*		                    "ID": "37c51c786dac5a0f",
*		                    "Code": "dau"
*		                }
*		            ],
*		            "Name": "+86.18217691434",
*		            "Phone": "+8618217691434",
*		            "Password": "8beec6537b3f743bbb857f8113bb9e9080a96301e77f610b192b58dd29f1ea49",
*		            "Nickname": "",
*		            "PictureURL": "",
*		            "Level": 1,
*		            "CreatedAt": "2018-11-30T10:36:31.626+08:00",
*		            "Binded": false,
*		            "QueryApply": false,
*		            "Order": false,
*		            "WhaleCard": "",
*		            "WhaleOrder": "",
*		            "SlotNum": 3,
*		            "Approved": false,
*		            "ParentID": "5c009cf97b7a052d94da7131",
*		            "Dvc": "",
*		            "WhaleDvc": "",
*		            "DragonBall": "",
*		            "Master": "",
*		            "WeChatUnionid": "",
*		            "WeChatAppOpenid": "",
*		            "WeChatAppToken": "",
*		            "WeChatAppRefreshToken": "",
*		            "V": 4
*		        },
*		        {
*		            "ID": "5c123bafda9bf30ce9bb3b96",
*					...
*		        },
*		    "message": "success"
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
/**
* @api {post} /v0/exorcist/user/:id 更新用户(ExorcistUserUpdate)
* @apiName ExorcistUserUpdate
* @apiGroup ExorcistUser
* @apiVersion  0.0.1
*
* @apiHeader {string} token user token
*
* @apiParam {bool} block 禁止访问
*
* @apiUse Success
* @apiSuccess (detail) {string} id Id
* @apiSuccess (detail) {string} other 参考返回Example
* @apiSuccessExample {json} Success-Response:
*		{
*		    "code": 0,
*		    "detail": [
*		        {
*		            "ID": "5c00a23099972d23e41e45e5",
*		            "Sn": [
*		                "0"
*		            ],
*		            "IpfsID": [
*		                "0"
*		            ],
*		            "QuestList": [
*		                {
*		                    "ID": "b3f288725c5cd80f",
*		                    "Code": "dau"
*		                },
*		                {
*		                    "ID": "2dc91542e7067a1d",
*		                    "Code": "dau"
*		                },
*		                {
*		                    "ID": "37c51c786dac5a0f",
*		                    "Code": "dau"
*		                }
*		            ],
*		            "Name": "+86.18217691434",
*		            "Phone": "+8618217691434",
*		            "Password": "8beec6537b3f743bbb857f8113bb9e9080a96301e77f610b192b58dd29f1ea49",
*		            "Nickname": "",
*		            "PictureURL": "",
*		            "Level": 1,
*		            "CreatedAt": "2018-11-30T10:36:31.626+08:00",
*		            "Binded": false,
*		            "QueryApply": false,
*		            "Order": false,
*		            "WhaleCard": "",
*		            "WhaleOrder": "",
*		            "SlotNum": 3,
*		            "Approved": false,
*		            "ParentID": "5c009cf97b7a052d94da7131",
*		            "Dvc": "",
*		            "WhaleDvc": "",
*		            "DragonBall": "",
*		            "Master": "",
*		            "WeChatUnionid": "",
*		            "WeChatAppOpenid": "",
*		            "WeChatAppToken": "",
*		            "WeChatAppRefreshToken": "",
*		            "V": 4
*		        },
*		        {
*		            "ID": "5c123bafda9bf30ce9bb3b96",
*					...
*		        },
*		    "message": "success"
*		}
*
* @apiUse Failed
* @apiSampleRequest /v0/exorcist/user
 */
func ExorcistUserUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
		user := model.NewExorcistUser()

		user.ID = model.ID(id)
		err := user.Find()

		if err != nil {
			failed(ctx, err.Error())
			return
		}
		user.Block, _ = strconv.ParseBool(ctx.PostForm("block"))
		err = user.Update()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, user)
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
