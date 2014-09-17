package models

import (
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/revel/revel"
	"time"
)

type AccountGroup struct {
	Id        int64  `json:"id,omitempty"`
	Version   int64  `json:"-"`
	Created   int64  `json:"-"`
	Updated   int64  `json:"-"`
	Name      string `json:"name"`
	Default   bool   `json:"default" db:"default_group"`
	AccountId int64  `json:"accountId" db:"account_id"`
}

func (a *AccountGroup) ToString() string {
	return fmt.Sprintf("id %d, name: %s", a.Id, a.Name)
}

func (a *AccountGroup) Validate(v *revel.Validation) {
	v.Required(a.Name).Key("name").Message("Account group name required")
	v.MinSize(a.Name, 3).Key("name").Message("Minimum length is 3 characters")
}

// implement the PreInsert and PreUpdate hooks
func (a *AccountGroup) PreInsert(s gorp.SqlExecutor) error {
	a.Created = time.Now().UnixNano()
	a.Updated = a.Created
	return nil
}

func (a *AccountGroup) PreUpdate(s gorp.SqlExecutor) error {
	a.Updated = time.Now().UnixNano()
	return nil
}
