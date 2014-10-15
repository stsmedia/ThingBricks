package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/services"
	"net/http"
)

type DataStreamGroupController struct {
	*revel.Controller
	DataStreamGroupService *services.DataStreamGroupService `inject:""`
}

func (c *DataStreamGroupController) FindByProject(accountId int64, projectId int64) revel.Result {
	return c.RenderJson(c.DataStreamGroupService.FindByProject(accountId, projectId))
}

func (c *DataStreamGroupController) One(accountId int64, projectId int64, id int64) revel.Result {
	return c.RenderJson(c.DataStreamGroupService.FindOne(accountId, projectId, id))
}

func (a *DataStreamGroupController) Add() revel.Result {
	var dataStreamGroup models.DataStreamGroup
	err := json.NewDecoder(a.Request.Body).Decode(&dataStreamGroup)
	if err != nil {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(err)
	}
	dataStreamGroup.Validate(a.Validation)
	if a.Validation.HasErrors() {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(a.Validation.Errors)
	}
	return a.RenderJson(a.DataStreamGroupService.Save(&dataStreamGroup))
}

func (a *DataStreamGroupController) Update(id int64) revel.Result {
	var dataStreamGroup models.DataStreamGroup
	err := json.NewDecoder(a.Request.Body).Decode(&dataStreamGroup)
	if err != nil {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(err)
	}
	revel.INFO.Println(&dataStreamGroup)
	if dataStreamGroup.Id != id {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(&models.Error{Error: "Entity ID and URL ID don't match"})
	}
	dataStreamGroup.Validate(a.Validation)
	if a.Validation.HasErrors() {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(a.Validation.Errors)
	}
	err = a.DataStreamGroupService.Update(&dataStreamGroup)
	if err != nil {
		a.Response.Status = http.StatusConflict
		return a.RenderJson(err)
	}
	return a.RenderJson(&dataStreamGroup)
}

func (c *DataStreamGroupController) Delete(accountId int64, projectId int64, id int64) revel.Result {
	return c.RenderJson(c.DataStreamGroupService.Delete(accountId, projectId, id))
}
