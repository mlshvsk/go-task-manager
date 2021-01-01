package services

import (
	"errors"
	"fmt"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories"
)

type ColumnService struct {
}

func GetColumns(projectId int) []*models.Column {
	res, _ := repositories.ColumnRepository.FindAll(projectId)
	return res
}

func GetColumn(projectId int) *models.Column {
	res, _ := repositories.ColumnRepository.Find(projectId)
	return res
}

func StoreColumn(c *models.Column) *models.Column {
	prevCol, _ := repositories.ColumnRepository.FindWithMaxPosition(c.ProjectId)
	c.Position = prevCol.Position + 1
	_ = repositories.ColumnRepository.Create(c)

	return c
}

func UpdateColumn(c *models.Column) *models.Column {
	err := repositories.ProjectRepository.Update(c.Id, map[string]interface{}{"name": c.Name})

	if err != nil {
		fmt.Println(err.Error())
	}

	return c
}

func DeleteColumn(columnId int, projectId int) error {
	res, _ := repositories.ColumnRepository.FindAll(projectId)

	if res == nil || len(res) <= 1 {
		return errors.New("cannot delete last column")
	}

	repositories.ColumnRepository.Delete(columnId)

	return nil
}

func MoveColumn(columnId int, projectId int, direction string) error {
	res, _ := repositories.ColumnRepository.Find(columnId)

	if direction == "right" {
		nextCol, _ := repositories.ColumnRepository.FindByNextPosition(projectId, res.Position)

		if nextCol == nil {
			return nil
		}

		nextCol.Position--
		res.Position++

		repositories.ColumnRepository.Update(nextCol)
		repositories.ColumnRepository.Update(res)
	}

	if direction == "left" {
		nextCol, _ := repositories.ColumnRepository.FindByPreviousPosition(projectId, res.Position)

		if nextCol == nil {
			return nil
		}

		nextCol.Position++
		res.Position--

		repositories.ColumnRepository.Update(nextCol)
		repositories.ColumnRepository.Update(res)
	}

	return nil
}
