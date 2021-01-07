package handlers

import (
	"github.com/mlshvsk/go-task-manager/logger"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) *AppError

type AppError struct {
	Error   error
	Message string
	Code    int
}

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		logger.ErrorLogger.Error(e.Error.Error())

		if e.Message == "" {
			switch e.Code {
			case http.StatusInternalServerError:
				e.Message = "Internal Server Error"
			case http.StatusNotFound:
				e.Message = "Model not found"
			case http.StatusBadRequest:
				e.Message = "Invalid input"
			default:
				e.Message = "Error"
			}
		}

		http.Error(w, e.Message, e.Code)
	}
}
