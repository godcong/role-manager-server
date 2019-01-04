package model

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
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

// NewUser ...
func NewUser() *User {
	return &User{
		Model: NewModel(),
	}
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
