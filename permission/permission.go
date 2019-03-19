package permission

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Permission interface {
	Slug() string
}

type permission struct {
	FuncName  string
	CRUD      string
	Dashboard string
	Method    string
	Model     string
	Version   string
	Prefix    string
}

func (p *permission) Slug() string {
	return strings.Join([]string{p.Dashboard, p.Model, p.CRUD}, ".")
}

func parseCRUD(funcName string) string {
	n := len(funcName)
	switch {
	case strings.LastIndex(funcName, "List") == (n - 4):
		return "list"
	case strings.LastIndex(funcName, "Get") == (n - 3):
		return "list"
	case strings.LastIndex(funcName, "Add") == (n - 3):
		return "add"
	case strings.LastIndex(funcName, "Delete") == (n - 6):
		return "delete"
	case strings.LastIndex(funcName, "Update") == (n - 6):
		return "update"
	case strings.LastIndex(funcName, "Show") == (n - 6):
		return "list"
	}
	return ""
}

func parseURI(uri string) []string {
	s := make([]string, 6)
	tmp := strings.Split(uri, "/")
	copy(s, tmp)
	return s
}

func ParseContext(ctx *gin.Context) Permission {
	p := &permission{}
	method := ctx.Request.Method

	uri := parseURI(ctx.Request.RequestURI)
	p.FuncName = handleFuncName(ctx)
	p.Version = strings.ToLower(uri[1])
	p.Dashboard = strings.ToLower(uri[2])
	p.Model = strings.ToLower(uri[3])
	p.Method = strings.ToLower(method)
	p.CRUD = strings.ToLower(parseCRUD(p.FuncName))
	log.Infof("%+v", p)
	return p
}

func handleFuncName(ctx *gin.Context) string {
	hn := strings.Split(ctx.HandlerName(), ".")
	if size := len(hn); size > 2 {
		return hn[size-2]
	}
	return ""
}
