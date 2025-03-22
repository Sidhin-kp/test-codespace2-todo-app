package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"todoapp/db"

	_ "github.com/lib/pq"
)

// Task represents a to-do item
type Task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

// TasksHandler handles GET and POST requests for tasks
func TasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getTasks(w, r)
	case "POST":
		addTask(w, r)
	case "DELETE":
		deleteTask(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// Fetch tasks from the database
func getTasks(w http.ResponseWriter, r *http.Request) {
	rows, err := db.DB.Query("SELECT id, task FROM tasks")
	if err != nil {
		log.Println("Error fetching tasks:", err)
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var t Task
		if err := rows.Scan(&t.ID, &t.Task); err != nil {
			log.Println("Error scanning task:", err)
			continue
		}
		tasks = append(tasks, t)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

// Add a new task to the database
func addTask(w http.ResponseWriter, r *http.Request) {
	var t Task
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Invalid JSON input", http.StatusBadRequest)
		return
	}

	err = db.DB.QueryRow("INSERT INTO tasks (task) VALUES ($1) RETURNING id", t.Task).Scan(&t.ID)
	if err != nil {
		log.Println("Error inserting task:", err)
		http.Error(w, "Error adding task", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

// Delete a task from the database
// Delete a task from the database
func deleteTask(w http.ResponseWriter, r *http.Request) {
	// Extract task ID from URL
	id := r.URL.Path[len("/tasks/"):]
	if id == "" {
		http.Error(w, "Missing task ID", http.StatusBadRequest)
		return
	}

	// Delete the task from the database
	result, err := db.DB.Exec("DELETE FROM tasks WHERE id = $1", id)
	if err != nil {
		log.Println("Error deleting task:", err)
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting affected rows:", err)
		http.Error(w, "Error processing request", http.StatusInternalServerError)
		return
	}

	if rowsAffected == 0 {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	// Send success response
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Task deleted successfully"}`))
}


