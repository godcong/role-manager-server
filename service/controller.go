package service

import (
	"github.com/gin-gonic/gin"
	"github.com/godcong/role-manager-server/model"
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

// LoginPOST ...
func LoginPOST(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.PostForm("username")
	}
}

func addUser(ctx *gin.Context) error {
	user := model.NewUser()
	user.Name = ctx.PostForm("name")
	user.Username = ctx.PostForm("username")
	user.Email = ctx.PostForm("email")
	user.Mobile = ctx.PostForm("mobile")
	user.IDCardFacade = ctx.PostForm("idCardFacade")
	user.IDCardObverse = ctx.PostForm("idCardObverse")
	user.Association = ctx.PostForm("association")
	user.SetPassword(ctx.PostForm("password"))
	return user.Create()

}
