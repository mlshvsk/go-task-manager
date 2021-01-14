package models

import "time"

type Task struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" validate:"required,min=1,max=500"`
	Description string    `json:"description" validate:"max=5000"`
	ColumnId    int64     `json:"column_id"`
	Position    int64     `json:"position"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskRepository interface {
	FindAll(offset int64, limit int64) ([]*Task, error)
	FindAllByColumn(columnId int64, offset int64, limit int64) ([]*Task, error)
	FindAllByColumnAndName(columnId int64, name string, offset int64, limit int64) ([]*Task, error)
	Find(id int64) (*Task, error)
	FindWithMaxPosition(columnId int64) (*Task, error)
	FindByNextPosition(columnId int64, position int64) (*Task, error)
	FindByPreviousPosition(columnId int64, position int64) (*Task, error)
	Create(t *Task) error
	Update(t *Task) error
	Delete(id int64) error
}

type TaskService interface {
	GetTasksByColumn(columnId int64, page int64, limit int64) ([]*Task, error)
	GetTasks(page int64, limit int64) ([]*Task, error)
	GetTask(id int64) (*Task, error)
	StoreTask(t *Task) error
	UpdateTask(t *Task) error
	DeleteTask(taskId int64) error
	MoveTaskWithinColumn(taskId int64, direction string) error
	MoveTaskToColumn(taskId int64, toColumnId int64) error
	MoveAllToColumn(fromColumn *Column, toColumn *Column) error
}
