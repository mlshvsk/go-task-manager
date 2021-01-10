package repositories

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mlshvsk/go-task-manager/database"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories/mysql"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
	"time"
)

func TestFindAllColumnsByProject(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getColumnRepository(db)

	columns := make([]*models.Column, 2)
	columns[0] = &models.Column{
		Id: 1,
		Name: "Test1",
		ProjectId: int64(1),
		Position: 1,
		CreatedAt: time.Now(),
	}

	columns[1] = &models.Column{
		Id: 2,
		Name: "Test2",
		ProjectId: int64(1),
		Position: 2,
		CreatedAt: time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM columns WHERE project_id=? ORDER BY position ASC")

	rows := sqlmock.NewRows([]string{"id", "name", "project_id", "position", "created_at"}).
		AddRow(columns[0].Id, columns[0].Name, columns[0].ProjectId, columns[0].Position, columns[0].CreatedAt).
		AddRow(columns[1].Id, columns[1].Name, columns[1].ProjectId, columns[1].Position, columns[1].CreatedAt)

	mock.ExpectQuery(query).WithArgs(int64(1)).WillReturnRows(rows)

	res, err := repo.FindAllByProject(int64(1))

	assert.Nil(t, err)
	assert.Equal(t, columns, res)
}

func TestFindColumn(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getColumnRepository(db)

	column := &models.Column{
		Id: 1,
		Name: "Test1",
		ProjectId: int64(1),
		Position: 1,
		CreatedAt: time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM columns WHERE id=?")

	rows := sqlmock.NewRows([]string{"id", "name", "project_id", "position", "created_at"}).
		AddRow(column.Id, column.Name, column.ProjectId, column.Position, column.CreatedAt)

	mock.ExpectQuery(query).WithArgs(column.Id).WillReturnRows(rows)

	res, err := repo.Find(column.Id)

	assert.Nil(t, err)
	assert.Equal(t, column, res)
}

func TestFindColumnByNextPosition(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getColumnRepository(db)

	column := &models.Column{
		Id: 1,
		Name: "Test1",
		ProjectId: int64(1),
		Position: 20,
		CreatedAt: time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM columns WHERE project_id=? and position>? ORDER BY position ASC")

	rows := sqlmock.NewRows([]string{"id", "name", "project_id", "position", "created_at"}).
		AddRow(column.Id, column.Name, column.ProjectId, column.Position, column.CreatedAt)

	mock.ExpectQuery(query).WithArgs(column.ProjectId, int64(10)).WillReturnRows(rows)

	res, err := repo.FindByNextPosition(column.ProjectId, int64(10))

	assert.Nil(t, err)
	assert.Equal(t, column, res)
}

func TestFindColumnByPreviousPosition(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getColumnRepository(db)

	column := &models.Column{
		Id: 1,
		Name: "Test1",
		ProjectId: int64(1),
		Position: 20,
		CreatedAt: time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM columns WHERE project_id=? and position<? ORDER BY position DESC")

	rows := sqlmock.NewRows([]string{"id", "name", "project_id", "position", "created_at"}).
		AddRow(column.Id, column.Name, column.ProjectId, column.Position, column.CreatedAt)

	mock.ExpectQuery(query).WithArgs(column.ProjectId, int64(30)).WillReturnRows(rows)

	res, err := repo.FindByPreviousPosition(column.ProjectId, int64(30))

	assert.Nil(t, err)
	assert.Equal(t, column, res)
}

func TestCreateColumn(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getColumnRepository(db)

	column := &models.Column{
		Name: "Test1",
		ProjectId: int64(1),
		Position: 20,
		CreatedAt: time.Now(),
	}

	query := regexp.QuoteMeta("INSERT INTO columns (created_at, name, position, project_id) VALUES (?, ?, ?, ?)")

	mock.ExpectExec(query).
		WithArgs(column.CreatedAt, column.Name, column.Position, column.ProjectId).
		WillReturnResult(sqlmock.NewResult(int64(10), 0))

	err := repo.Create(column)

	assert.Nil(t, err)
	assert.Equal(t, int64(10), column.Id)
}

func TestDeleteColumn(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getColumnRepository(db)

	deleteId := int64(1)

	query := regexp.QuoteMeta("DELETE FROM columns WHERE id=?")

	mock.ExpectExec(query).
		WithArgs(deleteId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(deleteId)

	assert.Nil(t, err)
}

func getColumnRepository(db *sql.DB) *columnRepository {
	baseRepo := &mysql.Repository{SqlDB: &database.SqlDB{Conn: db}, TableName: "test"}

	ColumnRepository = &columnRepository{
		base: baseRepo,
	}
	ColumnRepository.base.SetTableName("columns")

	return ColumnRepository
}

func newMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}
