package controllers

import (
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/http/handlers"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services/comment"
	"net/http"
)

func IndexComments(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	taskId, err := helpers.GetId(req, "taskId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	page, limit, err := helpers.GetPagination(req)
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	comments, err := comment.Service.GetCommentsByTask(taskId, page, limit)
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, comments)
}

func ShowComment(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	commentId, err := helpers.GetId(req, "commentId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	task, err := comment.Service.GetComment(commentId)
	if err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, task)
}

func StoreComment(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	taskId, err := helpers.GetId(req, "taskId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	var comment models.Comment
	if er := helpers.RetrieveModel(req.Body, &comment); er != nil {
		if _, ok := err.(*customErrors.ModelAlreadyExists); ok == true {
			return &handlers.AppError{Error: err, Message: "Model already exists", ResponseCode: http.StatusBadRequest}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
	}
	comment.TaskId = taskId

	if err := comment.CommentService.StoreComment(&comment); err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, comment)
}

func UpdateComment(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	var comment models.Comment
	id, err := helpers.GetId(req, "commentId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	if er := helpers.RetrieveModel(req.Body, &comment); er != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
	}
	comment.Id = id

	if err := comment.CommentService.UpdateComment(&comment); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, comment)
}

func DeleteComment(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "commentId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	if err := comment.Service.DeleteComment(id); err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return nil
}
