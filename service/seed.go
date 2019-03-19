package service

import (
	"github.com/Sirupsen/logrus"
	"github.com/godcong/role-manager-server/model"
)

func Seed() {
	for _, v := range Permissions() {
		e := model.InsertOne(v)
		if e != nil {
			return
		}
	}
}

func Permissions() []*model.Permission {
	var permissions []*model.Permission
	var p *model.Permission
	p = model.NewPermission()
	p.Slug = "dashboard.log.list"
	p.Name = "日志信息"
	p.Description = "日志信息"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "dashboard.permission.list"
	p.Name = "权限列表"
	p.Description = "权限列表"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "dashboard.user.list"
	p.Name = "管理用户列表"
	p.Description = "管理用户列表"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "dashboard.user.add"
	p.Name = "添加管理用户"
	p.Description = "添加管理用户"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "dashboard.user.update"
	p.Name = "更新管理用户"
	p.Description = "更新管理用户"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "dashboard.user.delete"
	p.Name = "删除管理用户"
	p.Description = "删除管理用户"
	permissions = append(permissions, p)

	//p = model.NewPermission()
	//p.Slug = "dashboard.user.delete"
	//p.Name = "管理用户信息"
	//p.Description = "管理用户信息"
	//permissions = append(permissions, p)

	logrus.Info(permissions)
	return permissions
}
