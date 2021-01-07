package repositories

import (
	"database/sql"
	"github.com/mlshvsk/go-task-manager/database"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories/mysql"
)

type taskRepository struct {
	mysql.Repository
}

var (
	TaskRepository *taskRepository
)

func InitTaskRepository(db *database.SqlDB) {
	TaskRepository = &taskRepository{
		Repository: mysql.Repository{
			SqlDB: db,
			TableName: "tasks",
		},
	}
}

func (tr *taskRepository) FindAll() ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.Repository.
		FindAll().
		OrderBy("position", "asc").
		Get(tr.scan(&tasks))

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr *taskRepository) FindAllByColumn(columnId int64) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.Repository.
		FindAll().
		Where("and", [][]interface{}{{"column_id", "=", columnId}}).
		Get(tr.scan(&tasks))

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr *taskRepository) FindAllByColumnAndName(columnId int64, name string) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.Repository.
		FindAll().
		Where("and", [][]interface{}{{"column_id", "=", columnId}, {"name", "=", name}}).
		Get(tr.scan(&tasks))

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr *taskRepository) Find(id int64) (*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.Repository.Find(id).Get(tr.scan(&tasks))

	if err != nil {
		return nil, err
	}

	if tasks == nil || len(tasks) == 0 {
		return nil, &customErrors.NotFoundError{Value: "task not found"}
	}

	return tasks[0], nil
}

func (tr *taskRepository) FindWithMaxPosition(columnId int64) (*models.Task, error) {
	var tasks = make([]*models.Task, 0)
	err := tr.Repository.
		FindAll().
		Where("and", [][]interface{}{{"column_id", "=", columnId}}).
		OrderBy("position", "desc").
		Get(tr.scan(&tasks))

	if err != nil {
		return nil, err
	}

	if tasks == nil || len(tasks) == 0 {
		return nil, nil
	}

	return tasks[0], nil
}

func (tr *taskRepository) FindByNextPosition(columnId int64, position int64) (*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.Repository.
		FindAll().Where("and", [][]interface{}{{"column_id", "=", columnId}, {"position", ">", position}}).
		OrderBy("position", "asc").
		Get(tr.scan(&tasks))

	if err != nil {
		return nil, err
	}

	if tasks == nil || len(tasks) == 0 {
		return nil, nil
	}

	return tasks[0], nil
}

func (tr *taskRepository) FindByPreviousPosition(columnId int64, position int64) (*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.Repository.
		FindAll().Where("and", [][]interface{}{{"column_id", "=", columnId}, {"position", "<", position}}).
		OrderBy("position", "desc").
		Get(tr.scan(&tasks))

	if err != nil {
		return nil, err
	}

	if tasks == nil || len(tasks) == 0 {
		return nil, nil
	}

	return tasks[0], nil
}

func (tr *taskRepository) Create(t *models.Task) error {
	id, err := tr.Repository.Create(map[string]interface{}{
		"name": t.Name,
		"description": t.Description,
		"column_id": t.ColumnId,
		"position": t.Position,
	})

	if err != nil {
		return err
	}

	t.Id = id
	return nil
}

func (tr *taskRepository) Update(t *models.Task) error {
	err := tr.Repository.Update(t.Id, map[string]interface{}{
		"name": &t.Name,
		"description": &t.Description,
		"column_id": &t.ColumnId,
		"position": &t.Position,
	})

	if err != nil {
		return err
	}

	return nil
}

func (tr *taskRepository) scan(tasks *[]*models.Task) func(rows *sql.Rows) error {
	return func(rows *sql.Rows) error {
		for rows.Next() {
			task := new(models.Task)
			if err := rows.Scan(&task.Id, &task.Name, &task.Description, &task.ColumnId, &task.Position, &task.CreatedAt); err != nil {
				return err
			}

			*tasks = append(*tasks, task)
		}

		defer rows.Close()

		return nil
	}
}
