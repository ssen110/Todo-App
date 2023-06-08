package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/ssen110/test/utilfunc"
	"github.com/ssen110/test/dbconnection"
	"github.com/ssen110/test/types"
	"github.com/gorilla/mux"
	"github.com/ssen110/test/controllers/queries"
	"fmt"
)

type JsonResponse struct {
	Type    string `json:"type"`
	Data    []types.TODODETAILS `json:"data"`
	Message string `json:"message"`
}

// Get all todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	db := dbconn.SetupDB()

	utilfunc.PrintMessage("Getting Todo Details...")
	rows, err := db.Query(queries.Get_all_todos)
	utilfunc.CheckErrWhileInterractingWithDB(err)

	//fmt.Println("Closing the DB connection")
	defer db.Close() // Closing the db connection
	//fmt.Println("Closed the DB connection")

	var todos []types.TODODETAILS
	for rows.Next() {
		var id int
		var title string
		var description string
		var completed bool

		err = rows.Scan(&id, &title, &description, &completed)

		utilfunc.CheckErrWhileScanning(err)

		todos = append(todos, types.TODODETAILS{Id: id, Title: title, Description: description, Completed: completed})
	}
	var response = JsonResponse{Type: "success", Data: todos}

	json.NewEncoder(w).Encode(response)
}

// Get one Todo
func GetOneTodo(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	todoID := params["todoid"]
	
	fmt.Println("todoid : " + todoID)
	db := dbconn.SetupDB()
	rows, err := db.Query(queries.Get_one_specific_todo, todoID)
	utilfunc.CheckErrWhileInterractingWithDB(err)
	defer db.Close() // Closing the db connection
	utilfunc.PrintMessage("Getting One Todo Details...")
	
	var todos []types.TODODETAILS
	for rows.Next() {
		var id int
		var title string
		var description string
		var completed bool

		err = rows.Scan(&id, &title, &description, &completed)
		utilfunc.CheckErrWhileScanning(err)
		todos = append(todos, types.TODODETAILS{Id: id, Title: title, Description: description, Completed: completed})
	}
	var response = JsonResponse{Type: "success", Data: todos}
	json.NewEncoder(w).Encode(response)
}