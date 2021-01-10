package mysql

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mlshvsk/go-task-manager/database"
	"github.com/mlshvsk/go-task-manager/repositories/base"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

func TestGetTableName(t *testing.T) {
	db, _ := NewMock()
	repo := initRepository(db)

	assert.Equal(t, "test", repo.GetTableName())
}

func TestSeTableName(t *testing.T) {
	db, _ := NewMock()
	repo := initRepository(db)
	repo.SetTableName("new")

	assert.Equal(t, "new", repo.GetTableName())
}

func TestFindAll(t *testing.T) {
	db, _ := NewMock()
	repo := initRepository(db)
	q := repo.FindAll()

	qExpected := getBaseQuery(db)
	qExpected.Main = "SELECT * FROM test"

	assert.Equal(t, qExpected, q)
}

func TestFind(t *testing.T) {
	db, _ := NewMock()
	repo := initRepository(db)
	q := repo.Find(1)

	qExpected := getBaseQuery(db)
	qExpected.Main = "SELECT * FROM test"
	qExpected.WhereClause = "WHERE id=?"
	qExpected.Values = []interface{}{int64(1)}

	assert.Equal(t, qExpected, q)
}

func TestCreate(t *testing.T) {
	db, mock := NewMock()
	repo := initRepository(db)

	qExpected := getBaseQuery(db)
	qExpected.Main = regexp.QuoteMeta("INSERT INTO test (a, b) VALUES (?, ?)  ")
	qExpected.Values = []interface{}{int64(1), "c"}

	mock.ExpectExec(qExpected.Main).
		WithArgs(1, "c").
		WillReturnResult(sqlmock.NewResult(int64(1), 1))

	_, err := repo.Create(map[string]interface{}{"a": 1, "b": "c"})

	assert.Nil(t, err)
}

func TestUpdateRepo(t *testing.T) {
	db, mock := NewMock()
	repo := initRepository(db)

	query := regexp.QuoteMeta("UPDATE test SET a=?, b=? WHERE id=?")

	mock.ExpectExec(query).
		WithArgs(1, "c", int64(1)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Update(int64(1), map[string]interface{}{"a": 1, "b": "c"})

	assert.Nil(t, err)
}

func TestDeleteRepo(t *testing.T) {
	db, mock := NewMock()
	repo := initRepository(db)

	query := regexp.QuoteMeta("DELETE FROM test WHERE id=?")

	mock.ExpectExec(query).
		WithArgs(int64(1)).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Delete(int64(1))

	assert.Nil(t, err)
}

func initRepository(db *sql.DB) base.Repository {
	return &Repository{SqlDB: &database.SqlDB{Conn: db}, TableName: "test"}
}
