package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Corrected function signature: uses *http.Request for the incoming request
func createUsers(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserMutex.Lock()
	defer UserMutex.Unlock()

	user.ID = NextUser
	NextUser++
	UserList[user.ID] = user

	w.Header().Set("Content-Type", "application/json")
	// Corrected: Uses w.WriteHeader() instead of w.writerHeader()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

// Corrected function signature: uses *http.Request for the incoming request
func getUsers(w http.ResponseWriter, r *http.Request) {
	// r.PathValue is correct for new Go 1.22+ routing
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserMutex.Lock()
	// Corrected: Uses UserMutex.Unlock() instead of UserMutex.unlock()
	defer UserMutex.Unlock()

	user, ok := UserList[id]
	if !ok {
		// Changed to StatusNotFound for more accurate HTTP status
		http.Error(w, "User not found", http.StatusNotFound) 
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

// Corrected function signature: uses *http.Request for the incoming request
func deleteUsers(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserMutex.Lock()
	// Corrected: Uses UserMutex.Unlock() instead of UserMutex.unlock()
	defer UserMutex.Unlock()

	_, ok := UserList[id]
	if !ok {
		// Changed to StatusNotFound for more accurate HTTP status
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(UserList, id)
	delete(TaskList, id)

	w.WriteHeader(http.StatusNoContent)
}

// Corrected function signature: uses *http.Request for the incoming request
func createToDoForUser(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.Atoi(r.PathValue("userID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	// Check if user exists (use RLock since we are only reading)
	UserMutex.RLock()
	_, ok := UserList[userID]
	UserMutex.RUnlock()
	
	if !ok {
		http.Error(w, "User Not Exists check the id you given", http.StatusBadRequest)
		return // Must return after sending error
	}
	
	var todo Tasks
	// Must redeclare 'err' or use a different variable name here if the previous 'err' is still needed,
	// but in this context, reusing 'err' with ':=' is safer.
	err = json.NewDecoder(r.Body).Decode(&todo) 
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if todo.Task == "" {
		http.Error(w, "task shouldn't be empty", http.StatusBadRequest)
		return
	}
	
	TaskMutex.Lock()
	defer TaskMutex.Unlock()

	todo.ID = NextTask
	NextTask++
	todo.UserID = userID
	todo.Status = false

	// Initialize the slice if the user has no existing tasks
	if _, exists := TaskList[userID]; !exists {
		TaskList[userID] = []Tasks{}
	}
	TaskList[userID] = append(TaskList[userID], todo)
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
