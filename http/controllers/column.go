package controllers

import (
	"github.com/mlshvsk/go-task-manager/domains"
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/http/handlers"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/http/transformers"
	"github.com/mlshvsk/go-task-manager/services"
	"net/http"
)

func IndexColumns(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	projectId, err := helpers.GetId(req, "projectId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	page, limit, e := helpers.GetPagination(req)
	if e != nil {
		return e
	}

	columns, err := services.ColumnService.GetColumns(projectId, page, limit)
	if err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	res, err := transformers.ExtendColumns(columns, req.URL.Query().Get("include_tasks"))
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, res)
}

func ShowColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	columnId, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	column, err := services.ColumnService.GetColumn(columnId)
	if err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	res, err := transformers.ExtendColumn(column, req.URL.Query().Get("include_tasks"))
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, res)
}

func StoreColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	projectId, err := helpers.GetId(req, "projectId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	var column domains.ColumnModel
	if er := helpers.RetrieveModel(req.Body, &column); er != nil {
		return er
	}
	column.ProjectId = projectId

	if err := services.ColumnService.StoreColumn(&column); err != nil {
		if _, ok := err.(*customErrors.ModelAlreadyExists); ok == true {
			return &handlers.AppError{Error: err, Message: "Model already exists", ResponseCode: http.StatusBadRequest}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, column)
}

func DeleteColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	if err := services.ColumnService.DeleteColumn(id); err != nil {
		if _, ok := err.(*customErrors.LastModelDeletion); ok == true {
			return &handlers.AppError{Error: err, Message: "Cannot delete last project column", ResponseCode: http.StatusBadRequest}
		}

		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return nil
}

func UpdateColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	var column domains.ColumnModel
	if er := helpers.RetrieveModel(req.Body, &column); er != nil {
		return er
	}
	column.Id = id

	if err := services.ColumnService.UpdateColumn(&column); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, column)
}

func MoveColumn(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "columnId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	body := struct {
		Direction string `json:"direction" validate:"required,oneof=left right"`
	}{}
	if er := helpers.RetrieveModel(req.Body, &body); er != nil {
		return er
	}

	if err := services.ColumnService.MoveColumn(id, body.Direction); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return nil
}
