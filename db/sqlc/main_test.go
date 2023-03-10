package db

import (
	"database/sql"
	"github.com/ZeinabAshjaei/go-master/utils"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("config cannot be loaded", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("Connection to the DB could not be established")
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
