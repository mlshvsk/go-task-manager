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

type Response struct {
	Data interface{} `json:"data"`
}

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

func GetPagination(req *http.Request) (int64, int64, error) {
	limit, ok := RequestVar(req, "limit")
	if ok == false {
		return 0, 0, errors.New("limit is not found in request")
	}

	page, ok := RequestVar(req, "page")
	if ok == false {
		return 0, 0, errors.New("page is not found in request")
	}

	limitParsed, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		return 0, 0, errors.New("cannot parse limit")
	}

	pageParsed, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		return 0, 0, errors.New("cannot parse page")
	}

	return pageParsed, limitParsed, nil
}

func GetId(req *http.Request, idString string) (int64, error) {
	id, ok := RequestVar(req, idString)
	if ok == false {
		return 0, errors.New("id is not found in request")
	}

	return strconv.ParseInt(id, 10, 64)
}

func PrepareResponse(rw http.ResponseWriter, res interface{}) *handlers.AppError {
	data := &Response{
		Data: res,
	}

	if err := json.NewEncoder(rw).Encode(data); err != nil {
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
