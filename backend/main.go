package main

import (
	"fmt"
	"log"
	"net/http"

	"todoapp/db"
	"todoapp/handlers"

	"github.com/rs/cors"
)

func main() {
	// Initialize the database
	db.InitDB()

	// Create a new HTTP multiplexer (router)
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/tasks", handlers.TasksHandler)
	mux.HandleFunc("/tasks/", handlers.TasksHandler) // Allow `/tasks/1` for DELETE operations

	// Enable CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Allow frontend origin
		AllowedMethods:   []string{"GET", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
	})

	// Wrap the multiplexer with CORS middleware
	handler := c.Handler(mux)

	// Start the HTTP server
	fmt.Println("✅ Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}

