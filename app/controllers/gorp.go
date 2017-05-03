package controllers

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	"github.com/revel/modules/db/app"
	"github.com/revel/revel"
	"github.com/akalenda/GolangRevelRBAC/app/models"
	"golang.org/x/crypto/bcrypt"
	"github.com/akalenda/GolangRevelRBAC/app/helpers"
)

var dbMap *gorp.DbMap

func InitDB() {
	db.Init()
	dbMap = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	models.User_AddTable(dbMap)
	models.UserProject_InitTable(dbMap)
	dbMap.AddTable(models.Hotel{}  ).SetKeys(true, "HotelId")
	dbMap.AddTable(models.Booking{}).SetKeys(true, "BookingId")

	dbMap.TraceOn("[gorp]", revel.INFO)
	helpers.CheckErr(dbMap.CreateTablesIfNotExists())

	initAdminAccount(dbMap)
	initHotels(dbMap)
	models.UserProject_InitAdminExample(dbMap)
}

func initAdminAccount(dbmap *gorp.DbMap) {
	var users []models.User
	dbmap.Select(&users, "select * from User where Username=admin")
	if len(users) == 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("admin"), bcrypt.DefaultCost)
		u := models.User{Name: "Administrator", Username: "admin", Roles: "ADMIN", HashedPassword: string(hashedPassword)}
		err = dbmap.Insert(&u)
		if err != nil && err.Error() != "User error: Username and password do not match" {
			panic(err)
		}
	}
}

func initHotels(dbmap *gorp.DbMap) {
	hotels := []*models.Hotel{
		{HotelId: 0, Name: "Marriott Courtyard", Address: "Tower Pl, Buckhead", City: "Atlanta", State: "GA", Zip:"30305", Country:"USA", Price:120},
		{HotelId: 0, Name: "W Hotel", Address: "Union Square, Manhattan", City: "New York",      State: "NY", Zip:"10011", Country:"USA", Price:450},
		{HotelId: 0, Name: "Hotel Rouge", Address: "1315 16th St NW", City: "Washington",        State: "DC", Zip:"20036", Country:"USA", Price:250},
	}
	for _, hotel := range hotels {
		helpers.CheckErr(dbmap.Insert(hotel))
	}
}

type GorpController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	txn, err := dbMap.Begin()
	if err != nil {
		panic(err)
	}
	c.Txn = txn
	return nil
}

func (c *GorpController) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Commit(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

func (c *GorpController) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
