package tests

import (
	"bytes"
	"encoding/json"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/http/routes"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
	"github.com/mlshvsk/go-task-manager/services/project"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetProjectById(t *testing.T) {
	InitLoggers(t)
	expectedProject := &models.Project{Id: 1, Name: "Test", Description: "Deesc", CreatedAt: time.Now()}
	f := func(columnId int64) (*models.Project, error) {
		return expectedProject, nil
	}
	services.ProjectService = &project.ServiceMock{GetProjectFunc: f}

	router := routes.NewRouter()
	r, err := http.NewRequest("GET", "/api/v1/projects/1", bytes.NewBuffer([]byte{0}))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedProject)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestGetProjects(t *testing.T) {
	InitLoggers(t)
	expectedProjects := make([]*models.Project, 2)
	expectedProjects[0] = &models.Project{Id: 1, Name: "Test"}
	expectedProjects[1] = &models.Project{Id: 2, Name: "Test2"}
	f := func(page int64, limit int64) ([]*models.Project, error) {
		assert.Equal(t, int64(2), page)
		assert.Equal(t, int64(10), limit)
		return expectedProjects, nil
	}
	services.ProjectService = &project.ServiceMock{GetProjectsFunc: f}

	router := routes.NewRouter()
	r, err := http.NewRequest("GET", "/api/v1/projects?page=2&limit=10", bytes.NewBuffer([]byte{0}))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedProjects)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestPostProject(t *testing.T) {
	InitLoggers(t)
	expectedTime := time.Now()
	expectedId := int64(100)
	expectedName := "Test"
	postComment := &models.Project{Name: expectedName}
	expectedProject := &models.Project{
		Id:        expectedId,
		Name:      expectedName,
		CreatedAt: expectedTime,
	}

	f := func(c *models.Project) error {
		c.Id = expectedId
		c.CreatedAt = expectedTime
		return nil
	}
	services.ProjectService = &project.ServiceMock{StoreProjectFunc: f}

	req, err := json.Marshal(postComment)
	assert.Nil(t, err)

	router := routes.NewRouter()
	r, err := http.NewRequest("POST", "/api/v1/projects", bytes.NewBuffer(req))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedProject)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestUpdateProject(t *testing.T) {
	InitLoggers(t)
	expectedName := "Test"
	postProject := &models.Project{Name: expectedName}
	expectedProject := &models.Project{
		Id:        int64(100),
		Name:      expectedName,
		CreatedAt: time.Now(),
	}

	f := func(c *models.Project) error {
		*c = *expectedProject
		return nil
	}
	services.ProjectService = &project.ServiceMock{UpdateProjectFunc: f}

	req, err := json.Marshal(postProject)
	assert.Nil(t, err)

	router := routes.NewRouter()
	r, err := http.NewRequest("PUT", "/api/v1/projects/100", bytes.NewBuffer(req))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedProject)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}
