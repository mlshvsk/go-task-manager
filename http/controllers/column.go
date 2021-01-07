package controllers

import (
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/http/handlers"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
	"net/http"
)

func IndexColumns(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	projectId, err := helpers.GetId(req, "projectId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	columns, err := services.GetColumns(projectId)
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, columns)
}

func ShowColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	columnId, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	column, err := services.GetColumn(columnId)
	if err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, Code: http.StatusNotFound}
		}

		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, column)
}

func StoreColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	projectId, err := helpers.GetId(req, "projectId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	var column models.Column
	if er := helpers.RetrieveModel(req.Body, &column); er != nil {
		return &handlers.AppError{Error: err, Code: http.StatusNotFound}
	}
	column.ProjectId = projectId


	if err := services.StoreColumn(&column); err != nil {
		if _, ok := err.(*customErrors.ModelAlreadyExists); ok == true {
			return &handlers.AppError{Error: err, Message: "Model already exists", Code: http.StatusBadRequest}
		}

		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, column)
}

func DeleteColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	if err := services.DeleteColumn(id); err != nil {
		if _, ok := err.(*customErrors.LastModelDeletion); ok == true {
			return &handlers.AppError{Error: err, Message: "Cannot delete last project column", Code: http.StatusBadRequest}
		}

		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return nil
}

func UpdateColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	var column models.Column
	if er := helpers.RetrieveModel(req.Body, &column); er != nil {
		return &handlers.AppError{Error: err, Code: http.StatusNotFound}
	}
	column.Id = id

	if err := services.UpdateColumn(&column); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, Code: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, column)
}

func MoveColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	body := struct{Direction string `json:"direction"`}{}
	if er := helpers.RetrieveModel(req.Body, &body); er != nil {
		return &handlers.AppError{Error: err, Code: http.StatusNotFound}
	}

	if err := services.MoveColumn(id, body.Direction); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, Code: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, Code: http.StatusInternalServerError}
	}

	return nil
}
