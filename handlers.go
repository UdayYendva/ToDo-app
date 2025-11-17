package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

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

func getUsers(w http.ResponseWriter, r *http.Request) { // to get users created

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserMutex.Lock()

	defer UserMutex.Unlock()

	user, ok := UserList[id]
	if !ok {

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

}

func deleteUsers(w http.ResponseWriter, r *http.Request) { // to delete users based on id
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	UserMutex.Lock()

	defer UserMutex.Unlock()

	_, ok := UserList[id]
	if !ok {

		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	delete(UserList, id)
	delete(TaskList, id)

	w.WriteHeader(http.StatusNoContent)
}
func getAllUsers(w http.ResponseWriter, r *http.Request) { // to get all users
	UserMutex.RLock()
	defer UserMutex.RUnlock()

	users := []User{}
	for _, user := range UserList {
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func createToDoForUser(w http.ResponseWriter, r *http.Request) { // to create tasks to users created
	userID, err := strconv.Atoi(r.PathValue("userID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UserMutex.RLock()
	_, ok := UserList[userID]
	UserMutex.RUnlock()

	if !ok {
		http.Error(w, "User Not Exists check the id you given", http.StatusBadRequest)
		return
	}

	var todo Tasks

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
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

	if _, exists := TaskList[userID]; !exists {
		TaskList[userID] = []Tasks{}
	}
	TaskList[userID] = append(TaskList[userID], todo)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}
func getTasksForUser(w http.ResponseWriter, r *http.Request) { // to get tasks for the users
	userID, err := strconv.Atoi(r.PathValue("userID"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	UserMutex.RLock()
	_, ok := UserList[userID]
	UserMutex.RUnlock()
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	TaskMutex.RLock()
	tasks, exists := TaskList[userID]
	TaskMutex.RUnlock()

	if !exists || len(tasks) == 0 {

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode([]Tasks{})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
