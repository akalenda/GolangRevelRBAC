package models

import (
	"fmt"
	"github.com/revel/revel"
	"regexp"
	"github.com/go-gorp/gorp"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

type User struct {
	UserId        	int
	Name          	string `db:",size:128"`
	Username      	string `db:",size:32"`
	Roles         	string `db:",size:256"`
	HashedPassword	string
}

func (u *User) MatchesHashedPasswordTo(unhashedPassword string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.HashedPassword), []byte(unhashedPassword)) == nil
}

func User_AddTable(dbmap *gorp.DbMap) {
	dbmap.AddTable(User{}).SetKeys(true, "UserId")
}

/**
NewUser attempts to construct a new User struct. It makes the assumption that the User is already in the system.
It will return an error if there was a problem pulling from the database, if no users match the username, or if for
some bizarre reason more than one User with that username was found.
 */
func GetUserFromDB(transaction *gorp.Transaction, username string) (*User, error) {
	var users []User
	_, errDB := transaction.Select(&users, "select * from User where Username=?", username)

	if errDB != nil {
		return nil, errDB
	} else if len(users) == 0 {
		return nil, errors.New("User error: Username not found")
	} else if len(users) > 1 {
		return &users[0], errors.New("Critical error: Multiple users found with same username")
	} else if &users[0] == nil {
		return nil, errors.New("Error: WAT")
	}
	return &users[0], nil
}

/**
RegisterNewUser creates a new User struct. It first ensures that a User with the given username does not already exist.
It inserts the new User's information into the database, and validates the information they've given in the sense of
its formatting. The password is not stored in the database, rather an encrypted version.
 */
func RegisterNewUser(transaction *gorp.Transaction, username string, password string, fullname string) (*User, error) {
	var users []User
	_, errDB := transaction.Select(&users, "select * from User where Username=?", username)
	hashedPassword, errHP := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if errDB != nil {
		return nil, errDB
	} else if len(users) > 0 {
		return &users[0], errors.New("User error: Username already claimed")
	} else if errHP != nil {
		return nil, errDB
	}

	u := User{Name: fullname, Username: username, Roles: "", HashedPassword: string(hashedPassword)}
	errDB = transaction.Insert(&u)

	if errDB != nil {
		return &u, errDB
	}
	return &u, nil
}

func (u *User) String() string {
	return fmt.Sprintf("User(%s)", u.Username)
}

var userRegex = regexp.MustCompile("^\\w*$")

func (u *User) Validate(v *revel.Validation) {
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
It reads like this: "A function on u, which is a User,
is a method named addRole that receives role, which is
a string, and returns a User."
 */
func (u *User) AddRole(role string) (*User, error) {
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

func (u *User) commitChangesToDB() {
	// TODO
}
