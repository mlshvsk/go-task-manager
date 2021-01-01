package services

import (
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories"
)

type TaskService struct {
}

func GetTasks(columnId int) []*models.Task {
	res, _ := repositories.TaskRepository.FindAll(columnId)
	return res
}

func StoreTask(t *models.Task) *models.Task {
	_ = repositories.TaskRepository.Create(t)

	return t
}

func GetTask(id int) (*models.Task, error) {
	return repositories.TaskRepository.Find(id)
}

func DeleteTask(taskId int) {
	repositories.TaskRepository.Delete(taskId)
}

func MoveTask(taskId int, projectId int, direction string) error {
	res, _ := repositories.TaskRepository.Find(taskId)

	if direction == "down" {
		nextTask, _ := repositories.TaskRepository.FindByNextPosition(projectId, res.Position)

		if nextTask == nil {
			return nil
		}

		nextTask.Position--
		res.Position++

		repositories.TaskRepository.Update(nextTask)
		repositories.TaskRepository.Update(res)
	}

	if direction == "up" {
		nextTask, _ := repositories.TaskRepository.FindByPreviousPosition(projectId, res.Position)

		if nextTask == nil {
			return nil
		}

		nextTask.Position++
		res.Position--

		repositories.TaskRepository.Update(nextTask)
		repositories.TaskRepository.Update(res)
	}

	return nil
}
