package column

import (
	"github.com/mlshvsk/go-task-manager/domains"
)

type ServiceMock struct {
	GetColumnsFunc   func(projectId int64, page int64, limit int64) ([]*domains.ColumnModel, error)
	GetColumnFunc    func(columnId int64) (*domains.ColumnModel, error)
	StoreColumnFunc  func(c *domains.ColumnModel) error
	UpdateColumnFunc func(c *domains.ColumnModel) error
	DeleteColumnFunc func(columnId int64) error
	MoveColumnFunc   func(columnId int64, direction string) error
}

func (s *ServiceMock) GetColumns(projectId int64, page int64, limit int64) ([]*domains.ColumnModel, error) {
	return s.GetColumnsFunc(projectId, page, limit)
}

func (s *ServiceMock) GetColumn(columnId int64) (*domains.ColumnModel, error) {
	return s.GetColumnFunc(columnId)
}

func (s *ServiceMock) StoreColumn(c *domains.ColumnModel) error {
	return s.StoreColumnFunc(c)
}

func (s *ServiceMock) UpdateColumn(c *domains.ColumnModel) error {
	return s.UpdateColumnFunc(c)
}

func (s *ServiceMock) DeleteColumn(columnId int64) error {
	return s.DeleteColumnFunc(columnId)
}

func (s *ServiceMock) MoveColumn(columnId int64, direction string) error {
	return s.MoveColumnFunc(columnId, direction)
}
