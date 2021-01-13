package task

import (
	"github.com/mlshvsk/go-task-manager/models"
)

type taskRepositoryMock struct {
	FindAllFunc                func(offset int64, limit int64) ([]*models.Task, error)
	FindAllByColumnFunc        func(columnId int64, offset int64, limit int64) ([]*models.Task, error)
	FindAllByColumnAndNameFunc func(columnId int64, name string, offset int64) ([]*models.Task, error)
	FindFunc                   func(id int64) (*models.Task, error)
	FindWithMaxPositionFunc    func(columnId int64) (*models.Task, error)
	FindByNextPositionFunc     func(columnId int64, position int64) (*models.Task, error)
	FindByPreviousPositionFunc func(columnId int64, position int64) (*models.Task, error)
	CreateFunc                 func(t *models.Task) error
	UpdateFunc                 func(t *models.Task) error
	DeleteFunc                 func(id int64) error
}

func (tr *taskRepositoryMock) FindAll(offset int64, limit int64) ([]*models.Task, error) {
	return tr.FindAllFunc(offset, limit)
}

func (tr *taskRepositoryMock) FindAllByColumn(columnId int64, offset int64, limit int64) ([]*models.Task, error) {
	return tr.FindAllByColumnFunc(columnId, offset, limit)
}

func (tr *taskRepositoryMock) FindAllByColumnAndName(columnId int64, name string, offset int64) ([]*models.Task, error) {
	return tr.FindAllByColumnAndNameFunc(columnId, name, offset)
}

func (tr *taskRepositoryMock) Find(id int64) (*models.Task, error) {
	return tr.FindFunc(id)
}

func (tr *taskRepositoryMock) FindWithMaxPosition(columnId int64) (*models.Task, error) {
	return tr.FindWithMaxPositionFunc(columnId)
}

func (tr *taskRepositoryMock) FindByNextPosition(columnId int64, position int64) (*models.Task, error) {
	return tr.FindByNextPositionFunc(columnId, position)
}

func (tr *taskRepositoryMock) FindByPreviousPosition(columnId int64, position int64) (*models.Task, error) {
	return tr.FindByPreviousPositionFunc(columnId, position)
}

func (tr *taskRepositoryMock) Create(t *models.Task) error {
	return tr.CreateFunc(t)
}

func (tr *taskRepositoryMock) Update(t *models.Task) error {
	return tr.UpdateFunc(t)
}

func (tr *taskRepositoryMock) Delete(id int64) error {
	return tr.DeleteFunc(id)
}
