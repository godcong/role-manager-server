package service

import "github.com/gin-gonic/gin"

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
func DashboardPermissionAdd(s string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardPermissionList ...
func DashboardPermissionList(s string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
	}
}

// DashboardPermission ...
func DashboardPermission(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		success(ctx, "")
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
