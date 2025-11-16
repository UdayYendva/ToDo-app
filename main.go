package main

import (
	"log"
	"net/http"
)

func main() {

	router := http.NewServeMux()

	router.HandleFunc("POST /users", createUsers)
	router.HandleFunc("GET /users/{id}", getUsers)
	router.HandleFunc("DELETE /users/{id}", deleteUsers)

	router.HandleFunc("POST /users/{userID}/todos", createToDoForUser)
	router.HandleFunc("GET /users/{userID}/tasks", getTasksForUser)

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	log.Println("Starting server on http://localhost:8081")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Could not listen on :8081: %v\n", err)
	}
}
