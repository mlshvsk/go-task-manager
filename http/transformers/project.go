package transformers

import (
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
)

type ExtendedProject struct {
	models.Project
	Columns []*models.Column `json:"columns"`
}

func ExtendProject(p *models.Project, includeColumns string) (interface{}, error) {
	if includeColumns == "true" {
		tc := &ExtendedProject{Project: *p}

		columns, err := services.ColumnService.GetColumns(p.Id, 0, -1)
		if err != nil {
			return nil, err
		}

		tc.Columns = columns

		return tc, nil
	}

	return p, nil
}

func ExtendProjects(c []*models.Project, includeTasks string) (interface{}, error) {
	tc := make([]interface{}, len(c))

	for i := range tc {
		res, err := ExtendProject(c[i], includeTasks)
		if err != nil {
			return nil, err
		}

		tc[i] = res
	}

	return tc, nil
}
