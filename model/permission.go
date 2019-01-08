package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

// Permission ...
type Permission struct {
	Model           `bson:",inline"`
	Name            string `bson:"name"`
	Slug            string `bson:"slug"`
	Description     string `bson:"description"`
	PermissionModel string `bson:"permission_model"`
}

// CreateIfNotExist ...
func (p *Permission) CreateIfNotExist() error {
	return CreateIfNotExist(p)
}

// NewPermission ...
func NewPermission() *Permission {
	return &Permission{
		Model: model(),
	}
}

// GetID ...
func (p *Permission) GetID() primitive.ObjectID {
	return p.ID
}

// SetID ...
func (p *Permission) SetID(id primitive.ObjectID) {
	p.ID = id
}

// Create ...
func (p *Permission) Create() error {
	return InsertOne(p)
}

// Update ...
func (p *Permission) Update() error {
	return UpdateOne(p)
}

// Delete ...
func (p *Permission) Delete() error {
	return DeleteByID(p)
}

// Find ...
func (p *Permission) Find() error {
	return FindByID(p)
}

func (p *Permission) _Name() string {
	return "permission"
}

// Roles ...
func (p *Permission) Roles() ([]*Role, error) {
	var roles []*Role
	pu := NewPermissionRole()
	err := Find(pu, bson.M{
		"role_id": pu.RoleID,
	}, func(cursor mongo.Cursor) error {
		pu := NewPermissionRole()
		err := cursor.Decode(pu)
		if err != nil {
			return err
		}
		role, err := pu.Role()
		if err != nil {
			return err
		}
		roles = append(roles, role)
		return nil
	})
	return roles, err
}

// Users ...
func (p *Permission) Users() ([]*User, error) {
	var users []*User
	pu := NewPermissionUser()
	err := Find(pu, bson.M{
		"user_id": pu.UserID,
	}, func(cursor mongo.Cursor) error {
		pu := NewPermissionUser()
		err := cursor.Decode(pu)
		if err != nil {
			return err
		}
		user, err := pu.User()
		if err != nil {
			return err
		}
		users = append(users, user)
		return nil
	})
	return users, err
}
