package model

import (
	"errors"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

// RoleUser ...
type RoleUser struct {
	Model  `bson:",inline"`
	UserID primitive.ObjectID `bson:"user_id"`
	RoleID primitive.ObjectID `bson:"role_id"`
	user   *User
	role   *Role
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
	if r.ID != primitive.NilObjectID {
		return IDExist(r)
	}
	return IsExist(r, bson.M{
		"role_id": r.RoleID,
		"user_id": r.UserID,
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
	} else if r.UserID != primitive.NilObjectID {
		return FindOne(r, bson.M{
			"user_id": r.UserID,
		})
	} else if r.RoleID != primitive.NilObjectID {
		return errors.New("user id could't null")
	}
	return nil
}

// NewRoleUser ...
func NewRoleUser() *RoleUser {
	return &RoleUser{
		Model: model(),
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
		return role, nil
	}
	return nil, errors.New("role not found")
}
