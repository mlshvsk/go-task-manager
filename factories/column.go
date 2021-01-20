package factories

import (
	"github.com/mlshvsk/go-task-manager/domains"
	"time"
)

func ColumnFactory(projectId int64, name string, position int64) (domains.ColumnModel, error) {
	c := new(domains.ColumnModel)

	c.ProjectId = projectId
	c.Name = name
	c.CreatedAt = time.Now()
	c.Position = position

	return *c, nil
}
