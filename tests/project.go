package tests

import (
	"bytes"
	"net/http"
	"encoding/json"
	"github.com/revel/revel"
	"github.com/stsmedia/thingbricks/app/models"
)

type ProjectTests struct {
	revel.TestSuite
}

const (
	project = "/projects"
)

func (t *ProjectTests) TestCreateProject() {
	body, _ := json.Marshal(&models.Project{Name:"foo1", Description:"Test Project"})
	t.Post(account + "/1" + project, contentType, bytes.NewBufferString(string(body)))
	t.AssertOk()
}

func (t *ProjectTests) TestCreateProjectInvalidJson() {
	t.Post(account + "/1" + project, contentType, bytes.NewBufferString("{\"name1\":\"foo\"}"))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Project name required")
}

func (t *ProjectTests) TestUpdateProject() {
	body, _ := json.Marshal(&models.Project{Id: 1, Name:"foo2", Description:"Test Project 1"})
	t.Put(account + "/1" + project + "/1", contentType, bytes.NewBufferString(string(body)))
	t.AssertOk()
	t.AssertContains("foo2")
}

func (t *ProjectTests) TestUpdateProjectNotMatchingIds() {
	body, _ := json.Marshal(&models.Project{Id:2, Name:"foo", Description:"bla"})
	t.Put(account + "/1" + project + "/1", contentType, bytes.NewBufferString(string(body)))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Entity ID and URL ID don't match")
}

func (t *ProjectTests) TestUpdateProjectInvalidJson() {
	t.Put(account + "/1" + project + "/1", contentType, bytes.NewBufferString("{\"id\":1, \"name1\":\"foo\"}"))
	t.AssertStatus(http.StatusNotAcceptable)
	t.AssertContains("Project name required")
}

func (t *ProjectTests) TestDeleteProject() {
	t.Delete(account + "/1" + project + "/1")
	t.AssertOk()
}

