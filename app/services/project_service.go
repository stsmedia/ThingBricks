package services

import (
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/persistence"
)

type ProjectService struct {
}

func (a *ProjectService) FindByAccount(accountId int64) []*models.Project {
	var projects []*models.Project
	_, err := persistence.Dbm.Select(&projects, "select * from projects where account_id = $1", accountId)
	checkErr(err, "no projects found")
	return projects
}

func (a *ProjectService) FindOne(accountId int64, id int64) *models.Project {
	var project *models.Project
	err := persistence.Dbm.SelectOne(&project, "select * from projects where id = $1 and account_id = $2", id, accountId)
	checkErr(err, "project not found")
	return project
}

func (a *ProjectService) Save(project *models.Project) *models.Project {
	revel.INFO.Println("saving project", project.ToString())
	persistence.Dbm.Insert(project)
	return project
}

func (a *ProjectService) Update(project *models.Project) error {
	revel.INFO.Println("updating project", project.ToString())
	count, err := persistence.Dbm.Update(project)
	checkErr(err, "project not updated "+project.ToString())
	if count == 1 {
		return nil
	}
	return err
}

func (a *ProjectService) Delete(accountId int64, id int64) bool {
	_, err := persistence.Dbm.Exec("delete from projects where id = $1 and account_id = $2", id, accountId)
	if err != nil {
		checkErr(err, "project not deleted")
		return false
	}
	return true
}
