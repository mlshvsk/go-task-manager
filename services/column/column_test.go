package column

import (
	"github.com/mlshvsk/go-task-manager/domains"
	"github.com/mlshvsk/go-task-manager/repositories/column"
	"github.com/mlshvsk/go-task-manager/services"
	"github.com/mlshvsk/go-task-manager/services/task"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestGetColumns(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedResult := []*domains.ColumnModel{
		{
			Id:   int64(rand.Int()),
			Name: "Test",
		},
		{
			Id:   int64(rand.Int()),
			Name: "Test2",
		},
	}
	expectedProjectId := int64(100)
	expectedPage := int64(0)
	expectedLimit := int64(2)

	cr.FindAllByProjectFunc = func(projectId int64, offset int64, limit int64) ([]*domains.ColumnModel, error) {
		assert.Equal(t, expectedProjectId, projectId)
		assert.Equal(t, expectedPage, offset)
		assert.Equal(t, expectedLimit, limit)
		return expectedResult, nil
	}

	InitColumnService(cr)
	res, err := services.ColumnService.GetColumns(expectedProjectId, expectedPage, expectedLimit)

	assert.Equal(t, expectedResult, res)
	assert.Nil(t, err)
}

func TestGetColumn(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedResult := &domains.ColumnModel{
		Id:   int64(rand.Int()),
		Name: "Test",
	}
	expectedProjectId := int64(100)

	cr.FindFunc = func(projectId int64) (*domains.ColumnModel, error) {
		assert.Equal(t, expectedProjectId, projectId)
		return expectedResult, nil
	}

	InitColumnService(cr)
	res, err := services.ColumnService.GetColumn(expectedProjectId)

	assert.Equal(t, expectedResult, res)
	assert.Nil(t, err)
}

func TestStoreFirstColumn(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedProjectId := int64(100)
	expectedPosition := int64(0)
	expectedId := int64(200)
	model := &domains.ColumnModel{
		Name:      "Test",
		ProjectId: expectedProjectId,
	}

	cr.FindAllByProjectAndNameFunc = func(projectId int64, name string) ([]*domains.ColumnModel, error) {
		assert.Equal(t, expectedProjectId, projectId)
		assert.Equal(t, model.Name, name)
		return nil, nil
	}
	cr.FindWithMaxPositionFunc = func(projectId int64) (*domains.ColumnModel, error) {
		assert.Equal(t, expectedProjectId, projectId)
		return nil, nil
	}
	cr.CreateFunc = func(c *domains.ColumnModel) error {
		model.Id = expectedId
		return nil
	}

	InitColumnService(cr)
	err := services.ColumnService.StoreColumn(model)

	assert.Nil(t, err)
	assert.NotNil(t, model.CreatedAt)
	assert.Equal(t, expectedPosition, model.Position)
	assert.Equal(t, expectedId, model.Id)
}

func TestStoreColumnAppendedPosition(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedProjectId := int64(100)
	expectedPosition := int64(100)
	model := &domains.ColumnModel{
		Name:      "Test",
		ProjectId: expectedProjectId,
	}

	cr.FindAllByProjectAndNameFunc = func(projectId int64, name string) ([]*domains.ColumnModel, error) {
		return nil, nil
	}
	cr.FindWithMaxPositionFunc = func(projectId int64) (*domains.ColumnModel, error) {
		maxModel := &domains.ColumnModel{
			Name:      "Test",
			ProjectId: expectedProjectId,
			Position:  expectedPosition - 1,
		}
		return maxModel, nil
	}
	cr.CreateFunc = func(c *domains.ColumnModel) error {
		return nil
	}

	InitColumnService(cr)
	err := services.ColumnService.StoreColumn(model)

	assert.Nil(t, err)
	assert.Equal(t, expectedPosition, model.Position)
}

func TestUpdateColumn(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedId := int64(100)
	expectedPosition := int64(200)
	expectedCreatedAt := time.Now()
	expectedProjectId := int64(100)

	expectedResult := &domains.ColumnModel{
		Id:        expectedId,
		Name:      "Test2",
		ProjectId: expectedProjectId,
		Position:  expectedPosition,
		CreatedAt: expectedCreatedAt,
	}

	updateModel := &domains.ColumnModel{
		Id:   expectedId,
		Name: "Test",
	}

	cr.FindFunc = func(projectId int64) (*domains.ColumnModel, error) {
		return expectedResult, nil
	}
	cr.UpdateFunc = func(model *domains.ColumnModel) error {
		return nil
	}

	InitColumnService(cr)
	err := services.ColumnService.UpdateColumn(updateModel)

	expectedResult.Name = updateModel.Name

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, updateModel)
}

func TestDeleteColumn(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedResult := &domains.ColumnModel{
		Id:       int64(rand.Int()),
		Name:     "Test",
		Position: int64(rand.Int()),
	}

	cr.FindAllByProjectFunc = func(projectId int64, offset int64, limit int64) ([]*domains.ColumnModel, error) {
		res := []*domains.ColumnModel{
			{
				Id: int64(rand.Int()),
			},
			{
				Id: int64(rand.Int()),
			},
		}
		return res, nil
	}

	cr.FindFunc = func(id int64) (*domains.ColumnModel, error) {
		return expectedResult, nil
	}

	cr.FindByNextPositionFunc = func(projectId int64, position int64) (*domains.ColumnModel, error) {
		model := &domains.ColumnModel{
			Id:       int64(rand.Int()),
			Name:     "Test",
			Position: position + 1,
		}

		return model, nil
	}

	cr.DeleteFunc = func(id int64) error {
		return nil
	}

	tsMock := &task.ServiceMock{}
	tsMock.MoveAllToColumnFunc = func(fromColumn *domains.ColumnModel, toColumn *domains.ColumnModel) error {
		return nil
	}
	services.TaskService = tsMock
	InitColumnService(cr)

	err := services.ColumnService.DeleteColumn(expectedResult.Id)

	assert.Nil(t, err)
}

func TestMoveColumnRight(t *testing.T) {
	expectedId := int64(100)
	expectedPosition := int64(200)
	expectedResult := &domains.ColumnModel{
		Id:       expectedId,
		Name:     "Test",
		Position: expectedPosition,
	}

	cr := column.InitColumnRepositoryMock()
	cr.FindFunc = func(id int64) (*domains.ColumnModel, error) {
		return expectedResult, nil
	}

	nextExpectedResult := &domains.ColumnModel{
		Id:       expectedId,
		Name:     "Test",
		Position: expectedPosition + 1,
	}
	cr.FindByNextPositionFunc = func(projectId int64, position int64) (*domains.ColumnModel, error) {
		return nextExpectedResult, nil
	}
	cr.UpdateFunc = func(model *domains.ColumnModel) error {
		return nil
	}

	InitColumnService(cr)
	err := services.ColumnService.MoveColumn(expectedId, "right")

	assert.Nil(t, err)
	assert.Equal(t, expectedPosition+1, expectedResult.Position)
	assert.Equal(t, expectedPosition, nextExpectedResult.Position)
}
