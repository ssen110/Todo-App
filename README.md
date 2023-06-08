# Todo-App
Todo  App using Go [Hornet Assignment]

# Requirements:
1. go version go1.20.5
  - with the default packages I also install gorilla/mux using "github.com/gorilla/mux"
  - also installed lib/pq using 'github.com/lib/pq'
  - module github.com/ssen110/test
2. postgres version 15.3
  - New DB is create as 'test1' with Password = 'root'
  - A table is created as 'todo_details' and the SQL query is there in 'database_sql_queries' folder
  - The 'todo_details' contains columns as : id int which is the PK, title text, description text, completed bool

main.go file is the file we need to run and the port is :8000

# Test:
run the program : using go run main.go
- To Get all the Todos, URL(get): localhost:8000/todos/
- To Get one specific Todo, URL(get): localhost:8000/todos/{id}
- To Save one entry to DB URL(post): localhost:8000/todos/
- To Delete one entry from DB URL(delete): localhost:8000/todos/{id}

To Test the program after running it we can use any API Platform like 'POSTMAN'

# Log file:
- The log file is also maintained if case we are getting any error in the /todoapp folder as 'myLOG.txt'



