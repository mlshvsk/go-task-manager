package factories

import (
	"github.com/mlshvsk/go-task-manager/domains"
	"time"
)

func CommentFactory(taskId int64, data string) (domains.CommentModel, error) {
	c := new(domains.CommentModel)

	c.TaskId = taskId
	c.Data = data
	c.CreatedAt = time.Now()

	return *c, nil
}
