package tests

import (
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
)

type AccountGroupTests struct {
	revel.TestSuite
}

func (t *AccountGroupTests) TestCreateAccountGroup() {
	body, _ := json.Marshal(&models.AccountGroup{Name:"foo", AccountId: 1})
	t.Post("/accounts/1/groups", contentType, bytes.NewBufferString(string(body)))
	t.AssertOk()
}

func (t *AccountGroupTests) TestCreateAccountGroupInvalidJson() {
	t.Post("/accounts/1/groups", contentType, bytes.NewBufferString("{\"name1\":\"foo\"}"))
	t.AssertStatus(http.StatusNotAcceptable)
//	t.AssertContains("Account group name required")
}

func (t *AccountGroupTests) TestUpdateAccountGroup() {
	body, _ := json.Marshal(&models.AccountGroup{Id: 1, Name:"foo1", AccountId: 1})
	t.Put("/accounts/1/groups/1", contentType, bytes.NewBufferString(string(body)))
	t.AssertOk()
	t.AssertContains("foo1")
}

func (t *AccountGroupTests) TestUpdateAccountGroupNotMatchingIds() {
	body, _ := json.Marshal(&models.AccountGroup{Id:2, Name:"foo", AccountId: 1})
	t.Put("/accounts/1/groups/1", contentType, bytes.NewBufferString(string(body)))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Entity ID and URL ID don't match")
}

func (t *AccountGroupTests) TestUpdateAccountGroupNotMatchingAccountIds() {
	body, _ := json.Marshal(&models.AccountGroup{Id:2, Name:"foo", AccountId: 2})
	t.Put("/accounts/1/groups/1", contentType, bytes.NewBufferString(string(body)))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Entity ID and URL ID don't match")
}

func (t *AccountGroupTests) TestUpdateAccountGroupInvalidJson() {
	t.Put("/accounts/1/groups/1", contentType, bytes.NewBufferString("{\"id\":1, \"name1\":\"foo\"}"))
	t.AssertStatus(http.StatusNotAcceptable)
//	t.AssertContains("Account group name required")
}

func (t *AccountGroupTests) TestDeleteAccountGroup() {
	t.Delete("/accounts/1/groups/1")
	t.AssertOk()
}

