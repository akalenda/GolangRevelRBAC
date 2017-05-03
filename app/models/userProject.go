package models

import (
	"github.com/go-gorp/gorp"
	"github.com/akalenda/GolangRevelRBAC/app/helpers"
)

type UserProject struct {
	UserProjectId int
	OwnerUsername string `db:",size:20"`
	Name          string `db:",size:32"`
	Description   string `db:",size:256"`
}

func UserProject_InitTable(dbmap *gorp.DbMap) {
	dbmap.AddTable(UserProject{}).SetKeys(true, "UserProjectId")
}

func UserProject_InitAdminExample(dbmap *gorp.DbMap) {
	var query []UserProject
	dbmap.Select(&query, "select * from UserProject where OwnerUsername=admin")
	if len(query) == 0 {
		exampleProject := UserProject{
			OwnerUsername: "admin",
			Name: "Example project",
			Description: "The database is initialized with this project",
		}
		helpers.CheckErr(dbmap.Insert(&exampleProject))
	}
}
