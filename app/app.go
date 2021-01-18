package app

import (
	"fmt"
	"github.com/mlshvsk/go-task-manager/database"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mlshvsk/go-task-manager/config"
	"github.com/rs/cors"
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
	handler := cors.Default().Handler(r)
	log.Fatal(http.ListenAndServe(addr, handler))
}

func (app *App) IsProd() bool {
	return app.Config.Env == "prod"
}
