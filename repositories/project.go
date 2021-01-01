package repositories

import (
	"database/sql"
	"fmt"
	"github.com/mlshvsk/go-task-manager/database"
	"github.com/mlshvsk/go-task-manager/models"
)

type projectRepository struct {
	mysqlRepository
}

const tableName = "projects"

var ProjectRepository *projectRepository

func InitProjectRepository(db *database.Sqldb) {
	ProjectRepository = &projectRepository{mysqlRepository{db, tableName, models.Project{}}}
}

func (r *projectRepository) FindAll() ([]*models.Project, error) {
	var projects = make([]*models.Project, 0)
	err := r.mysqlRepository.FindAll().Get(r.Scan(&projects))

	if err != nil {
		fmt.Printf("Error retrieving all projects: %v", err.Error())
		return nil, err
	}

	return projects, nil
}

func (r *projectRepository) Find(id int) (*models.Project, error) {
	var projects = make([]*models.Project, 0)
	err := r.mysqlRepository.Find(id).Get(r.Scan(&projects))

	if err != nil {
		fmt.Printf("Error retrieving all projects: %v", err.Error())
		return nil, err
	}

	return projects[0], nil
}

func (r *projectRepository) Create(p *models.Project) error {
	data := make(map[string]interface{})
	data["name"] = &p.Name
	data["description"] = &p.Description
	id, err := r.mysqlRepository.Create(data)

	if err != nil {
		return err
	}

	p.Id = id
	return nil
}

/*func (r *projectRepository) Scan(rows *sql.Rows) ([]*models.Project, error) {
	var projects []*models.Project

	for rows.Next() {
		project := new(models.Project)
		if err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.CreatedAt); err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	defer rows.Close()

	return projects, nil
}*/

func (r *projectRepository) Scan(projects *[]*models.Project) func(rows *sql.Rows) error {
	return func(rows *sql.Rows) error {
		for rows.Next() {
			project := new(models.Project)
			if err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.CreatedAt); err != nil {
				return err
			}

			*projects = append(*projects, project)
		}

		defer rows.Close()

		return nil
	}
}
