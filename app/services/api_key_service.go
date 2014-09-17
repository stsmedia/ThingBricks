package services

import (
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/persistence"
	"github.com/nu7hatch/gouuid"
)

type ApiKeyService struct {
}

func (a *ApiKeyService) FindOne(loginId int64, accessKey string) *models.KeyPair {
	var keyPair *models.KeyPair
	persistence.Dbm.SelectOne(&keyPair, "select * from api_keys where access_key = $1 and login_id = $2", accessKey, loginId)
	return keyPair
}

func (a *ApiKeyService) Create(label string, loginId int64) *models.KeyPair {
	key := &models.KeyPair{Label: label, AccessKey: createUUID(), Secret: createUUID(), Active: true, LoginId: loginId}
	persistence.Dbm.Insert(key)
	return key
}

func (a *ApiKeyService) Delete(loginId int64, accessKey string) bool {
	_, err := persistence.Dbm.Exec("delete from api_keys where access_key = $1 and login_id = $2", accessKey, loginId)
	if err != nil {
		checkErr(err, "account not deleted")
		return false
	}
	return true
}

func createUUID() string {
	id, err := uuid.NewV4();
	checkErr(err, "Unable to generate UUID")
	return id.String()
}
