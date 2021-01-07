package factories

import (
	"github.com/mlshvsk/go-task-manager/models"
	"time"
)

func TaskFactory(columnId int64, name string, description string, position int64) (models.Task, error) {
	t := new(models.Task)

	t.ColumnId = columnId
	t.Name = name
	t.Description = description
	t.Position = position
	t.CreatedAt = time.Now()

	return *t, nil
}
