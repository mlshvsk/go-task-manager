package task

import (
	"errors"
	"github.com/mlshvsk/go-task-manager/factories"
	"github.com/mlshvsk/go-task-manager/domains"
	"github.com/mlshvsk/go-task-manager/services"
	"sync"
)

type taskService struct {
	r domains.TaskRepository
}

const (
	directionUp = "up"
	directionDown = "down"
)

func InitTaskService(r domains.TaskRepository) {
	(&sync.Once{}).Do(func() {
		services.TaskService = &taskService{r}
	})
}

func (s *taskService) GetTasksByColumn(columnId int64, page int64, limit int64) ([]*domains.TaskModel, error) {
	if _, err := services.ColumnService.GetColumn(columnId); err != nil {
		return nil, err
	}

	return s.r.FindAllByColumn(columnId, page, limit)
}

func (s *taskService) GetTasks(page int64, limit int64) ([]*domains.TaskModel, error) {
	return s.r.FindAll(page, limit)
}

func (s *taskService) GetTask(id int64) (*domains.TaskModel, error) {
	return s.r.Find(id)
}

func (s *taskService) StoreTask(t *domains.TaskModel) error {
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

func (s *taskService) UpdateTask(t *domains.TaskModel) error {
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
	nextTask := new(domains.TaskModel)
	task, err := s.r.Find(taskId)
	if err != nil {
		return err
	}

	if direction == directionDown {
		nextTask, err = s.r.FindByNextPosition(task.ColumnId, task.Position)
		if err != nil {
			return err
		}

		if nextTask == nil {
			return nil
		}

		nextTask.Position, task.Position = task.Position, nextTask.Position
	} else if direction == directionUp {
		nextTask, err = s.r.FindByPreviousPosition(task.ColumnId, task.Position)
		if err != nil {
			return err
		}

		if nextTask == nil {
			return nil
		}

		nextTask.Position, task.Position = task.Position, nextTask.Position
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

	_, err = services.ColumnService.GetColumn(toColumnId)
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

func (s *taskService) MoveAllToColumn(fromColumn *domains.ColumnModel, toColumn *domains.ColumnModel) error {
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
