package factories

import (
	"github.com/mlshvsk/go-task-manager/domains"
	"time"
)

func ProjectFactory(name string, description string) (domains.ProjectModel, error) {
	p := new(domains.ProjectModel)

	p.Name = name
	p.Description = description
	p.CreatedAt = time.Now()

	return *p, nil
}
