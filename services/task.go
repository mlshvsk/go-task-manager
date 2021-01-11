package services

import (
	"errors"
	"github.com/mlshvsk/go-task-manager/factories"
	"github.com/mlshvsk/go-task-manager/models"
	"sync"
)

type taskService struct {
	r models.TaskRepository
}

var TaskService taskService

func InitTaskService(r models.TaskRepository) {
	(&sync.Once{}).Do(func() {
		TaskService = taskService{r}
	})
}

func (s *taskService) GetTasksByColumn(columnId int64, page int64, limit int64) ([]*models.Task, error) {
	return s.r.FindAllByColumn(columnId, page, limit)
}

func (s *taskService) GetTasks(page int64, limit int64) ([]*models.Task, error) {
	return s.r.FindAll(page, limit)
}

func (s *taskService) GetTask(id int64) (*models.Task, error) {
	return s.r.Find(id)
}

func (s *taskService) StoreTask(t *models.Task) error {
	previousTask, err := s.r.FindWithMaxPosition(t.ColumnId)
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

	return s.r.Create(t)
}

func (s *taskService) UpdateTask(t *models.Task) error {
	taskFromDB, err := s.r.Find(t.Id)
	if err != nil {
		return err
	}

	t.Position = taskFromDB.Position
	t.CreatedAt = taskFromDB.CreatedAt

	return s.r.Update(t)
}

func (s *taskService) DeleteTask(taskId int64) error {
	_, err := s.r.Find(taskId)
	if err != nil {
		return err
	}

	return s.r.Delete(taskId)
}

func (s *taskService) MoveTaskWithinColumn(taskId int64, direction string) error {
	nextTask := new(models.Task)
	task, err := s.r.Find(taskId)
	if err != nil {
		return err
	}

	if direction == "down" {
		nextTask, err = s.r.FindByNextPosition(task.ColumnId, task.Position)
		if err != nil {
			return err
		}

		if nextTask == nil {
			return nil
		}

		nextTask.Position--
		task.Position++

		s.r.Update(nextTask)
		s.r.Update(task)
	} else if direction == "up" {
		nextTask, err = s.r.FindByPreviousPosition(task.ColumnId, task.Position)
		if err != nil {
			return err
		}

		if nextTask == nil {
			return nil
		}

		nextTask.Position++
		task.Position--

		s.r.Update(nextTask)
		s.r.Update(task)
	} else {
		return errors.New("invalid direction: " + direction)
	}

	if err := s.r.Update(nextTask); err != nil {
		return err
	}

	return s.r.Update(task)
}

func (s *taskService) MoveTaskToColumn(taskId int64, toColumnId int64) error {
	task, err := s.r.Find(taskId)
	if err != nil {
		return err
	}

	_, err = ColumnService.GetColumn(toColumnId)
	if err != nil {
		return err
	}

	toColumnMaxPosition, err := s.r.FindWithMaxPosition(toColumnId)
	if err != nil {
		return err
	}

	if toColumnMaxPosition != nil {
		task.Position = toColumnMaxPosition.Position + 1
	} else {
		task.Position = 0
	}

	return s.r.Update(task)
}

func (s *taskService) moveAllToColumn(fromColumn *models.Column, toColumn *models.Column) error {
	tasks, err := s.r.FindAllByColumn(fromColumn.Id, 0, -1)
	if err != nil {
		return err
	}

	nextPosition := int64(0)
	toColumnMaxPosition, err := s.r.FindWithMaxPosition(toColumn.Id)
	if err != nil {
		return err
	} else if toColumnMaxPosition != nil {
		nextPosition = toColumnMaxPosition.Position + 1
	}

	for _, v := range tasks {
		v.ColumnId = toColumn.Id
		v.Position = nextPosition

		err := s.r.Update(v)
		if err != nil {
			return err
		}

		nextPosition++
	}

	return nil
}
