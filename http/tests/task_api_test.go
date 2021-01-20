package tests

import (
	"bytes"
	"encoding/json"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/http/routes"
	"github.com/mlshvsk/go-task-manager/domains"
	"github.com/mlshvsk/go-task-manager/services"
	"github.com/mlshvsk/go-task-manager/services/task"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetTaskById(t *testing.T) {
	InitLoggers(t)
	expectedTask := &domains.TaskModel{Id: 1, Name: "Test"}
	f := func(taskId int64) (*domains.TaskModel, error) {
		return expectedTask, nil
	}
	services.TaskService = &task.ServiceMock{GetTaskFunc: f}

	router := routes.NewRouter()
	r, err := http.NewRequest("GET", "/api/v1/tasks/1", bytes.NewBuffer([]byte{0}))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedTask)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestGetTasks(t *testing.T) {
	InitLoggers(t)
	expectedTasks := make([]*domains.TaskModel, 2)
	expectedTasks[0] = &domains.TaskModel{Id: 1, Name: "Test"}
	expectedTasks[1] = &domains.TaskModel{Id: 2, Name: "Test2"}
	f := func(page int64, limit int64) ([]*domains.TaskModel, error) {
		assert.Equal(t, int64(2), page)
		assert.Equal(t, int64(10), limit)
		return expectedTasks, nil
	}
	services.TaskService = &task.ServiceMock{GetTasksFunc: f}

	router := routes.NewRouter()
	r, err := http.NewRequest("GET", "/api/v1/tasks?page=2&limit=10", bytes.NewBuffer([]byte{0}))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedTasks)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestPostTask(t *testing.T) {
	InitLoggers(t)
	expectedTime := time.Now()
	expectedId := int64(100)
	expectedName := "Test"
	expectedColumnId := int64(200)
	postColumn := &domains.TaskModel{Name: expectedName}
	expectedColumn := &domains.TaskModel{
		Id:        expectedId,
		Name:      expectedName,
		ColumnId: expectedColumnId,
		CreatedAt: expectedTime,
	}

	f := func(c *domains.TaskModel) error {
		c.Id = expectedId
		c.CreatedAt = expectedTime
		return nil
	}
	services.TaskService = &task.ServiceMock{StoreTaskFunc: f}

	req, err := json.Marshal(postColumn)
	assert.Nil(t, err)

	router := routes.NewRouter()
	r, err := http.NewRequest("POST", "/api/v1/columns/200/tasks", bytes.NewBuffer(req))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedColumn)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}
