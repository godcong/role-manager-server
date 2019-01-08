package service

import (
	"github.com/godcong/role-manager-server/model"
	"log"
)

// Migrate ...
func Migrate() {
	createRoles()
}

func createRoles() {
	models := []model.Modeler{
		model.NewGenesisRole(),
		model.NewAdminRole(),
		model.NewOrgRole(),
		model.NewMonitorRole(),
		model.NewGodRole(),
	}

	for _, m := range models {
		err := m.CreateIfNotExist()
		log.Println(err)
	}
}
