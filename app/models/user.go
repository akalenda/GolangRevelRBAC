package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
	"github.com/go-gorp/gorp"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type user struct {
	UserId        	int
	Name          	string `db:",size:128"`
	Username      	string `db:",size:32"`
	Roles         	string `db:",size:256"`
	HashedPassword	string
}

func User_AddTable(dbmap *gorp.DbMap) {
	dbmap.AddTable(user{}).SetKeys(true, "UserId")
}

/**
NewUser attempts to construct a new user struct. It makes the assumption that the user is already in the system.
It will return an error if there was a problem pulling from the database, if no users match the username, or if for
some bizarre reason more than one user with that username was found.
 */
func NewRegisteredUser(dbmap *gorp.DbMap, username string, password string) (*user, error) {
	var users []user
	_, errDB := dbmap.Select(&users, "select * from user where Username=?", username)
	hashedPassword, errHP := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if errDB != nil {
		return nil, errDB
	} else if len(users) == 0 {
		return nil, errors.New("User error: Username not found")
	} else if len(users) > 1 {
		return &users[0], errors.New("Critical error: Multiple users found with same username")
	} else if errHP != nil {
		return &users[0], errDB
	} else if string(hashedPassword) == users[0].HashedPassword {
		return &users[0], errors.New("User error: Username and password do not match")
	}
	return &users[0], nil
}

/**
RegisterNewUser creates a new user struct. It first ensures that a user with the given username does not already exist.
It inserts the new user's information into the database, and validates the information they've given in the sense of
its formatting. The password is not stored in the database, rather an encrypted version.
 */
func RegisterNewUser(dbmap *gorp.DbMap, username string, password string, fullname string) (*user, error) {
	var users []user
	_, errDB := dbmap.Select(&users, "select * from user where Username=?", username)
	hashedPassword, errHP := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if errDB != nil {
		return nil, errDB
	} else if len(users) > 0 {
		return &users[0], errors.New("User error: Username already claimed")
	} else if errHP != nil {
		return nil, errDB
	}

	u := user{Name: fullname, Username: username, Roles: "", HashedPassword: string(hashedPassword)}
	errDB = dbmap.Insert(&u)

	if errDB != nil {
		return &u, errDB
	}
	return &u, nil
}

func (u *user) String() string {
	return fmt.Sprintf("user(%s)", u.Username)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (u *user) Validate(v *revel.Validation) {
	v.Check(u.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{4},
		revel.Match{userRegex},
	)

	v.Check(u.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{Max: 15},
		revel.MinSize{Min: 5},
	)
}

/**
AddRole is a method. My very first Golang method!
It reads like this: "A function on u, which is a user,
is a method named addRole that receives role, which is
a string, and returns a user."
 */
func (u *user) AddRole(role string) (*user, error) {
	if Role_isKnown(role) {
		if len(u.Roles) == 0 {
			u.Roles = role
		} else {
			u.Roles += "|" + role
		}
		return u, nil
	}
	return u, errors.New("Internal design error: '" + role + "' is not a known role.")
}

func (u *user) commitChangesToDB() {
	// TODO
}
