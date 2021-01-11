package main

import (
	"github.com/mlshvsk/go-task-manager/app"
	"github.com/mlshvsk/go-task-manager/config"
	"github.com/mlshvsk/go-task-manager/http/routes"
	"github.com/mlshvsk/go-task-manager/logger"
	"github.com/mlshvsk/go-task-manager/repositories"
	"github.com/mlshvsk/go-task-manager/repositories/base"
	"github.com/mlshvsk/go-task-manager/repositories/mysql"
	"github.com/mlshvsk/go-task-manager/services"
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

	cr := repositories.InitCommentRepository(initBaseRepository(a))
	services.InitCommentService(cr)

	tr := repositories.InitTaskRepository(initBaseRepository(a))
	services.InitTaskService(tr)

	colR := repositories.InitColumnRepository(initBaseRepository(a))
	services.InitColumnService(colR)

	pr := repositories.InitProjectRepository(initBaseRepository(a))
	services.InitProjectService(pr)
}

func initBaseRepository(a *app.App) base.Repository {
	switch a.Config.Sql.Driver {
	case "mysql":
		return &mysql.Repository{SqlDB: a.Database}
	}

	return nil
}
