package repositories

import (
	"database/sql"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories/base"
)

type taskRepository struct {
	base base.Repository
}

func InitTaskRepository(baseRepo base.Repository) *taskRepository {
	tr := &taskRepository{
		base: baseRepo,
	}

	tr.base.SetTableName("tasks")
	return tr
}

func (tr *taskRepository) FindAll(offset int64, limit int64) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.base.
		FindAll().
		OrderBy("position", "asc").
		Limit(offset, limit).
		Get(tr.scan(&tasks))

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr *taskRepository) FindAllByColumn(columnId int64, offset int64, limit int64) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.base.
		FindAll().
		Where("and", [][]interface{}{{"column_id", "=", columnId}}).
		Limit(offset, limit).
		Get(tr.scan(&tasks))

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr *taskRepository) FindAllByColumnAndName(columnId int64, name string, offset int64) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.base.
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
	err := tr.base.Find(id).Get(tr.scan(&tasks))

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
	err := tr.base.
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
	err := tr.base.
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
	err := tr.base.
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
	id, err := tr.base.Create(map[string]interface{}{
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
	err := tr.base.Update(t.Id, map[string]interface{}{
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

func (tr *taskRepository) Delete(id int64) error {
	return tr.base.Delete(id)
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
