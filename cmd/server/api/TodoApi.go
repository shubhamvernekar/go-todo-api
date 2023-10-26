package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
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

func getTasks(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(task)
	if err != nil {
		log.Error().Msgf("error encoding json %v", err)
	}
}

func getTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "error parsing id", http.StatusInternalServerError)
		log.Error().Msgf("error parsing id %v", err)
		return
	}

	t := findTaskByID(taskID)

	if t == nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(t)

	if err != nil {
		log.Error().Msgf("error encoding json %v", err)
	}
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask models.Task
	err := json.NewDecoder(r.Body).Decode(&newTask)
	if err != nil {
		http.Error(w, "error decoding request body", http.StatusInternalServerError)
		log.Error().Msgf("error decoding request body %v", err)
		return
	}

	newTask.ID = generateID()
	task = append(task, newTask)

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(newTask)

	if err != nil {
		log.Error().Msgf("error encoding json %v", err)
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "error parsing id", http.StatusInternalServerError)
		log.Error().Msgf("error parsing id %v", err)
		return
	}

	found := false

	for i := range task {
		if task[i].ID == taskID {
			task = append(task[:i], task[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(`{
		"message" : "Successfully deleted"
	}`)

	if err != nil {
		log.Error().Msgf("error encoding json : %v", err)
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	var updateTask models.Task
	err := json.NewDecoder(r.Body).Decode(&updateTask)
	if err != nil {
		http.Error(w, "error decoding request body", http.StatusInternalServerError)
		log.Error().Msgf("error decoding request body %v", err)
		return
	}

	t := findTaskByID(updateTask.ID)

	if t == nil {
		http.NotFound(w, r)
		return
	}

	t.Desc = updateTask.Desc
	t.IsDone = updateTask.IsDone
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(t)

	if err != nil {
		log.Error().Msgf("error encoding json : %v", err)
	}
}

func markDone(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	taskID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "error parsing id", http.StatusInternalServerError)
		log.Error().Msgf("error parding id %v", err)
		return
	}

	t := findTaskByID(taskID)

	if t == nil {
		http.NotFound(w, r)
		return
	}

	t.IsDone = !t.IsDone

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(t)

	if err != nil {
		log.Error().Msgf("error encoding json : %v", err)
	}
}

func findTaskByID(id int) *models.Task {
	for i := range task {
		if task[i].ID == id {
			return &task[i]
		}
	}

	return nil
}

func generateID() int {
	if len(task) == 0 {
		return 1
	}
	return task[len(task)-1].ID + 1
}
