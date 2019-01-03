package model

import "testing"

func TestUser_Create(t *testing.T) {
	user := User{
		Model:         nil,
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
	t.Log(user.Create())
}
