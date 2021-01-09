package logger

import (
	"github.com/mlshvsk/go-task-manager/config"
	"go.uber.org/zap"
	"log"
)

var RequestLogger *zap.SugaredLogger
var ErrorLogger *zap.SugaredLogger

func InitRequestLogger(cfg config.Config) {
	loggerConf := zap.NewProductionConfig()
	loggerConf.OutputPaths = []string{
		cfg.RequestLog.OutputPath,
	}
	logger, err := loggerConf.Build()
	if err != nil {
		log.Fatal(err)
	}
	RequestLogger = logger.Sugar()
}

func InitErrorLogger(cfg config.Config) {
	loggerConf := zap.NewProductionConfig()
	loggerConf.OutputPaths = []string{
		cfg.ErrorLog.OutputPath,
	}
	logger, err := loggerConf.Build()
	if err != nil {
		log.Fatal(err)
	}
	ErrorLogger = logger.Sugar()
}
