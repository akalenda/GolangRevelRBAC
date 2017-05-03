package controllers

import (
	"github.com/revel/revel"
	"github.com/akalenda/GolangRevelRBAC/app/models"
	"github.com/akalenda/GolangRevelRBAC/app/routes"
	"github.com/akalenda/GolangRevelRBAC/app/helpers"
)

type Application struct {
	GorpController
}

func (c Application) AddUser() revel.Result {
	if u := c.connected(); u != nil {
		c.ViewArgs["user"] = u
	}
	return nil
}

func (c Application) connected() *models.User {
	if c.ViewArgs["user"] != nil {
		return c.ViewArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		u, err := models.GetUserFromDB(c.Txn, username)
		helpers.CheckErr(err)
		return u
	}
	return nil
}

func (c Application) Index() revel.Result {
	if c.connected() != nil {
		return c.Redirect(routes.UserProjects.GETIndex())
	}
	c.Flash.Error("Please log in first")
	return c.Render()
}

func (c Application) Register() revel.Result {
	return c.Render()
}

func (c Application) SaveUser(username string, verifyPassword string) revel.Result {
	u, err := models.RegisterNewUser(c.Txn, username, verifyPassword, "placeholder")
	helpers.CheckErr(err)
	c.Session["user"] = username
	c.Flash.Success("Welcome, " + u.Name)
	return c.Redirect(routes.Hotels.Index())
}

func (c Application) Login(username string, password string, remember bool) revel.Result {
	u, err := models.GetUserFromDB(c.Txn, username)
	helpers.CheckErr(err)
	if u.MatchesHashedPasswordTo(password) {
		c.Session["user"] = username
		if remember {
			c.Session.SetDefaultExpiration()
		} else {
			c.Session.SetNoExpiration()
		}
		c.Flash.Success("Welcome, " + username)
		return c.Redirect(routes.Hotels.Index())
	}
	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(routes.Application.Index())
}

func (c Application) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(routes.Application.Index())
}
