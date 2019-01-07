package model

import (
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

// User ...
type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string
	Username      string
	Email         string
	Mobile        string
	IDCardFacade  string
	IDCardObverse string
	Organization  string
	Password      string
	Token         string
	*Model
}

// CreateIfNotExist ...
func (u *User) CreateIfNotExist() error {
	return CreateIfNotExist(u)
}

// GetID ...
func (u *User) GetID() primitive.ObjectID {
	return u.ID
}

// SetID ...
func (u *User) SetID(id primitive.ObjectID) {
	u.ID = id
}

func (u *User) _Name() string {
	return "user"
}

// Update ...
func (u *User) Update() error {
	return UpdateOne(u)
}

// Delete ...
func (u *User) Delete() error {
	return DeleteByID(u)

}

// Create ...
func (u *User) Create() error {
	return InsertOne(u)

}

// Find ...
func (u *User) Find() error {
	return FindByID(u)
}

// SetPassword ...
func (u *User) SetPassword(pwd string) {
	u.Password = pwd
}

// ValidatePassword ...
func (u *User) ValidatePassword(pwd string) bool {
	return u.Password == pwd
}

// Role ...
func (u *User) Role() (*Role, error) {
	ru := NewRoleUser()
	err := FindOne(ru, bson.M{
		"userid": u.ID,
	})
	log.Println(*ru)
	if err != nil {
		return nil, err
	}
	role := NewRole()
	role.ID = ru.RoleID
	err = role.Find()
	log.Println(role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

// Roles ...
func (u *User) Roles() ([]*Role, error) {
	var list []*RoleUser

	ru := NewRoleUser()
	err := Find(ru, bson.M{
		"userid": u.ID,
	}, func(cursor mongo.Cursor) error {
		ru := NewRoleUser()
		e := cursor.Decode(&ru)
		if e == nil {
			list = append(list, ru)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	var roles []*Role
	for _, val := range list {
		role := NewRole()
		role.ID = val.RoleID
		err := role.Find()
		if err != nil {
			log.Println(err)
			continue
		}
		roles = append(roles, role)
	}
	return roles, nil
}

// NewUser ...
func NewUser() *User {
	return &User{
		Model: NewModel(),
	}
}
