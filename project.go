package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"` // "pending" or "completed"
}

var tasks []Task
var nID = 101

func createHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	//creating first task
	var task1 Task
	json.NewDecoder(r.Body).Decode(&task1)
	task1.ID = nID
	task1.Title = "Walking"
	task1.Description = "Walk in garden for 5kms."
	task1.Status = "pending"
	nID++
	tasks = append(tasks, task1)

	//creating second task
	var task2 Task
	json.NewDecoder(r.Body).Decode(&task2)
	task2.ID = nID
	task2.Title = "Make Food"
	task2.Description = "Make Pizza with all sauses"
	task2.Status = "pending"
	nID++
	tasks = append(tasks, task2)

}
func getAllTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	//getting all tasks
	json.NewEncoder(w).Encode(tasks)
}

func getTaskByIDHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for _, task := range tasks {
		if task.ID == id {
			json.NewEncoder(w).Encode(task)
			return
		}
	}
	http.NotFound(w, r)
}

func updateTaskHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			var updatedTask Task
			json.NewDecoder(r.Body).Decode(&updatedTask)
			updatedTask.ID = task.ID
			updatedTask.Title = task.Title
			updatedTask.Description = task.Description
			updatedTask.Status = "completed"
			tasks = append(tasks, updatedTask)
			json.NewEncoder(w).Encode(updatedTask)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/tasks", createHandler).Methods("POST")
	router.HandleFunc("/tasks", getAllTaskHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", getTaskByIDHandler).Methods("GET")
	router.HandleFunc("/tasks/{id}", updateTaskHandler).Methods("PUT")
	router.HandleFunc("/tasks/{id}", deleteTask).Methods("DELETE")
	pNumber := ":8091"
	fmt.Printf("Server is running on the port: %s\n", pNumber)
	http.ListenAndServe(pNumber, router)
}
