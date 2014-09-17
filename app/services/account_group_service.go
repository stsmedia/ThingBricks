package services

import (
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/persistence"
)

type AccountGroupService struct {
}

func (a *AccountGroupService) FindByAccount(accountId int64) []*models.AccountGroup {
	var accountGroups []*models.AccountGroup
	_, err := persistence.Dbm.Select(&accountGroups, "select * from accounts where account_id = $1", accountId)
	checkErr(err, "no account groups found")
	return accountGroups
}

func (a *AccountGroupService) FindOne(accountId int64, id int64) *models.AccountGroup {
	obj, err := persistence.Dbm.Get(models.AccountGroup{}, id)
	checkErr(err, "account group not found")
	if obj == nil {
		return nil
	}
	return obj.(*models.AccountGroup)
}

func (a *AccountGroupService) FindDefault(accountId int64) (*models.AccountGroup, error) {
	var accountGroup *models.AccountGroup
	err := persistence.Dbm.SelectOne(&accountGroup, "select * from account_groups where default_group = true and account_id = $1", accountId)
	if err != nil {
		return nil, err
	}
	return accountGroup, nil
}

func (a *AccountGroupService) Save(accountGroup *models.AccountGroup) *models.AccountGroup {
	revel.INFO.Println("saving account group", accountGroup.ToString())
	// there can only be one default group per account
	_, err := a.FindDefault(accountGroup.AccountId)
	if err != nil {
		accountGroup.Default = false
	} else {
		accountGroup.Default = true
	}
	persistence.Dbm.Insert(accountGroup)
	return accountGroup
}

func (a *AccountGroupService) Update(accountGroup *models.AccountGroup) error {
	revel.INFO.Println("updating account group", accountGroup.ToString())
	count, err := persistence.Dbm.Update(accountGroup)
	checkErr(err, "account not updated "+accountGroup.ToString())
	if count == 1 {
		return nil
	}
	return err
}

func (a *AccountGroupService) Delete(accountId int64, id int64) bool {
	accountGroup, err := a.FindDefault(accountId)
	if err == nil && accountGroup.Id == id {
		revel.ERROR.Println("default account group cannot be deleted")
		return false
	}
	persistence.Dbm.Exec("delete from account_groups where account_id $1 and id=$2", accountId, id)

	return true
}
