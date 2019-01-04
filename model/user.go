package model

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"log"
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
	Model
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
	_, err := UpdateOne(u, u.ID, u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	one, err := DeleteByID(u, u.ID)
	log.Println(one)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Create() error {
	_, err := InsertOne(u, u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindByID(id string) error {
	return FindByID(u, id, u)
}
