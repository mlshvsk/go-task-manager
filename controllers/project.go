package controllers

import (
	"encoding/json"
	"github.com/mlshvsk/go-task-manager/helpers"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
	"net/http"
	"strconv"
)

type ProjectController struct {
}

func IndexProjects(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	p, _ := services.GetProjects()
	json.NewEncoder(rw).Encode(p)
}

func StoreProject(rw http.ResponseWriter, req *http.Request) {
	var body models.Project
	requestBody, err := helpers.RequestBody(req.Body)
	err = json.Unmarshal(requestBody, &body)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	services.StoreProject(&body)
	json.NewEncoder(rw).Encode(body)
}

func ShowProject(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "projectId"), 10, 10)
	p, _ := services.GetProject(int(i))
	json.NewEncoder(rw).Encode(p)
}

func UpdateProject(rw http.ResponseWriter, req *http.Request) {
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "projectId"), 10, 10)
	var body models.Project
	requestBody, err := helpers.RequestBody(req.Body)
	err = json.Unmarshal(requestBody, &body)
	body.Id = int(i)
	services.UpdateProject(&body)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode(body)
}

func DeleteProject(rw http.ResponseWriter, req *http.Request) {
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "projectId"), 10, 10)
	services.DeleteProject(int(i))
}
