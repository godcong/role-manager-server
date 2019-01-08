package service

import (
	"github.com/godcong/role-manager-server/model"
)

// Migrate ...
func Migrate() {
	createRoles()
}

func createRoles() {
	models := []model.Modeler{
		model.NewGenesis(),
		model.NewAdmin(),
		model.NewOrg(),
		model.NewMonitor(),
		model.NewGod(),
	}

	for _, m := range models {

		m.Create()
	}

}
