package main

import (
	"github.com/mlshvsk/go-task-manager/app"
	"github.com/mlshvsk/go-task-manager/config"
	"github.com/mlshvsk/go-task-manager/http/routes"
	"github.com/mlshvsk/go-task-manager/logger"
	"github.com/mlshvsk/go-task-manager/repositories"
	"log"
)

func main() {
	cfg, err := config.Load("config/app.json")
	if err != nil {
		log.Fatal(err)
	}

	a := app.New(cfg)
	initServices(a)
	a.Run(routes.NewRouter())
}

func initServices(a *app.App) {
	logger.InitRequestLogger(a.Config)
	logger.InitErrorLogger(a.Config)
	repositories.InitCommentRepository(a.Database)
	repositories.InitTaskRepository(a.Database)
	repositories.InitColumnRepository(a.Database)
	repositories.InitProjectRepository(a.Database)
}
