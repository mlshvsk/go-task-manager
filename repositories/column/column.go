package column

import (
	"database/sql"
	"github.com/mlshvsk/go-task-manager/domains"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/repositories/base"
)

type columnRepository struct {
	base base.Repository
}

func InitColumnRepository(baseRepo base.Repository) *columnRepository {
	r := &columnRepository{
		base: baseRepo,
	}

	r.base.SetTableName("columns")
	return r
}

func (cr *columnRepository) FindAll(offset int64, limit int64) ([]*domains.ColumnModel, error) {
	var columns = make([]*domains.ColumnModel, 0)
	err := cr.base.
		FindAll().
		OrderBy("project_id", "asc").
		Limit(offset, limit).
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	return columns, nil
}

func (cr *columnRepository) FindAllByProject(projectId int64, offset int64, limit int64) ([]*domains.ColumnModel, error) {
	var columns = make([]*domains.ColumnModel, 0)
	err := cr.base.
		FindAll().
		Where("and", [][]interface{}{{"project_id", "=", projectId}}).
		OrderBy("position", "asc").
		Limit(offset, limit).
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	return columns, nil
}

func (cr *columnRepository) FindAllByProjectAndName(projectId int64, name string) ([]*domains.ColumnModel, error) {
	var columns = make([]*domains.ColumnModel, 0)
	err := cr.base.
		FindAll().
		Where("and", [][]interface{}{{"project_id", "=", projectId}, {"name", "=", name}}).
		OrderBy("position", "asc").
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	return columns, nil
}

func (cr *columnRepository) Find(id int64) (*domains.ColumnModel, error) {
	var columns = make([]*domains.ColumnModel, 0)
	err := cr.base.Find(id).Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	if len(columns) == 0 {
		return nil, &customErrors.NotFoundError{Value: "column not found"}
	}

	return columns[0], nil
}

func (cr *columnRepository) FindByNextPosition(projectId int64, position int64) (*domains.ColumnModel, error) {
	var columns = make([]*domains.ColumnModel, 0)

	err := cr.base.
		FindAll().
		Where("and", [][]interface{}{{"project_id", "=", projectId}, {"position", ">", position}}).
		OrderBy("position", "asc").
		Limit(0, 1).
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	if len(columns) == 0 {
		return nil, nil
	}

	return columns[0], nil
}

func (cr *columnRepository) FindByPreviousPosition(projectId int64, position int64) (*domains.ColumnModel, error) {
	var columns = make([]*domains.ColumnModel, 0)

	err := cr.base.
		FindAll().
		Where("and", [][]interface{}{{"project_id", "=", projectId}, {"position", "<", position}}).
		OrderBy("position", "desc").
		Limit(0, 1).
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	if len(columns) == 0 {
		return nil, nil
	}

	return columns[0], nil
}

func (cr *columnRepository) FindWithMaxPosition(projectId int64) (*domains.ColumnModel, error) {
	var columns = make([]*domains.ColumnModel, 0)
	err := cr.base.
		FindAll().
		Where("and", [][]interface{}{{"project_id", "=", projectId}}).
		OrderBy("position", "desc").
		Limit(0, 1).
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	if len(columns) == 0 {
		return nil, nil
	}

	return columns[0], nil
}

func (cr *columnRepository) Create(c *domains.ColumnModel) error {
	data := make(map[string]interface{})
	data["name"] = &c.Name
	data["project_id"] = &c.ProjectId
	data["position"] = &c.Position
	data["created_at"] = &c.CreatedAt

	id, err := cr.base.Create(data)
	c.Id = id

	return err
}

func (cr *columnRepository) Update(c *domains.ColumnModel) error {
	data := make(map[string]interface{})
	data["name"] = &c.Name
	data["project_id"] = &c.ProjectId
	data["position"] = &c.Position

	err := cr.base.Update(c.Id, data)

	return err
}

func (cr *columnRepository) Delete(id int64) error {
	return cr.base.Delete(id)
}

func (cr *columnRepository) scan(columns *[]*domains.ColumnModel) func(rows *sql.Rows) error {
	return func(rows *sql.Rows) error {
		for rows.Next() {
			column := new(domains.ColumnModel)
			if err := rows.Scan(&column.Id, &column.Name, &column.ProjectId, &column.Position, &column.CreatedAt); err != nil {
				return err
			}

			*columns = append(*columns, column)
		}

		defer rows.Close()

		return nil
	}
}
