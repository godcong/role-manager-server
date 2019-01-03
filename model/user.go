package model

import (
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
)

type User struct {
	*Model
	Name          string
	Username      string
	Email         string
	MobilePhone   string
	IDCardFacade  string
	IDCardObverse string
	Association   string
	Password      string
	Token         string
}

func (u *User) Update() error {
	panic("implement me")
}

func (u *User) Delete() error {
	panic("implement me")
}

func (u *User) Find() error {
	panic("implement me")
}

func (u *User) Collection() *mongo.Collection {
	return Collection("user")
}

func (u *User) Create() error {
	one, err := InsertOne("user", u)
	log.Println(one)
	if err != nil {
		return err
	}
	return nil
}
