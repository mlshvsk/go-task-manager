package comment

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

func TestFindAllCommentsByTask(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getCommentRepository(db)

	comments := make([]*models.Comment, 2)
	comments[0] = &models.Comment{
		Id:        1,
		Data:      "Test1",
		CreatedAt: time.Now(),
	}
	comments[1] = &models.Comment{
		Id:        2,
		Data:      "Test2",
		CreatedAt: time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM comments WHERE task_id=? ORDER BY created_at DESC")

	rows := sqlmock.NewRows([]string{"id", "data", "task_id", "created_at"}).
		AddRow(comments[0].Id, comments[0].Data, comments[0].TaskId, comments[0].CreatedAt).
		AddRow(comments[1].Id, comments[1].Data, comments[1].TaskId, comments[1].CreatedAt)

	mock.ExpectQuery(query).WithArgs(int64(1)).WillReturnRows(rows)

	res, err := repo.FindAllByTask(int64(1), 0, -1)

	assert.Nil(t, err)
	assert.Equal(t, comments, res)
}

func TestFindComment(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getCommentRepository(db)

	comment := &models.Comment{
		Id:        1,
		Data:      "Test1",
		TaskId:    int64(1),
		CreatedAt: time.Now(),
	}

	query := regexp.QuoteMeta("SELECT * FROM comments WHERE id=?")

	rows := sqlmock.NewRows([]string{"id", "data", "task_id", "created_at"}).
		AddRow(comment.Id, comment.Data, comment.TaskId, comment.CreatedAt)

	mock.ExpectQuery(query).WithArgs(comment.Id).WillReturnRows(rows)

	res, err := repo.Find(comment.Id)

	assert.Nil(t, err)
	assert.Equal(t, comment, res)
}

func TestCreateComment(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getCommentRepository(db)

	comment := &models.Comment{
		Data:      "Test1",
		TaskId:    int64(1),
		CreatedAt: time.Now(),
	}

	query := regexp.QuoteMeta("INSERT INTO comments (created_at, data, task_id) VALUES (?, ?, ?)")

	mock.ExpectExec(query).
		WithArgs(comment.CreatedAt, comment.Data, comment.TaskId).
		WillReturnResult(sqlmock.NewResult(int64(10), 0))

	err := repo.Create(comment)

	assert.Nil(t, err)
	assert.Equal(t, int64(10), comment.Id)
}

func TestUpdateComment(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getCommentRepository(db)

	comment := &models.Comment{
		Id:        int64(100),
		Data:      "Test1",
		TaskId:    int64(1),
		CreatedAt: time.Now(),
	}

	query := regexp.QuoteMeta("UPDATE comments SET data=?, task_id=? WHERE id=?")

	mock.ExpectExec(query).
		WithArgs(comment.Data, comment.TaskId, comment.Id).
		WillReturnResult(sqlmock.NewResult(int64(10), 0))

	err := repo.Update(comment)

	assert.Nil(t, err)
}

func TestDeleteComment(t *testing.T) {
	db, mock := newMock()
	defer db.Close()

	repo := getCommentRepository(db)

	deleteId := int64(1)

	query := regexp.QuoteMeta("DELETE FROM comments WHERE id=?")

	mock.ExpectExec(query).
		WithArgs(deleteId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(deleteId)

	assert.Nil(t, err)
}

func getCommentRepository(db *sql.DB) *commentRepository {
	baseRepo := &mysql.Repository{SqlDB: &database.SqlDB{Conn: db}, TableName: "comments"}

	return &commentRepository{
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
