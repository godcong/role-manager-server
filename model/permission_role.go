package model

import "github.com/mongodb/mongo-go-driver/bson/primitive"

// PermissionRole ...
type PermissionRole struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	PermissionID primitive.ObjectID
	RoleID       primitive.ObjectID
	permission   *Permission
	role         *Role
	*Model
}

// Role ...
func (r *PermissionRole) Role() *Role {
	return r.role
}

// SetRole ...
func (r *PermissionRole) SetRole(role *Role) {
	r.role = role
	r.RoleID = role.ID
}

// Permission ...
func (r *PermissionRole) Permission() *Permission {
	return r.permission
}

// SetPermission ...
func (r *PermissionRole) SetPermission(permission *Permission) {
	r.permission = permission
	r.PermissionID = permission.ID
}

// GetID ...
func (r *PermissionRole) GetID() primitive.ObjectID {
	return r.ID
}

// SetID ...
func (r *PermissionRole) SetID(id primitive.ObjectID) {
	r.ID = id
}

// Create ...
func (r *PermissionRole) Create() error {
	return InsertOne(r)
}

// Update ...
func (r *PermissionRole) Update() error {
	return UpdateOne(r)
}

// Delete ...
func (r *PermissionRole) Delete() error {
	return DeleteByID(r)
}

// Find ...
func (r *PermissionRole) Find() error {
	return FindByID(r)
}

// NewPermissionRole ...
func NewPermissionRole() *PermissionRole {
	return &PermissionRole{
		Model: NewModel(),
	}
}

func (r *PermissionRole) _Name() string {
	return "user_role"
}
