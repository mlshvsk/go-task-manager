package comment

import (
	"github.com/mlshvsk/go-task-manager/models"
)

type ServiceMock struct {
	GetCommentsByTaskFunc func(taskId int64, page int64, limit int64) ([]*models.Comment, error)
	GetCommentFunc        func(commentId int64) (*models.Comment, error)
	StoreCommentFunc      func(c *models.Comment) error
	UpdateCommentFunc     func(c *models.Comment) error
	DeleteCommentFunc     func(commentId int64) error
}

func (s *ServiceMock) GetCommentsByTask(taskId int64, page int64, limit int64) ([]*models.Comment, error) {
	return s.GetCommentsByTaskFunc(taskId, page, limit)
}

func (s *ServiceMock) GetComment(commentId int64) (*models.Comment, error) {
	return s.GetCommentFunc(commentId)
}

func (s *ServiceMock) StoreComment(c *models.Comment) error {
	return s.StoreCommentFunc(c)
}

func (s *ServiceMock) UpdateComment(c *models.Comment) error {
	return s.UpdateCommentFunc(c)
}

func (s *ServiceMock) DeleteComment(commentId int64) error {
	return s.DeleteCommentFunc(commentId)
}
