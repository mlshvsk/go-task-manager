package task

import (
	"github.com/mlshvsk/go-task-manager/models"
)

type ServiceMock struct {
	GetTasksByColumnFunc     func(columnId int64, page int64, limit int64) ([]*models.Task, error)
	GetTasksFunc             func(page int64, limit int64) ([]*models.Task, error)
	GetTaskFunc              func(id int64) (*models.Task, error)
	StoreTaskFunc            func(t *models.Task) error
	UpdateTaskFunc           func(t *models.Task) error
	DeleteTaskFunc           func(taskId int64) error
	MoveTaskWithinColumnFunc func(taskId int64, direction string) error
	MoveTaskToColumnFunc     func(taskId int64, toColumnId int64) error
	MoveAllToColumnFunc      func(fromColumn *models.Column, toColumn *models.Column) error
}

func (s *ServiceMock) GetTasksByColumn(columnId int64, page int64, limit int64) ([]*models.Task, error) {
	return s.GetTasksByColumnFunc(columnId, page, limit)
}

func (s *ServiceMock) GetTasks(page int64, limit int64) ([]*models.Task, error) {
	return s.GetTasksFunc(page, limit)
}

func (s *ServiceMock) GetTask(id int64) (*models.Task, error) {
	return s.GetTaskFunc(id)
}

func (s *ServiceMock) StoreTask(t *models.Task) error {
	return s.StoreTaskFunc(t)
}

func (s *ServiceMock) UpdateTask(t *models.Task) error {
	return s.UpdateTaskFunc(t)
}

func (s *ServiceMock) DeleteTask(taskId int64) error {
	return s.DeleteTaskFunc(taskId)
}

func (s *ServiceMock) MoveTaskWithinColumn(taskId int64, direction string) error {
	return s.MoveTaskWithinColumnFunc(taskId, direction)
}

func (s *ServiceMock) MoveTaskToColumn(taskId int64, toColumnId int64) error {
	return s.MoveTaskToColumnFunc(taskId, toColumnId)
}

func (s *ServiceMock) MoveAllToColumn(fromColumn *models.Column, toColumn *models.Column) error {
	return s.MoveAllToColumnFunc(fromColumn, toColumn)
}
