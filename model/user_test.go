package model

import (
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"testing"
)

var user = User{
	Name:          "godcong",
	Username:      "ungodcong",
	Email:         "godcong@ggg.com",
	MobilePhone:   "123456",
	IDCardFacade:  "/d/d/e/e/d/c/",
	IDCardObverse: "/f/g/h/j/a",
	Association:   "yelion",
	Password:      "godcong0910",
	Token:         "1212133333",
}

func TestUser_Create(t *testing.T) {
	t.Log(user.Create())
	t.Log(user)
}

func TestUser_Delete(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("5c2dd3c5819e895f7c1af1d4")
	user := User{
		ID: id,
	}
	//t.Log(user.Delete())
	e := user.FindByID("5c2dd3c5819e895f7c1af1d4")
	t.Log(e)
	t.Log(user)
}
