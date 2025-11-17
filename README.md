This project is a simple REST API built using Go.  
It allows you to manage **Users** and their **To-Do Tasks** with basic CRUD operations.

---



📘 Go User & Task Management API

A lightweight, concurrency-safe REST API built using pure Go, designed for managing Users and their associated To-Do Tasks.
This project demonstrates clean API structuring, mutex-based synchronization, JSON handling, routing logic, and in-memory data storage.

⭐ Features
👤 User Management

Create a new user

Retrieve a user by ID

Delete a user

📝 Task (To-Do) Management

Create a task for a specific user

Fetch all tasks of a user

Tasks stored user-wise

Auto-increment task IDs

🔒 Thread-Safe Operations

Uses sync.Mutex and sync.RWMutex

Prevents race conditions under concurrent requests

⚡ Lightweight + No Dependencies

Standard library only (net/http, encoding/json)

