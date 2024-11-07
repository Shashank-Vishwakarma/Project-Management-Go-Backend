package main

import (
	"log"
	"net/http"

	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/config"
	"github.com/Shashank-Vishwakarma/Project-Management-Go-Backend/database"
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

	log.Println("Server is running on port 8080")
	log.Println("Database initialized")
	http.ListenAndServe(":8080", router)
}