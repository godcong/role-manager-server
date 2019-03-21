package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
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
		return IDExist(r)
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

// All ...
func (r *Role) All() ([]*Role, error) {
	var roles []*Role
	m := bson.M{}
	err := Find(r, m, func(cursor mongo.Cursor) error {
		log.Println(cursor.DecodeBytes())
		var r Role
		err := cursor.Decode(&r)
		if err != nil {
			return err
		}
		roles = append(roles, &r)
		return nil
	})
	return roles, err
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

// NewGenesisRole ...
func NewGenesisRole() *Role {
	role := NewRole()
	role.Slug = SlugGenesis
	role.Name = NameGenesis
	role.Description = NameGenesis
	return role
}

// NewAdminRole ...
func NewAdminRole() *Role {
	role := NewRole()
	role.Slug = SlugAdmin
	role.Name = NameAdmin
	role.Description = NameAdmin
	return role
}

// NewOrgRole ...
func NewOrgRole() *Role {
	role := NewRole()
	role.Slug = SlugOrg
	role.Name = NameOrg
	role.Description = NameOrg
	return role
}

// NewMonitorRole ...
func NewMonitorRole() *Role {
	role := NewRole()
	role.Slug = SlugMonitor
	role.Name = NameMonitor
	role.Description = NameMonitor
	return role
}

// NewGodRole 用户就是上帝
func NewGodRole() *Role {
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
