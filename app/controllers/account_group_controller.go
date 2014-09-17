package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/services"
	"net/http"
)

type AccountGroupController struct {
	*revel.Controller
	AccountGroupService *services.AccountGroupService `inject:""`
}

func (c *AccountGroupController) FindByAccount(accountId int64) revel.Result {
	return c.RenderJson(c.AccountGroupService.FindByAccount(accountId))
}

func (c *AccountGroupController) One(accountId int64, id int64) revel.Result {
	return c.RenderJson(c.AccountGroupService.FindOne(accountId, id))
}

func (a *AccountGroupController) Add(accountId int64) revel.Result {
	var accountGroup models.AccountGroup
	err := json.NewDecoder(a.Request.Body).Decode(&accountGroup)
	if err != nil {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(err)
	}
	if accountGroup.AccountId != accountId {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(&models.Error{Error: "Entity ID and URL ID don't match"})
	}
	accountGroup.Validate(a.Validation)
	if a.Validation.HasErrors() {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(a.Validation.Errors)
	}
	return a.RenderJson(a.AccountGroupService.Save(&accountGroup))
}

func (a *AccountGroupController) Update(accountId int64, id int64) revel.Result {
	var accountGroup models.AccountGroup
	err := json.NewDecoder(a.Request.Body).Decode(&accountGroup)
	if err != nil {
		a.Response.Status = http.StatusNotAcceptable
		revel.INFO.Println(err)
		return a.RenderJson(err)
	}
	if accountGroup.Id != id || accountGroup.AccountId != accountId {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(&models.Error{Error: "Entity ID and URL ID don't match"})
	}
	accountGroup.Validate(a.Validation)
	if a.Validation.HasErrors() {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(a.Validation.Errors)
	}
	err = a.AccountGroupService.Update(&accountGroup)
	if err != nil {
		a.Response.Status = http.StatusConflict
		return a.RenderJson(err)
	}
	return a.RenderJson(&accountGroup)
}

func (a *AccountGroupController) Delete(accountId int64, id int64) revel.Result {
	success := a.AccountGroupService.Delete(accountId, id)
	if !success {
		a.Response.Status = http.StatusNotAcceptable
		return a.RenderJson(&models.Error{Error: "Cannot delete default account group"})
	}
	return a.RenderJson(id)
}
