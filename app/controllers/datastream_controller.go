package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/services"
	"net/http"
)

type DataStreamController struct {
	*revel.Controller
	DataStreamService *services.DataStreamService `inject:""`
}

func (c *DataStreamController) FindByProject(accountId int64, projectId int64) revel.Result {
	return c.RenderJson(c.DataStreamService.FindByProject(accountId, projectId))
}

func (c *DataStreamController) One(accountId int64, projectId int64, id int64) revel.Result {
	return c.RenderJson(c.DataStreamService.FindOne(accountId, projectId, id))
}

func (a *DataStreamController) Add() revel.Result {
	var dataStream models.DataStream
	err := json.NewDecoder(a.Request.Body).Decode(&dataStream)
	if err != nil {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(err)
	}
	dataStream.Validate(a.Validation)
	if a.Validation.HasErrors() {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(a.Validation.Errors)
	}
	return a.RenderJson(a.DataStreamService.Save(&dataStream))
}

func (a *DataStreamController) Update(id int64) revel.Result {
	var dataStream models.DataStream
	err := json.NewDecoder(a.Request.Body).Decode(&dataStream)
	if err != nil {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(err)
	}
	revel.INFO.Println(&dataStream)
	if dataStream.Id != id {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(&models.Error{Error: "Entity ID and URL ID don't match"})
	}
	dataStream.Validate(a.Validation)
	if a.Validation.HasErrors() {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(a.Validation.Errors)
	}
	err = a.DataStreamService.Update(&dataStream)
	if err != nil {
		a.Response.Status = http.StatusConflict
		return a.RenderJson(err)
	}
	return a.RenderJson(&dataStream)
}

func (c *DataStreamController) Delete(accountId int64, projectId int64, id int64) revel.Result {
	return c.RenderJson(c.DataStreamService.Delete(accountId, projectId, id))
}
