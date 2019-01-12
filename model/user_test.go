package model

import (
	"github.com/godcong/role-manager-server/util"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"testing"
	"time"
)

// TestUser_Create ...
func TestUser_Create(t *testing.T) {
	user := NewUser()
	user.Username = "godcong"
	user.Name = util.GenerateRandomString(32)
	user.SetPassword("DBD978CCDBBE8B6DE77F6B37B5DF9B5B62A7E892A501C3B53EAA16B0838BD5ED")
	t.Log(user.Create())
	t.Log(user)
}

// TestUser_Delete ...
func TestUser_Delete(t *testing.T) {
	id, _ := primitive.ObjectIDFromHex("5c2eeb5bb69c469e69c79a26")
	user := User{
		Model: Model{
			ID:         id,
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
			DeletedAt:  nil,
			Version:    0,
			softDelete: false,
		},
		Name:          "",
		Username:      "",
		Email:         "",
		Mobile:        "",
		IDCardFacade:  "",
		IDCardObverse: "",
		Password:      "",
		Token:         "",
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

// TestUser_Update ...
func TestUser_Update(t *testing.T) {
	user := User{
		Model: Model{
			ID:         primitive.ObjectID{},
			CreatedAt:  time.Time{},
			UpdatedAt:  time.Time{},
			DeletedAt:  nil,
			Version:    0,
			softDelete: false,
		},
		Name:          "",
		Username:      "",
		Email:         "",
		Mobile:        "",
		IDCardFacade:  "",
		IDCardObverse: "",
		Organization:  "",
		Password:      "",
		Token:         "",
	}
	user.softDelete = true
	user.Find()

	user.Username = "SSSSSSSSSSSSSSSSSSSS"
	err := user.Update()
	t.Log(err)
	t.Log(user)
}

// TestUser_Find ...
func TestUser_Find(t *testing.T) {
	user := NewUser()
	user.ID = ID("5c2eeb95761de4f5a13b3b83")
	e := user.Find()
	t.Log(user, e)
}

// TestRole_Create ...
func TestRole_Create(t *testing.T) {
	g := NewGenesisRole()
	e := g.Create()
	t.Log(g, e)
}

// TestRoleUser_Find ...
func TestRoleUser_Find(t *testing.T) {
	ru := NewRoleUser()
	ru.UserID = ID("5c2eeb95761de4f5a13b3b83")
	ru.RoleID = ID("5c2f2864451279e9ff6f2128")
	e := ru.Find()
	t.Log(e, ru)
	t.Log(ru.User())

}

// TestRoleUser_Create ...
func TestRoleUser_Create(t *testing.T) {
	ru := NewRoleUser()
	ru.RoleID = ID("5c344d138efb3aefa1341d0c")
	ru.UserID = ID("5c343d3ddfbfa08c879d01a2")
	e := ru.CreateIfNotExist()
	t.Log(ru, e)
}
