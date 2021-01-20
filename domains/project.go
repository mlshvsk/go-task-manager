package domains

import "time"

type ProjectModel struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name,omitempty" validate:"required,max=500"`
	Description string    `json:"description,omitempty" validate:"max=1000"`
	CreatedAt   time.Time `json:"created_at"`
}

type ProjectRepository interface {
	FindAll(offset int64, limit int64) ([]*ProjectModel, error)
	FindAllByName(name string) ([]*ProjectModel, error)
	Find(id int64) (*ProjectModel, error)
	Create(p *ProjectModel) error
	Update(p *ProjectModel) error
	Delete(id int64) error
}

type ProjectService interface {
	GetProjects(page int64, limit int64) ([]*ProjectModel, error)
	GetProject(id int64) (*ProjectModel, error)
	StoreProject(p *ProjectModel) error
	UpdateProject(p *ProjectModel) error
	DeleteProject(id int64) error
}
