package controllers

import (
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/http/handlers"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
	"net/http"
)

type ProjectController struct {
}

func IndexProjects(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	projects, err := services.GetProjects()

	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, projects)
}

func StoreProject(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	var project models.Project
	if er := helpers.RetrieveModel(req.Body, &project); er != nil {
		return er
	}

	if err := services.StoreProject(&project); err != nil {
		if _, ok := err.(*customErrors.ModelAlreadyExists); ok == true {
			return &handlers.AppError{Error: err, Message: "Model already exists", ResponseCode: http.StatusBadRequest}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, project)
}

func ShowProject(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "projectId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	project, err := services.GetProject(id)
	if err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, project)
}

func UpdateProject(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "projectId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	var project models.Project
	if er := helpers.RetrieveModel(req.Body, &project); er != nil {
		return er
	}
	project.Id = id

	if err := services.UpdateProject(&project); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.EncodeResponse(rw, project)
}

func DeleteProject(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "projectId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	if err := services.DeleteProject(id); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return nil
}
