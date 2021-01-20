package factories

import (
	"github.com/mlshvsk/go-task-manager/domains"
	"time"
)

func TaskFactory(columnId int64, name string, description string, position int64) (domains.TaskModel, error) {
	t := new(domains.TaskModel)

	t.ColumnId = columnId
	t.Name = name
	t.Description = description
	t.Position = position
	t.CreatedAt = time.Now()

	return *t, nil
}
