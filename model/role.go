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
)

// SlugAdmin ...
const (
	SlugAdmin = "admin"
	NameAdmin = "节点管理员"
)

// SlugOrg ...
const (
	SlugOrg = "organization"
	NameOrg = "组织管理员"
)

// SlugMonitor ...
const (
	SlugMonitor = "monitor"
	NameMonitor = "监督"
)

// SlugUser ...
const (
	SlugUser = "user"
	NameUser = "普通用户"
)

// Role ...
type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string
	Slug        string
	Description string
	Level       int
	*Model
}

// FindGenesis ...
func FindGenesis() ([]*Role, error) {
	var roles []*Role
	role := NewRole()
	e := Find(role, bson.M{
		"slug": SlugGenesis,
	}, func(cursor mongo.Cursor) error {
		role := NewRole()
		e := cursor.Decode(&role)
		if e == nil {
			roles = append(roles, role)
		}
		return nil
	})
	return roles, e
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
	return FindByID(r)
}

func (r *Role) _Name() string {
	return "role"
}

// GetID ...
func (r *Role) GetID() primitive.ObjectID {
	return r.ID
}

// RoleBySlug ...
func RoleBySlug(slug string) (*Role, error) {
	r := NewRole()
	err := FindOne(r, bson.M{
		"slug": slug,
	})
	return r, err
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
		Model: NewModel(),
	}
}
