package sqldb

import (
	"database/sql"
)

// DB is a global variable to hold db connection
var DB *sql.DB

// ConnectDB opens a connection to the database
func ConnectDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root123@tcp(127.0.0.1:3306)/Sendify")
	if err != nil {
		panic(err.Error())
	}
	DB = db
	return DB
}
