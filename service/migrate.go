package service

import (
	"github.com/godcong/role-manager-server/model"
	"github.com/mongodb/mongo-go-driver/bson"
	log "github.com/sirupsen/logrus"
)

// Migrate ...
func Migrate() {
	createRoles()
	createGenesis()
}

func createRoles() {
	models := []model.Modeler{
		model.NewGenesisRole(),
		model.NewAdminRole(),
		model.NewOrgRole(),
		model.NewMonitorRole(),
		model.NewGodRole(),
		model.NewMenu(),
	}

	for _, m := range models {
		err := m.CreateIfNotExist()
		log.Println(err)
	}
}

func createGenesis() {
	role := model.NewGenesisRole()
	err := model.FindOne(role, bson.M{
		"slug": role.Slug,
	})
	if err != nil {
		panic(err)
	}
	ru := model.NewRoleUser()
	if model.IsExist(ru, bson.M{
		"role_id": role.ID,
	}) {
		panic("role exist")
		return
	}

	passwd := "123456"
	user := model.NewUser()
	user.Name = "genesis"
	user.SetPassword(PWD(passwd))
	err = makeUser(user, role)
	if err != nil {
		panic(err)
	}
}
