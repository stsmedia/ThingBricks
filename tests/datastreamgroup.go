package tests

import (
	"bytes"
	"encoding/json"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
	"net/http"
)

type DataStreamGroupTests struct {
	revel.TestSuite
}

const (
	datastreamgroup = "/datastreamgroups"
)

func (t *DataStreamGroupTests) TestCreateDataStreamGroup() {
	body, _ := json.Marshal(&models.DataStreamGroup{Name: "foo1", Description: "Test Stream"})
	t.Post(account+"/1"+project+"/1"+datastreamgroup, contentType, bytes.NewBufferString(string(body)))
	t.AssertOk()
}

func (t *DataStreamGroupTests) TestCreateDataStreamGroupInvalidJson() {
	t.Post(account+"/1"+project+"/1"+datastreamgroup, contentType, bytes.NewBufferString("{\"name1\":\"foo\"}"))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Data stream group name required")
}

func (t *DataStreamGroupTests) TestUpdateDataStreamGroup() {
	body, _ := json.Marshal(&models.DataStreamGroup{Id: 1, Name: "foo2", Description: "Test Project 1"})
	t.Put(account+"/1"+project+"/1"+datastreamgroup+"/1", contentType, bytes.NewBufferString(string(body)))
	t.AssertOk()
	t.AssertContains("foo2")
}

func (t *DataStreamGroupTests) TestUpdateDataStreamGroupNotMatchingIds() {
	body, _ := json.Marshal(&models.DataStreamGroup{Id: 2, Name: "foo", Description: "bla"})
	t.Put(account+"/1"+project+"/1"+datastreamgroup+"/1", contentType, bytes.NewBufferString(string(body)))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Entity ID and URL ID don't match")
}

func (t *DataStreamGroupTests) TestUpdateDataStreamGroupInvalidJson() {
	t.Put(account+"/1"+project+"/1"+datastreamgroup+"/1", contentType, bytes.NewBufferString("{\"id\":1, \"name1\":\"foo\"}"))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Data stream group name required")
}

func (t *DataStreamGroupTests) TestDeleteDataStreamGroup() {
	t.Delete(account + "/1" + project + "/1" + datastreamgroup + "/1")
	t.AssertOk()
}
