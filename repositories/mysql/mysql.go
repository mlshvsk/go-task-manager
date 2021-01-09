package mysql

import (
	"database/sql"
	"github.com/mlshvsk/go-task-manager/repositories/base"
)

type Repository base.BaseRepository

func (r *Repository) GetTableName() string {
	return r.TableName
}

func (r *Repository) SetTableName(tableName string) {
	r.TableName = tableName
}

func (r *Repository) SqlDb() *sql.DB {
	return r.Conn
}

func (r *Repository) FindAll() base.Query {
	q := &Query{Repository: r}

	return q.Select([]string{"*"})
}

func (r *Repository) Find(id int64) base.Query {
	q := &Query{Repository: r}

	return q.Select([]string{"*"}).Where("and", [][]interface{}{{"id", "=", id}})
}

func (r *Repository) Create(data map[string]interface{}) (int64, error) {
	q := &Query{Repository: r}

	res, err := q.Insert(data).Exec()

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *Repository) Update(id int64, data map[string]interface{}) error {
	q := &Query{Repository: r}

	_, err := q.Update(data).Where("and", [][]interface{}{{"id", "=", id}}).Exec()

	return err
}

func (r *Repository) Delete(id int64) error {
	q := &Query{Repository: r}

	_, err := q.Delete().Where("and", [][]interface{}{{"id", "=", id}}).Exec()

	return err
}
