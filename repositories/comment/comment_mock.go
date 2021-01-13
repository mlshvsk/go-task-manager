package comment

import (
	"github.com/mlshvsk/go-task-manager/models"
)

type commentRepositoryMock struct {
	FindAllByTaskFunc func(taskId int64, offset int64, limit int64) ([]*models.Comment, error)
	FindFunc          func(id int64) (*models.Comment, error)
	CreateFunc        func(c *models.Comment) error
	UpdateFunc        func(c *models.Comment) error
	DeleteFunc        func(id int64) error
}

func (cr *commentRepositoryMock) FindAllByTask(taskId int64, offset int64, limit int64) ([]*models.Comment, error) {
	return cr.FindAllByTaskFunc(taskId, offset, limit)
}

func (cr *commentRepositoryMock) Find(id int64) (*models.Comment, error) {
	return cr.FindFunc(id)
}

func (cr *commentRepositoryMock) Create(c *models.Comment) error {
	return cr.CreateFunc(c)
}

func (cr *commentRepositoryMock) Update(c *models.Comment) error {
	return cr.UpdateFunc(c)
}

func (cr *commentRepositoryMock) Delete(id int64) error {
	return cr.DeleteFunc(id)
}
