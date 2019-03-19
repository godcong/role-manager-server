package permission

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

type permission struct {
	Method  string
	Model   string
	Version string
	Prefix  string
}

func ParseContext(ctx *gin.Context) {
	method := ctx.Request.Method

	uri := ctx.Request.RequestURI

	log.Info(method, uri)
}
