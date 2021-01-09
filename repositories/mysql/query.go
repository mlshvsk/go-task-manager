package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/repositories/base"
	"strings"
)

type Query base.BaseQuery

func (q *Query) Select(columns []string) base.Query {
	q.Main = fmt.Sprintf("SELECT %v FROM %v", strings.Join(columns, ", "), q.Repository.GetTableName())

	return q
}

func (q *Query) Delete() base.Query {
	q.Main = fmt.Sprintf("DELETE FROM %v", q.Repository.GetTableName())

	return q
}

func (q *Query) Update(data map[string]interface{}) base.Query {
	var fieldPlaceholders string
	var pl []string

	for i, v := range data {
		pl = append(pl, i + " = ?")
		q.Values = append(q.Values, v)
	}

	fieldPlaceholders = strings.Join(pl, ", ")

	q.Main = fmt.Sprintf("UPDATE %v SET %v", q.Repository.GetTableName(), fieldPlaceholders)

	return q
}

func (q *Query) Insert(data map[string]interface{}) base.Query {
	var fieldNames, placeholders string
	var values []interface{}
	var pl []string
	var fn []string

	for i, v := range data {
		fn = append(fn, i)
		pl = append(pl, "?")
		values = append(values, v)
	}

	placeholders = strings.Join(pl, ", ")
	fieldNames = strings.Join(fn, ", ")

	q.Main = fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", q.Repository.GetTableName(), fieldNames, placeholders)
	q.Values = values

	return q
}

func (q *Query) OrderBy(colName string, order string) base.Query {
	q.OrderByClause += "ORDER BY " + colName + " " + strings.ToUpper(order)

	return q
}

func (q *Query) Where(logicalOperator string, data [][]interface{}) base.Query {
	var placeholders []string

	for _, v := range data {
		if len(v) != 3 {
			q.Error = errors.New("where: array len of where clause must be equal to 3")
			return q
		}

		col, ok := v[0].(string)
		if ok == false {
			q.Error = errors.New("where: cannot convert column name to string")
			return q
		}

		operator, ok := v[1].(string)
		if ok == false {
			q.Error = errors.New("where: cannot convert operator to string")
			return q
		}

		placeholders = append(placeholders, col + operator + "?")
		q.Values = append(q.Values, v[2])
	}

	q.WhereClause = "WHERE " + strings.Join(placeholders, " " + logicalOperator + " ")

	return q
}

func (q *Query) Get(callback func(rows *sql.Rows) error) error {
	if q.Error != nil {
		return &customErrors.QueryError{Err: q.Error}
	}

	rows, err := q.Repository.SqlDb().Query(q.compoundQuery(), q.Values...)

	if err != nil {
		return &customErrors.QueryExecError{Value: err.Error(), Query: q.compoundQuery()}
	}

	return callback(rows)
}

func (q *Query) Exec() (sql.Result, error) {
	if q.Error != nil {
		return nil, &customErrors.QueryError{Err: q.Error}
	}

	res, err := q.Repository.SqlDb().Exec(q.compoundQuery(), q.Values...)

	if err != nil {
		err = &customErrors.QueryExecError{Value: err.Error(), Query: q.compoundQuery()}
	}

	return res, err
}

func (q *Query) compoundQuery() string {
	return fmt.Sprintf("%s %s %s", q.Main, q.WhereClause, q.OrderByClause)
}


