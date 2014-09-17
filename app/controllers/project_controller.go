package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/services"
	"net/http"
)

type ProjectController struct {
	*revel.Controller
	ProjectService *services.ProjectService `inject:""`
}

func (c *ProjectController) FindByAccount(accountId int64) revel.Result {
	return c.RenderJson(c.ProjectService.FindByAccount(accountId))
}

func (c *ProjectController) One(accountId int64, id int64) revel.Result {
	return c.RenderJson(c.ProjectService.FindOne(accountId, id))
}

func (a *ProjectController) Add() revel.Result {
	var project models.Project
	err := json.NewDecoder(a.Request.Body).Decode(&project)
	if err != nil {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(err)
	}
	project.Validate(a.Validation)
	if a.Validation.HasErrors() {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(a.Validation.Errors)
	}
	return a.RenderJson(a.ProjectService.Save(&project))
}

func (a *ProjectController) Update(id int64) revel.Result {
	var project models.Project
	err := json.NewDecoder(a.Request.Body).Decode(&project)
	if err != nil {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(err)
	}
	revel.INFO.Println(&project)
	if project.Id != id {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(&models.Error{Error: "Entity ID and URL ID don't match"})
	}
	project.Validate(a.Validation)
	if a.Validation.HasErrors() {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(a.Validation.Errors)
	}
	err = a.ProjectService.Update(&project)
	if err != nil {
		a.Response.Status = http.StatusConflict
		return a.RenderJson(err)
	}
	return a.RenderJson(&project)
}

func (c *ProjectController) Delete(accountId int64, id int64) revel.Result {
	return c.RenderJson(c.ProjectService.Delete(accountId, id))
}
