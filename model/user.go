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

func (u *User) _Name() string {
	return "user"
}

func (u *User) Update() error {
	_, err := UpdateOne(C(u._Name()), u.ID, u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	one, err := DeleteByID(C(u._Name()), u.ID)
	log.Println(one)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Create() error {
	u.BeforeInsert()
	_, err := InsertOne(C(u._Name()), "user", u)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindByID(id string) error {
	return FindByID(C(u._Name()), id, u)
}
