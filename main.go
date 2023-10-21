package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shubhamvernekar/go-todo-api/api"
)

func rootRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello ")
}

func main() {
	router := mux.NewRouter()

	api.SetupRoutes(router)
	http.Handle("/", router)

	fmt.Println("Server listining to port 3000")
	http.ListenAndServe(":3000", nil)
}