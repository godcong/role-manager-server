package model

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string
	Username      string
	Email         string
	MobilePhone   string
	IDCardFacade  string
	IDCardObverse string
	Association   string
	Password      string
	Token         string
	*Model
}

func NewUser() *User {
	return &User{
		Model: NewModel(),
	}
}

func (u *User) GetID() primitive.ObjectID {
	return u.ID
}

func (u *User) SetID(id primitive.ObjectID) {
	u.ID = id
}

func (u *User) _Name() string {
	return "user"
}

func (u *User) Update() error {
	return UpdateOne(u)
}

func (u *User) Delete() error {
	return DeleteByID(u)

}

func (u *User) Create() error {
	return InsertOne(u)

}

func (u *User) Find() error {
	return FindByID(u)
}
