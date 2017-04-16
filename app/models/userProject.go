package models

import "github.com/go-gorp/gorp"

type UserProject struct {
	UserProjectID int
	OwnerUsername string `db:",size:20"`
	Name          string `db:",size:32"`
	Description   string `db:",size:128"`
}

func UserProject_AddTable(dbmap *gorp.DbMap) {
	dbmap.AddTable(UserProject{}).SetKeys(true, "UserProjectId")
}


