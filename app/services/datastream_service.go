package services

import (
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"github.com/stsmedia/thingbricks/app/persistence"
)

type DataStreamService struct {
}

func (a *DataStreamService) FindByProject(accountId int64, projectId int64, dataStreamGroupId int64) []*models.DataStream {
	var dataStreams []*models.DataStream
	_, err := persistence.Dbm.Select(&dataStreams, "select * from data_streams where project_id = $1 and account_id = $2 and data_stream_group_id = $3", projectId, accountId, dataStreamGroupId)
	checkErr(err, "no data streams found")
	return dataStreams
}

func (a *DataStreamService) FindOne(accountId int64, projectId int64, dataStreamGroupId int64, id int64) *models.DataStream {
	var dataStream *models.DataStream
	err := persistence.Dbm.SelectOne(&dataStream, "select * from data_streams where id = $1 and project_id = $2 and account_id = $3 and data_stream_group_id = $4", id, projectId, accountId, dataStreamGroupId)
	checkErr(err, "data stream not found")
	return dataStream
}

func (a *DataStreamService) Save(dataStream *models.DataStream) *models.DataStream {
	revel.INFO.Println("saving data stream", dataStream.ToString())
	persistence.Dbm.Insert(dataStream)
	return dataStream
}

func (a *DataStreamService) Update(dataStream *models.DataStream) error {
	revel.INFO.Println("updating data stream", dataStream.ToString())
	count, err := persistence.Dbm.Update(dataStream)
	checkErr(err, "data stream not updated "+dataStream.ToString())
	if count == 1 {
		return nil
	}
	return err
}

func (a *DataStreamService) Delete(accountId int64, projectId int64, dataStreamGroupId int64, id int64) bool {
	_, err := persistence.Dbm.Exec("delete from data_streams where id = $1 and account_id = $2 and project_id = $3 and data_stream_group_id = $4", id, accountId, projectId, dataStreamGroupId)
	if err != nil {
		checkErr(err, "data stream not deleted")
		return false
	}
	return true
}
