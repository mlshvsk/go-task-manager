package domains

import "time"

type CommentModel struct {
	Id        int64     `json:"id"`
	Data      string    `json:"data" validate:"required,max=5000"`
	TaskId    int64     `json:"task_id"`
	CreatedAt time.Time `json:"created_at"`
}

type CommentRepository interface {
	FindAllByTask(taskId int64, offset int64, limit int64) ([]*CommentModel, error)
	Find(id int64) (*CommentModel, error)
	Create(c *CommentModel) error
	Update(t *CommentModel) error
	Delete(id int64) error
}

type CommentService interface {
	GetCommentsByTask(taskId int64, page int64, limit int64) ([]*CommentModel, error)
	GetComment(commentId int64) (*CommentModel, error)
	StoreComment(c *CommentModel) error
	UpdateComment(c *CommentModel) error
	DeleteComment(commentId int64) error
}
