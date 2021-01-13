package column

import (
	"errors"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/factories"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services/task"
	"sync"
)

type columnService struct {
	r models.ColumnRepository
}

var Service models.ColumnService

func InitColumnService(r models.ColumnRepository) {
	(&sync.Once{}).Do(func() {
		Service = &columnService{r}
	})
}

func (s *columnService) GetColumns(projectId int64, page int64, limit int64) ([]*models.Column, error) {
	return s.r.FindAllByProject(projectId, page, limit)
}

func (s *columnService) GetColumn(columnId int64) (*models.Column, error) {
	return s.r.Find(columnId)
}

func (s *columnService) StoreColumn(c *models.Column) error {
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

func (s *columnService) UpdateColumn(c *models.Column) error {
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

	//TODO: rewrite it, its bullshit
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

	if err := task.Service.MoveAllToColumn(column, nextColumn); err != nil {
		return err
	}

	return s.r.Delete(columnId)
}

func (s *columnService) MoveColumn(columnId int64, direction string) error {
	nextColumn := new(models.Column)
	column, err := s.r.Find(columnId)
	if err != nil {
		return err
	}

	if direction == "right" {
		nextColumn, err = s.r.FindByNextPosition(column.ProjectId, column.Position)
		if err != nil {
			return err
		} else if nextColumn == nil {
			return nil
		}

		nextColumn.Position--
		column.Position++
	} else if direction == "left" {
		nextColumn, err = s.r.FindByPreviousPosition(column.ProjectId, column.Position)
		if err != nil {
			return err
		} else if nextColumn == nil {
			return nil
		}

		nextColumn.Position++
		column.Position--
	} else {
		return errors.New("invalid direction: " + direction)
	}

	if err := s.r.Update(nextColumn); err != nil {
		return err
	}

	return s.r.Update(column)
}
