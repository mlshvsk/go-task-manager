package base

import (
	"database/sql"
	"github.com/mlshvsk/go-task-manager/database"
)

type BaseRepository struct {
	*database.SqlDB
	TableName string
}

type Repository interface {
	GetTableName() string
	SetTableName(name string)
	SqlDb() *sql.DB
	FindAll() Query
	Find(id int64) Query
	Create(data map[string]interface{}) (int64, error)
	Update(id int64, data map[string]interface{}) error
	Delete(id int64) error
}

type BaseQuery struct {
	Main          string
	OrderByClause string
	WhereClause   string
	OffsetClause        string
	Values        []interface{}
	Repository    Repository
	Error         error
}

type Query interface {
	Select(columns []string) Query
	Delete() Query
	Update(data map[string]interface{}) Query
	Insert(data map[string]interface{}) Query
	OrderBy(colName string, order string) Query
	Limit(page int64, limit int64) Query
	Where(logicalOperator string, data [][]interface{}) Query
	Get(callback func(rows *sql.Rows) error) error
	Exec() (sql.Result, error)
}
