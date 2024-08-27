package database

import (
	"database/sql" // Importing the database/sql package to interact with the SQL database
	"log"          // Importing the log package to log errors or information

	_ "github.com/lib/pq" // Importing the PostgreSQL driver (lib/pq) anonymously; it registers itself with database/sql
)

var DB *sql.DB // Declaring a global variable `DB` of type `*sql.DB`, which represents a database connection pool

// InitDB initializes the database connection using the provided connection string.
func InitDB(connectionString string) {
	var err error // Declaring a variable `err` to capture any errors during the process

	// Open a new database connection using the PostgreSQL driver and the connection string.
	// `sql.Open` does not establish a connection immediately, it just prepares the connection.
	DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err) // If there is an error, log it and stop the program
	}

	// `DB.Ping` actually establishes a connection to the database and verifies that it is reachable.
	if err = DB.Ping(); err != nil {
		log.Fatal(err) // If there is an error, log it and stop the program
	}
}
