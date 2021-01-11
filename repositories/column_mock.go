package repositories

import (
	"github.com/mlshvsk/go-task-manager/models"
)

type columnRepositoryMock struct {
	FinAllFunc                  func(offset int64, limit int64) ([]*models.Column, error)
	FindAllByProjectFunc        func(projectId int64, offset int64, limit int64) ([]*models.Column, error)
	FindAllByProjectAndNameFunc func(projectId int64, name string) ([]*models.Column, error)
	FindFunc                    func(id int64) (*models.Column, error)
	FindByNextPositionFunc      func(projectId int64, position int64) (*models.Column, error)
	FindByPreviousPositionFunc  func(projectId int64, position int64) (*models.Column, error)
	FindWithMaxPositionFunc     func(projectId int64) (*models.Column, error)
	CreateFunc                  func(c *models.Column) error
	UpdateFunc                  func(c *models.Column) error
	DeleteFunc                  func(id int64) error
}

func InitColumnRepositoryMock() *columnRepositoryMock {
	return &columnRepositoryMock{}
}

func (cr *columnRepositoryMock) FindAll(offset int64, limit int64) ([]*models.Column, error) {
	return cr.FinAllFunc(offset, limit)
}

func (cr *columnRepositoryMock) FindAllByProject(projectId int64, offset int64, limit int64) ([]*models.Column, error) {
	return cr.FindAllByProjectFunc(projectId, offset, limit)
}

func (cr *columnRepositoryMock) FindAllByProjectAndName(projectId int64, name string) ([]*models.Column, error) {
	return cr.FindAllByProjectAndNameFunc(projectId, name)
}

func (cr *columnRepositoryMock) Find(id int64) (*models.Column, error) {
	return cr.FindFunc(id)
}

func (cr *columnRepositoryMock) FindByNextPosition(projectId int64, position int64) (*models.Column, error) {
	return cr.FindByNextPositionFunc(projectId, position)
}

func (cr *columnRepositoryMock) FindByPreviousPosition(projectId int64, position int64) (*models.Column, error) {
	return cr.FindByPreviousPositionFunc(projectId, position)
}

func (cr *columnRepositoryMock) FindWithMaxPosition(projectId int64) (*models.Column, error) {
	return cr.FindWithMaxPositionFunc(projectId)
}

func (cr *columnRepositoryMock) Create(c *models.Column) error {
	return cr.CreateFunc(c)
}

func (cr *columnRepositoryMock) Update(c *models.Column) error {
	return cr.UpdateFunc(c)
}

func (cr *columnRepositoryMock) Delete(id int64) error {
	return cr.DeleteFunc(id)
}
