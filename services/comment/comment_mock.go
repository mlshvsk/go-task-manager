package comment

import (
	"github.com/mlshvsk/go-task-manager/domains"
)

type ServiceMock struct {
	GetCommentsByTaskFunc func(taskId int64, page int64, limit int64) ([]*domains.CommentModel, error)
	GetCommentFunc        func(commentId int64) (*domains.CommentModel, error)
	StoreCommentFunc      func(c *domains.CommentModel) error
	UpdateCommentFunc     func(c *domains.CommentModel) error
	DeleteCommentFunc     func(commentId int64) error
}

func (s *ServiceMock) GetCommentsByTask(taskId int64, page int64, limit int64) ([]*domains.CommentModel, error) {
	return s.GetCommentsByTaskFunc(taskId, page, limit)
}

func (s *ServiceMock) GetComment(commentId int64) (*domains.CommentModel, error) {
	return s.GetCommentFunc(commentId)
}

func (s *ServiceMock) StoreComment(c *domains.CommentModel) error {
	return s.StoreCommentFunc(c)
}

func (s *ServiceMock) UpdateComment(c *domains.CommentModel) error {
	return s.UpdateCommentFunc(c)
}

func (s *ServiceMock) DeleteComment(commentId int64) error {
	return s.DeleteCommentFunc(commentId)
}
