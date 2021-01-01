package repositories

import (
	"fmt"
	"github.com/mlshvsk/go-task-manager/database"
	"github.com/mlshvsk/go-task-manager/models"
)

type mysqlRepository struct {
	*database.Sqldb
	tableName string
	baseModel models.Model
}

func (r *mysqlRepository) FindAll() *Query {
	q := &Query{r: r}

	return q.Select([]string{"*"})
}

func (r *mysqlRepository) Find(id int) *Query {
	q := &Query{r: r}

	return q.Select([]string{"*"}).Where("and", [][]interface{}{{"id", "=", id}})
}

func (r *mysqlRepository) FindAllWhere(data [][]interface{}) *Query {
	q := &Query{r: r}

	return q.Select([]string{"*"}).Where("AND", data)
}

func (r *mysqlRepository) Create(data map[string]interface{}) (int, error) {
	q := &Query{r: r}

	res, err := q.Insert(data).Exec()

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	return int(id), err
}

func (r *mysqlRepository) Update(id int, data map[string]interface{}) error {
	q := &Query{r: r}

	_, err := q.Update(data).Where("AND", [][]interface{}{{"id", "=", id}}).Exec()

	return err
}

func (r *mysqlRepository) Delete(id int) error {
	q := &Query{r: r}

	_, err := q.Delete().Where("and", [][]interface{}{{"id", "=", id}}).Exec()

	return err
}
