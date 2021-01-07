package app

import (
	"fmt"
	"github.com/mlshvsk/go-task-manager/database"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/mlshvsk/go-task-manager/config"
	//"github.com/mlshvsk/go-task-manager/database"
)

type App struct {
	Config   config.Config
	Database *database.SqlDB
}

func New(cfg config.Config) *App {
	db, err := database.Load(cfg.Sql)

	if err != nil {
		log.Fatal(err)
	}

	return &App{cfg, db}
}

func (app *App) Run(r *mux.Router) {
	port := app.Config.Port
	addr := fmt.Sprintf(":%v", port)
	fmt.Printf("APP is listening on port: %d\n", port)
	log.Fatal(http.ListenAndServe(addr, handlers.CORS()(r)))
}

func (app *App) IsProd() bool {
	return app.Config.Env == "prod"
}
