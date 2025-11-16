package main

import (
	"sync"
)

// --- Struct Definitions ---

// User struct represents a user in the system.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Tasks struct represents a single ToDo item.
type Tasks struct {
	ID     int    `json:"id"`
	UserID int    `json:"userId"`
	Task   string `json:"task"`
	Status bool   `json:"status"` // false for incomplete, true for complete
}

// --- Global Data Stores and Synchronization ---

// UserList stores users, indexed by their ID.
var UserList = make(map[int]User)

// TaskList stores tasks, indexed by UserID. Each user ID maps to a slice of their Tasks.
var TaskList = make(map[int][]Tasks)

// NextUser is used to generate unique IDs for new users.
var NextUser = 1

// NextTask is used to generate unique IDs for new tasks across all users.
var NextTask = 1

// UserMutex protects access to UserList and NextUser.
var UserMutex = sync.RWMutex{}

// TaskMutex protects access to TaskList and NextTask.
var TaskMutex = sync.RWMutex{}
