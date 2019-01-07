package model

import (
	"errors"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// RoleUser ...
type RoleUser struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	UserID primitive.ObjectID
	RoleID primitive.ObjectID
	user   *User
	role   *Role
	*Model
}

// SetRole ...
func (r *RoleUser) SetRole(role *Role) {
	r.role = role
	r.RoleID = role.ID
}

// SetUser ...
func (r *RoleUser) SetUser(user *User) {
	r.user = user
	r.UserID = user.ID
}

// GetID ...
func (r *RoleUser) GetID() primitive.ObjectID {
	return r.ID
}

// SetID ...
func (r *RoleUser) SetID(id primitive.ObjectID) {
	r.ID = id
}

// Create ...
func (r *RoleUser) Create() error {
	return InsertOne(r)
}

// CreateIfNotExist ...
func (r *RoleUser) CreateIfNotExist() error {
	return CreateIfNotExist(r)
}

// IsExist ...
func (r *RoleUser) IsExist() bool {
	return IsExist(r, bson.M{
		"roleid": r.RoleID,
		"userid": r.UserID,
	})
}

// Update ...
func (r *RoleUser) Update() error {
	return UpdateOne(r)
}

// Delete ...
func (r *RoleUser) Delete() error {
	return DeleteByID(r)
}

// Find ...
func (r *RoleUser) Find() error {
	if r.ID != primitive.NilObjectID {
		return FindByID(r)
	} else if r.RoleID != primitive.NilObjectID && r.UserID != primitive.NilObjectID {
		return FindOne(r, bson.M{
			"roleid": r.RoleID,
			"userid": r.UserID,
		})
	}
	return nil
}

// NewRoleUser ...
func NewRoleUser() *RoleUser {
	return &RoleUser{
		Model: NewModel(),
	}
}

func (r *RoleUser) _Name() string {
	return "role_user"
}

// User ...
func (r *RoleUser) User() (*User, error) {
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

// Role ...
func (r *RoleUser) Role() (*Role, error) {

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
		r.RoleID = role.ID
		return role, nil
	}
	return nil, errors.New("role not found")
}
