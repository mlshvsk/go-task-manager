package controllers

import (
	customErrors "github.com/mlshvsk/go-task-manager/errors"
	"github.com/mlshvsk/go-task-manager/http/handlers"
	"github.com/mlshvsk/go-task-manager/http/helpers"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services/project"
	"net/http"
)

type ProjectController struct {
}

func IndexProjects(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	page, limit, err := helpers.GetPagination(req)
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	projects, err := project.Service.GetProjects(page, limit)

	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, projects)
}

func StoreProject(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	var project models.Project
	if er := helpers.RetrieveModel(req.Body, &project); er != nil {
		return er
	}

	if err := project.ProjectService.StoreProject(&project); err != nil {
		if _, ok := err.(*customErrors.ModelAlreadyExists); ok == true {
			return &handlers.AppError{Error: err, Message: "Model already exists", ResponseCode: http.StatusBadRequest}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, project)
}

func ShowProject(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "projectId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	project, err := project.Service.GetProject(id)
	if err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}

		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, project)
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

	if err := project.ProjectService.UpdateProject(&project); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return helpers.PrepareResponse(rw, project)
}

func DeleteProject(rw http.ResponseWriter, req *http.Request) *handlers.AppError {
	id, err := helpers.GetId(req, "projectId")
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	if err := project.Service.DeleteProject(id); err != nil {
		if _, ok := err.(*customErrors.NotFoundError); ok == true {
			return &handlers.AppError{Error: err, ResponseCode: http.StatusNotFound}
		}
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return nil
}
