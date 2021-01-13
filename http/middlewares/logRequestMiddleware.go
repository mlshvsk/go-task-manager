package middlewares

import (
	"bytes"
	"github.com/mlshvsk/go-task-manager/logger"
	"io/ioutil"
	"net/http"
)

func LogRequestMiddlewareFunc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.RequestLogger.Infow("Incoming request:",
			"url", r.RequestURI,
			"method", r.Method,
			"user-agent", r.UserAgent(),
			"body", getBody(r),
		)

		next.ServeHTTP(w, r)
	})
}

func getBody(r *http.Request) string {
	body, _ := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(body))
	return string(body)
}
