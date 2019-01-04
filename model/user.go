package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"log"
	"strings"
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
	Association   string
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
	u.Password = PWD(pwd)
}

// ValidatePassword ...
func (u *User) ValidatePassword(pwd string) bool {
	return u.Password == PWD(pwd)
}

// PWD ...
func PWD(pwd string) string {
	salt := []byte("22c77682334a55f41b6cdbdf5ca27a830a4241a0e13c101f6dc5bd2dde86e3b7")
	m := hmac.New(sha256.New, []byte("3htASeUkrx5LcnuTENQIZQPCVlwdnvIJ7bYtSpoJYq38MgUJnx1CQIR1gjZ8HJxAEcN4gqugBg"))
	m.Write([]byte(pwd))
	m.Write(salt)
	return strings.ToUpper(fmt.Sprintf("%x", m.Sum(nil)))
}

// Roles ...
func (u *User) Roles() ([]*Role, error) {
	var list []*RoleUser

	ru := NewRoleUser()
	err := Find(ru, bson.M{
		"userid": u.ID,
	}, &list)
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
