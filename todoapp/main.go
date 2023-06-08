package main

import (
	"log"
	"net/http"
	"github.com/ssen110/test/controllers"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()

	// Get all todos
	router.HandleFunc("/todos/", controllers.GetTodos).Methods("GET")

	// Get One todo by todoid
	router.HandleFunc("/todos/{todoid}", controllers.GetOneTodo).Methods("GET")

	// Create a todo
	router.HandleFunc("/todos/", controllers.CreateTodos).Methods("POST")

	// Delete a specific todo by the todoid
	router.HandleFunc("/todos/{todoid}", controllers.DeleteTodo).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}