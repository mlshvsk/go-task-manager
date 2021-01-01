package main

import (
	"github.com/mlshvsk/go-task-manager/app"
	"github.com/mlshvsk/go-task-manager/config"
	"github.com/mlshvsk/go-task-manager/repositories"
	"github.com/mlshvsk/go-task-manager/routes"
	"log"
)

func main() {
	cfg, err := config.Load("config/app.json")
	if err != nil {
		log.Fatal(err)
	}
	app := app.New(cfg)
	initServices(app)
	app.Run(routes.NewRouter())
}

func initServices(a *app.App) {
	repositories.InitCommentRepository(a.Database)
	repositories.InitTaskRepository(a.Database)
	repositories.InitColumnRepository(a.Database)
	repositories.InitProjectRepository(a.Database)
}
