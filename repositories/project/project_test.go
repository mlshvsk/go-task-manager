package project

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mlshvsk/go-task-manager/database"
	"github.com/mlshvsk/go-task-manager/domains"
	"github.com/mlshvsk/go-task-manager/repositories/mysql"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
	"time"
)

func TestFindAllProjects(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getProjectRepository(db)

	projects := make([]*domains.ProjectModel, 2)
	projects[0] = &domains.ProjectModel{
		Id:          1,
		Name:        "Test1",
		Description: "Descr1",
		CreatedAt:   time.Now(),
	}

	projects[1] = &domains.ProjectModel{
		Id:          2,
		Name:        "Test2",
		Description: "Descr2",
		CreatedAt:   time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM projects ORDER BY name ASC")

	rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
		AddRow(projects[0].Id, projects[0].Name, projects[0].Description, projects[0].CreatedAt).
		AddRow(projects[1].Id, projects[1].Name, projects[1].Description, projects[1].CreatedAt)

	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	res, err := repo.FindAll(0, -1)

	assert.Nil(t, err)
	assert.Equal(t, projects, res)
}

func TestFindProject(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getProjectRepository(db)

	project := &domains.ProjectModel{
		Id:          1,
		Name:        "Test1",
		Description: "TestD",
		CreatedAt:   time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM projects WHERE id=?")

	rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
		AddRow(project.Id, project.Name, project.Description, project.CreatedAt)

	mock.ExpectQuery(query).WithArgs(project.Id).WillReturnRows(rows)

	res, err := repo.Find(project.Id)

	assert.Nil(t, err)
	assert.Equal(t, project, res)
}

func TestFindProjectByName(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getProjectRepository(db)

	projects := make([]*domains.ProjectModel, 1)
	projects[0] = &domains.ProjectModel{
		Id:          1,
		Name:        "Test1",
		Description: "Descr1",
		CreatedAt:   time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM projects WHERE name=? ORDER BY id ASC")

	rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
		AddRow(projects[0].Id, projects[0].Name, projects[0].Description, projects[0].CreatedAt)

	mock.ExpectQuery(query).WithArgs(projects[0].Name).WillReturnRows(rows)

	res, err := repo.FindAllByName(projects[0].Name)

	assert.Nil(t, err)
	assert.Equal(t, projects, res)
}

func TestCreateProject(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getProjectRepository(db)

	project := &domains.ProjectModel{
		Name:        "Test1",
		Description: "TestD",
		CreatedAt:   time.Now(),
	}

	query := regexp.QuoteMeta("INSERT INTO projects (created_at, description, name) VALUES (?, ?, ?)")

	mock.ExpectExec(query).
		WithArgs(project.CreatedAt, project.Description, project.Name).
		WillReturnResult(sqlmock.NewResult(int64(10), 0))

	err := repo.Create(project)

	assert.Nil(t, err)
	assert.Equal(t, int64(10), project.Id)
}

func TestDeleteProject(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getProjectRepository(db)

	deleteId := int64(1)

	query := regexp.QuoteMeta("DELETE FROM projects WHERE id=?")

	mock.ExpectExec(query).
		WithArgs(deleteId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(deleteId)

	assert.Nil(t, err)
}

func getProjectRepository(db *sql.DB) *projectRepository {
	baseRepo := &mysql.Repository{SqlDB: &database.SqlDB{Conn: db}, TableName: "projects"}

	return &projectRepository{
		base: baseRepo,
	}
}

func newMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
