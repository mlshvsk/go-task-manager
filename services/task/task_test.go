package task

import (
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories/task"
	"github.com/mlshvsk/go-task-manager/services"
	columnService "github.com/mlshvsk/go-task-manager/services/column"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestGetTasks(t *testing.T) {
	tr := task.InitTaskRepositoryMock()
	expectedResult := []*models.Task{
		{
			Id:          int64(rand.Int()),
			Name:        "Test",
			Description: "TestD",
		},
		{
			Id:          int64(rand.Int()),
			Name:        "Test2",
			Description: "TestD2",
		},
	}
	expectedPage := int64(0)
	expectedLimit := int64(2)

	tr.FindAllFunc = func(offset int64, limit int64) ([]*models.Task, error) {
		assert.Equal(t, expectedPage, offset)
		assert.Equal(t, expectedLimit, limit)
		return expectedResult, nil
	}

	InitTaskService(tr)
	res, err := services.TaskService.GetTasks(expectedPage, expectedLimit)

	assert.Equal(t, expectedResult, res)
	assert.Nil(t, err)
}

func TestGetTasksByColumn(t *testing.T) {
	tr := task.InitTaskRepositoryMock()
	expectedResult := []*models.Task{
		{
			Id:          int64(rand.Int()),
			Name:        "Test",
			Description: "TestD",
		},
		{
			Id:          int64(rand.Int()),
			Name:        "Test2",
			Description: "TestD2",
		},
	}
	expectedPage := int64(0)
	expectedLimit := int64(2)
	expectedColumn := int64(100)

	tr.FindAllByColumnFunc = func(columnId int64, offset int64, limit int64) ([]*models.Task, error) {
		assert.Equal(t, expectedColumn, columnId)
		assert.Equal(t, expectedPage, offset)
		assert.Equal(t, expectedLimit, limit)
		return expectedResult, nil
	}

	InitTaskService(tr)
	res, err := services.TaskService.GetTasksByColumn(expectedColumn, expectedPage, expectedLimit)

	assert.Equal(t, expectedResult, res)
	assert.Nil(t, err)
}

func TestMoveTaskDownWithinColumn(t *testing.T) {
	tr := task.InitTaskRepositoryMock()
	columnId := int64(rand.Int())
	initialPosition := int64(10)
	nextPosition := int64(11)
	initialUpdated := false
	nextUpdated := false

	taskDB := &models.Task{
		Id:          int64(rand.Int()),
		Name:        "Task",
		Description: "Task",
		ColumnId:    columnId,
		Position:    initialPosition,
	}

	taskNextPosition := &models.Task{
		Id:          int64(rand.Int()),
		Name:        "Task2",
		Description: "Task2",
		ColumnId:    columnId,
		Position:    nextPosition,
	}

	tr.FindFunc = func(taskId int64) (*models.Task, error) {
		return taskDB, nil
	}

	tr.FindByNextPositionFunc = func(columnId int64, position int64) (*models.Task, error) {
		return taskNextPosition, nil
	}

	tr.UpdateFunc = func(task *models.Task) error {
		if task.Id == taskDB.Id {
			assert.Equal(t, nextPosition, task.Position)
			initialUpdated = true
			return nil
		}

		if task.Id == taskNextPosition.Id {
			assert.Equal(t, initialPosition, taskNextPosition.Position)
			nextUpdated = true
			return nil
		}

		return nil
	}

	InitTaskService(tr)
	err := services.TaskService.MoveTaskWithinColumn(taskDB.Id, "down")

	assert.Nil(t, err)
	assert.True(t, initialUpdated)
	assert.True(t, nextUpdated)
}

func TestMoveTaskUpWithinColumn(t *testing.T) {
	tr := task.InitTaskRepositoryMock()
	columnId := int64(rand.Int())
	initialPosition := int64(10)
	nextPosition := int64(11)
	initialUpdated := false
	nextUpdated := false

	taskDB := &models.Task{
		Id:          int64(rand.Int()),
		Name:        "Task",
		Description: "Task",
		ColumnId:    columnId,
		Position:    initialPosition,
	}

	taskNextPosition := &models.Task{
		Id:          int64(rand.Int()),
		Name:        "Task2",
		Description: "Task2",
		ColumnId:    columnId,
		Position:    nextPosition,
	}

	tr.FindFunc = func(taskId int64) (*models.Task, error) {
		return taskDB, nil
	}

	tr.FindByPreviousPositionFunc = func(columnId int64, position int64) (*models.Task, error) {
		return taskNextPosition, nil
	}

	tr.UpdateFunc = func(task *models.Task) error {
		if task.Id == taskDB.Id {
			assert.Equal(t, nextPosition, task.Position)
			initialUpdated = true
			return nil
		}

		if task.Id == taskNextPosition.Id {
			assert.Equal(t, initialPosition, taskNextPosition.Position)
			nextUpdated = true
			return nil
		}

		return nil
	}

	InitTaskService(tr)
	err := services.TaskService.MoveTaskWithinColumn(taskDB.Id, "up")

	assert.Nil(t, err)
	assert.True(t, initialUpdated)
	assert.True(t, nextUpdated)
}

func TestMoveTaskToAnotherColumn(t *testing.T) {
	tr := task.InitTaskRepositoryMock()
	columnId := int64(rand.Int())
	toColumnId := int64(rand.Int())
	columnMaxPosition := int64(rand.Int())

	taskDB := &models.Task{
		Id:          int64(rand.Int()),
		Name:        "Task",
		Description: "Task",
		ColumnId:    columnId,
	}

	taskMaxPosition := &models.Task{
		Id:          int64(rand.Int()),
		Name:        "Task2",
		Description: "Task2",
		ColumnId:    toColumnId,
		Position:    columnMaxPosition,
	}

	tr.FindFunc = func(taskId int64) (*models.Task, error) {
		return taskDB, nil
	}

	tr.FindWithMaxPositionFunc = func(columnId int64) (*models.Task, error) {
		return taskMaxPosition, nil
	}

	tr.UpdateFunc = func(task *models.Task) error {
		assert.Equal(t, columnMaxPosition+1, task.Position)

		return nil
	}

	getColFunc := func(columnId int64) (*models.Column, error) {
		return nil, nil
	}
	services.ColumnService = &columnService.ServiceMock{GetColumnFunc: getColFunc}

	InitTaskService(tr)
	err := services.TaskService.MoveTaskToColumn(taskDB.Id, toColumnId)

	assert.Nil(t, err)
}

func TestMoveTaskToAnotherColumnFirstPosition(t *testing.T) {
	tr := task.InitTaskRepositoryMock()
	columnId := int64(rand.Int())
	toColumnId := int64(rand.Int())

	taskDB := &models.Task{
		Id:          int64(rand.Int()),
		Name:        "Task",
		Description: "Task",
		ColumnId:    columnId,
	}

	tr.FindFunc = func(taskId int64) (*models.Task, error) {
		return taskDB, nil
	}

	tr.FindWithMaxPositionFunc = func(columnId int64) (*models.Task, error) {
		return nil, nil
	}

	tr.UpdateFunc = func(task *models.Task) error {
		assert.Equal(t, int64(0), task.Position)

		return nil
	}

	getColFunc := func(columnId int64) (*models.Column, error) {
		return nil, nil
	}
	services.ColumnService = &columnService.ServiceMock{GetColumnFunc: getColFunc}

	InitTaskService(tr)
	err := services.TaskService.MoveTaskToColumn(taskDB.Id, toColumnId)

	assert.Nil(t, err)
}
