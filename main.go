package main

import (
	"log"
	"net/http"

	"root/config"
	"root/controllers"

	"github.com/gorilla/mux"
)

func main() {
	// Initialize the database
	db, err := config.InitDatabase()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Set up router
	router := mux.NewRouter()
	controllers.RegisterUserRoutes(router)

	// Start the server
	log.Println("Server running on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
