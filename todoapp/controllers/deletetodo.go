package controllers

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ssen110/test/dbconnection"
	"github.com/ssen110/test/utilfunc"
	"encoding/json"
	"github.com/ssen110/test/controllers/queries"
)

// Delete one Todo
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	todoID := params["todoid"]

	var response = JsonResponse{}

	if todoID == "" {
		response = JsonResponse{Type: "error", Message: "You are missing todoID parameter."}
	} else {
		db := dbconn.SetupDB()
		utilfunc.PrintMessage("Deleting todo_details from DB")
		_, err := db.Exec(queries.Delete_one_specific_todo, todoID)
		utilfunc.CheckErrWhileInterractingWithDB(err)
		defer db.Close() // Closing the db connection
		response = JsonResponse{Type: "success", Message: "The todo has been deleted successfully!"}
	}
	json.NewEncoder(w).Encode(response)
}
