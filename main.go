package main

import (
	"github.com/mlshvsk/go-task-manager/app"
	"github.com/mlshvsk/go-task-manager/config"
	"github.com/mlshvsk/go-task-manager/http/routes"
	"github.com/mlshvsk/go-task-manager/logger"
	"github.com/mlshvsk/go-task-manager/repositories/base"
	"github.com/mlshvsk/go-task-manager/repositories/column"
	"github.com/mlshvsk/go-task-manager/repositories/comment"
	"github.com/mlshvsk/go-task-manager/repositories/mysql"
	"github.com/mlshvsk/go-task-manager/repositories/project"
	"github.com/mlshvsk/go-task-manager/repositories/task"
	column2 "github.com/mlshvsk/go-task-manager/services/column"
	comment2 "github.com/mlshvsk/go-task-manager/services/comment"
	project2 "github.com/mlshvsk/go-task-manager/services/project"
	task2 "github.com/mlshvsk/go-task-manager/services/task"
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

	cr := comment.InitCommentRepository(initBaseRepository(a))
	comment2.InitCommentService(cr)

	tr := task.InitTaskRepository(initBaseRepository(a))
	task2.InitTaskService(tr)

	colR := column.InitColumnRepository(initBaseRepository(a))
	column2.InitColumnService(colR)

	pr := project.InitProjectRepository(initBaseRepository(a))
	project2.InitProjectService(pr)
}

func initBaseRepository(a *app.App) base.Repository {
	switch a.Config.Sql.Driver {
	case "mysql":
		return &mysql.Repository{SqlDB: a.Database}
	}

	return nil
}
