package service

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func result(ctx *gin.Context, code int, message string, detail interface{}) {
	h := gin.H{
		"code":    code,
		"message": message,
		"detail":  detail,
	}
	ctx.JSON(http.StatusOK, h)
}

func success(ctx *gin.Context, detail interface{}) {
	result(ctx, 0, "success", detail)
}

func failed(ctx *gin.Context, message string) {
	result(ctx, -1, message, nil)
}

func LoginPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.PostForm("username")
	}
}
