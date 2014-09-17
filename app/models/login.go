package models

import (
	"fmt"
	"github.com/coopernurse/gorp"
	"time"
)

type Login struct {
	Id             int64  `json:"id,omitempty"`
	Version        int64  `json:"-"`
	Created        int64  `json:"-"`
	Updated        int64  `json:"-"`
	Email          string `json:"email"`
	FirstName      string `json:"firstName" db:"first_name"`
	LastName       string `json:"lastName" db:"last_name"`
	Gender         string `json:"gender"`
	Picture        string `json:"picture"`
	AccessToken    string `json:"-" db:"access_token"`
	Network        string `json:"-"`
	AccountGroupId int64  `json:"accountGroupId" db:"account_group_id"`
}

func (l *Login) ToString() string {
	return fmt.Sprintf("id %d, first name: %s, last name: %s, email: %s, picture: %s", l.Id, l.FirstName, l.LastName, l.Email, l.Picture)
}

// implement the PreInsert and PreUpdate hooks
func (l *Login) PreInsert(s gorp.SqlExecutor) error {
	l.Created = time.Now().UnixNano()
	l.Updated = l.Created
	return nil
}
