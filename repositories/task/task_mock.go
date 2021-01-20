package task

import (
	"github.com/mlshvsk/go-task-manager/domains"
)

type taskRepositoryMock struct {
	FindAllFunc                func(offset int64, limit int64) ([]*domains.TaskModel, error)
	FindAllByColumnFunc        func(columnId int64, offset int64, limit int64) ([]*domains.TaskModel, error)
	FindAllByColumnAndNameFunc func(columnId int64, name string, offset int64, limit int64) ([]*domains.TaskModel, error)
	FindFunc                   func(id int64) (*domains.TaskModel, error)
	FindWithMaxPositionFunc    func(columnId int64) (*domains.TaskModel, error)
	FindByNextPositionFunc     func(columnId int64, position int64) (*domains.TaskModel, error)
	FindByPreviousPositionFunc func(columnId int64, position int64) (*domains.TaskModel, error)
	CreateFunc                 func(t *domains.TaskModel) error
	UpdateFunc                 func(t *domains.TaskModel) error
	DeleteFunc                 func(id int64) error
}

func InitTaskRepositoryMock() *taskRepositoryMock {
	return &taskRepositoryMock{}
}

func (tr *taskRepositoryMock) FindAll(offset int64, limit int64) ([]*domains.TaskModel, error) {
	return tr.FindAllFunc(offset, limit)
}

func (tr *taskRepositoryMock) FindAllByColumn(columnId int64, offset int64, limit int64) ([]*domains.TaskModel, error) {
	return tr.FindAllByColumnFunc(columnId, offset, limit)
}

func (tr *taskRepositoryMock) FindAllByColumnAndName(columnId int64, name string, offset int64, limit int64) ([]*domains.TaskModel, error) {
	return tr.FindAllByColumnAndNameFunc(columnId, name, offset, limit)
}

func (tr *taskRepositoryMock) Find(id int64) (*domains.TaskModel, error) {
	return tr.FindFunc(id)
}

func (tr *taskRepositoryMock) FindWithMaxPosition(columnId int64) (*domains.TaskModel, error) {
	return tr.FindWithMaxPositionFunc(columnId)
}

func (tr *taskRepositoryMock) FindByNextPosition(columnId int64, position int64) (*domains.TaskModel, error) {
	return tr.FindByNextPositionFunc(columnId, position)
}

func (tr *taskRepositoryMock) FindByPreviousPosition(columnId int64, position int64) (*domains.TaskModel, error) {
	return tr.FindByPreviousPositionFunc(columnId, position)
}

func (tr *taskRepositoryMock) Create(t *domains.TaskModel) error {
	return tr.CreateFunc(t)
}

func (tr *taskRepositoryMock) Update(t *domains.TaskModel) error {
	return tr.UpdateFunc(t)
}

func (tr *taskRepositoryMock) Delete(id int64) error {
	return tr.DeleteFunc(id)
}
