package service

import (
	"github.com/godcong/role-manager-server/model"
	"github.com/sirupsen/logrus"
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
	p.Slug = "dashboard.permission.add"
	p.Name = "添加权限"
	p.Description = "添加权限"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "dashboard.permission.update"
	p.Name = "更新权限"
	p.Description = "更新权限"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "dashboard.permission.delete"
	p.Name = "删除权限"
	p.Description = "删除权限"
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

	p = model.NewPermission()
	p.Slug = "dashboard.role.list"
	p.Name = "角色列表"
	p.Description = "角色列表"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "dashboard.role.add"
	p.Name = "添加角色权限"
	p.Description = "添加角色权限"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "dashboard.role.delete"
	p.Name = "删除权限"
	p.Description = "删除权限"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "dashboard.role.update"
	p.Name = "更新权限"
	p.Description = "更新权限"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "admin.organization.list"
	p.Name = "组织列表"
	p.Description = "组织列表"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "admin.organization.add"
	p.Name = "添加组织管理用户"
	p.Description = "添加组织管理用户"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "admin.organization.update"
	p.Name = "更新组织"
	p.Description = "更新组织"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "admin.organization.delete"
	p.Name = "删除组织"
	p.Description = "删除组织"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "org.media.list"
	p.Name = "视频列表"
	p.Description = "视频列表"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "org.media.add"
	p.Name = "视频添加"
	p.Description = "视频添加"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "org.media.update"
	p.Name = "视频更新"
	p.Description = "视频更新"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "org.media.update"
	p.Name = "视频更新"
	p.Description = "视频更新"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "user.media.list"
	p.Name = "用户视频列表"
	p.Description = "用户视频列表"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "user.permission.list"
	p.Name = "用户权限列表"
	p.Description = "用户权限列表"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "user.role.list"
	p.Name = "用户角色列表"
	p.Description = "用户角色列表"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "user.report.list"
	p.Name = "用户举报列表"
	p.Description = "用户举报列表"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "exorcist.user.list"
	p.Name = "用户列表"
	p.Description = "用户列表"
	permissions = append(permissions, p)

	p = model.NewPermission()
	p.Slug = "exorcist.user.update"
	p.Name = "更新用户"
	p.Description = "更新用户"
	permissions = append(permissions, p)
	//p = model.NewPermission()
	//p.Slug = "dashboard.user.delete"
	//p.Name = "管理用户信息"
	//p.Description = "管理用户信息"
	//permissions = append(permissions, p)

	logrus.Info(permissions)
	return permissions
}
