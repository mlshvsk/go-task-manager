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
	"net/url"
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

func GetPagination(req *http.Request) (int64, int64, *handlers.AppError) {
	query := req.URL.Query()
	page, err := getPage(query)
	if err != nil {
		return 0, 0, &handlers.AppError{Error: err, Message: "Invalid page parameter", ResponseCode: http.StatusBadRequest}
	}
	limit, err := getLimit(query)
	if err != nil {
		return 0, 0, &handlers.AppError{Error: err, Message: "Invalid limit parameter", ResponseCode: http.StatusBadRequest}
	}

	return page, limit, nil
}

func getPage(query url.Values) (int64, error) {
	pageParam := query.Get("page")

	if len(pageParam) == 0 {
		return 0, nil
	}

	pageParsed, err := strconv.ParseInt(pageParam, 10, 64)
	if err != nil {
		return 0, errors.New("cannot parse page parameter")
	}
	if pageParsed <= 0 {
		return 0, errors.New("invalid page number")
	}

	return pageParsed - 1, nil
}

func getLimit(query url.Values) (int64, error) {
	limitParam := query.Get("limit")

	if len(limitParam) == 0 {
		return 10, nil
	}

	limitParsed, err := strconv.ParseInt(limitParam, 10, 64)
	if err != nil {
		return 0, errors.New("cannot parse limit parameter")
	}
	if limitParsed <= 0 {
		return 0, errors.New("invalid limit number")
	}

	return limitParsed, nil
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
