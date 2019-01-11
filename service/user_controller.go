package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
)

// UserReport ...
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
