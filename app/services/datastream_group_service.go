package services

import (
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/persistence"
)

type DataStreamGroupService struct {
}

func (a *DataStreamGroupService) FindByProject(accountId int64, projectId int64) []*models.DataStreamGroup {
	var dataStreamGroups []*models.DataStreamGroup
	_, err := persistence.Dbm.Select(&dataStreamGroups, "select * from data_stream_groups where project_id = $1 and account_id = $2", projectId, accountId)
	checkErr(err, "no data stream groups found")
	return dataStreamGroups
}

func (a *DataStreamGroupService) FindOne(accountId int64, projectId int64, id int64) *models.DataStreamGroup {
	var dataStreamGroups *models.DataStreamGroup
	err := persistence.Dbm.SelectOne(&dataStreamGroups, "select * from data_stream_groups where id = $1 and project_id = $2 and account_id = $3", id, projectId, accountId)
	checkErr(err, "data stream group not found")
	return dataStreamGroups
}

func (a *DataStreamGroupService) Save(dataStreamGroup *models.DataStreamGroup) *models.DataStreamGroup {
	revel.INFO.Println("saving data stream group", dataStreamGroup.ToString())
	persistence.Dbm.Insert(dataStreamGroup)
	return dataStreamGroup
}

func (a *DataStreamGroupService) Update(dataStreamGroups *models.DataStreamGroup) error {
	revel.INFO.Println("updating data stream group", dataStreamGroups.ToString())
	count, err := persistence.Dbm.Update(dataStreamGroups)
	checkErr(err, "data stream group not updated "+dataStreamGroups.ToString())
	if count == 1 {
		return nil
	}
	return err
}

func (a *DataStreamGroupService) Delete(accountId int64, projectId int64, id int64) bool {
	_, err := persistence.Dbm.Exec("delete from data_stream_groups where id = $1 and account_id = $2 and project_id = $3", id, accountId, projectId)
	if err != nil {
		checkErr(err, "data stream group not deleted")
		return false
	}
	return true
}
