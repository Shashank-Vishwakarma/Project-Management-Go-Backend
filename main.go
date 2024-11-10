package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/database"
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

	// register routes
	routes.AuthRoutes(router)

	port := config.Config.Port
	log.Printf("Server is running on port %s", port)
	log.Println("Database initialized")
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}