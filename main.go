package main

import (
	"log"
	"net/http"

	"horstito/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/sessions", handlers.CreateSession).Methods("POST")
	r.HandleFunc("/sessions", handlers.GetAllSessions).Methods("GET")
	r.HandleFunc("/sessions/{id}", handlers.GetSessionByID).Methods("GET")
	r.HandleFunc("/sessions/{id}", handlers.DeleteSession).Methods("DELETE")

	log.Println("🚀 Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
