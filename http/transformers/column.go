package transformers

import (
	"github.com/mlshvsk/go-task-manager/domains"
	"github.com/mlshvsk/go-task-manager/services"
)

type ExtendedColumn struct {
	domains.ColumnModel
	Tasks []*domains.TaskModel `json:"tasks"`
}

func ExtendColumn(c *domains.ColumnModel, includeTasks string) (interface{}, error) {
	if includeTasks == "true" {
		tc := &ExtendedColumn{ColumnModel: *c}

		tasks, err := services.TaskService.GetTasksByColumn(c.Id, 0, -1)
		if err != nil {
			return nil, err
		}

		tc.Tasks = tasks

		return tc, nil
	}

	return c, nil
}

func ExtendColumns(c []*domains.ColumnModel, includeTasks string) (interface{}, error) {
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
