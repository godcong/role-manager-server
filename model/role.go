package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// SlugGenesis ...
const (
	SlugGenesis = "genesis"
	NameGenesis = "超级管理员"
	SlugAdmin   = "admin"
	NameAdmin   = "节点管理员"
	SlugOrg     = "organization"
	NameOrg     = "组织管理员"
	SlugMonitor = "monitor"
	NameMonitor = "监督"
	SlugUser    = "user"
	NameUser    = "普通用户"
)

// Role ...
type Role struct {
	Model       `bson:",inline"`
	Name        string `bson:"name"`
	Slug        string `bson:"slug"`
	Description string `bson:"description"`
	Level       int    `bson:"level"`
}

// IsExist ...
func (r *Role) IsExist() bool {
	if r.ID != primitive.NilObjectID {
		return IsExist(r, bson.M{
			"_id": r.ID,
		})
	}
	return IsExist(r, bson.M{
		"slug": r.Slug,
	})
}

// CreateIfNotExist ...
func (r *Role) CreateIfNotExist() error {
	return CreateIfNotExist(r)
}

// SetID ...
func (r *Role) SetID(id primitive.ObjectID) {
	r.ID = id
}

// Create ...
func (r *Role) Create() error {
	return InsertOne(r)
}

// Update ...
func (r *Role) Update() error {
	return UpdateOne(r)
}

// Delete ...
func (r *Role) Delete() error {
	return DeleteByID(r)
}

// Find ...
func (r *Role) Find() error {

	if r.ID != primitive.NilObjectID {
		return FindByID(r)
	}
	return FindOne(r, bson.M{
		"slug": r.Slug,
	})
}

func (r *Role) _Name() string {
	return "role"
}

// GetID ...
func (r *Role) GetID() primitive.ObjectID {
	return r.ID
}

// Users ...
func (r *Role) Users() ([]*User, error) {
	var users []*User
	ru := NewRoleUser()
	err := Find(ru, bson.M{
		"role_id": r.ID,
	}, func(cursor mongo.Cursor) error {
		ru := NewRoleUser()
		err := cursor.Decode(ru)
		if err != nil {
			return err
		}
		user, err := ru.User()
		if err != nil {
			return err
		}
		users = append(users, user)
		return nil
	})
	return users, err
}

// Permissions ...
func (r *Role) Permissions() ([]*Permission, error) {
	var ps []*Permission
	pr := NewPermissionRole()
	err := Find(pr, bson.M{
		"role_id": r.ID,
	}, func(cursor mongo.Cursor) error {
		ru := NewPermissionRole()
		err := cursor.Decode(ru)
		if err != nil {
			return err
		}
		p, err := ru.Permission()
		if err != nil {
			return err
		}
		ps = append(ps, p)
		return nil
	})
	return ps, err
}

// CheckPermission ...
func (r *Role) CheckPermission(permission *Permission) error {
	pr := NewPermissionRole()
	err := FindOne(pr, bson.M{
		"role_id":       r.ID,
		"permission_id": permission.ID,
	})
	if err != nil {
		return err
	}
	return nil
}

// NewGenesis ...
func NewGenesis() *Role {
	role := NewRole()
	role.Slug = SlugGenesis
	role.Name = NameGenesis
	role.Description = NameGenesis
	return role
}

// NewAdmin ...
func NewAdmin() *Role {
	role := NewRole()
	role.Slug = SlugAdmin
	role.Name = NameAdmin
	role.Description = NameAdmin
	return role
}

// NewOrg ...
func NewOrg() *Role {
	role := NewRole()
	role.Slug = SlugOrg
	role.Name = NameOrg
	role.Description = NameOrg
	return role
}

// NewMonitor ...
func NewMonitor() *Role {
	role := NewRole()
	role.Slug = SlugMonitor
	role.Name = NameMonitor
	role.Description = NameMonitor
	return role
}

// NewGod 用户就是上帝
func NewGod() *Role {
	role := NewRole()
	role.Slug = SlugUser
	role.Name = NameUser
	role.Description = NameUser
	return role
}

// NewRole ...
func NewRole() *Role {
	return &Role{
		Model: model(),
	}
}
