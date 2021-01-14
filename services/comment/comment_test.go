package comment

import (
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories/comment"
	"github.com/mlshvsk/go-task-manager/services"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestGetComments(t *testing.T) {
	cr := comment.InitCommentRepositoryMock()
	expectedResult := []*models.Comment{
		{
			Id:   int64(rand.Int()),
			Data: "Test",
		},
		{
			Id:   int64(rand.Int()),
			Data: "Test2",
		},
	}
	expectedTaskId := int64(rand.Uint64())
	expectedPage := int64(rand.Uint64())
	expectedLimit := int64(rand.Uint64())

	cr.FindAllByTaskFunc = func(taskId int64, offset int64, limit int64) ([]*models.Comment, error) {
		assert.Equal(t, expectedTaskId, taskId)
		assert.Equal(t, expectedPage, offset)
		assert.Equal(t, expectedLimit, limit)
		return expectedResult, nil
	}

	InitCommentService(cr)
	res, err := services.CommentService.GetCommentsByTask(expectedTaskId, expectedPage, expectedLimit)

	assert.Equal(t, expectedResult, res)
	assert.Nil(t, err)
}

func TestGetComment(t *testing.T) {
	cr := comment.InitCommentRepositoryMock()
	expectedResult := &models.Comment{
		Id:   int64(rand.Int()),
		Data: "Test",
	}
	expectedCommentId := int64(rand.Uint64())

	cr.FindFunc = func(commentId int64) (*models.Comment, error) {
		assert.Equal(t, expectedCommentId, commentId)
		return expectedResult, nil
	}

	InitCommentService(cr)
	res, err := services.CommentService.GetComment(expectedCommentId)

	assert.Equal(t, expectedResult, res)
	assert.Nil(t, err)
}
