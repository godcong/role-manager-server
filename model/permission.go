package model

type Permission struct {
	*Model
	Name            string
	Slug            string
	Description     string
	PermissionModel string
}
