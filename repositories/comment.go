package repositories

import (
	"database/sql"
	"fmt"
	"github.com/mlshvsk/go-task-manager/database"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories/mysql"
)

type commentRepository struct {
	mysql.Repository
}

var CommentRepository *commentRepository

func InitCommentRepository(db *database.SqlDB) {
	CommentRepository = &commentRepository{
		Repository: mysql.Repository{
			SqlDB: db,
			TableName: "comments",
		},
	}
}

func (cr *commentRepository) FindAllByTask(taskId int64) ([]*models.Comment, error) {
	comments := make([]*models.Comment, 0)

	err := cr.Repository.
		FindAll().
		Where("and", [][]interface{}{{"task_id", "=", taskId}}).
		OrderBy("created_at", "desc").
		Get(cr.scan(&comments))

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (cr *commentRepository) Find(id int64) (*models.Comment, error) {
	comments := make([]*models.Comment, 0)
	err := cr.Repository.Find(id).Get(cr.scan(&comments))

	if err != nil {
		return nil, err
	}

	if comments == nil || len(comments) == 0 {
		return nil, &customErrors.NotFoundError{Value: "comment not found"}
	}

	return comments[0], nil
}

func (cr *commentRepository) Create(c *models.Comment) error {
	id, err := cr.Repository.Create(map[string]interface{}{"data": &c.Data, "task_id": &c.TaskId, "created_at": &c.CreatedAt})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	c.Id = id
	return nil
}

func (cr *commentRepository) Update(t *models.Comment) error {
	err := cr.Repository.Update(t.Id, map[string]interface{}{
		"data": &t.Data,
		"task_id": &t.TaskId,
	})

	if err != nil {
		return err
	}

	return nil
}

func (cr *commentRepository) scan(comments *[]*models.Comment) func(rows *sql.Rows) error {
	return func(rows *sql.Rows) error {
		for rows.Next() {
			comment := new(models.Comment)
			if err := rows.Scan(&comment.Id, &comment.Data, &comment.TaskId, &comment.CreatedAt); err != nil {
				return err
			}

			*comments = append(*comments, comment)
		}

		defer rows.Close()

		return nil
	}
}
