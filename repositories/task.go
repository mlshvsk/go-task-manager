package repositories

import (
	"database/sql"
	"fmt"
	"github.com/mlshvsk/go-task-manager/database"
	"github.com/mlshvsk/go-task-manager/models"
)

type taskRepository struct {
	mysqlRepository
}

var TaskRepository *taskRepository

func InitTaskRepository(db *database.Sqldb) {
	TaskRepository = &taskRepository{mysqlRepository{db, "tasks", models.Task{}}}
}

func (tr *taskRepository) FindAll(columnId int) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0)

	err := tr.mysqlRepository.FindAllWhere([][]interface{}{{"column_id", "=", columnId}}).Get(tr.Scan(&tasks))

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tr *taskRepository) Find(id int) (*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.mysqlRepository.Find(id).Get(tr.Scan(&tasks))

	if err != nil {
		fmt.Printf("Error retrieving all projects: %v", err.Error())
		return nil, err
	}

	return tasks[0], nil
}

func (tr *taskRepository) FindByNextPosition(columnId int, position int) (*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.mysqlRepository.
		FindAllWhere([][]interface{}{{"column_id", "=", columnId}, {"position", ">", position}}).
		OrderBy("position", "asc").
		Get(tr.Scan(&tasks))

	if err != nil {
		fmt.Printf("Error retrieving all columns: %v", err.Error())
		return nil, err
	}

	if tasks == nil || len(tasks) == 0 {
		return nil, nil
	}

	return tasks[0], nil
}

func (tr *taskRepository) FindByPreviousPosition(columnId int, position int) (*models.Task, error) {
	tasks := make([]*models.Task, 0)
	err := tr.mysqlRepository.
		FindAllWhere([][]interface{}{{"column_id", "=", columnId}, {"position", "<", position}}).
		OrderBy("position", "desc").
		Get(tr.Scan(&tasks))

	if err != nil {
		fmt.Printf("Error retrieving all columns: %v", err.Error())
		return nil, err
	}

	if tasks == nil || len(tasks) == 0 {
		return nil, nil
	}

	return tasks[0], nil
}

func (tr *taskRepository) Create(t *models.Task) error {
	id, err := tr.mysqlRepository.Create(map[string]interface{}{"name": &t.Name, "description": &t.Description, "column_id": &t.ColumnId})

	if err != nil {
		return err
	}

	t.Id = id
	return nil
}

func (tr *taskRepository) Update(t *models.Task) error {
	err := tr.mysqlRepository.Update(t.Id, map[string]interface{}{"name": &t.Name, "description": &t.Description, "column_id": &t.ColumnId, "position": &t.Position})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (tr *taskRepository) Scan(tasks *[]*models.Task) func(rows *sql.Rows) error {
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
