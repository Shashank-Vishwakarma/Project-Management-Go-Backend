package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/constants"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/database"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/middlewares"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/routes"
	"github.com/gorilla/mux"
)

func init() {
	// load config
	err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Database initialization
	err = database.InitDB()
	if err != nil {
		panic(err)
	}
}

func main() {
	router := mux.NewRouter()

	// create subrouters
	authRouter := router.PathPrefix(constants.AUTH_BASE_ENDPOINT).Subrouter()
	projectRouter := router.PathPrefix(constants.PROJECTS_BASE_ENDPOINT).Subrouter()
	taskRouter := router.PathPrefix(constants.TASKS_BASE_ENDPOINT).Subrouter()

	// register middlewares
	router.Use(middlewares.RequestLogger)
	authRouter.Use(middlewares.VerifyToken, middlewares.VerifyRole)
	projectRouter.Use(middlewares.VerifyToken, middlewares.VerifyRole)
	taskRouter.Use(middlewares.VerifyToken, middlewares.VerifyRole)

	// register routes
	routes.AuthRoutes(authRouter)
	routes.ProjectRoutes(projectRouter)
	routes.TaskRoutes(taskRouter)

	port := config.Config.Port
	log.Printf("Server is running on port %s", port)
	log.Println("Database initialized")
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}