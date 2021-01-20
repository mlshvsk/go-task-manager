package project

import (
	"github.com/mlshvsk/go-task-manager/domains"
)

type projectRepositoryMock struct {
	FindAllFunc       func(offset int64, limit int64) ([]*domains.ProjectModel, error)
	FindAllByNameFunc func(name string) ([]*domains.ProjectModel, error)
	FindFunc          func(id int64) (*domains.ProjectModel, error)
	CreateFunc        func(p *domains.ProjectModel) error
	UpdateFunc        func(p *domains.ProjectModel) error
	DeleteFunc        func(id int64) error
}

func InitProjectRepositoryMock() *projectRepositoryMock {
	return &projectRepositoryMock{}
}

func (r *projectRepositoryMock) FindAll(offset int64, limit int64) ([]*domains.ProjectModel, error) {
	return r.FindAllFunc(offset, limit)
}

func (r *projectRepositoryMock) FindAllByName(name string) ([]*domains.ProjectModel, error) {
	return r.FindAllByNameFunc(name)
}

func (r *projectRepositoryMock) Find(id int64) (*domains.ProjectModel, error) {
	return r.FindFunc(id)
}

func (r *projectRepositoryMock) Create(p *domains.ProjectModel) error {
	return r.CreateFunc(p)
}

func (r *projectRepositoryMock) Update(p *domains.ProjectModel) error {
	return r.UpdateFunc(p)
}

func (r *projectRepositoryMock) Delete(id int64) error {
	return r.DeleteFunc(id)
}
