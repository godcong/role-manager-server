package model

import (
	"errors"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// PermissionRole ...
type PermissionRole struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	PermissionID primitive.ObjectID
	RoleID       primitive.ObjectID
	permission   *Permission
	role         *Role
	*Model
}

// CreateIfNotExist ...
func (r *PermissionRole) CreateIfNotExist() error {
	return CreateIfNotExist(r)
}

// IsExist ...
func (r *PermissionRole) IsExist() bool {
	return IsExist(r, bson.M{
		"permissionid": r.PermissionID,
		"roleid":       r.RoleID,
	})
}

// Role ...
func (r *PermissionRole) Role() (*Role, error) {
	if r.ID == primitive.NilObjectID {
		return nil, errors.New("id is null")
	}
	if r.RoleID != primitive.NilObjectID {
		role := NewRole()
		role.ID = r.RoleID
		err := role.Find()
		if err != nil {
			return nil, err
		}
		r.role = role
		return role, nil
	}
	return nil, errors.New("role not found")
}

// SetRole ...
func (r *PermissionRole) SetRole(role *Role) {
	r.role = role
	r.RoleID = role.ID
}

// Permission ...
func (r *PermissionRole) Permission() (*Permission, error) {
	if r.ID == primitive.NilObjectID {
		return nil, errors.New("id is null")
	}
	if r.PermissionID != primitive.NilObjectID {
		per := NewPermission()
		per.ID = r.PermissionID
		err := per.Find()
		if err != nil {
			return nil, err
		}
		r.permission = per
		return per, nil
	}
	return nil, errors.New("permission not found")
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
