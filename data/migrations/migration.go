package migrations

import (
	_ "github.com/lib/pq" // Postgres driver
	"github.com/highdream0828/smallapp/data/dbspeeds"
)

func Up() {
	// Create the user table
	stmt := `
	CREATE TABLE IF NOT EXISTS "user" (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	);`
		
	dbspeeds.Exec(stmt)         
}

func Down() {
	// Delete the user table
	stmt := `DROP TABLE "user";`
	dbspeeds.Exec(stmt)         
}