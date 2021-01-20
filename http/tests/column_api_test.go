package tests

import (
	"bytes"
	"encoding/json"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/http/routes"
	"github.com/mlshvsk/go-task-manager/domains"
	"github.com/mlshvsk/go-task-manager/services"
	"github.com/mlshvsk/go-task-manager/services/column"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetColumnById(t *testing.T) {
	InitLoggers(t)
	expectedColumn := &domains.ColumnModel{Id: 1, Name: "Test"}
	f := func(columnId int64) (*domains.ColumnModel, error) {
		return expectedColumn, nil
	}
	services.ColumnService = &column.ServiceMock{GetColumnFunc: f}

	router := routes.NewRouter()
	r, err := http.NewRequest("GET", "/api/v1/columns/1", bytes.NewBuffer([]byte{0}))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedColumn)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestGetColumns(t *testing.T) {
	InitLoggers(t)
	expectedColumns := make([]*domains.ColumnModel, 2)
	expectedColumns[0] = &domains.ColumnModel{Id: 1, Name: "Test"}
	expectedColumns[1] = &domains.ColumnModel{Id: 2, Name: "Test2"}
	f := func(projectId int64, page int64, limit int64) ([]*domains.ColumnModel, error) {
		assert.Equal(t, int64(1), projectId)
		assert.Equal(t, int64(2), page)
		assert.Equal(t, int64(10), limit)
		return expectedColumns, nil
	}
	services.ColumnService = &column.ServiceMock{GetColumnsFunc: f}

	router := routes.NewRouter()
	r, err := http.NewRequest("GET", "/api/v1/projects/1/columns?page=2&limit=10", bytes.NewBuffer([]byte{0}))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedColumns)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestPostColumn(t *testing.T) {
	InitLoggers(t)
	expectedTime := time.Now()
	expectedId := int64(100)
	expectedName := "Test"
	expectedProjectId := int64(200)
	postColumn := &domains.ColumnModel{Name: expectedName}
	expectedColumn := &domains.ColumnModel{
		Id:        expectedId,
		Name:      expectedName,
		ProjectId: expectedProjectId,
		CreatedAt: expectedTime,
	}

	f := func(c *domains.ColumnModel) error {
		c.Id = expectedId
		c.CreatedAt = expectedTime
		return nil
	}
	services.ColumnService = &column.ServiceMock{StoreColumnFunc: f}

	req, err := json.Marshal(postColumn)
	assert.Nil(t, err)

	router := routes.NewRouter()
	r, err := http.NewRequest("POST", "/api/v1/projects/200/columns", bytes.NewBuffer(req))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedColumn)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}
