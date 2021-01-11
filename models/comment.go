package models

import "time"

type Comment struct {
	Id 			int64		`json:"id"`
	Data 		string		`json:"data" validate:"required,max=5000"`
	TaskId 		int64		`json:"task_id"`
	CreatedAt 	time.Time	`json:"created_at"`
}

type CommentRepository interface {
	FindAllByTask(taskId int64, offset int64, limit int64) ([]*Comment, error)
	Find(id int64) (*Comment, error)
	Create(c *Comment) error
	Update(t *Comment) error
	Delete(id int64) error
}

type CommentService interface {
	GetCommentsByTask(taskId int64, page int64, limit int64) ([]*Comment, error)
	GetComment(commentId int64) (*Comment, error)
	StoreComment(c *Comment) error
	UpdateComment(c *Comment) error
	DeleteComment(commentId int64) error
}
