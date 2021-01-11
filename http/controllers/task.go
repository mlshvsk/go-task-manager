package controllers

import (
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/http/handlers"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
	"net/http"
)

func IndexTasksByColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	columnId, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	page, limit, err := helpers.GetPagination(req)
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	tasks, err := services.TaskService.GetTasksByColumn(columnId, page, limit)
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, tasks)
}

func IndexTasks(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	page, limit, err := helpers.GetPagination(req)
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	tasks, err := services.TaskService.GetTasks(page, limit)
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, tasks)
}

func StoreTask(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	columnId, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	var task models.Task
	if er := helpers.RetrieveModel(req.Body, &task); er != nil {
		if _, ok := err.(*customErrors.ModelAlreadyExists); ok == true {
			return &handlers.AppError{Error: err, Message: "Model already exists", ResponseCode: http.StatusBadRequest}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
	}
	task.ColumnId = columnId

	if err := services.TaskService.StoreTask(&task); err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, task)
}

func ShowTask(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	taskId, err := helpers.GetId(req, "taskId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	task, err := services.TaskService.GetTask(taskId)
	if err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, task)
}

func UpdateTask(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	var task models.Task
	id, err := helpers.GetId(req, "taskId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	if er := helpers.RetrieveModel(req.Body, &task); er != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
	}
	task.Id = id

	if err := services.TaskService.UpdateTask(&task); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, task)
}

func MoveTask(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "taskId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	body := struct{Direction string `json:"direction"`}{}
	if er := helpers.RetrieveModel(req.Body, &body); er != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
	}

	if err := services.TaskService.MoveTaskWithinColumn(id, body.Direction); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return nil
}

func MoveTaskColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "taskId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	columnId, err := helpers.GetId(req, "newColumnId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	if err := services.TaskService.MoveTaskToColumn(id, columnId); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return nil
}

func DeleteTask(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	if err := services.TaskService.DeleteTask(id); err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return nil
}
