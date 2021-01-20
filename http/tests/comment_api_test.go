package tests

import (
	"bytes"
	"encoding/json"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/http/routes"
	"github.com/mlshvsk/go-task-manager/domains"
	"github.com/mlshvsk/go-task-manager/services"
	"github.com/mlshvsk/go-task-manager/services/comment"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestGetCommentById(t *testing.T) {
	InitLoggers(t)
	expectedComment := &domains.CommentModel{Id: 1, Data: "Test"}
	f := func(columnId int64) (*domains.CommentModel, error) {
		return expectedComment, nil
	}
	services.CommentService = &comment.ServiceMock{GetCommentFunc: f}

	router := routes.NewRouter()
	r, err := http.NewRequest("GET", "/api/v1/comments/1", bytes.NewBuffer([]byte{0}))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedComment)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestGetComments(t *testing.T) {
	InitLoggers(t)
	expectedComments := make([]*domains.CommentModel, 2)
	expectedComments[0] = &domains.CommentModel{Id: 1, Data: "Test"}
	expectedComments[1] = &domains.CommentModel{Id: 2, Data: "Test2"}
	f := func(projectId int64, page int64, limit int64) ([]*domains.CommentModel, error) {
		assert.Equal(t, int64(1), projectId)
		assert.Equal(t, int64(2), page)
		assert.Equal(t, int64(10), limit)
		return expectedComments, nil
	}
	services.CommentService = &comment.ServiceMock{GetCommentsByTaskFunc: f}

	router := routes.NewRouter()
	r, err := http.NewRequest("GET", "/api/v1/tasks/1/comments?page=2&limit=10", bytes.NewBuffer([]byte{0}))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedComments)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestPostComment(t *testing.T) {
	InitLoggers(t)
	expectedTime := time.Now()
	expectedId := int64(100)
	expectedData := "Test"
	expectedTaskId := int64(200)
	postComment := &domains.CommentModel{Data: expectedData}
	expectedComment := &domains.CommentModel{
		Id:        expectedId,
		Data:      expectedData,
		TaskId: expectedTaskId,
		CreatedAt: expectedTime,
	}

	f := func(c *domains.CommentModel) error {
		c.Id = expectedId
		c.CreatedAt = expectedTime
		return nil
	}
	services.CommentService = &comment.ServiceMock{StoreCommentFunc: f}

	req, err := json.Marshal(postComment)
	assert.Nil(t, err)

	router := routes.NewRouter()
	r, err := http.NewRequest("POST", "/api/v1/tasks/200/comments", bytes.NewBuffer(req))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedComment)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}

func TestUpdateComment(t *testing.T) {
	InitLoggers(t)
	expectedData := "Test"
	postComment := &domains.CommentModel{Data: expectedData}
	expectedComment := &domains.CommentModel{
		Id:        int64(100),
		Data:      expectedData,
		TaskId: int64(100),
		CreatedAt: time.Now(),
	}

	f := func(c *domains.CommentModel) error {
		*c = *expectedComment
		return nil
	}
	services.CommentService = &comment.ServiceMock{UpdateCommentFunc: f}

	req, err := json.Marshal(postComment)
	assert.Nil(t, err)

	router := routes.NewRouter()
	r, err := http.NewRequest("PUT", "/api/v1/comments/100", bytes.NewBuffer(req))

	assert.Nil(t, err)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, r)

	res, err := helpers.RequestBody(w.Result().Body)
	assert.Nil(t, err)

	expectedRes, err := ExpectedOkResponse(expectedComment)
	assert.Nil(t, err)

	assert.Equal(t, string(expectedRes)+"\n", string(res))
}
