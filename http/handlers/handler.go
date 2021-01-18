package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/mlshvsk/go-task-manager/logger"
	"net/http"
)

type Handler func(http.ResponseWriter, *http.Request) *AppError

type AppError struct {
	Error        error  `json:"-"`
	Message      string `json:"message"`
	ResponseCode int    `json:"-"`
}

func (fn Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil {
		fmt.Println(e)
		logger.ErrorLogger.Error(e.Error.Error())

		if e.Message == "" {
			switch e.ResponseCode {
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

		w.WriteHeader(e.ResponseCode)
		if err := json.NewEncoder(w).Encode(e); err != nil {
			http.Error(w, "", http.StatusInternalServerError)
		}
	}
}
