package comment

import (
	"github.com/mlshvsk/go-task-manager/factories"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
	"sync"
)

type commentService struct {
	r models.CommentRepository
}

func InitCommentService(r models.CommentRepository) {
	(&sync.Once{}).Do(func() {
		services.CommentService = &commentService{r}
	})
}

func (s *commentService) GetCommentsByTask(taskId int64, page int64, limit int64) ([]*models.Comment, error) {
	return s.r.FindAllByTask(taskId, page, limit)
}

func (s *commentService) GetComment(commentId int64) (*models.Comment, error) {
	return s.r.Find(commentId)
}

func (s *commentService) StoreComment(c *models.Comment) error {
	var err error
	*c, err = factories.CommentFactory(c.TaskId, c.Data)
	if err != nil {
		return err
	}

	return s.r.Create(c)
}

func (s *commentService) UpdateComment(c *models.Comment) error {
	commentFromDB, err := s.r.Find(c.Id)
	if err != nil {
		return err
	}

	c.TaskId = commentFromDB.TaskId
	c.CreatedAt = commentFromDB.CreatedAt
	return s.r.Update(c)
}

func (s *commentService) DeleteComment(commentId int64) error {
	_, err := s.r.Find(commentId)
	if err != nil {
		return err
	}

	return s.r.Delete(commentId)
}
