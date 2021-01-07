package services

import (
	"errors"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/factories"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories"
)

type ColumnService struct {
}

func GetColumns(projectId int64) ([]*models.Column, error) {
	return repositories.ColumnRepository.FindAllByProject(projectId)
}

func GetColumn(columnId int64) (*models.Column, error) {
	return repositories.ColumnRepository.Find(columnId)
}

func StoreColumn(c *models.Column) error {
	columns, err := repositories.ColumnRepository.FindAllByProjectAndName(c.ProjectId, c.Name)
	if err != nil {
		return err
	}
	if columns != nil && len(columns) > 0 {
		return &customErrors.ModelAlreadyExists{}
	}

	prevCol, err := repositories.ColumnRepository.FindWithMaxPosition(c.ProjectId)
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

	return repositories.ColumnRepository.Create(c)
}

func UpdateColumn(c *models.Column) error {
	columnFromDB, err := repositories.ColumnRepository.Find(c.Id)
	if err != nil {
		return err
	}

	c.ProjectId = columnFromDB.ProjectId
	c.Position = columnFromDB.Position
	c.CreatedAt = columnFromDB.CreatedAt
	return repositories.ColumnRepository.Update(c)
}

func DeleteColumn(columnId int64) error {
	column, err := repositories.ColumnRepository.Find(columnId)
	if err != nil {
		return err
	}

	columns, err := repositories.ColumnRepository.FindAllByProject(column.ProjectId)
	if err != nil {
		return err
	} else if columns == nil || len(columns) <= 1 {
		return errors.New("cannot delete last column")
	}

	//TODO: rewrite it, its bullshit
	nextColumn, err := repositories.ColumnRepository.FindByNextPosition(column.ProjectId, column.Position)
	if err != nil {
		return err
	} else if nextColumn == nil {
		nextColumn, err = repositories.ColumnRepository.FindByPreviousPosition(column.ProjectId, column.Position)
		if err != nil {
			return err
		} else if nextColumn == nil {
			return errors.New("cannot find column to move tasks before deletion")
		}
	}

	if err := moveAllToColumn(column, nextColumn); err != nil {
		return err
	}

	return repositories.ColumnRepository.Delete(columnId)
}

func MoveColumn(columnId int64, direction string) error {
	nextColumn := new(models.Column)
	column, err := repositories.ColumnRepository.Find(columnId)
	if err != nil {
		return err
	}

	if direction == "right" {
		nextColumn, err = repositories.ColumnRepository.FindByNextPosition(column.ProjectId, column.Position)
		if err != nil {
			return err
		} else if nextColumn == nil {
			return nil
		}

		nextColumn.Position--
		column.Position++
	} else if direction == "left" {
		nextColumn, err = repositories.ColumnRepository.FindByPreviousPosition(column.ProjectId, column.Position)
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

	if err := repositories.ColumnRepository.Update(nextColumn); err != nil {
		return err
	}

	return repositories.ColumnRepository.Update(column)
}
