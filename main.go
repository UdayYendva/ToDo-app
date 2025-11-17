package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("POST /users", createUsers)
	router.HandleFunc("GET /users/{id}", getUsers)
	router.HandleFunc("DELETE /users/{id}", deleteUsers)
	router.HandleFunc("GET /users", getAllUsers)

	router.HandleFunc("POST /users/{userID}/todos", createToDoForUser)
	router.HandleFunc("GET /users/{userID}/tasks", getTasksForUser)

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}
	fmt.Println("=====================================")
	fmt.Println("       ToDo App Server Starting       ")
	fmt.Println("=====================================")
	fmt.Printf("Listening on port: 8081\n")
	fmt.Println("Available endpoints:")
	fmt.Println("  GET    /users            -> List all users")
	fmt.Println("  POST   /users            -> Create a user")
	fmt.Println("  GET    /users/{id}       -> Get user by ID")
	fmt.Println("  DELETE /users/{id}       -> Delete user")
	fmt.Println("  POST   /users/{id}/tasks -> Create task for user")
	fmt.Println("  GET    /users/{id}/tasks -> List tasks for user")
	fmt.Println("=====================================")

	log.Println("⚡ Starting server on http://localhost:8081")

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not listen on :8081: %v\n", err)
	}
}
