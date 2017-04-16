package controllers

import (
	"database/sql"
	"github.com/go-gorp/gorp"
	"github.com/revel/examples/booking2/app/models"
	"github.com/revel/modules/db/app"
	"github.com/revel/revel"
)

var (
	dbmap *gorp.DbMap
)

func InitDB() {
	db.Init()
	dbmap = &gorp.DbMap{Db: db.Db, Dialect: gorp.SqliteDialect{}}

	models.User_AddTable(dbmap)
	models.UserProject_AddTable(dbmap)
	dbmap.AddTable(models.Hotel{}  ).SetKeys(true, "HotelId")
	dbmap.AddTable(models.Booking{}).SetKeys(true, "BookingId")

	dbmap.TraceOn("[gorp]", revel.INFO)
	checkErr(dbmap.CreateTablesIfNotExists())

	_, err := models.RegisterNewUser(dbmap, "admin", "admin", "administrator")
	if err != nil && err.Error() != "User error: Username and password do not match" {
		panic(err)
	}

	hotels := []*models.Hotel{
		{HotelId: 0, Name: "Marriott Courtyard", Address: "Tower Pl, Buckhead", City: "Atlanta", State: "GA", Zip:"30305", Country:"USA", Price:120},
		{HotelId: 0, Name: "W Hotel", Address: "Union Square, Manhattan", City: "New York",      State: "NY", Zip:"10011", Country:"USA", Price:450},
		{HotelId: 0, Name: "Hotel Rouge", Address: "1315 16th St NW", City: "Washington",        State: "DC", Zip:"20036", Country:"USA", Price:250},
	}
	for _, hotel := range hotels {
		checkErr(dbmap.Insert(hotel))
	}
}

type GorpController struct {
	*revel.Controller
	Txn *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result {
	txn, err := dbmap.Begin()
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

/* ********************* Helpers ******************* */
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
