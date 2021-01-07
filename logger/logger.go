package logger

import (
	"github.com/mlshvsk/go-task-manager/config"
	"go.uber.org/zap"
)

var RequestLogger *zap.SugaredLogger
var ErrorLogger *zap.SugaredLogger

func InitRequestLogger(cfg config.Config) {
	loggerConf := zap.NewProductionConfig()
	loggerConf.OutputPaths = []string{
		cfg.RequestLog.OutputPath,
	}
	logger, _ := loggerConf.Build()
	//defer logger.Sync() // flushes buffer, if any
	RequestLogger = logger.Sugar()
}

func InitErrorLogger(cfg config.Config) {
	loggerConf := zap.NewProductionConfig()
	loggerConf.OutputPaths = []string{
		cfg.ErrorLog.OutputPath,
	}
	logger, _ := loggerConf.Build()
	//defer logger.Sync() // flushes buffer, if any
	ErrorLogger = logger.Sugar()
}
