package controllers

import (
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/http/handlers"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
	"net/http"
)

func IndexTasks(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	columnId, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	tasks, err := services.GetTasksByColumn(columnId)
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, tasks)
}

func StoreTask(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	columnId, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	var task models.Task
	if er := helpers.RetrieveModel(req.Body, &task); er != nil {
		if _, ok := err.(*customErrors.ModelAlreadyExists); ok == true {
			return &handlers.AppError{Error: err, Message: "Model already exists", Code: http.StatusBadRequest}
		}

		return &handlers.AppError{Error: err, Code: http.StatusNotFound}
	}
	task.ColumnId = columnId

	if err := services.StoreTask(&task); err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, task)
}

func ShowTask(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	taskId, err := helpers.GetId(req, "taskId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	task, err := services.GetTask(taskId)
	if err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, Code: http.StatusNotFound}
		}

		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, task)
}

func UpdateTask(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	var task models.Task
	id, err := helpers.GetId(req, "taskId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	if er := helpers.RetrieveModel(req.Body, &task); er != nil {
		return &handlers.AppError{Error: err, Code: http.StatusNotFound}
	}
	task.Id = id

	if err := services.UpdateTask(&task); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, Code: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, task)
}

func MoveTask(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "taskId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	body := struct{Direction string `json:"direction"`}{}
	if er := helpers.RetrieveModel(req.Body, &body); er != nil {
		return &handlers.AppError{Error: err, Code: http.StatusNotFound}
	}

	if err := services.MoveTaskWithinColumn(id, body.Direction); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, Code: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return nil
}

func MoveTaskColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "taskId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	columnId, err := helpers.GetId(req, "newColumnId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	if err := services.MoveTaskToColumn(id, columnId); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, Code: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return nil
}

func DeleteTask(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	if err := services.DeleteTask(id); err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return nil
}