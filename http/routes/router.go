package routes

import (
	"github.com/gorilla/mux"
	"github.com/mlshvsk/go-task-manager/http/controllers"
	"github.com/mlshvsk/go-task-manager/http/handlers"
	"github.com/mlshvsk/go-task-manager/http/middlewares"
	"net/http"
)


func NewRouter() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()
	api.Use(middlewares.LogRequestMiddlewareFunc)

	projects := api.PathPrefix("/projects").Subrouter()

	projects.Handle("", handlers.Handler(controllers.IndexProjects)).Methods(http.MethodGet)
	projects.Handle("", handlers.Handler(controllers.StoreProject)).Methods(http.MethodPost)
	projects.Handle("/{projectId:[0-9]*}", handlers.Handler(controllers.ShowProject)).Methods(http.MethodGet)
	projects.Handle("/{projectId:[0-9]*}", handlers.Handler(controllers.DeleteProject)).Methods(http.MethodDelete)
	projects.Handle("/{projectId:[0-9]*}", handlers.Handler(controllers.UpdateProject)).Methods(http.MethodPut)
	projects.Handle("/{projectId:[0-9]*}/columns", handlers.Handler(controllers.IndexColumns)).Methods(http.MethodGet)
	projects.Handle("/{projectId:[0-9]*}/columns", handlers.Handler(controllers.StoreColumn)).Methods(http.MethodPost)


	columns := api.PathPrefix("/columns").Subrouter()

	columns.Handle("/{columnId:[0-9]*}", handlers.Handler(controllers.ShowColumn)).Methods(http.MethodGet)
	columns.Handle("/{columnId:[0-9]*}", handlers.Handler(controllers.DeleteColumn)).Methods(http.MethodDelete)
	columns.Handle("/{columnId:[0-9]*}", handlers.Handler(controllers.UpdateColumn)).Methods(http.MethodPut)
	columns.Handle("/{columnId:[0-9]*}/move", handlers.Handler(controllers.MoveColumn)).Methods(http.MethodPost)
	columns.Handle("/{columnId:[0-9]*}/tasks", handlers.Handler(controllers.IndexTasksByColumn)).Methods(http.MethodGet)
	columns.Handle("/{columnId:[0-9]*}/tasks", handlers.Handler(controllers.StoreTask)).Methods(http.MethodPost)

	tasks := columns.PathPrefix("/tasks").Subrouter()

	tasks.Handle("", handlers.Handler(controllers.IndexTasks)).Methods(http.MethodGet)
	tasks.Handle("/{taskId:[0-9]*}", handlers.Handler(controllers.DeleteTask)).Methods(http.MethodDelete)
	tasks.Handle("/{taskId:[0-9]*}", handlers.Handler(controllers.ShowTask)).Methods(http.MethodGet)
	tasks.Handle("/{taskId:[0-9]*}", handlers.Handler(controllers.UpdateTask)).Methods(http.MethodPut)
	tasks.Handle("/{taskId:[0-9]*}/move", handlers.Handler(controllers.MoveTask)).Methods(http.MethodPost)
	tasks.Handle("/{taskId:[0-9]*}/column/{newColumnId:[0-9]*}", handlers.Handler(controllers.MoveTaskColumn)).Methods(http.MethodPost)
	tasks.Handle("/{taskId:[0-9]*}/comments", handlers.Handler(controllers.IndexComments)).Methods(http.MethodGet)
	tasks.Handle("/{taskId:[0-9]*}/comments", handlers.Handler(controllers.StoreComment)).Methods(http.MethodPost)

	comments := tasks.PathPrefix("/comments").Subrouter()

	comments.Handle("/{commentId:[0-9]*}", handlers.Handler(controllers.ShowComment)).Methods(http.MethodGet)
	comments.Handle("/{commentId:[0-9]*}", handlers.Handler(controllers.UpdateComment)).Methods(http.MethodGet)
	comments.Handle("/{commentId:[0-9]*}", handlers.Handler(controllers.DeleteComment)).Methods(http.MethodDelete)

	return r
}
