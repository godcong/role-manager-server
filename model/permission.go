package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	log "github.com/sirupsen/logrus"
)

// Permission ...
type Permission struct {
	Model           `bson:",inline"`
	Name            string `bson:"name"`
	Slug            string `bson:"slug"`
	Description     string `bson:"description"`
	PermissionModel string `bson:"permission_model"`
}

// IsExist ...
func (p *Permission) IsExist() bool {
	if p.ID != primitive.NilObjectID {
		return IDExist(p)
	}
	return IsExist(p, bson.M{
		"slug": p.Slug,
	})
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
	if p.ID != primitive.NilObjectID {
		return FindByID(p)
	}
	return FindOne(p, bson.M{
		"slug": p.Slug,
	})
}

// ALL ...
func (p *Permission) ALL() ([]*Permission, error) {
	var permissions []*Permission
	m := bson.M{}
	err := Find(p, m, func(cursor mongo.Cursor) error {
		log.Println(cursor.DecodeBytes())
		var p Permission
		err := cursor.Decode(&p)
		if err != nil {
			return err
		}
		permissions = append(permissions, &p)
		return nil
	})
	return permissions, err
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
