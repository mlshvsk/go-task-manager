package mysql

import (
	"github.com/mlshvsk/go-task-manager/database"
)

type Repository struct {
	*database.SqlDB
	TableName string
}

func (r *Repository) FindAll() *Query {
	q := &Query{r: r}

	return q.Select([]string{"*"})
}

func (r *Repository) Find(id int64) *Query {
	q := &Query{r: r}

	return q.Select([]string{"*"}).Where("and", [][]interface{}{{"id", "=", id}})
}

func (r *Repository) Create(data map[string]interface{}) (int64, error) {
	q := &Query{r: r}

	res, err := q.Insert(data).Exec()

	if err != nil {
		return 0, err
	}

	return res.LastInsertId()
}

func (r *Repository) Update(id int64, data map[string]interface{}) error {
	q := &Query{r: r}

	_, err := q.Update(data).Where("and", [][]interface{}{{"id", "=", id}}).Exec()

	return err
}

func (r *Repository) Delete(id int64) error {
	q := &Query{r: r}

	_, err := q.Delete().Where("and", [][]interface{}{{"id", "=", id}}).Exec()

	return err
}
