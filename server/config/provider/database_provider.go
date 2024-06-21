package provider

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func NewDbProvider(env *EnvProvider) *sql.DB {
	db, err := sql.Open("postgres", env.databaseUrl)
	if err != nil {
		log.Fatal("Unable to connect to database")
	}

	db.SetMaxOpenConns(env.databaseMaxConns)

	return db
}
