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
	id, _ := primitive.ObjectIDFromHex("5c2eeb5bb69c469e69c79a26")
	user := User{
		ID: id,
	}
	//t.Log(user.Delete())
	user.SetSoftDelete(true)
	e := user.Delete()
	t.Log(e)
	t.Log(user)

	e = user.Find()
	t.Log(e)
	t.Log(user)

}

func TestUser_Update(t *testing.T) {
	user := User{
		ID: ID("5c2eea9a3db6598a9c25c65c"),
	}
	user.softDelete = true
	user.Find()

	user.Username = "SSSSSSSSSSSSSSSSSSSS"
	err := user.Update()
	t.Log(err)
	t.Log(user)
}
