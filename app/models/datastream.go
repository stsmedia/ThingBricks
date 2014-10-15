package models

import (
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/revel/revel"
	"time"
)

type DataStream struct {
	Id                int64  `json:"id,omitempty"`
	Version           int64  `json:"-"`
	Created           int64  `json:"-"`
	Updated           int64  `json:"-"`
	Name              string `json:"name" binding:"required"`
	Description       string `json:"description,omitempty"`
	ProjectId         int64  `json:"projectId" db:"project_id"`
	AccountId         int64  `json:"accountId" db:"account_id"`
	DataStreamGroupId int64  `json:"dataStreamGroupId" db:"data_stream_group_id"`
}

func (a *DataStream) ToString() string {
	return fmt.Sprintf("id %d, name: %s, description: %s", a.Id, a.Name, a.Description)
}

func (d *DataStream) Validate(v *revel.Validation) {
	v.Required(d.Name).Key("name").Message("Data stream name required")
	v.MinSize(d.Name, 3).Key("name").Message("Minimum length is 3 characters")
	v.Required(d.AccountId).Key("accountId").Message("Account ID required")
	v.Required(d.ProjectId).Key("projectId").Message("Project ID required")
	v.Required(d.DataStreamGroupId).Key("dataStreamGroupId").Message("Data Stream Group ID required")
}

// implement the PreInsert and PreUpdate hooks
func (a *DataStream) PreInsert(s gorp.SqlExecutor) error {
	a.Created = time.Now().UnixNano()
	a.Updated = a.Created
	return nil
}

func (a *DataStream) PreUpdate(s gorp.SqlExecutor) error {
	a.Updated = time.Now().UnixNano()
	return nil
}
