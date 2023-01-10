package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq" //for driver detected as postgres
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:abc123@localhost:5432/simple_bank?sslmode=disable"
)

// global var, but only global under same package because the lower case
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	//setup connection
	var err error
	testDB, err = sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	//create Query
	testQueries = New(testDB)
	os.Exit(m.Run())
}
