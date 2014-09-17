package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/services"
	"net/http"
)

type AccountController struct {
	*revel.Controller
	AccountService *services.AccountService `inject:""`
}

func (c *AccountController) All() revel.Result {
	return c.RenderJson(c.AccountService.FindAll())
}

func (c *AccountController) One(id int64) revel.Result {
	return c.RenderJson(c.AccountService.FindOne(id))
}

func (a *AccountController) Add() revel.Result {
	var account models.Account
	err := json.NewDecoder(a.Request.Body).Decode(&account)
	if err != nil {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(err)
	}
	account.Validate(a.Validation)
	if a.Validation.HasErrors() {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(a.Validation.Errors)
	}
	return a.RenderJson(a.AccountService.Save(&account))
}

func (a *AccountController) Update(id int64) revel.Result {
	var account models.Account
	err := json.NewDecoder(a.Request.Body).Decode(&account)
	if err != nil {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(err)
	}
	revel.INFO.Println(&account)
	if account.Id != id {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(&models.Error{Error: "Entity ID and URL ID don't match"})
	}
	account.Validate(a.Validation)
	if a.Validation.HasErrors() {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(a.Validation.Errors)
	}
	err = a.AccountService.Update(&account)
	if err != nil {
		a.Response.Status = http.StatusConflict
		return a.RenderJson(err)
	}
	return a.RenderJson(&account)
}

func (c *AccountController) Delete(id int64) revel.Result {
	return c.RenderJson(c.AccountService.Delete(id))
}
