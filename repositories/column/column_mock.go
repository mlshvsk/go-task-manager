package column

import (
	"github.com/mlshvsk/go-task-manager/domains"
)

type columnRepositoryMock struct {
	FinAllFunc                  func(offset int64, limit int64) ([]*domains.ColumnModel, error)
	FindAllByProjectFunc        func(projectId int64, offset int64, limit int64) ([]*domains.ColumnModel, error)
	FindAllByProjectAndNameFunc func(projectId int64, name string) ([]*domains.ColumnModel, error)
	FindFunc                    func(id int64) (*domains.ColumnModel, error)
	FindByNextPositionFunc      func(projectId int64, position int64) (*domains.ColumnModel, error)
	FindByPreviousPositionFunc  func(projectId int64, position int64) (*domains.ColumnModel, error)
	FindWithMaxPositionFunc     func(projectId int64) (*domains.ColumnModel, error)
	CreateFunc                  func(c *domains.ColumnModel) error
	UpdateFunc                  func(c *domains.ColumnModel) error
	DeleteFunc                  func(id int64) error
}

func InitColumnRepositoryMock() *columnRepositoryMock {
	return &columnRepositoryMock{}
}

func (cr *columnRepositoryMock) FindAll(offset int64, limit int64) ([]*domains.ColumnModel, error) {
	return cr.FinAllFunc(offset, limit)
}

func (cr *columnRepositoryMock) FindAllByProject(projectId int64, offset int64, limit int64) ([]*domains.ColumnModel, error) {
	return cr.FindAllByProjectFunc(projectId, offset, limit)
}

func (cr *columnRepositoryMock) FindAllByProjectAndName(projectId int64, name string) ([]*domains.ColumnModel, error) {
	return cr.FindAllByProjectAndNameFunc(projectId, name)
}

func (cr *columnRepositoryMock) Find(id int64) (*domains.ColumnModel, error) {
	return cr.FindFunc(id)
}

func (cr *columnRepositoryMock) FindByNextPosition(projectId int64, position int64) (*domains.ColumnModel, error) {
	return cr.FindByNextPositionFunc(projectId, position)
}

func (cr *columnRepositoryMock) FindByPreviousPosition(projectId int64, position int64) (*domains.ColumnModel, error) {
	return cr.FindByPreviousPositionFunc(projectId, position)
}

func (cr *columnRepositoryMock) FindWithMaxPosition(projectId int64) (*domains.ColumnModel, error) {
	return cr.FindWithMaxPositionFunc(projectId)
}

func (cr *columnRepositoryMock) Create(c *domains.ColumnModel) error {
	return cr.CreateFunc(c)
}

func (cr *columnRepositoryMock) Update(c *domains.ColumnModel) error {
	return cr.UpdateFunc(c)
}

func (cr *columnRepositoryMock) Delete(id int64) error {
	return cr.DeleteFunc(id)
}
