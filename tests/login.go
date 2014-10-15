package tests

import (
	"github.com/revel/revel"
)

type LoginTests struct {
	revel.TestSuite
}

func (t *LoginTests) TestFindByAccountGroup() {
	t.Get("/accounts/1/groups/1/logins")
	t.AssertOk()
}

func (t *LoginTests) TestDeleteAccountGroup() {
	t.Delete("/accounts/1/groups/1/logins/1")
	t.AssertOk()
}
