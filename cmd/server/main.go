package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/shubhamvernekar/go-todo-api/cmd/server/api"
)

func main() {
	router := mux.NewRouter()
	api.SetupRoutes(router)

	server := &http.Server{
		Addr:           ":3000",
		Handler:        router,
		ReadTimeout:    10 * time.Second, // Set your read timeout
		WriteTimeout:   10 * time.Second, // Set your write timeout
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Server listining to port 3000")
	if err := server.ListenAndServe(); err != nil {
		log.Println("error running server ", err)
	}
}
