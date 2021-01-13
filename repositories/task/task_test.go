package task

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

func TestFindAllTasks(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getTaskRepository(db)

	tasks := make([]*models.Task, 2)
	tasks[0] = &models.Task{
		Id:          1,
		Name:        "Test1",
		Description: "Descr1",
		ColumnId: int64(1),
		Position: int64(1),
		CreatedAt:   time.Now(),
	}

	tasks[1] = &models.Task{
		Id:          2,
		Name:        "Test2",
		Description: "Descr2",
		ColumnId: int64(2),
		Position: int64(1),
		CreatedAt:   time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM tasks ORDER BY created_at ASC")

	rows := sqlmock.NewRows([]string{"id", "name", "description", "column_id", "position", "created_at"}).
		AddRow(tasks[0].Id, tasks[0].Name, tasks[0].Description, tasks[0].ColumnId, tasks[0].Position, tasks[0].CreatedAt).
		AddRow(tasks[1].Id, tasks[1].Name, tasks[1].Description, tasks[1].ColumnId, tasks[1].Position, tasks[1].CreatedAt)

	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	res, err := repo.FindAll(0, -1)

	assert.Nil(t, err)
	assert.Equal(t, tasks, res)
}

func TestFindAllTasksByColumn(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getTaskRepository(db)

	tasks := make([]*models.Task, 1)
	tasks[0] = &models.Task{
		Id:          1,
		Name:        "Test1",
		Description: "Descr1",
		ColumnId: int64(1),
		Position: int64(1),
		CreatedAt:   time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM tasks WHERE column_id=? ORDER BY position ASC")

	rows := sqlmock.NewRows([]string{"id", "name", "description", "column_id", "position", "created_at"}).
		AddRow(tasks[0].Id, tasks[0].Name, tasks[0].Description, tasks[0].ColumnId, tasks[0].Position, tasks[0].CreatedAt)

	mock.ExpectQuery(query).WithArgs(tasks[0].ColumnId).WillReturnRows(rows)

	res, err := repo.FindAllByColumn(tasks[0].ColumnId,0, -1)

	assert.Nil(t, err)
	assert.Equal(t, tasks, res)
}

func TestFindWithMaxPosition(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getTaskRepository(db)

	tasks := make([]*models.Task, 2)
	tasks[0] = &models.Task{
		Id:          1,
		Name:        "Test1",
		Description: "Descr1",
		ColumnId: int64(1),
		Position: int64(2),
		CreatedAt:   time.Now(),
	}
	tasks[1] = &models.Task{
		Id:          2,
		Name:        "Test2",
		Description: "Descr2",
		ColumnId: int64(1),
		Position: int64(1),
		CreatedAt:   time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM tasks WHERE column_id=? ORDER BY position DESC")

	rows := sqlmock.NewRows([]string{"id", "name", "description", "column_id", "position", "created_at"}).
		AddRow(tasks[0].Id, tasks[0].Name, tasks[0].Description, tasks[0].ColumnId, tasks[0].Position, tasks[0].CreatedAt).
		AddRow(tasks[1].Id, tasks[1].Name, tasks[1].Description, tasks[1].ColumnId, tasks[1].Position, tasks[1].CreatedAt)


	mock.ExpectQuery(query).WithArgs(tasks[0].ColumnId).WillReturnRows(rows)

	res, err := repo.FindWithMaxPosition(tasks[0].ColumnId)

	assert.Nil(t, err)
	assert.Equal(t, tasks[0], res)
}

func TestFindByNextPosition(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getTaskRepository(db)

	tasks := make([]*models.Task, 2)
	tasks[0] = &models.Task{
		Id:          1,
		Name:        "Test1",
		Description: "Descr1",
		ColumnId: int64(1),
		Position: int64(2),
		CreatedAt:   time.Now(),
	}
	tasks[1] = &models.Task{
		Id:          2,
		Name:        "Test2",
		Description: "Descr2",
		ColumnId: int64(1),
		Position: int64(1),
		CreatedAt:   time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM tasks WHERE column_id=? and position>? ORDER BY position ASC")

	rows := sqlmock.NewRows([]string{"id", "name", "description", "column_id", "position", "created_at"}).
		AddRow(tasks[0].Id, tasks[0].Name, tasks[0].Description, tasks[0].ColumnId, tasks[0].Position, tasks[0].CreatedAt).
		AddRow(tasks[1].Id, tasks[1].Name, tasks[1].Description, tasks[1].ColumnId, tasks[1].Position, tasks[1].CreatedAt)


	mock.ExpectQuery(query).WithArgs(tasks[0].ColumnId, tasks[0].Position).WillReturnRows(rows)

	res, err := repo.FindByNextPosition(tasks[0].ColumnId, tasks[0].Position)

	assert.Nil(t, err)
	assert.Equal(t, tasks[0], res)
}

func TestFindByPreviousPosition(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getTaskRepository(db)

	tasks := make([]*models.Task, 2)
	tasks[0] = &models.Task{
		Id:          1,
		Name:        "Test1",
		Description: "Descr1",
		ColumnId: int64(1),
		Position: int64(2),
		CreatedAt:   time.Now(),
	}
	tasks[1] = &models.Task{
		Id:          2,
		Name:        "Test2",
		Description: "Descr2",
		ColumnId: int64(1),
		Position: int64(1),
		CreatedAt:   time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM tasks WHERE column_id=? and position<? ORDER BY position DESC")

	rows := sqlmock.NewRows([]string{"id", "name", "description", "column_id", "position", "created_at"}).
		AddRow(tasks[0].Id, tasks[0].Name, tasks[0].Description, tasks[0].ColumnId, tasks[0].Position, tasks[0].CreatedAt).
		AddRow(tasks[1].Id, tasks[1].Name, tasks[1].Description, tasks[1].ColumnId, tasks[1].Position, tasks[1].CreatedAt)


	mock.ExpectQuery(query).WithArgs(tasks[0].ColumnId, tasks[0].Position).WillReturnRows(rows)

	res, err := repo.FindByPreviousPosition(tasks[0].ColumnId, tasks[0].Position)

	assert.Nil(t, err)
	assert.Equal(t, tasks[0], res)
}

func TestCreateTask(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getTaskRepository(db)

	task := &models.Task{
		Name:        "Test1",
		Description: "TestD",
		ColumnId: int64(10),
		Position: int64(20),
		CreatedAt:   time.Now(),
	}

	query := regexp.QuoteMeta("INSERT INTO tasks (column_id, created_at, description, name, position) VALUES (?, ?, ?, ?, ?)")

	mock.ExpectExec(query).
		WithArgs(task.ColumnId, task.CreatedAt, task.Description, task.Name, task.Position).
		WillReturnResult(sqlmock.NewResult(int64(10), 0))

	err := repo.Create(task)

	assert.Nil(t, err)
	assert.Equal(t, int64(10), task.Id)
}

func TestDeleteTask(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getTaskRepository(db)

	deleteId := int64(1)

	query := regexp.QuoteMeta("DELETE FROM tasks WHERE id=?")

	mock.ExpectExec(query).
		WithArgs(deleteId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(deleteId)

	assert.Nil(t, err)
}

func getTaskRepository(db *sql.DB) *taskRepository {
	baseRepo := &mysql.Repository{SqlDB: &database.SqlDB{Conn: db}, TableName: "tasks"}

	return &taskRepository{
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
