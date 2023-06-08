package queries

var Get_all_todos = "SELECT * FROM todo_details"
var Get_one_specific_todo = "SELECT * FROM todo_details where id = $1"
var Insert_one_todo = "INSERT INTO todo_details(title, description, completed) VALUES($1, $2, $3) returning id;"
var Delete_one_specific_todo = "DELETE FROM todo_details where id = $1"

