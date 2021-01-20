package domains

import "time"

type TaskModel struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name" validate:"required,min=1,max=500"`
	Description string    `json:"description" validate:"max=5000"`
	ColumnId    int64     `json:"column_id"`
	Position    int64     `json:"position"`
	CreatedAt   time.Time `json:"created_at"`
}

type TaskRepository interface {
	FindAll(offset int64, limit int64) ([]*TaskModel, error)
	FindAllByColumn(columnId int64, offset int64, limit int64) ([]*TaskModel, error)
	FindAllByColumnAndName(columnId int64, name string, offset int64, limit int64) ([]*TaskModel, error)
	Find(id int64) (*TaskModel, error)
	FindWithMaxPosition(columnId int64) (*TaskModel, error)
	FindByNextPosition(columnId int64, position int64) (*TaskModel, error)
	FindByPreviousPosition(columnId int64, position int64) (*TaskModel, error)
	Create(t *TaskModel) error
	Update(t *TaskModel) error
	Delete(id int64) error
}

type TaskService interface {
	GetTasksByColumn(columnId int64, page int64, limit int64) ([]*TaskModel, error)
	GetTasks(page int64, limit int64) ([]*TaskModel, error)
	GetTask(id int64) (*TaskModel, error)
	StoreTask(t *TaskModel) error
	UpdateTask(t *TaskModel) error
	DeleteTask(taskId int64) error
	MoveTaskWithinColumn(taskId int64, direction string) error
	MoveTaskToColumn(taskId int64, toColumnId int64) error
	MoveAllToColumn(fromColumn *ColumnModel, toColumn *ColumnModel) error
}
