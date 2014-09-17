package tests

import (
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
)

type AccountTests struct {
	revel.TestSuite
}

const (
	account = "/accounts"
	contentType = "application/json"
)

func (t *AccountTests) TestCreateAccount() {
	body, _ := json.Marshal(&models.Account{Name:"foo1", Email:"bla@foo.com", Description:"Test Account"})
	t.Post(account, contentType, bytes.NewBufferString(string(body)))
	t.AssertOk()
}

func (t *AccountTests) TestCreateAccountInvalidJson() {
	t.Post(account, contentType, bytes.NewBufferString("{\"name1\":\"foo\"}"))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Account name required")
	t.AssertContains("Must be a valid email address")
}

func (t *AccountTests) TestUpdateAccount() {
	body, _ := json.Marshal(&models.Account{Id: 1, Name:"foo1", Email:"bla@foo.com"})
	t.Put(account + "/1", contentType, bytes.NewBufferString(string(body)))
	t.AssertOk()
	t.AssertContains("foo1")
}

func (t *AccountTests) TestUpdateAccountNotMatchingIds() {
	body, _ := json.Marshal(&models.Account{Id:2, Name:"foo", Email:"bla@foo.com"})
	t.Put(account + "/1", contentType, bytes.NewBufferString(string(body)))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Entity ID and URL ID don't match")
}

func (t *AccountTests) TestUpdateAccountInvalidJson() {
	t.Put(account + "/1", contentType, bytes.NewBufferString("{\"id\":1, \"name1\":\"foo\"}"))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Account name required")
}

func (t *AccountTests) TestDeleteAccount() {
	t.Delete(account + "/1")
	t.AssertOk()
}

