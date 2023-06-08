package controllers

import (
	"net/http"
	"github.com/ssen110/test/dbconnection"
	"github.com/ssen110/test/utilfunc"
	"github.com/ssen110/test/controllers/queries"
	"fmt"
	"encoding/json"
)

// Create a CreateTodos
func CreateTodos(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")
	completed := r.FormValue("completed")

	var response = JsonResponse{}
	var  com bool
	if completed == "0"{
		com = false
	}else{
		com = true
	}

	if title == "" {
		response = JsonResponse{Type: "error", Message: "You are missing title parameter. Which is Needed "}
	} else {
		db := dbconn.SetupDB()

		utilfunc.PrintMessage("Inserting Todos into DB")

		fmt.Println("Inserting new Todo with Title: " + title + " and description: " + description)

		var lastInsertID int
		err := db.QueryRow(queries.Insert_one_todo, title, description, com).Scan(&lastInsertID)

		utilfunc.CheckErrWhileInterractingWithDB(err)
		defer db.Close() // Closing the db connection
		response = JsonResponse{Type: "success", Message: "The Todo has been inserted successfully!"}
	}

	json.NewEncoder(w).Encode(response)
}