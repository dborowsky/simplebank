package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/dborowsky/simplebank/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connection to DB", err)
	}

	defer testDB.Close()

	testQueries = New(testDB)

	exitCode := m.Run()

	os.Exit(exitCode)
}
