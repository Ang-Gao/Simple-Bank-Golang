package db

import (
	"database/sql"
	"log"
	util "myproject/util"
	"os"
	"testing"

	_ "github.com/lib/pq" //for driver detected as postgres
)

// global var, but only global under same package because the lower case
var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	//load configuration
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("config error")
	}

	//setup connection
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	//create Query
	testQueries = New(testDB)
	os.Exit(m.Run())
}
