package tests

import (
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
)

type DataStreamTests struct {
	revel.TestSuite
}

const (
	datastream = "/datastreams"
)

func (t *DataStreamTests) TestCreateDataStream() {
	body, _ := json.Marshal(&models.DataStream{Name:"foo1", Description:"Test Stream"})
	t.Post(account + "/1" + project + "/1" + datastream, contentType, bytes.NewBufferString(string(body)))
	t.AssertOk()
}

func (t *DataStreamTests) TestCreateDataStreamInvalidJson() {
	t.Post(account + "/1" + project + "/1" + datastream, contentType, bytes.NewBufferString("{\"name1\":\"foo\"}"))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Data stream name required")
}

func (t *DataStreamTests) TestUpdateDataStream() {
	body, _ := json.Marshal(&models.DataStream{Id: 1, Name:"foo2", Description:"Test Project 1"})
	t.Put(account + "/1" + project + "/1" + datastream + "/1", contentType, bytes.NewBufferString(string(body)))
	t.AssertOk()
	t.AssertContains("foo2")
}

func (t *DataStreamTests) TestUpdateDataStreamNotMatchingIds() {
	body, _ := json.Marshal(&models.DataStream{Id:2, Name:"foo", Description:"bla"})
	t.Put(account + "/1" + project + "/1" + datastream + "/1", contentType, bytes.NewBufferString(string(body)))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Entity ID and URL ID don't match")
}

func (t *DataStreamTests) TestUpdateDataStreamInvalidJson() {
	t.Put(account + "/1" + project + "/1" + datastream + "/1", contentType, bytes.NewBufferString("{\"id\":1, \"name1\":\"foo\"}"))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Data stream name required")
}

func (t *DataStreamTests) TestDeleteDataStream() {
	t.Delete(account + "/1" + project + "/1" + datastream + "/1")
	t.AssertOk()
}

