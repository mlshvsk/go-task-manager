package repositories

import (
	"database/sql"
	"fmt"
	"github.com/mlshvsk/go-task-manager/database"
	"github.com/mlshvsk/go-task-manager/models"
)

type columnRepository struct {
	mysqlRepository
}

var ColumnRepository *columnRepository

func InitColumnRepository(db *database.Sqldb) {
	ColumnRepository = &columnRepository{mysqlRepository{db, "columns", models.Column{}}}
}

func (cr *columnRepository) FindAll(projectId int) ([]*models.Column, error) {
	var columns = make([]*models.Column, 0)
	err := cr.mysqlRepository.FindAllWhere([][]interface{}{{"project_id", "=", projectId}}).Get(cr.Scan(&columns))

	if err != nil {
		fmt.Printf("Error retrieving all columns: %v", err.Error())
		return nil, err
	}

	return columns, nil
}

func (cr *columnRepository) Find(id int) (*models.Column, error) {
	var columns = make([]*models.Column, 0)
	err := cr.mysqlRepository.Find(id).Get(cr.Scan(&columns))

	if err != nil {
		fmt.Printf("Error retrieving all columns: %v", err.Error())
		return nil, err
	}

	return columns[0], nil
}

func (cr *columnRepository) FindByNextPosition(projectId int, position int) (*models.Column, error) {
	var columns = make([]*models.Column, 0)

	err := cr.mysqlRepository.
		FindAllWhere([][]interface{}{{"project_id", "=", projectId}, {"position", ">", position}}).
		OrderBy("position", "asc").
		Get(cr.Scan(&columns))

	if err != nil {
		fmt.Printf("Error retrieving all columns: %v", err.Error())
		return nil, err
	}

	if columns == nil || len(columns) == 0 {
		return nil, nil
	}

	return columns[0], nil
}

func (cr *columnRepository) FindByPreviousPosition(projectId int, position int) (*models.Column, error) {
	var columns = make([]*models.Column, 0)

	err := cr.mysqlRepository.
		FindAllWhere([][]interface{}{{"project_id", "=", projectId}, {"position", "<", position}}).
		OrderBy("position", "desc").
		Get(cr.Scan(&columns))


	if err != nil {
		fmt.Printf("Error scanning all columns: %v", err.Error())
		return nil, err
	}

	if columns == nil || len(columns) == 0 {
		return nil, nil
	}

	return columns[0], nil
}

func (cr *columnRepository) FindWithMaxPosition(projectId int) (*models.Column, error) {
	var columns = make([]*models.Column, 0)
	err := cr.mysqlRepository.
		FindAllWhere([][]interface{}{{"project_id", "=", projectId}}).
		OrderBy("position", "desc").
		Get(cr.Scan(&columns))

	if err != nil {
		fmt.Printf("Error retrieving all columns: %v", err.Error())
		return nil, err
	}

	return columns[0], nil
}

func (cr *columnRepository) Create(c *models.Column) error {
	id, err := cr.mysqlRepository.Create(map[string]interface{}{"name": &c.Name, "project_id": &c.ProjectId, "position": &c.Position})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	c.Id = id
	return nil
}

func (cr *columnRepository) Update(c *models.Column) error {
	err := cr.mysqlRepository.Update(c.Id, map[string]interface{}{"name": &c.Name, "project_id": &c.ProjectId, "position": &c.Position})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

/*func (cr *columnRepository) Scan(rows *sql.Rows) ([]*models.Column, error) {
	var columns []*models.Column

	for rows.Next() {
		column := new(models.Column)
		if err := rows.Scan(&column.Id, &column.Name, &column.ProjectId, &column.Position, &column.CreatedAt); err != nil {
			return nil, err
		}

		columns = append(columns, column)
	}

	defer rows.Close()

	return columns, nil
}*/

func (cr *columnRepository) Scan(columns *[]*models.Column) func(rows *sql.Rows) error {
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
