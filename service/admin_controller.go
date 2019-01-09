package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
)

// AdminOrganizationDelete ...
func AdminOrganizationDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// AdminOrganizationUpdate ...
/**
* @api {post} /v0/admin/organization/:id 更新组织
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

// AdminOrganizationList ...
func AdminOrganizationList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		org := model.NewOrganization()
		organizations, err := org.ALL()
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
		name := ctx.PostForm("applyName")         //商户名称
		code := ctx.PostForm("applyCode")         //社会统一信用代码
		contact := ctx.PostForm("applyContact")   //商户联系人
		position := ctx.PostForm("applyPosition") //联系人职位
		phone := ctx.PostForm("applyPhone")       //联系人手机号
		mail := ctx.PostForm("applyMailbox")      //联系人邮箱
		org := model.NewOrganization()
		org.Name = name
		org.Code = code
		org.Contact = contact
		org.Position = position
		org.Phone = phone
		org.Mailbox = mail
		org.Verify = model.VerifyPass //申请中
		err := org.CreateIfNotExist()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, org)
	}
}

// AdminOrganizationShow ...
func AdminOrganizationShow(ver string) gin.HandlerFunc {
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
