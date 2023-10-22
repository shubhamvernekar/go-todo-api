package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shubhamvernekar/go-todo-api/models"
)

var task []models.Task

func SetupRoutes(router *mux.Router) {
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/task/{id:[0-9]+}", getTask).Methods("GET")
	router.HandleFunc("/task", createTask).Methods("POST")
	router.HandleFunc("/task/{id:[0-9]+}", deleteTask).Methods("DELETE")
	router.HandleFunc("/task", updateTask).Methods("PUT")
	router.HandleFunc("/task/markDone/{id:[0-9]+}", markDone).Methods("GET")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Println("error parsing id")
		http.Error(w, "error parsing id", http.StatusInternalServerError)
		return
	}

	t := findTaskByID(taskId)

	if t == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)

	if err != nil {
		fmt.Println("error decoding request body")
		http.Error(w, "error decoding request body", http.StatusInternalServerError)
		return
	}

	newTask.ID = generateId()
	task = append(task, newTask)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTask)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Println("error parsing id")
		http.Error(w, "error parsing id", http.StatusInternalServerError)
		return
	}

	for i := range task {
		if task[i].ID == taskId {
			task = append(task[:i], task[i+1:]...)
			break
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(`{
		"message" : "Successfuly deleted"
	}`)

}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var updateTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updateTask)

	if err != nil {
		fmt.Println("error decoding request body")
		http.Error(w, "error decoding request body", http.StatusInternalServerError)
		return
	}

	t := findTaskByID(updateTask.ID)

	if t == nil {
		http.NotFound(w, r)
	}

	t.Desc = "hello"
	t.IsDone = updateTask.IsDone
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)

}

func markDone(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskId, err := strconv.Atoi(params["id"])

	if err != nil {
		fmt.Println("error parsing id")
		http.Error(w, "error parsing id", http.StatusInternalServerError)
		return
	}

	t := findTaskByID(taskId)

	if t == nil {
		http.NotFound(w, r)
		return
	}

	t.IsDone = !t.IsDone

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(t)
}

func findTaskByID(ID int) *models.Task {
	for i := range task {
		if task[i].ID == ID {
			return &task[i]
		}
	}

	return nil
}

func generateId() int {
	if len(task) == 0 {
		return 1
	}
	return task[len(task)-1].ID + 1
}

func printTask() {
	fmt.Println(task)
}
