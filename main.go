package main

import (
	"log"
	"net/http"
)

func main() {
	// 1. Create a new router
	router := http.NewServeMux()

	// 2. Define API Endpoints

	// User Endpoints
	router.HandleFunc("POST /users", createUsers)        // Create a new user
	router.HandleFunc("GET /users/{id}", getUsers)       // Get a user by ID
	router.HandleFunc("DELETE /users/{id}", deleteUsers) // Delete a user by ID

	// Task Endpoints
	router.HandleFunc("POST /users/{userID}/todos", createToDoForUser) // Create a new todo for a user

	// 3. Configure the HTTP server
	server := &http.Server{
		Addr:    ":8081", 
		Handler: router,
	}

	// 4. Start the server
	log.Println("Starting server on http://localhost:8081")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on :8081: %v\n", err)
	}
}
