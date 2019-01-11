package service

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"net/url"
)

// CensorServer ...
var CensorServer = "localhost:7789"

// OrgMediaUpload ...
func OrgMediaUpload(ver string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		key := ctx.PostForm("object_key")
		resp, err := http.PostForm(CensorServer,
			url.Values{
				"name": []string{key},
			})
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		bytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			failed(ctx, err.Error())
			return
		}
		success(ctx, bytes)

	}
}
