package controllers

import (
	"encoding/json"
	"github.com/mlshvsk/go-task-manager/helpers"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
	"net/http"
	"strconv"
)

func IndexColumns(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "projectId"), 10, 10)

	json.NewEncoder(rw).Encode(services.GetColumns(int(i)))
}

func ShowColumn(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "columnId"), 10, 10)

	json.NewEncoder(rw).Encode(services.GetColumn(int(i)))
}

func StoreColumn(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "projectId"), 10, 10)
	requestBody, err := helpers.RequestBody(req.Body)
	var body *models.Column

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	_ = json.Unmarshal(requestBody, &body)
	body.ProjectId = int(i)

	services.StoreColumn(body)
}

func DeleteColumn(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	id, _ := strconv.ParseInt(helpers.RequestVar(req, "columnId"), 10, 10)
	projectId, _ := strconv.ParseInt(helpers.RequestVar(req, "projectId"), 10, 10)

	if err := services.DeleteColumn(int(id), int(projectId)); err != nil {
		json.NewEncoder(rw).Encode(err.Error())
	}
}

func UpdateColumn(rw http.ResponseWriter, req *http.Request) {
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "columnId"), 10, 10)
	var body models.Column
	requestBody, err := helpers.RequestBody(req.Body)
	err = json.Unmarshal(requestBody, &body)
	body.Id = int(i)
	services.UpdateColumn(&body)

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(rw).Encode(body)
}

func MoveColumn(rw http.ResponseWriter, req *http.Request) {
	body := struct{Direction string `json:"direction"`}{}
	rw.Header().Set("Content-Type", "application/json")
	id, _ := strconv.ParseInt(helpers.RequestVar(req, "columnId"), 10, 10)
	projectId, _ := strconv.ParseInt(helpers.RequestVar(req, "projectId"), 10, 10)
	requestBody, _ := helpers.RequestBody(req.Body)
	_ = json.Unmarshal(requestBody, &body)

	if err := services.MoveColumn(int(id), int(projectId), body.Direction); err != nil {
		json.NewEncoder(rw).Encode(err.Error())
	}
}
