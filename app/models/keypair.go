package models

import (
	"code.google.com/p/go.crypto/bcrypt"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/revel/revel"
	"time"
)

type KeyPair struct {
	Id        int64  `json:"id,omitempty"`
	Version   int64  `json:"-"`
	Created   int64  `json:"-"`
	Updated   int64  `json:"-"`
	AccessKey string `json:"accessKey" db:"access_key"`
	Secret    string `json:"secret" db:"-"`
	Label     string `json:"label"`
	Active    bool   `json:"active"`
	Hashed    []byte `json:"-"`
	LoginId   int64  `json:"loginId" db:"login_id"`
}

func (k *KeyPair) ToString() string {
	return fmt.Sprintf("id %d, label: %s, active: %b, access key: %s", k.Id, k.Label, k.Active, k.AccessKey)
}

// implement the PreInsert and PreUpdate hooks
func (k *KeyPair) PreInsert(s gorp.SqlExecutor) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(k.Secret), bcrypt.DefaultCost); if err != nil {
		revel.ERROR.Println("unable to create has for password")
	}
	k.Created = time.Now().UnixNano()
	k.Updated = k.Created
	k.Hashed = hash
	return nil
}
