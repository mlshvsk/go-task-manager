package controllers

import (
	"encoding/json"
	"github.com/mlshvsk/go-task-manager/helpers"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
	"net/http"
	"strconv"
)

func IndexTasks(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "columnId"), 10, 10)

	json.NewEncoder(rw).Encode(services.GetTasks(int(i)))
}

func StoreTask(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "columnId"), 10, 10)
	requestBody, err := helpers.RequestBody(req.Body)
	var body *models.Task

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	_ = json.Unmarshal(requestBody, &body)
	body.ColumnId = int(i)

	services.StoreTask(body)
}

func ShowTask(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "taskId"), 10, 10)
	p, _ := services.GetTask(int(i))
	json.NewEncoder(rw).Encode(p)
}

func UpdateTask(rw http.ResponseWriter, req *http.Request) {
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "taskId"), 10, 10)
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

func MoveTask(rw http.ResponseWriter, req *http.Request) {
	body := struct{Direction string `json:"direction"`}{}
	rw.Header().Set("Content-Type", "application/json")
	id, _ := strconv.ParseInt(helpers.RequestVar(req, "taskId"), 10, 10)
	columnId, _ := strconv.ParseInt(helpers.RequestVar(req, "columnId"), 10, 10)
	requestBody, _ := helpers.RequestBody(req.Body)
	_ = json.Unmarshal(requestBody, &body)

	if err := services.MoveTask(int(id), int(columnId), body.Direction); err != nil {
		json.NewEncoder(rw).Encode(err.Error())
	}
}

func DeleteTask(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	id, _ := strconv.ParseInt(helpers.RequestVar(req, "taskId"), 10, 10)

	services.DeleteTask(int(id))
}
