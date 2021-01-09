package helpers

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/mlshvsk/go-task-manager/http/handlers"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func RequestBody(body io.ReadCloser) (reqBody []byte, err error) {
	reqBody, err = ioutil.ReadAll(body)
	defer body.Close()

	return
}

func RequestVar(req *http.Request, key string) (string, bool) {
	vars := mux.Vars(req)
	val, ok := vars[key]

	return val, ok
}

func GetId(req *http.Request, idString string) (int64, error) {
	id, ok := RequestVar(req, idString)
	if ok == false {
		return 0, errors.New("id is not found in request")
	}

	return strconv.ParseInt(id, 10, 64)
}

func EncodeResponse(rw http.ResponseWriter, res interface{}) *handlers.AppError {
	if err := json.NewEncoder(rw).Encode(res); err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	return nil
}

func RetrieveModel(reqBody io.ReadCloser, model interface{}) *handlers.AppError {
	requestBody, err := RequestBody(reqBody)
	if err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	if err := json.Unmarshal(requestBody, &model); err != nil {
		return &handlers.AppError{Error: err, ResponseCode: http.StatusInternalServerError}
	}

	if err := validator.New().Struct(model); err != nil {
		return &handlers.AppError{Error: err, Message: err.Error(), ResponseCode: http.StatusBadRequest}
	}

	return nil
}
