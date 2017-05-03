package controllers

import (
	"github.com/revel/revel"
	"github.com/akalenda/GolangRevelRBAC/app/models"
	"github.com/akalenda/GolangRevelRBAC/app/helpers"
)

type UserProjects struct {
	Application
}

func (c UserProjects) GETIndex() revel.Result {
	username := c.connected().UserId
	results, err := c.Txn.Select(models.UserProject{}, `select * from UserProject`, username)
	helpers.CheckErr(err)
	var projectList []*models.UserProject
	for _, result := range results{
		projectItem := result.(*models.UserProject)
		projectList = append(projectList, projectItem)
		println("***PI:", projectItem.UserProjectId, projectItem.OwnerUsername, projectItem.Name, projectItem.Description)
	}
	return c.Render(projectList)
}

func (c UserProjects) POSTIndex(projectName string, projectDescription string) revel.Result {
	username := c.connected().Username
	println("INSERTING: ", username, " ", projectName, " ", projectDescription)
	c.Txn.Insert(&models.UserProject{
		OwnerUsername: username,
		Name: projectName,
		Description: projectDescription,
	})
	return c.Render()
}