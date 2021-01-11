package models

import "time"

type Project struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name,omitempty" validate:"required,max=500"`
	Description string    `json:"description,omitempty" validate:"max=1000"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProjectRepository interface {
	FindAll(offset int64, limit int64) ([]*Project, error)
	FindAllByName(name string) ([]*Project, error)
	Find(id int64) (*Project, error)
	Create(p *Project) error
	Update(p *Project) error
	Delete(id int64) error
}

type ProjectService interface {
	GetProjects(page int64, limit int64) ([]*Project, error)
	GetProject(id int64) (*Project, error)
	StoreProject(p *Project) error
	UpdateProject(p *Project) error
	DeleteProject(id int64) error
}
