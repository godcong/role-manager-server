package model

import (
	"log"
	"testing"
)

// TestPermission_Create ...
func TestPermission_Create(t *testing.T) {
	p := NewPermission()
	p.Slug = ".v0.user.add"
	err := p.CreateIfNotExist()
	log.Println(err)
	role := NewGenesis()
	err = role.Find()
	log.Println(*role, err)
	user := NewUser()
	user.ID = ID("5c343d3ddfbfa08c879d01a2")
	err = user.Find()
	log.Println(*user, err)
	err = RelateMaker(func() (modeler Modeler, e error) {
		//err := p.Create()
		return p, nil
	}, func() (modeler Modeler, e error) {
		return role, nil
	}, func(a, b Modeler) error {
		pr := NewPermissionRole()
		pr.SetRole(role)
		pr.SetPermission(p)
		return pr.Create()
	})
	t.Log(err)

	err = RelateMaker(func() (modeler Modeler, e error) {
		//err := p.Create()
		return p, nil
	}, func() (modeler Modeler, e error) {
		return role, nil
	}, func(a, b Modeler) error {
		pu := NewPermissionUser()
		pu.SetUser(user)
		pu.SetPermission(p)
		return pu.Create()
	})
	t.Log(err)
}
