package services

import (
	"fmt"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/repositories"
)

type CommentService struct {
}

func GetComments(taskId int) []*models.Comment {
	res, _ := repositories.CommentRepository.FindAll(taskId)
	return res
}

func GetComment(commentId int) *models.Comment {
	res, _ := repositories.CommentRepository.Find(commentId)
	return res
}

func StoreComment(c *models.Comment) *models.Comment {
	err := repositories.CommentRepository.Create(c)

	if err != nil {
		fmt.Println(err.Error())
	}

	return c
}

func DeleteComment(commentId int) {
	repositories.CommentRepository.Delete(commentId)
}
