package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	connection, err := sql.Open("postgres", "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable")

	if err != nil {
		log.Fatal("Connection to the DB could not be established")
	}

	testQueries = New(connection)

	os.Exit(m.Run())
}
