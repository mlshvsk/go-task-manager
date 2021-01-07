package repositories

import (
	"database/sql"
	"github.com/mlshvsk/go-task-manager/database"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories/mysql"
)

type projectRepository struct {
	mysql.Repository
}

var ProjectRepository *projectRepository

func InitProjectRepository(db *database.SqlDB) {
	ProjectRepository = &projectRepository{
		Repository: mysql.Repository{
			SqlDB:     db,
			TableName: "projects",
		},
	}
}

func (r *projectRepository) FindAll() ([]*models.Project, error) {
	var projects = make([]*models.Project, 0)
	err := r.Repository.
		FindAll().
		OrderBy("name", "asc").
		Get(r.scan(&projects))

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectRepository) FindAllByName(name string) ([]*models.Project, error) {
	var projects = make([]*models.Project, 0)
	err := r.Repository.
		FindAll().
		Where("and", [][]interface{}{{"name", "=", name}}).
		OrderBy("name", "asc").
		Get(r.scan(&projects))

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectRepository) Find(id int64) (*models.Project, error) {
	var projects = make([]*models.Project, 0)

	if err := r.Repository.Find(id).Get(r.scan(&projects)); err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, &customErrors.NotFoundError{Value: "project not found"}
	}

	return projects[0], nil
}

func (r *projectRepository) Create(p *models.Project) error {
	data := make(map[string]interface{})
	data["name"] = &p.Name
	data["description"] = &p.Description
	data["created_at"] = &p.CreatedAt

	id, err := r.Repository.Create(data)

	if err != nil {
		return err
	}

	p.Id = id

	return nil
}

func (r *projectRepository) Update(p *models.Project) error {
	data := make(map[string]interface{})
	data["name"] = &p.Name
	data["description"] = &p.Description

	if err := r.Repository.Update(p.Id, data); err != nil {
		return err
	}

	return nil
}

func (r *projectRepository) scan(projects *[]*models.Project) func(rows *sql.Rows) error {
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
