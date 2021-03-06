package project

import (
	"database/sql"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/domains"
	"github.com/mlshvsk/go-task-manager/repositories/base"
)

type projectRepository struct {
	base base.Repository
}

func InitProjectRepository(baseRepo base.Repository) *projectRepository {
	pr := &projectRepository{
		base: baseRepo,
	}

	pr.base.SetTableName("projects")
	return pr
}

func (r *projectRepository) FindAll(offset int64, limit int64) ([]*domains.ProjectModel, error) {
	var projects = make([]*domains.ProjectModel, 0)
	err := r.base.
		FindAll().
		OrderBy("name", "asc").
		Limit(offset, limit).
		Get(r.scan(&projects))

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectRepository) FindAllByName(name string) ([]*domains.ProjectModel, error) {
	var projects = make([]*domains.ProjectModel, 0)
	err := r.base.
		FindAll().
		Where("and", [][]interface{}{{"name", "=", name}}).
		OrderBy("id", "asc").
		Get(r.scan(&projects))

	if err != nil {
		return nil, err
	}

	return projects, nil
}

func (r *projectRepository) Find(id int64) (*domains.ProjectModel, error) {
	var projects = make([]*domains.ProjectModel, 0)

	if err := r.base.Find(id).Get(r.scan(&projects)); err != nil {
		return nil, err
	}

	if len(projects) == 0 {
		return nil, &customErrors.NotFoundError{Value: "project not found"}
	}

	return projects[0], nil
}

func (r *projectRepository) Create(p *domains.ProjectModel) error {
	data := make(map[string]interface{})
	data["name"] = &p.Name
	data["description"] = &p.Description
	data["created_at"] = &p.CreatedAt

	id, err := r.base.Create(data)
	if err != nil {
		return err
	}

	p.Id = id
	return nil
}

func (r *projectRepository) Update(p *domains.ProjectModel) error {
	data := make(map[string]interface{})
	data["name"] = &p.Name
	data["description"] = &p.Description

	if err := r.base.Update(p.Id, data); err != nil {
		return err
	}

	return nil
}

func (r *projectRepository) Delete(id int64) error {
	return r.base.Delete(id)
}

func (r *projectRepository) scan(projects *[]*domains.ProjectModel) func(rows *sql.Rows) error {
	return func(rows *sql.Rows) error {
		for rows.Next() {
			project := new(domains.ProjectModel)
			if err := rows.Scan(&project.Id, &project.Name, &project.Description, &project.CreatedAt); err != nil {
				return err
			}

			*projects = append(*projects, project)
		}

		defer rows.Close()

		return nil
	}
}
