package routes

import (
	"github.com/gorilla/mux"
	"github.com/mlshvsk/go-task-manager/controllers"
	"net/http"
)


func NewRouter() *mux.Router {
	r := mux.NewRouter()

	api := r.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/projects", controllers.IndexProjects).Methods(http.MethodGet)
	api.HandleFunc("/projects", controllers.StoreProject).Methods(http.MethodPost)
	api.HandleFunc("/projects/{projectId:[0-9]*}", controllers.ShowProject).Methods(http.MethodGet)
	api.HandleFunc("/projects/{projectId:[0-9]*}", controllers.DeleteProject).Methods(http.MethodDelete)
	api.HandleFunc("/projects/{projectId:[0-9]*}", controllers.UpdateProject).Methods(http.MethodPut)

	columns := api.PathPrefix("/projects/{projectId:[0-9]*}/columns").Subrouter()

	columns.HandleFunc("", controllers.IndexColumns).Methods(http.MethodGet)
	columns.HandleFunc("", controllers.StoreColumn).Methods(http.MethodPost)
	columns.HandleFunc("/{columnId:[0-9]*}", controllers.ShowColumn).Methods(http.MethodGet)
	columns.HandleFunc("/{columnId:[0-9]*}", controllers.DeleteColumn).Methods(http.MethodDelete)
	columns.HandleFunc("/{columnId:[0-9]*}", controllers.UpdateColumn).Methods(http.MethodPut)
	columns.HandleFunc("/{columnId:[0-9]*}/move", controllers.MoveColumn).Methods(http.MethodPost)

	tasks := columns.PathPrefix("/{columnId:[0-9]*}/tasks").Subrouter()

	tasks.HandleFunc("", controllers.IndexTasks).Methods(http.MethodGet)
	tasks.HandleFunc("", controllers.StoreTask).Methods(http.MethodPost)
	tasks.HandleFunc("/{taskId:[0-9]*}", controllers.DeleteTask).Methods(http.MethodDelete)
	tasks.HandleFunc("/{taskId:[0-9]*}", controllers.ShowTask).Methods(http.MethodGet)
	tasks.HandleFunc("/{taskId:[0-9]*}", controllers.UpdateTask).Methods(http.MethodPut)
	tasks.HandleFunc("/{taskId:[0-9]*}/move", controllers.MoveTask).Methods(http.MethodPost)
	//tasks.HandleFunc("/{id}/move/{newColumnId}", controllers.MoveTaskColumn).Methods(http.MethodPost)

	comments := tasks.PathPrefix("/{taskId:[0-9]*}/comments").Subrouter()

	comments.HandleFunc("", controllers.IndexComments).Methods(http.MethodGet)
	comments.HandleFunc("", controllers.StoreComment).Methods(http.MethodPost)
	comments.HandleFunc("/{commentId:[0-9]*}", controllers.ShowComment).Methods(http.MethodGet)
	comments.HandleFunc("/{commentId:[0-9]*}", controllers.DeleteComment).Methods(http.MethodDelete)

	return r
}
