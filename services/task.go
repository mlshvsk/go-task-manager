package services

import (
	"errors"
	"github.com/mlshvsk/go-task-manager/factories"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories"
)

type TaskService struct {
}

func GetTasksByColumn(columnId int64, page int64, limit int64) ([]*models.Task, error) {
	return repositories.TaskRepository.FindAllByColumn(columnId, page, limit)
}

func GetTasks(page int64, limit int64) ([]*models.Task, error) {
	return repositories.TaskRepository.FindAll(page, limit)
}

func GetTask(id int64) (*models.Task, error) {
	return repositories.TaskRepository.Find(id)
}

func StoreTask(t *models.Task) error {
	previousTask, err := repositories.TaskRepository.FindWithMaxPosition(t.ColumnId)
	if err != nil {
		return err
	}

	if previousTask == nil {
		t.Position = 0
	} else {
		t.Position = previousTask.Position + 1
	}

	*t, err = factories.TaskFactory(t.ColumnId, t.Name, t.Description, t.Position)
	if err != nil {
		return err
	}

	return repositories.TaskRepository.Create(t)
}

func UpdateTask(t *models.Task) error {
	taskFromDB, err := repositories.TaskRepository.Find(t.Id)
	if err != nil {
		return err
	}

	t.Position = taskFromDB.Position
	t.CreatedAt = taskFromDB.CreatedAt

	return repositories.TaskRepository.Update(t)
}

func DeleteTask(taskId int64) error {
	_, err := repositories.TaskRepository.Find(taskId)
	if err != nil {
		return err
	}

	return repositories.TaskRepository.Delete(taskId)
}

func MoveTaskWithinColumn(taskId int64, direction string) error {
	nextTask := new(models.Task)
	task, err := repositories.TaskRepository.Find(taskId)
	if err != nil {
		return err
	}

	if direction == "down" {
		nextTask, err = repositories.TaskRepository.FindByNextPosition(task.ColumnId, task.Position)
		if err != nil {
			return err
		}

		if nextTask == nil {
			return nil
		}

		nextTask.Position--
		task.Position++

		repositories.TaskRepository.Update(nextTask)
		repositories.TaskRepository.Update(task)
	} else if direction == "up" {
		nextTask, err = repositories.TaskRepository.FindByPreviousPosition(task.ColumnId, task.Position)
		if err != nil {
			return err
		}

		if nextTask == nil {
			return nil
		}

		nextTask.Position++
		task.Position--

		repositories.TaskRepository.Update(nextTask)
		repositories.TaskRepository.Update(task)
	} else {
		return errors.New("invalid direction: " + direction)
	}

	if err := repositories.TaskRepository.Update(nextTask); err != nil {
		return err
	}

	return repositories.TaskRepository.Update(task)
}

func MoveTaskToColumn(taskId int64, toColumnId int64) error {
	task, err := repositories.TaskRepository.Find(taskId)
	if err != nil {
		return err
	}

	_, err = repositories.ColumnRepository.Find(toColumnId)
	if err != nil {
		return err
	}

	toColumnMaxPosition, err := repositories.TaskRepository.FindWithMaxPosition(toColumnId)
	if err != nil {
		return err
	}

	if toColumnMaxPosition != nil {
		task.Position = toColumnMaxPosition.Position + 1
	} else {
		task.Position = 0
	}

	return repositories.TaskRepository.Update(task)
}

func moveAllToColumn(fromColumn *models.Column, toColumn *models.Column) error {
	tasks, err := repositories.TaskRepository.FindAllByColumn(fromColumn.Id, 0, -1)
	if err != nil {
		return err
	}

	nextPosition := int64(0)
	toColumnMaxPosition, err := repositories.TaskRepository.FindWithMaxPosition(toColumn.Id)
	if err != nil {
		return err
	} else if toColumnMaxPosition != nil {
		nextPosition = toColumnMaxPosition.Position + 1
	}

	for _, v := range tasks {
		v.ColumnId = toColumn.Id
		v.Position = nextPosition

		err := repositories.TaskRepository.Update(v)
		if err != nil {
			return err
		}

		nextPosition++
	}

	return nil
}
