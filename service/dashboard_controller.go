package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
)

// DashboardRoleDelete ...
func DashboardRoleDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardRoleUpdate ...
func DashboardRoleUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardRoleAdd ...
func DashboardRoleAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardRoleList ...
func DashboardRoleList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardRoleShow ...
func DashboardRoleShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardPermissionDelete ...
func DashboardPermissionDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardPermissionUpdate ...
func DashboardPermissionUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardPermissionAdd ...
func DashboardPermissionAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := model.NewPermission()
		slug := ctx.PostForm("slug")
		name := ctx.PostForm("name")
		des := ctx.PostForm("description")
		if des == "" {
			des = name
		}
		p.Slug = slug
		p.Name = name
		p.Description = des
		err := p.Create()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, p)
	}
}

// DashboardPermissionList ...
func DashboardPermissionList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := model.NewPermission()
		permissions, err := p.ALL()
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, permissions)
	}
}

// DashboardUserDelete ...
func DashboardUserDelete(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardUserUpdate ...
func DashboardUserUpdate(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardUserList ...
func DashboardUserList(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardUserAdd ...
func DashboardUserAdd(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		addUser(ctx)

		success(ctx, "")
	}
}

// DashboardUserShow 查看用户信息
func DashboardUserShow(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		addUser(ctx)

		success(ctx, "")
	}
}

// DashboardAdd ...
func DashboardAdd(s string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

// DashboardListGet ...
func DashboardListGet(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
