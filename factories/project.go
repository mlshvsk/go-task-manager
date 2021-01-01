package factories

import (
	"github.com/mlshvsk/go-task-manager/models"
	"time"
)

func ProjectFactory(name string, description string) (*models.Project, error) {
	p := new(models.Project)

	p.Name = name
	p.Description = description
	p.CreatedAt = time.Now()

	return p, nil
}
