package column

import (
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories/column"
	"github.com/mlshvsk/go-task-manager/services/task"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestGetColumns(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedResult := []*models.Column{
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

	cr.FindAllByProjectFunc = func(projectId int64, offset int64, limit int64) ([]*models.Column, error) {
		assert.Equal(t, expectedProjectId, projectId)
		assert.Equal(t, expectedPage, offset)
		assert.Equal(t, expectedLimit, limit)
		return expectedResult, nil
	}

	InitColumnService(cr)
	res, err := Service.GetColumns(expectedProjectId, expectedPage, expectedLimit)

	assert.Equal(t, expectedResult, res)
	assert.Nil(t, err)
}

func TestGetColumn(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedResult := &models.Column{
		Id:   int64(rand.Int()),
		Name: "Test",
	}
	expectedProjectId := int64(100)

	cr.FindFunc = func(projectId int64) (*models.Column, error) {
		assert.Equal(t, expectedProjectId, projectId)
		return expectedResult, nil
	}

	InitColumnService(cr)
	res, err := Service.GetColumn(expectedProjectId)

	assert.Equal(t, expectedResult, res)
	assert.Nil(t, err)
}

func TestStoreFirstColumn(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedProjectId := int64(100)
	expectedPosition := int64(0)
	expectedId := int64(200)
	model := &models.Column{
		Name:      "Test",
		ProjectId: expectedProjectId,
	}

	cr.FindAllByProjectAndNameFunc = func(projectId int64, name string) ([]*models.Column, error) {
		assert.Equal(t, expectedProjectId, projectId)
		assert.Equal(t, model.Name, name)
		return nil, nil
	}
	cr.FindWithMaxPositionFunc = func(projectId int64) (*models.Column, error) {
		assert.Equal(t, expectedProjectId, projectId)
		return nil, nil
	}
	cr.CreateFunc = func(c *models.Column) error {
		model.Id = expectedId
		return nil
	}

	InitColumnService(cr)
	err := Service.StoreColumn(model)

	assert.Nil(t, err)
	assert.NotNil(t, model.CreatedAt)
	assert.Equal(t, expectedPosition, model.Position)
	assert.Equal(t, expectedId, model.Id)
}

func TestStoreColumnAppendedPosition(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedProjectId := int64(100)
	expectedPosition := int64(100)
	model := &models.Column{
		Name:      "Test",
		ProjectId: expectedProjectId,
	}

	cr.FindAllByProjectAndNameFunc = func(projectId int64, name string) ([]*models.Column, error) {
		return nil, nil
	}
	cr.FindWithMaxPositionFunc = func(projectId int64) (*models.Column, error) {
		maxModel := &models.Column{
			Name:      "Test",
			ProjectId: expectedProjectId,
			Position:  expectedPosition - 1,
		}
		return maxModel, nil
	}
	cr.CreateFunc = func(c *models.Column) error {
		return nil
	}

	InitColumnService(cr)
	err := Service.StoreColumn(model)

	assert.Nil(t, err)
	assert.Equal(t, expectedPosition, model.Position)
}

func TestUpdateColumn(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedId := int64(100)
	expectedPosition := int64(200)
	expectedCreatedAt := time.Now()
	expectedProjectId := int64(100)

	expectedResult := &models.Column{
		Id:        expectedId,
		Name:      "Test2",
		ProjectId: expectedProjectId,
		Position:  expectedPosition,
		CreatedAt: expectedCreatedAt,
	}

	updateModel := &models.Column{
		Id:   expectedId,
		Name: "Test",
	}

	cr.FindFunc = func(projectId int64) (*models.Column, error) {
		return expectedResult, nil
	}
	cr.UpdateFunc = func(model *models.Column) error {
		return nil
	}

	InitColumnService(cr)
	err := Service.UpdateColumn(updateModel)

	expectedResult.Name = updateModel.Name

	assert.Nil(t, err)
	assert.Equal(t, expectedResult, updateModel)
}

func TestDeleteColumn(t *testing.T) {
	cr := column.InitColumnRepositoryMock()
	expectedResult := &models.Column{
		Id:       int64(rand.Int()),
		Name:     "Test",
		Position: int64(rand.Int()),
	}

	cr.FindAllByProjectFunc = func(projectId int64, offset int64, limit int64) ([]*models.Column, error) {
		res := []*models.Column{
			{
				Id: int64(rand.Int()),
			},
			{
				Id: int64(rand.Int()),
			},
		}
		return res, nil
	}

	cr.FindFunc = func(id int64) (*models.Column, error) {
		return expectedResult, nil
	}

	cr.FindByNextPositionFunc = func(projectId int64, position int64) (*models.Column, error) {
		model := &models.Column{
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
	tsMock.MoveAllToColumnFunc = func(fromColumn *models.Column, toColumn *models.Column) error {
		return nil
	}
	task.Service = tsMock
	InitColumnService(cr)

	err := Service.DeleteColumn(expectedResult.Id)

	assert.Nil(t, err)
}

func TestMoveColumnRight(t *testing.T) {
	expectedId := int64(100)
	expectedPosition := int64(200)
	expectedResult := &models.Column{
		Id:       expectedId,
		Name:     "Test",
		Position: expectedPosition,
	}

	cr := column.InitColumnRepositoryMock()
	cr.FindFunc = func(id int64) (*models.Column, error) {
		return expectedResult, nil
	}

	nextExpectedResult := &models.Column{
		Id:       expectedId,
		Name:     "Test",
		Position: expectedPosition + 1,
	}
	cr.FindByNextPositionFunc = func(projectId int64, position int64) (*models.Column, error) {
		return nextExpectedResult, nil
	}
	cr.UpdateFunc = func(model *models.Column) error {
		return nil
	}

	InitColumnService(cr)
	err := Service.MoveColumn(expectedId, "right")

	assert.Nil(t, err)
	assert.Equal(t, expectedPosition+1, expectedResult.Position)
	assert.Equal(t, expectedPosition, nextExpectedResult.Position)
}
