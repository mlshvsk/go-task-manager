package project

import (
	"github.com/mlshvsk/go-task-manager/domains"
)

type ServiceMock struct {
	GetProjectsFunc   func(page int64, limit int64) ([]*domains.ProjectModel, error)
	GetProjectFunc    func(id int64) (*domains.ProjectModel, error)
	StoreProjectFunc  func(p *domains.ProjectModel) error
	UpdateProjectFunc func(p *domains.ProjectModel) error
	DeleteProjectFunc func(id int64) error
}

func (s *ServiceMock) GetProjects(page int64, limit int64) ([]*domains.ProjectModel, error) {
	return s.GetProjectsFunc(page, limit)
}

func (s *ServiceMock) GetProject(id int64) (*domains.ProjectModel, error) {
	return s.GetProjectFunc(id)
}

func (s *ServiceMock) StoreProject(p *domains.ProjectModel) error {
	return s.StoreProjectFunc(p)
}

func (s *ServiceMock) UpdateProject(p *domains.ProjectModel) error {
	return s.UpdateProjectFunc(p)
}

func (s *ServiceMock) DeleteProject(id int64) error {
	return s.DeleteProjectFunc(id)
}
