package mysql

import (
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/mlshvsk/go-task-manager/database"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/stretchr/testify/assert"
	"log"
	"regexp"
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	db, _ := NewMock()
	q := getBaseQuery(db)
	qExpected := getBaseQuery(db)
	qExpected.Main = "SELECT id, name FROM test"


	q.Select([]string{"id", "name"})

	assert.Equal(t, qExpected, q)
}

func TestDelete(t *testing.T) {
	db, _ := NewMock()
	q := getBaseQuery(db)
	qExpected := getBaseQuery(db)
	qExpected.Main = "DELETE FROM test"

	q.Delete()

	assert.Equal(t, qExpected, q)
}

func TestUpdate(t *testing.T) {
	db, _ := NewMock()
	q := getBaseQuery(db)

	qExpected := getBaseQuery(db)
	qExpected.Main = "UPDATE test SET age=?, name=?"
	qExpected.Values = []interface{}{10, "Mike"}

	updates := make(map[string]interface{})
	updates["name"] = "Mike"
	updates["age"] = 10

	q.Update(updates)

	assert.Equal(t, qExpected, q)
}

func TestInsert(t *testing.T) {
	db, _ := NewMock()
	q := getBaseQuery(db)

	qExpected := getBaseQuery(db)
	qExpected.Main = "INSERT INTO test (age, name) VALUES (?, ?)"
	qExpected.Values = []interface{}{10, "Mike"}

	values := make(map[string]interface{})
	values["name"] = "Mike"
	values["age"] = 10

	q.Insert(values)

	assert.Equal(t, qExpected, q)
}

func TestOrderBy(t *testing.T) {
	db, _ := NewMock()
	q := getBaseQuery(db)

	qExpected := getBaseQuery(db)
	qExpected.OrderByClause = "ORDER BY id ASC"

	q.OrderBy("id", "asc")

	assert.Equal(t, qExpected, q)
}

func TestWhere(t *testing.T) {
	db, _ := NewMock()
	q := getBaseQuery(db)

	qExpected := getBaseQuery(db)
	qExpected.WhereClause = "WHERE name=? AND position<?"
	qExpected.Values = []interface{}{"Mike", 2}

	where := [][]interface{}{{"name", "=", "Mike"}, {"position", "<", 2}}

	q.Where("AND", where)

	assert.Equal(t, qExpected, q)
}

func TestWhereColumnNameError(t *testing.T) {
	db, _ := NewMock()
	q := getBaseQuery(db)

	qExpected := getBaseQuery(db)
	qExpected.Error = errors.New("where: cannot convert column name to string")

	where := [][]interface{}{{28, "=", "Mike"}, {"position", "<", 2}}

	q.Where("AND", where)

	assert.Equal(t, &qExpected.Error, &q.Error)
}

func TestWhereArrayLenError(t *testing.T) {
	db, _ := NewMock()
	q := getBaseQuery(db)

	qExpected := getBaseQuery(db)
	qExpected.Error = errors.New("where: array len of where clause must be equal to 3")

	where := [][]interface{}{{"name", "=", "Mike"}, {"position", "<"}}

	q.Where("AND", where)

	assert.Equal(t, qExpected.Error, q.Error)
}

func TestWhereOperatorError(t *testing.T) {
	db, _ := NewMock()
	q := getBaseQuery(db)

	qExpected := getBaseQuery(db)
	qExpected.Error = errors.New("where: cannot convert operator to string")

	where := [][]interface{}{{"name", "=", "Mike"}, {"position", 32, 2}}

	q.Where("AND", where)

	assert.Equal(t, qExpected.Error, q.Error)
}

func TestCompoundQuery(t *testing.T) {
	db, _ := NewMock()
	q := getBaseQuery(db)

	expectedQuery := "SELECT id, name FROM test  ORDER BY id DESC "

	q.Select([]string{"id", "name"}).OrderBy("id", "desc")

	assert.Equal(t, expectedQuery, q.CompoundQuery())
}

func TestGetQuery(t *testing.T) {
	db, mock := NewMock()
	q := getBaseQuery(db)
	callbackArg := 1
	where := [][]interface{}{{"name", "=", "Test"}}
	expectedQuery := "SELECT id, name FROM test WHERE name=\\? ORDER BY id DESC"

	p := &models.Project{
		Id: 123,
		Name: "Test",
		Description: "Test",
		CreatedAt: time.Now(),
	}

	rows := sqlmock.NewRows([]string{"id", "name", "description", "created_at"}).
		AddRow(p.Id, p.Name, p.Description, p.CreatedAt)
	mock.ExpectQuery(expectedQuery).WithArgs(p.Name).WillReturnRows(rows)

	err := q.Select([]string{"id", "name"}).
		Where("and", where).
		OrderBy("id", "desc").
		Get(func(rows *sql.Rows) error {callbackArg = 2; return nil})

	assert.NoError(t, err)
	assert.Equal(t, 2, callbackArg)
}

func TestExecQuery(t *testing.T) {
	db, mock := NewMock()
	q := getBaseQuery(db)

	expectedQuery := regexp.QuoteMeta("INSERT INTO test (created_at, description, name) VALUES (?, ?, ?)")
	p := &models.Project{
		Id: 123,
		Name: "Test",
		Description: "Test",
		CreatedAt: time.Now(),
	}

	mock.ExpectExec(expectedQuery).
		WithArgs(p.CreatedAt, p.Description, p.Name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	insert := map[string]interface{}{
		"name": p.Name,
		"description": p.Description,
		"created_at": p.CreatedAt,
	}

	_, err := q.Insert(insert).Exec()

	assert.NoError(t, err)
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func getBaseQuery(db *sql.DB) *Query {
	q := new(Query)
	q.Repository = &Repository{SqlDB: &database.SqlDB{Conn: db}, TableName: "test"}

	return q
}
