package helpers

import (
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
)

func RequestBody(body io.ReadCloser) (reqBody []byte, err error) {
	reqBody, err = ioutil.ReadAll(body)
	defer body.Close()

	return
}

func RequestVar(req *http.Request, key string) string {
	vars := mux.Vars(req)
	res, _ := vars[key]

	return res
}
