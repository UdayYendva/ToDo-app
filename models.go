package main

import (
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Tasks struct {
	ID     int    `json:"id"`
	UserID int    `json:"usersId"`
	Task   string `json:"task"`
	Status bool   `json:"status"`
}

var UserList = make(map[int]User)

var TaskList = make(map[int][]Tasks)

var NextUser = 1
var NextTask = 1
var UserMutex = sync.RWMutex{}
var TaskMutex = sync.RWMutex{}
