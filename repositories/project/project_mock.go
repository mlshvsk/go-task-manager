package project

import (
	"github.com/mlshvsk/go-task-manager/models"
)

type projectRepositoryMock struct {
	FindAllFunc       func(offset int64, limit int64) ([]*models.Project, error)
	FindAllByNameFunc func(name string) ([]*models.Project, error)
	FindFunc          func(id int64) (*models.Project, error)
	CreateFunc        func(p *models.Project) error
	UpdateFunc        func(p *models.Project) error
	DeleteFunc        func(id int64) error
}

func (r *projectRepositoryMock) FindAll(offset int64, limit int64) ([]*models.Project, error) {
	return r.FindAllFunc(offset, limit)
}

func (r *projectRepositoryMock) FindAllByName(name string) ([]*models.Project, error) {
	return r.FindAllByNameFunc(name)
}

func (r *projectRepositoryMock) Find(id int64) (*models.Project, error) {
	return r.FindFunc(id)
}

func (r *projectRepositoryMock) Create(p *models.Project) error {
	return r.CreateFunc(p)
}

func (r *projectRepositoryMock) Update(p *models.Project) error {
	return r.UpdateFunc(p)
}

func (r *projectRepositoryMock) Delete(id int64) error {
	return r.DeleteFunc(id)
}
