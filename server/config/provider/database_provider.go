package provider

import (
	"database/sql"
	"log"
	"os"

	"github.com/DATA-DOG/go-txdb"
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

func RegisterTestTxDb() {
	databaseUrl := os.Getenv("DATABASE_URL")
	txdb.Register("txdb", "postgres", databaseUrl)
}

func NewTestDbProvider(env *EnvProvider) *sql.DB {
	db, error := sql.Open("txdb", "TestTransactionDB")

	if error != nil {
		log.Fatal("Unable to connect to database")
	}

	return db
}
