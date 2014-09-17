package models

import (
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/revel/revel"
	"time"
)

type Account struct {
	Id          int64  `json:"id,omitempty"`
	Version     int64  `json:"-"`
	Created     int64  `json:"-"`
	Updated     int64  `json:"-"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	Email       string `json:"email" binding:"required"`
}

func (a *Account) ToString() string {
	return fmt.Sprintf("id %d, name: %s, email: %s", a.Id, a.Name, a.Email)
}

func (account *Account) Validate(v *revel.Validation) {
	v.Required(account.Name).Key("name").Message("Account name required")
	v.MinSize(account.Name, 3).Key("name").Message("Minimum length is 3 characters")
	v.Required(account.Email).Key("email").Message("Contact email required")
	v.Email(account.Email).Key("email").Message("Must be a valid email address")
}

// implement the PreInsert and PreUpdate hooks
func (a *Account) PreInsert(s gorp.SqlExecutor) error {
	a.Created = time.Now().UnixNano()
	a.Updated = a.Created
	return nil
}

func (a *Account) PreUpdate(s gorp.SqlExecutor) error {
	a.Updated = time.Now().UnixNano()
	return nil
}
