package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
	"log"
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
		log.Println(verify)
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

	}
}
