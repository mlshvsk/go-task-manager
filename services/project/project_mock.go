package project

import (
	"github.com/mlshvsk/go-task-manager/models"
)

type ServiceMock struct {
	GetProjectsFunc   func(page int64, limit int64) ([]*models.Project, error)
	GetProjectFunc    func(id int64) (*models.Project, error)
	StoreProjectFunc  func(p *models.Project) error
	UpdateProjectFunc func(p *models.Project) error
	DeleteProjectFunc func(id int64) error
}

func (s *ServiceMock) GetProjects(page int64, limit int64) ([]*models.Project, error) {
	return s.GetProjectsFunc(page, limit)
}

func (s *ServiceMock) GetProject(id int64) (*models.Project, error) {
	return s.GetProjectFunc(id)
}

func (s *ServiceMock) StoreProject(p *models.Project) error {
	return s.StoreProjectFunc(p)
}

func (s *ServiceMock) UpdateProject(p *models.Project) error {
	return s.UpdateProjectFunc(p)
}

func (s *ServiceMock) DeleteProject(id int64) error {
	return s.DeleteProjectFunc(id)
}
