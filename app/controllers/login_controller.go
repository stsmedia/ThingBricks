package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/services"
	"net/http"
	"time"
)

type LoginController struct {
	*revel.Controller
	AuthService *services.AuthService `inject:""`
}

type credentials struct {
	Network     string `json:"network"`
	AccessToken string `json:"accessToken"`
}

func (l *LoginController) Signin() revel.Result {
	var signin credentials
	err := json.NewDecoder(l.Request.Body).Decode(&signin)
	if err != nil {
		return l.RenderJson(&models.Error{Error: "Invalid json formatting"})
	}
	login, err := l.AuthService.Verify(&models.Login{Network: signin.Network, AccessToken: signin.AccessToken})
	if err != nil || login == nil {
		l.Response.Status = http.StatusForbidden
		revel.ERROR.Println(err)
		return l.RenderJson(err)
	}

	expire := time.Now().AddDate(0, 0, 1)
	cookie := http.Cookie{"tbsec", "foo", "/", "localhost:9001", expire, expire.Format(time.UnixDate), 86400, false, false, "tbsec=foo", []string{"tbsec=foo"}}
	l.SetCookie(&cookie)
	return l.RenderJson(login)
}

func (c *LoginController) FindByAccountGroup(accountGroupId int64) revel.Result {
	return c.RenderJson(c.AuthService.FindByAccountGroup(accountGroupId))
}

func (c *LoginController) Delete(accountGroupId int64, id int64) revel.Result {
	return c.RenderJson(c.AuthService.Delete(accountGroupId, id))
}
