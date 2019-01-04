package model

import (
	"errors"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// PermissionUser ...
type PermissionUser struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	PermissionID primitive.ObjectID
	UserID       primitive.ObjectID
	permission   *Permission
	user         *User
	*Model
}

// CreateIfNotExist ...
func (r *PermissionUser) CreateIfNotExist() error {
	return CreateIfNotExist(r)
}

// IsExist ...
func (r *PermissionUser) IsExist() bool {
	return IsExist(r, bson.M{
		"permissionid": r.PermissionID,
		"userid":       r.UserID,
	})
}

// User ...
func (r *PermissionUser) User() (*User, error) {
	if r.ID == primitive.NilObjectID {
		return nil, errors.New("id is null")
	}
	if r.UserID != primitive.NilObjectID {
		user := NewUser()
		user.ID = r.UserID
		err := user.Find()
		if err != nil {
			return nil, err
		}
		r.user = user
		return user, nil
	}
	return nil, errors.New("user not found")
}

// SetUser ...
func (r *PermissionUser) SetUser(user *User) {
	r.user = user
	r.UserID = user.ID
}

// Permission ...
func (r *PermissionUser) Permission() (*Permission, error) {
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
func (r *PermissionUser) SetPermission(permission *Permission) {
	r.permission = permission
	r.PermissionID = permission.ID
}

// GetID ...
func (r *PermissionUser) GetID() primitive.ObjectID {
	return r.ID
}

// SetID ...
func (r *PermissionUser) SetID(id primitive.ObjectID) {
	r.ID = id
}

// Create ...
func (r *PermissionUser) Create() error {
	return InsertOne(r)
}

// Update ...
func (r *PermissionUser) Update() error {
	return UpdateOne(r)
}

// Delete ...
func (r *PermissionUser) Delete() error {
	return DeleteByID(r)
}

// Find ...
func (r *PermissionUser) Find() error {
	return FindByID(r)
}

// NewPermissionUser ...
func NewPermissionUser() *PermissionUser {
	return &PermissionUser{
		Model: NewModel(),
	}
}

func (r *PermissionUser) _Name() string {
	return "user_role"
}
