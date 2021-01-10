package services

import (
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/factories"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories"
)

type ProjectService struct {
}

func GetProjects(page int64, limit int64) ([]*models.Project, error) {
	return repositories.ProjectRepository.FindAll(page, limit)
}

func GetProject(id int64) (*models.Project, error) {
	return repositories.ProjectRepository.Find(id)
}

func StoreProject(p *models.Project) error {
	projects, err := repositories.ProjectRepository.FindAllByName(p.Name)
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

	if err := repositories.ProjectRepository.Create(p); err != nil {
		return err
	}

	column, err := factories.ColumnFactory(p.Id, "New", 0)
	if err != nil {
		return err
	}

	return repositories.ColumnRepository.Create(&column)
}

func UpdateProject(p *models.Project) error {
	projectFromDB, err := repositories.ProjectRepository.Find(p.Id)
	if err != nil {
		return err
	}

	p.CreatedAt = projectFromDB.CreatedAt

	return repositories.ProjectRepository.Update(p)
}

func DeleteProject(id int64) error {
	if _, err := repositories.ProjectRepository.Find(id); err != nil {
		return err
	}

	return repositories.ProjectRepository.Delete(id)
}
