package models

import (
	"time"
)

type Column struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name,omitempty" validate:"required,max=255"`
	ProjectId int64     `json:"project_id"`
	Position  int64     `json:"position"`
	CreatedAt time.Time `json:"created_at"`
}

type ColumnRepository interface {
	FindAll(offset int64, limit int64) ([]*Column, error)
	FindAllByProject(projectId int64, offset int64, limit int64) ([]*Column, error)
	FindAllByProjectAndName(projectId int64, name string) ([]*Column, error)
	Find(id int64) (*Column, error)
	FindByNextPosition(projectId int64, position int64) (*Column, error)
	FindByPreviousPosition(projectId int64, position int64) (*Column, error)
	FindWithMaxPosition(projectId int64) (*Column, error)
	Create(c *Column) error
	Update(c *Column) error
	Delete(id int64) error
}

type ColumnService interface {
	GetColumns(projectId int64, page int64, limit int64) ([]*Column, error)
	GetColumn(columnId int64) (*Column, error)
	StoreColumn(c *Column) error
	UpdateColumn(c *Column) error
	DeleteColumn(columnId int64) error
	MoveColumn(columnId int64, direction string) error
}
