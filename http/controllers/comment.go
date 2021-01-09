package controllers

import (
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/http/handlers"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
	"net/http"
)

func IndexComments(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	commentId, err := helpers.GetId(req, "commentId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	comments, err := services.GetCommentsByTask(commentId)
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, comments)
}

func ShowComment(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	commentId, err := helpers.GetId(req, "commentId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	task, err := services.GetComment(commentId)
	if err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, task)
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

	if err := services.StoreComment(&comment); err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, comment)
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

	if err := services.UpdateComment(&comment); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, comment)
}

func DeleteComment(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "commentId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	if err := services.DeleteComment(id); err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return nil
}
