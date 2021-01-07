package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	errors2 "github.com/mlshvsk/go-task-manager/errors"
	"strings"
)

type Query struct {
	main string
	orderBy string
	where string
	values []interface{}
	r *Repository
	error error
}

func (q *Query) Select(columns []string) *Query {
	q.main = fmt.Sprintf("SELECT %v FROM %v ", strings.Join(columns, ", "), q.r.TableName)

	return q
}

func (q *Query) Delete() *Query {
	q.main = fmt.Sprintf("DELETE FROM %v ", q.r.TableName)

	return q
}

func (q *Query) Update(data map[string]interface{}) *Query {
	var fieldPlaceholders string
	var pl []string

	for i, v := range data {
		pl = append(pl, i + " = ?")
		q.values = append(q.values, v)
	}

	fieldPlaceholders = strings.Join(pl, ", ")

	q.main = fmt.Sprintf("UPDATE %v SET %v", q.r.TableName, fieldPlaceholders)

	return q
}

func (q *Query) Insert(data map[string]interface{}) *Query {
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

	q.main = fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", q.r.TableName, fieldNames, placeholders)
	q.values = values

	return q
}

func (q *Query) OrderBy(colName string, order string) *Query {
	q.orderBy += " ORDER BY " + colName + " " + strings.ToUpper(order)

	return q
}

func (q *Query) Where(logicalOperator string, data [][]interface{}) *Query {
	var placeholders []string

	for _, v := range data {
		if len(v) != 3 {
			q.error = errors.New("where: array len of where clause must be equal to 3")
			return q
		}

		col, ok := v[0].(string)
		if ok == false {
			q.error = errors.New("where: cannot convert column name to string")
			return q
		}

		operator, ok := v[1].(string)
		if ok == false {
			q.error = errors.New("where: cannot convert operator to string")
			return q
		}

		placeholders = append(placeholders, col + operator + "?")
		q.values = append(q.values, v[2])
	}

	q.where = " WHERE " + strings.Join(placeholders, " " + logicalOperator + " ")

	return q
}

func (q *Query) Get(callback func(rows *sql.Rows) error) error {
	if q.error != nil {
		return &errors2.QueryError{Err: q.error}
	}

	rows, err := q.r.Conn.Query(q.compoundQuery(), q.values...)

	if err != nil {
		return &errors2.ExecError{Value: err.Error(), Query: q.compoundQuery()}
	}

	return callback(rows)
}

func (q *Query) Exec() (sql.Result, error) {
	if q.error != nil {
		return nil, &errors2.QueryError{Err: q.error}
	}

	res, err := q.r.Conn.Exec(q.compoundQuery(), q.values...)

	if err != nil {
		err = &errors2.ExecError{Value: err.Error(), Query: q.compoundQuery()}
	}

	return res, err
}

func (q *Query) compoundQuery() string {
	return q.main + q.where + q.orderBy
}


