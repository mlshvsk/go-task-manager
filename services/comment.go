package services

import (
	"github.com/mlshvsk/go-task-manager/factories"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories"
)

type CommentService struct {
}

func GetCommentsByTask(taskId int64) ([]*models.Comment, error) {
	return repositories.CommentRepository.FindAllByTask(taskId)
}

func GetComment(commentId int64) (*models.Comment, error) {
	return repositories.CommentRepository.Find(commentId)
}

func StoreComment(c *models.Comment) error {
	var err error
	*c, err = factories.CommentFactory(c.TaskId, c.Data)
	if err != nil {
		return err
	}

	return repositories.CommentRepository.Create(c)
}

func UpdateComment(c *models.Comment) error {
	commentFromDB, err := repositories.CommentRepository.Find(c.Id)
	if err != nil {
		return err
	}

	c.TaskId = commentFromDB.TaskId
	c.CreatedAt = commentFromDB.CreatedAt
	return repositories.CommentRepository.Update(c)
}

func DeleteComment(commentId int64) error {
	_, err := repositories.CommentRepository.Find(commentId)
	if err != nil {
		return err
	}

	return repositories.CommentRepository.Delete(commentId)
}
