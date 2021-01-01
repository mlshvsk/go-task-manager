package controllers

import (
	"encoding/json"
	"github.com/mlshvsk/go-task-manager/helpers"
	"github.com/mlshvsk/go-task-manager/models"
	"github.com/mlshvsk/go-task-manager/services"
	"net/http"
	"strconv"
)

func IndexComments(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "taskId"), 10, 10)

	json.NewEncoder(rw).Encode(services.GetComments(int(i)))
}

func ShowComment(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "commentId"), 10, 10)
	p := services.GetComment(int(i))
	json.NewEncoder(rw).Encode(p)
}

func StoreComment(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	i, _ := strconv.ParseInt(helpers.RequestVar(req, "taskId"), 10, 10)
	requestBody, err := helpers.RequestBody(req.Body)
	var body *models.Comment

	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
	}

	_ = json.Unmarshal(requestBody, &body)
	body.TaskId = int(i)

	services.StoreComment(body)
}

func DeleteComment(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	id, _ := strconv.ParseInt(helpers.RequestVar(req, "commentId"), 10, 10)

	services.DeleteComment(int(id))
}
