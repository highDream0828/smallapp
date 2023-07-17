package dbspeeds

import (
"database/sql"
_ "github.com/lib/pq" // Postgres driver
)

var ( DB *sql.DB )

// Connect DB
func Connect() {

	// Connection string 
	src := "postgres://username:password@host:port/database"

	// Connect to database 
	var err error
	DB, err = sql.Open("postgres", src)  
	if err != nil {
		panic(err)
	}

	// Ping the database to check connection
	DB.Ping()
}

func Deinitialize() {
	DB.Close()
}

// execute queries that retrieve data from the database, like SELECT statements.
func Query(sql string, args ...interface{}) (*sql.Rows, error) {
	return DB.Query(sql, args...)
}

// Execute commands that modify the database, like INSERT, UPDATE and DELETE statements.
func Exec(sql string, args ...interface{}) (sql.Result, error) {
	return DB.Exec(sql, args...)
}
