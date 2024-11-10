package routes

import (
	"net/http"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/handlers"
	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router) {
	router.HandleFunc("projects/{project_id}/tasks", handlers.CreateTaskHandler).Methods(http.MethodPost)
	router.HandleFunc("projects/{project_id}/tasks", handlers.GetAllTasksHandler).Methods(http.MethodGet)
	router.HandleFunc("projects/{project_id}/tasks/{task_id}", handlers.GetTaskDetailsHandler).Methods(http.MethodGet)
	router.HandleFunc("projects/{project_id}/tasks/{task_id}", handlers.UpdateTaskHandler).Methods(http.MethodPut)
	router.HandleFunc("projects/{project_id}/tasks/{task_id}", handlers.DeleteTaskHandler).Methods(http.MethodDelete)
}