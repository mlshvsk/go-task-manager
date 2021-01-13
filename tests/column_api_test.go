package tests

import (
	"bytes"
	"encoding/json"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/http/routes"
	"github.com/mlshvsk/go-task-manager/logger"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services/column"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetColumnById(t *testing.T) {
	initLoggers(t)
	expectedColumn := &models.Column{Id: 1, Name: "Test"}
	f := func(columnId int64) (*models.Column, error) {
		return expectedColumn, nil
	}
	column.Service = &column.ServiceMock{GetColumnFunc: f}

	router := routes.NewRouter()
	r, err := http.NewRequest("GET", "/api/v1/columns/1", bytes.NewBuffer([]byte{0}))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := expectedOkResponse(expectedColumn)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestGetColumns(t *testing.T) {
	initLoggers(t)
	expectedColumns := make([]*models.Column, 2)
	expectedColumns[0] = &models.Column{Id: 1, Name: "Test"}
	expectedColumns[1] = &models.Column{Id: 2, Name: "Test2"}
	f := func(projectId int64, page int64, limit int64) ([]*models.Column, error) {
		assert.Equal(t, int64(1), projectId)
		assert.Equal(t, int64(2), page)
		assert.Equal(t, int64(10), limit)
		return expectedColumns, nil
	}
	column.Service = &column.ServiceMock{GetColumnsFunc: f}

	router := routes.NewRouter()
	r, err := http.NewRequest("GET", "/api/v1/projects/1/columns?page=2&limit=10", bytes.NewBuffer([]byte{0}))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := expectedOkResponse(expectedColumns)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestPostColumn(t *testing.T) {
	initLoggers(t)
	expectedTime := time.Now()
	expectedId := int64(100)
	expectedName := "Test"
	expectedProjectId := int64(200)
	postColumn := &models.Column{Name: expectedName}
	expectedColumn := &models.Column{
		Id:        expectedId,
		Name:      expectedName,
		ProjectId: expectedProjectId,
		CreatedAt: expectedTime,
	}

	f := func(c *models.Column) error {
		c.Id = expectedId
		c.CreatedAt = expectedTime
		return nil
	}
	column.Service = &column.ServiceMock{StoreColumnFunc: f}

	req, err := json.Marshal(postColumn)
	assert.Nil(t, err)

	router := routes.NewRouter()
	r, err := http.NewRequest("POST", "/api/v1/projects/200/columns", bytes.NewBuffer(req))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := expectedOkResponse(expectedColumn)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func initLoggers(t *testing.T) {
	logger.RequestLogger = zaptest.NewLogger(t).Sugar()
	logger.ErrorLogger = zaptest.NewLogger(t).Sugar()
}

func expectedOkResponse(data interface{}) ([]byte, error) {
	responseStruct := &struct {
		Data interface{} `json:"data"`
	}{Data: data}

	return json.Marshal(responseStruct)
}
