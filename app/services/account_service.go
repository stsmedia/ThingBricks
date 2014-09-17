package services

import (
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/persistence"
)

type AccountService struct {
}

func (a *AccountService) FindAll() []*models.Account {
	var accounts []*models.Account
	_, err := persistence.Dbm.Select(&accounts, "select * from accounts")
	checkErr(err, "no accounts found")
	return accounts
}

func (a *AccountService) FindOne(id int64) *models.Account {
	obj, err := persistence.Dbm.Get(models.Account{}, id)
	checkErr(err, "account not found")
	if obj == nil {
		return nil
	}
	return obj.(*models.Account)
}

func (a *AccountService) Save(account *models.Account) *models.Account {
	revel.INFO.Println("saving account", account.ToString())
	tx, err := persistence.Dbm.Begin()
	checkErr(err, "unable to start transaction")
	tx.Insert(account)
	defaultGroup := &models.AccountGroup{Name: "Default", AccountId: account.Id, Default: true}
	tx.Insert(defaultGroup)
	tx.Commit()
	return account
}

func (a *AccountService) Update(account *models.Account) error {
	revel.INFO.Println("updating account", account.ToString())
	count, err := persistence.Dbm.Update(account)
	checkErr(err, "account not updated "+account.ToString())
	if count == 1 {
		return nil
	}
	return err
}

func (a *AccountService) Delete(id int64) bool {
	_, err := persistence.Dbm.Exec("delete from accounts where id=$1", id)
	if err != nil {
		checkErr(err, "account not deleted")
		return false
	}
	return true
}

func checkErr(err error, msg string) {
	if err != nil {
		revel.ERROR.Fatalln(msg, err)
	}
}
