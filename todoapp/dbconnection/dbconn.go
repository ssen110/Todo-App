package dbconn

import (
	"github.com/ssen110/test/constants"
	"database/sql"
	"fmt"

)
// DB set up
func SetupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", constants.DB_USER, constants.DB_PASSWORD, constants.DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

// Function for handling errors
func checkErr(err error) {
	if err != nil{
		panic(err)
	}
}
