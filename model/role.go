package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

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

// Role ...
type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string
	Slug        string
	Description string
	Level       int
	*Model
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

// NewRole ...
func NewRole() *Role {
	return &Role{
		Model: NewModel(),
	}
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
