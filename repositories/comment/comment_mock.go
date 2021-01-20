package comment

import (
	"github.com/mlshvsk/go-task-manager/domains"
)

type commentRepositoryMock struct {
	FindAllByTaskFunc func(taskId int64, offset int64, limit int64) ([]*domains.CommentModel, error)
	FindFunc          func(id int64) (*domains.CommentModel, error)
	CreateFunc        func(c *domains.CommentModel) error
	UpdateFunc        func(c *domains.CommentModel) error
	DeleteFunc        func(id int64) error
}

func InitCommentRepositoryMock() *commentRepositoryMock {
	return &commentRepositoryMock{}
}

func (cr *commentRepositoryMock) FindAllByTask(taskId int64, offset int64, limit int64) ([]*domains.CommentModel, error) {
	return cr.FindAllByTaskFunc(taskId, offset, limit)
}

func (cr *commentRepositoryMock) Find(id int64) (*domains.CommentModel, error) {
	return cr.FindFunc(id)
}

func (cr *commentRepositoryMock) Create(c *domains.CommentModel) error {
	return cr.CreateFunc(c)
}

func (cr *commentRepositoryMock) Update(c *domains.CommentModel) error {
	return cr.UpdateFunc(c)
}

func (cr *commentRepositoryMock) Delete(id int64) error {
	return cr.DeleteFunc(id)
}
