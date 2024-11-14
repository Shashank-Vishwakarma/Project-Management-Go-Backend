package routes

import (
	"net/http"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/handlers"
	"github.com/gorilla/mux"
)

func TaskRoutes(router *mux.Router) {
	router.HandleFunc("", handlers.CreateTaskHandler).Methods(http.MethodPost)
	router.HandleFunc("", handlers.GetAllTasksHandler).Methods(http.MethodGet)
	router.HandleFunc("/{task_id}", handlers.GetTaskDetailsHandler).Methods(http.MethodGet)
	router.HandleFunc("/{task_id}", handlers.UpdateTaskHandler).Methods(http.MethodPut)
	router.HandleFunc("/{task_id}", handlers.DeleteTaskHandler).Methods(http.MethodDelete)
}