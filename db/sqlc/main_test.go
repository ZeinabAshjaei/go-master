package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable")

	if err != nil {
		log.Fatal("Connection to the DB could not be established")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
