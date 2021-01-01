package repositories

import (
	"database/sql"
	"fmt"
	"github.com/mlshvsk/go-task-manager/database"
	"github.com/mlshvsk/go-task-manager/models"
)

type commentRepository struct {
	mysqlRepository
}

var CommentRepository *commentRepository

func InitCommentRepository(db *database.Sqldb) {
	CommentRepository = &commentRepository{mysqlRepository{db, "comments", models.Comment{}}}
}

func (cr *commentRepository) FindAll(taskId int) ([]*models.Comment, error) {
	comments := make([]*models.Comment, 0)

	err := cr.mysqlRepository.FindAllWhere([][]interface{}{{"task_id", "=", taskId}}).Get(cr.Scan(&comments))

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (cr *commentRepository) Find(id int) (*models.Comment, error) {
	comments := make([]*models.Comment, 0)
	err := cr.mysqlRepository.Find(id).Get(cr.Scan(&comments))

	if err != nil {
		fmt.Printf("Error retrieving all projects: %v", err.Error())
		return nil, err
	}

	return comments[0], nil
}

func (cr *commentRepository) Create(c *models.Comment) error {
	id, err := cr.mysqlRepository.Create(map[string]interface{}{"data": &c.Data, "task_id": &c.TaskId})

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	c.Id = id
	return nil
}

func (cr *commentRepository) Scan(comments *[]*models.Comment) func(rows *sql.Rows) error {
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
