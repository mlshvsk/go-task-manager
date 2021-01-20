package domains

import (
	"time"
)

type ColumnModel struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name,omitempty" validate:"required,max=255"`
	ProjectId int64     `json:"project_id"`
	Position  int64     `json:"position"`
	CreatedAt time.Time `json:"created_at"`
}

type ColumnRepository interface {
	FindAll(offset int64, limit int64) ([]*ColumnModel, error)
	FindAllByProject(projectId int64, offset int64, limit int64) ([]*ColumnModel, error)
	FindAllByProjectAndName(projectId int64, name string) ([]*ColumnModel, error)
	Find(id int64) (*ColumnModel, error)
	FindByNextPosition(projectId int64, position int64) (*ColumnModel, error)
	FindByPreviousPosition(projectId int64, position int64) (*ColumnModel, error)
	FindWithMaxPosition(projectId int64) (*ColumnModel, error)
	Create(c *ColumnModel) error
	Update(c *ColumnModel) error
	Delete(id int64) error
}

type ColumnService interface {
	GetColumns(projectId int64, page int64, limit int64) ([]*ColumnModel, error)
	GetColumn(columnId int64) (*ColumnModel, error)
	StoreColumn(c *ColumnModel) error
	UpdateColumn(c *ColumnModel) error
	DeleteColumn(columnId int64) error
	MoveColumn(columnId int64, direction string) error
}
