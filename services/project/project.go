package project

import (
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/factories"
	"github.com/mlshvsk/go-task-manager/models"
	column2 "github.com/mlshvsk/go-task-manager/services/column"
	"sync"
)

type projectService struct {
	r models.ProjectRepository
}

var Service models.ProjectService

func InitProjectService(r models.ProjectRepository) {
	(&sync.Once{}).Do(func() {
		Service = &projectService{r}
	})
}

func (s *projectService) GetProjects(page int64, limit int64) ([]*models.Project, error) {
	return s.r.FindAll(page, limit)
}

func (s *projectService) GetProject(id int64) (*models.Project, error) {
	return s.r.Find(id)
}

func (s *projectService) StoreProject(p *models.Project) error {
	projects, err := s.r.FindAllByName(p.Name)
	if err != nil {
		return err
	}
	if projects != nil && len(projects) > 0 {
		return &customErrors.ModelAlreadyExists{}
	}

	*p, err = factories.ProjectFactory(p.Name, p.Description)
	if err != nil {
		return err
	}

	if err := s.r.Create(p); err != nil {
		return err
	}

	column, err := factories.ColumnFactory(p.Id, "New", 0)
	if err != nil {
		return err
	}

	return column2.Service.StoreColumn(&column)
}

func (s *projectService) UpdateProject(p *models.Project) error {
	projectFromDB, err := s.r.Find(p.Id)
	if err != nil {
		return err
	}

	p.CreatedAt = projectFromDB.CreatedAt

	return s.r.Update(p)
}

func (s *projectService) DeleteProject(id int64) error {
	if _, err := s.r.Find(id); err != nil {
		return err
	}

	return s.r.Delete(id)
}