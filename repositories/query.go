package repositories

import (
	"database/sql"
	"fmt"
	"strings"
)

type Query struct {
	main string
	query string
	orderBy string
	where string
	values []interface{}
	r *mysqlRepository
	errors []error
}

func (q *Query) Select(columns []string) *Query {
	q.main = fmt.Sprintf("SELECT %v FROM %v ", strings.Join(columns, ", "), q.r.tableName)

	return q
}

func (q *Query) Delete() *Query {
	q.main = fmt.Sprintf("DELETE FROM %v ", q.r.tableName)

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

	q.main = fmt.Sprintf("UPDATE %v SET %v", q.r.tableName, fieldPlaceholders)

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
		q.values = append(values, v)
	}

	placeholders = strings.Join(pl, ", ")
	fieldNames = strings.Join(fn, ", ")

	q.main = fmt.Sprintf("INSERT INTO %v (%v) VALUES (%v)", q.r.tableName, fieldNames, placeholders)

	return q
}

func (q *Query) OrderBy(colName string, order string) *Query {
	q.orderBy += " ORDER BY " + colName + " " + strings.ToUpper(order)

	return q
}

func (q *Query) Where(logicalOperator string, data [][]interface{}) *Query {
	var placeholders []string

	for _, v := range data {
		col, _ := v[0].(string)
		operator, _ := v[1].(string)

		placeholders = append(placeholders, col + operator + "?")
		q.values = append(q.values, v[2])
	}

	q.where = " WHERE " + strings.Join(placeholders, " " + logicalOperator + " ")

	return q
}

func (q *Query) Get(callback func(rows *sql.Rows) error) error {
	rows, err := q.r.Conn.Query(q.main + q.where + q.orderBy, q.values...)

	if err != nil {
		return err
	}

	return callback(rows)
}

func (q *Query) Exec() (sql.Result, error) {
	return q.r.Conn.Exec(q.query, q.values...)
}


