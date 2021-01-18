package transformers

import (
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
)

type ExtendedColumn struct {
	models.Column
	Tasks []*models.Task `json:"tasks"`
}

func ExtendColumn(c *models.Column, includeTasks string) (interface{}, error) {
	if includeTasks == "true" {
		tc := &ExtendedColumn{Column: *c}

		tasks, err := services.TaskService.GetTasksByColumn(c.Id, 0, -1)
		if err != nil {
			return nil, err
		}

		tc.Tasks = tasks

		return tc, nil
	}

	return c, nil
}

func ExtendColumns(c []*models.Column, includeTasks string) (interface{}, error) {
	tc := make([]interface{}, len(c))

	for i := range tc {
		res, err := ExtendColumn(c[i], includeTasks)
		if err != nil {
			return nil, err
		}

		tc[i] = res
	}

	return tc, nil
}
