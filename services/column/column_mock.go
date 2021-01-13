package column

import (
	"github.com/mlshvsk/go-task-manager/models"
)

type ServiceMock struct {
	GetColumnsFunc   func(projectId int64, page int64, limit int64) ([]*models.Column, error)
	GetColumnFunc    func(columnId int64) (*models.Column, error)
	StoreColumnFunc  func(c *models.Column) error
	UpdateColumnFunc func(c *models.Column) error
	DeleteColumnFunc func(columnId int64) error
	MoveColumnFunc   func(columnId int64, direction string) error
}

func (s *ServiceMock) GetColumns(projectId int64, page int64, limit int64) ([]*models.Column, error) {
	return s.GetColumnsFunc(projectId, page, limit)
}

func (s *ServiceMock) GetColumn(columnId int64) (*models.Column, error) {
	return s.GetColumnFunc(columnId)
}

func (s *ServiceMock) StoreColumn(c *models.Column) error {
	return s.StoreColumnFunc(c)
}

func (s *ServiceMock) UpdateColumn(c *models.Column) error {
	return s.UpdateColumnFunc(c)
}

func (s *ServiceMock) DeleteColumn(columnId int64) error {
	return s.DeleteColumnFunc(columnId)
}

func (s *ServiceMock) MoveColumn(columnId int64, direction string) error {
	return s.MoveColumnFunc(columnId, direction)
}
