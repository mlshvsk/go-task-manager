package factories

import (
	"github.com/mlshvsk/go-task-manager/models"
	"time"
)

func CommentFactory(taskId int64, data string) (models.Comment, error) {
	c := new(models.Comment)

	c.TaskId = taskId
	c.Data = data
	c.CreatedAt = time.Now()

	return *c, nil
}
