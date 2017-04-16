package models

import (
	"gopkg.in/mikespook/gorbac.v2"
)

var urlsToRolesThatCanAccess = map[string][]string{
	"/welcome"       : {"CONSUMER"},
	"/admin"         : {"ADMIN"},
	"/createProject" : {"PROJECT_MANAGER"},
	"/debug"         : {"DEVELOPER"},
}

var maxLengthOfRoleList = 0
var idToRole = produceSetOfRoles()
var roleToId = reverseMap(idToRole)

/* *********************** Public **************************************** */

// As roles can be listed as, for example, "USER|ADMIN|PROJECT_MANAGER", this gives what the longest possible such list
// would be based on the roles available. In particular it's used when defining the width of a column in the database.
func Role_getMaxLengthOfRoleList() int {
	return maxLengthOfRoleList
}

func Role_prepareRbac(rbac *gorbac.RBAC){
	for url, listOfRoles := range urlsToRolesThatCanAccess {
		permission := gorbac.NewStdPermission(url)
		for i := range listOfRoles {
			role := gorbac.NewStdRole(listOfRoles[i])
			role.Assign(permission)
		}
	}
}

func Role_isKnown(role string) bool {
	// We guarantee that no valid roles will fall on index zero by priming the list with an empty string
	if len(role) == 0{
		return false
	}
	return roleToId[role] != 0
}

/* ********************* Helpers *********************************/

// This scans the urlsToRolesThatCanAccess mapping and builds a handy set of possible roles.
// As a side effect, this also sets maxLengthOfRoleList.
func produceSetOfRoles() []string {
	mapOfRoles := make(map[string]bool)
	for _, roles := range urlsToRolesThatCanAccess {
		for _, role := range roles {
			mapOfRoles[role] = true
		}
	}
	setOfRoles := []string{""}
	for role := range mapOfRoles {
		setOfRoles = append(setOfRoles, role)
		maxLengthOfRoleList += len(role) + 1 // The +1 is for the "|" delimiter
	}
	maxLengthOfRoleList += 3 // And a bit of wiggle room.
	return setOfRoles
}

func reverseMap(m []string) map[string]int {
	rm := make(map[string]int)
	for index, element := range m {
		rm[element] = index
	}
	return rm
}
