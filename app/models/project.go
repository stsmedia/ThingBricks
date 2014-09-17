package models

import (
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/revel/revel"
	"time"
)

type Project struct {
	Id          int64  `json:"id,omitempty"`
	Version     int64  `json:"-"`
	Created     int64  `json:"-"`
	Updated     int64  `json:"-"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty"`
	AccountId   int64  `json:"accountId" db:"account_id"`
}

func (a *Project) ToString() string {
	return fmt.Sprintf("id %d, name: %s, description: %s", a.Id, a.Name, a.Description)
}

func (project *Project) Validate(v *revel.Validation) {
	v.Required(project.Name).Key("name").Message("Project name required")
	v.MinSize(project.Name, 3).Key("name").Message("Minimum length is 3 characters")
	v.Required(project.AccountId).Key("accountId").Message("Account ID required")
}

// implement the PreInsert and PreUpdate hooks
func (a *Project) PreInsert(s gorp.SqlExecutor) error {
	a.Created = time.Now().UnixNano()
	a.Updated = a.Created
	return nil
}

func (a *Project) PreUpdate(s gorp.SqlExecutor) error {
	a.Updated = time.Now().UnixNano()
	return nil
}
