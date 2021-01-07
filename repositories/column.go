package repositories

import (
	"database/sql"
	"fmt"
	"github.com/mlshvsk/go-task-manager/database"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories/mysql"
)

type columnRepository struct {
	mysql.Repository
}

var ColumnRepository *columnRepository

func InitColumnRepository(db *database.SqlDB) {
	ColumnRepository = &columnRepository{
		Repository: mysql.Repository{
			SqlDB:     db,
			TableName: "columns",
		},
	}
}

func (cr *columnRepository) FindAll() ([]*models.Column, error) {
	var columns = make([]*models.Column, 0)
	err := cr.Repository.
		FindAll().
		OrderBy("project_id", "asc").
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	return columns, nil
}

func (cr *columnRepository) FindAllByProject(projectId int64) ([]*models.Column, error) {
	var columns = make([]*models.Column, 0)
	err := cr.Repository.
		FindAll().
		Where("and", [][]interface{}{{"project_id", "=", projectId}}).
		OrderBy("position", "asc").
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	return columns, nil
}

func (cr *columnRepository) FindAllByProjectAndName(projectId int64, name string) ([]*models.Column, error) {
	var columns = make([]*models.Column, 0)
	err := cr.Repository.
		FindAll().
		Where("and", [][]interface{}{{"project_id", "=", projectId}, {"name", "=", name}}).
		OrderBy("position", "asc").
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	return columns, nil
}

func (cr *columnRepository) Find(id int64) (*models.Column, error) {
	var columns = make([]*models.Column, 0)
	err := cr.Repository.Find(id).Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	if columns == nil || len(columns) == 0 {
		return nil, &customErrors.NotFoundError{Value: "column not found"}
	}

	return columns[0], nil
}

func (cr *columnRepository) FindByNextPosition(projectId int64, position int64) (*models.Column, error) {
	var columns = make([]*models.Column, 0)

	err := cr.Repository.
		FindAll().
		Where("and", [][]interface{}{{"project_id", "=", projectId}, {"position", ">", position}}).
		OrderBy("position", "asc").
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	if columns == nil || len(columns) == 0 {
		return nil, nil
	}

	return columns[0], nil
}

func (cr *columnRepository) FindByPreviousPosition(projectId int64, position int64) (*models.Column, error) {
	var columns = make([]*models.Column, 0)

	err := cr.Repository.
		FindAll().
		Where("and", [][]interface{}{{"project_id", "=", projectId}, {"position", "<", position}}).
		OrderBy("position", "desc").
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	if columns == nil || len(columns) == 0 {
		return nil, nil
	}

	return columns[0], nil
}

func (cr *columnRepository) FindWithMaxPosition(projectId int64) (*models.Column, error) {
	var columns = make([]*models.Column, 0)
	err := cr.Repository.
		FindAll().
		Where("and", [][]interface{}{{"project_id", "=", projectId}}).
		OrderBy("position", "desc").
		Get(cr.scan(&columns))

	if err != nil {
		return nil, err
	}

	return columns[0], nil
}

func (cr *columnRepository) Create(c *models.Column) error {
	data := make(map[string]interface{})
	data["name"] = &c.Name
	data["project_id"] = &c.ProjectId
	data["position"] = &c.Position
	data["created_at"] = &c.CreatedAt

	id, err := cr.Repository.Create(data)
	c.Id = id

	return err
}

func (cr *columnRepository) Update(c *models.Column) error {
	data := make(map[string]interface{})
	data["name"] = &c.Name
	data["project_id"] = &c.ProjectId
	data["position"] = &c.Position

	err := cr.Repository.Update(c.Id, data)

	return err
}

func (cr *columnRepository) scan(columns *[]*models.Column) func(rows *sql.Rows) error {
	return func(rows *sql.Rows) error {
		for rows.Next() {
			column := new(models.Column)
			if err := rows.Scan(&column.Id, &column.Name, &column.ProjectId, &column.Position, &column.CreatedAt); err != nil {
				fmt.Println(err.Error())
				return err
			}

			*columns = append(*columns, column)
		}

		defer rows.Close()

		return nil
	}
}
