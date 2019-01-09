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
		success(ctx, "")
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
