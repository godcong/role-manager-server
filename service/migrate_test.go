package service

import (
	"github.com/godcong/role-manager-server/model"
	"testing"
)

// TestMigrate ...
func TestMigrate(t *testing.T) {
	Migrate()
	model.NewRoleUser()
}
