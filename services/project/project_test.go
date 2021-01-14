package project

import (
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories/project"
	"github.com/mlshvsk/go-task-manager/services"
	columnService "github.com/mlshvsk/go-task-manager/services/column"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestGetProjects(t *testing.T) {
	pr := project.InitProjectRepositoryMock()
	expectedResult := []*models.Project{
		{
			Id:   int64(rand.Int()),
			Name: "Test",
			Description: "TestD",
		},
		{
			Id:   int64(rand.Int()),
			Name: "Test2",
			Description: "TestD2",
		},
	}
	expectedPage := int64(0)
	expectedLimit := int64(2)

	pr.FindAllFunc = func(offset int64, limit int64) ([]*models.Project, error) {
		assert.Equal(t, expectedPage, offset)
		assert.Equal(t, expectedLimit, limit)
		return expectedResult, nil
	}

	InitProjectService(pr)
	res, err := services.ProjectService.GetProjects(expectedPage, expectedLimit)

	assert.Equal(t, expectedResult, res)
	assert.Nil(t, err)
}

func TestGetProject(t *testing.T) {
	pr := project.InitProjectRepositoryMock()
	expectedResult := &models.Project{
		Id:   int64(rand.Int()),
		Name: "Test",
		Description: "TestD",
	}
	expectedId := int64(100)

	pr.FindFunc = func(id int64) (*models.Project, error) {
		assert.Equal(t, expectedId, id)
		return expectedResult, nil
	}

	InitProjectService(pr)
	res, err := services.ProjectService.GetProject(expectedId)

	assert.Equal(t, expectedResult, res)
	assert.Nil(t, err)
}

func TestStoreProject(t *testing.T) {
	columnCreateIsCalled := false
	pr := project.InitProjectRepositoryMock()

	p := &models.Project{
		Name: "Test",
		Description: "TestD",
	}
	expectedProjectId := int64(rand.Int())

	pr.CreateFunc = func(p *models.Project) error {
		p.Id = expectedProjectId
		return nil
	}
	pr.FindAllByNameFunc = func(name string) ([]*models.Project, error) {
		return nil, nil
	}
	createFunc := func(c *models.Column) error {
		columnCreateIsCalled = true
		assert.Equal(t, expectedProjectId, c.ProjectId)
		assert.Equal(t, "New", c.Name)
		return nil
	}

	services.ColumnService = &columnService.ServiceMock{StoreColumnFunc: createFunc}

	InitProjectService(pr)
	err := services.ProjectService.StoreProject(p)

	assert.Nil(t, err)
	assert.True(t, columnCreateIsCalled)
	assert.Equal(t, expectedProjectId, p.Id)
	assert.Equal(t, expectedProjectId, p.Id)
}

func TestDeleteProject(t *testing.T) {
	pr := project.InitProjectRepositoryMock()
	expectedProjectId := int64(rand.Int())

	pr.DeleteFunc = func(id int64) error {
		assert.Equal(t, expectedProjectId, id)
		return nil
	}
	pr.FindFunc = func(id int64) (*models.Project, error) {
		return nil, nil
	}

	InitProjectService(pr)
	err := services.ProjectService.DeleteProject(expectedProjectId)

	assert.Nil(t, err)
}
