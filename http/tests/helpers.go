package tests

import (
	"encoding/json"
	"github.com/mlshvsk/go-task-manager/logger"
	"go.uber.org/zap/zaptest"
	"testing"
)

func InitLoggers(t *testing.T) {
	logger.RequestLogger = zaptest.NewLogger(t).Sugar()
	logger.ErrorLogger = zaptest.NewLogger(t).Sugar()
}

func ExpectedOkResponse(data interface{}) ([]byte, error) {
	responseStruct := &struct {
		Data interface{} `json:"data"`
	}{Data: data}

	return json.Marshal(responseStruct)
}
