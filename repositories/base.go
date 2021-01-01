package repositories

import (
	"database/sql"
	"fmt"
	"github.com/mlshvsk/go-task-manager/database"
	"github.com/mlshvsk/go-task-manager/models"
	"strings"
)

type Repository interface {
	FindAll() ([]models.Model, error)
	Find(id int64)
	Create(m *models.Model) (*models.Model, error)
	Delete(id int64) error
}

type baseRepository struct {
	*database.Sqldb
	tableName string
	baseModel models.Model
}

func (r *baseRepository) FindAll() (*sql.Rows, error) {
	return r.Conn.Query("SELECT * FROM " + r.tableName)
}

func (r *baseRepository) Find(id int) (*sql.Rows, error) {
	return r.Conn.Query("SELECT * FROM " + r.tableName + " WHERE id = ?", id)
}

func (r *baseRepository) FindAllWhere(constraints string, values ...interface{}) (*sql.Rows, error) {
	return r.Conn.Query("SELECT * FROM " + r.tableName + " WHERE ", values...)
}

func (r *baseRepository) Create(fields []string, values ...interface{}) (int, error) {
	fn, pl := r.prepareInsertQuery(fields)

	res, err := r.Conn.Exec("INSERT INTO " + r.tableName + "(" + fn + ") VALUES (" + pl + ");", values...)

	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	return int(id), err
}

func (r *baseRepository) Update(id int, fields []string, values ...interface{}) error {
	pl := r.prepareUpdateQuery(fields)
	params := append(values, id)
	fmt.Println("UPDATE " + r.tableName + " SET " + pl + " WHERE id = ?")
	fmt.Println(params)

	_, err := r.Conn.Exec("UPDATE " + r.tableName + " SET " + pl + " WHERE id = ?", params...)

	return err
}

func (r *baseRepository) Delete(id int) error {
	_, err := r.Conn.Exec("DELETE FROM " + r.tableName + " WHERE id = ?", id)

	return err
}

func (r *baseRepository) prepareInsertQuery(fields []string) (fieldNames string, placeholders string) {
	fieldNames = strings.Join(fields, ", ")

	pl := make([]string, len(fields))
	for i := range pl {
		pl[i] = "?"
	}
	placeholders = strings.Join(pl, ", ")

	return
}

func (r *baseRepository) prepareUpdateQuery(fields []string) string {
	pl := make([]string, len(fields))
	for i, v := range fields {
		pl[i] = v + " = ?"

		if (i + 1) != len(fields) {
			pl[i] += ","
		}
	}

	return strings.Join(pl, " ")
}
