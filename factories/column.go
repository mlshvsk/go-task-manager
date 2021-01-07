package factories

import (
	"github.com/mlshvsk/go-task-manager/models"
	"time"
)

func ColumnFactory(projectId int64, name string, position int64) (models.Column, error) {
	c := new(models.Column)

	c.ProjectId = projectId
	c.Name = name
	c.CreatedAt = time.Now()
	c.Position = position

	return *c, nil
}
