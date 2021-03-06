package comment

import (
	"database/sql"
	"github.com/mlshvsk/go-task-manager/domains"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/repositories/base"
)

type commentRepository struct {
	base base.Repository
}

func InitCommentRepository(baseRepo base.Repository) *commentRepository {
	r := &commentRepository{
		base: baseRepo,
	}

	r.base.SetTableName("comments")

	return r
}

func (cr *commentRepository) FindAllByTask(taskId int64, offset int64, limit int64) ([]*domains.CommentModel, error) {
	comments := make([]*domains.CommentModel, 0)

	err := cr.base.
		FindAll().
		Where("and", [][]interface{}{{"task_id", "=", taskId}}).
		OrderBy("created_at", "desc").
		Limit(offset, limit).
		Get(cr.scan(&comments))

	if err != nil {
		return nil, err
	}

	return comments, nil
}

func (cr *commentRepository) Find(id int64) (*domains.CommentModel, error) {
	comments := make([]*domains.CommentModel, 0)
	err := cr.base.Find(id).Get(cr.scan(&comments))

	if err != nil {
		return nil, err
	}

	if len(comments) == 0 {
		return nil, &customErrors.NotFoundError{Value: "comment not found"}
	}

	return comments[0], nil
}

func (cr *commentRepository) Create(c *domains.CommentModel) error {
	id, err := cr.base.Create(map[string]interface{}{"data": &c.Data, "task_id": &c.TaskId, "created_at": &c.CreatedAt})

	if err != nil {
		return err
	}

	c.Id = id
	return nil
}

func (cr *commentRepository) Update(c *domains.CommentModel) error {
	err := cr.base.Update(c.Id, map[string]interface{}{
		"data":    &c.Data,
		"task_id": &c.TaskId,
	})

	if err != nil {
		return err
	}

	return nil
}

func (cr *commentRepository) Delete(id int64) error {
	return cr.base.Delete(id)
}

func (cr *commentRepository) scan(comments *[]*domains.CommentModel) func(rows *sql.Rows) error {
	return func(rows *sql.Rows) error {
		for rows.Next() {
			comment := new(domains.CommentModel)
			if err := rows.Scan(&comment.Id, &comment.Data, &comment.TaskId, &comment.CreatedAt); err != nil {
				return err
			}

			*comments = append(*comments, comment)
		}

		defer rows.Close()

		return nil
	}
}
