package column

import (
	"errors"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/factories"
	"github.com/mlshvsk/go-task-manager/domains"
	"github.com/mlshvsk/go-task-manager/services"
	"sync"
)

type columnService struct {
	r domains.ColumnRepository
}

const (
	directionRight = "right"
	directionLeft = "left"
)

func InitColumnService(r domains.ColumnRepository) {
	(&sync.Once{}).Do(func() {
		services.ColumnService = &columnService{r}
	})
}

func (s *columnService) GetColumns(projectId int64, page int64, limit int64) ([]*domains.ColumnModel, error) {
	return s.r.FindAllByProject(projectId, page, limit)
}

func (s *columnService) GetColumn(columnId int64) (*domains.ColumnModel, error) {
	return s.r.Find(columnId)
}

func (s *columnService) StoreColumn(c *domains.ColumnModel) error {
	columns, err := s.r.FindAllByProjectAndName(c.ProjectId, c.Name)
	if err != nil {
		return err
	}
	if columns != nil && len(columns) > 0 {
		return &customErrors.ModelAlreadyExists{}
	}

	prevCol, err := s.r.FindWithMaxPosition(c.ProjectId)
	if err != nil {
		return err
	} else if prevCol == nil {
		c.Position = 0
	} else {
		c.Position = prevCol.Position + 1
	}

	*c, err = factories.ColumnFactory(c.ProjectId, c.Name, c.Position)
	if err != nil {
		return err
	}

	return s.r.Create(c)
}

func (s *columnService) UpdateColumn(c *domains.ColumnModel) error {
	columnFromDB, err := s.r.Find(c.Id)
	if err != nil {
		return err
	}

	c.ProjectId = columnFromDB.ProjectId
	c.Position = columnFromDB.Position
	c.CreatedAt = columnFromDB.CreatedAt
	return s.r.Update(c)
}

func (s *columnService) DeleteColumn(columnId int64) error {
	column, err := s.r.Find(columnId)
	if err != nil {
		return err
	}

	columns, err := s.r.FindAllByProject(column.ProjectId, 0, -1)
	if err != nil {
		return err
	} else if columns == nil || len(columns) <= 1 {
		return errors.New("cannot delete last column")
	}

	nextColumn, err := s.r.FindByNextPosition(column.ProjectId, column.Position)
	if err != nil {
		return err
	} else if nextColumn == nil {
		nextColumn, err = s.r.FindByPreviousPosition(column.ProjectId, column.Position)
		if err != nil {
			return err
		} else if nextColumn == nil {
			return errors.New("cannot find column to move tasks to before deletion")
		}
	}

	if err := services.TaskService.MoveAllToColumn(column, nextColumn); err != nil {
		return err
	}

	return s.r.Delete(columnId)
}

func (s *columnService) MoveColumn(columnId int64, direction string) error {
	nextColumn := new(domains.ColumnModel)
	column, err := s.r.Find(columnId)
	if err != nil {
		return err
	}

	if direction == directionRight {
		nextColumn, err = s.r.FindByNextPosition(column.ProjectId, column.Position)
		if err != nil {
			return err
		} else if nextColumn == nil {
			return nil
		}

		nextColumn.Position, column.Position = column.Position, nextColumn.Position
	} else if direction == directionLeft {
		nextColumn, err = s.r.FindByPreviousPosition(column.ProjectId, column.Position)
		if err != nil {
			return err
		} else if nextColumn == nil {
			return nil
		}

		nextColumn.Position, column.Position = column.Position, nextColumn.Position
	} else {
		return errors.New("invalid direction: " + direction)
	}

	if err := s.r.Update(nextColumn); err != nil {
		return err
	}

	return s.r.Update(column)
}
