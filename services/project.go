package services

import (
	"fmt"
	"github.com/mlshvsk/go-task-manager/factories"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories"
)

type ProjectService struct {
}

func GetProjects() ([]*models.Project, error) {
	return repositories.ProjectRepository.FindAll()
}

func GetProject(id int) (*models.Project, error) {
	return repositories.ProjectRepository.Find(id)
}

func StoreProject(p *models.Project) *models.Project {
	_ = repositories.ProjectRepository.Create(p)

	c, _ := factories.ColumnFactory(p.Id, "New", 0)
	fmt.Println(c)
	_ = repositories.ColumnRepository.Create(c)

	return p
}

func UpdateProject(p *models.Project) *models.Project {
	err := repositories.ProjectRepository.Update(p.Id, map[string]interface{}{"name": p.Name, "description": p.Description})

	if err != nil {
		fmt.Println(err.Error())
	}

	return p
}

func DeleteProject(id int) {
	repositories.ProjectRepository.Delete(id)
}
