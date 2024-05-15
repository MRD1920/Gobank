package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/MRD1920/Gobank.git/util"
	_ "github.com/lib/pq" //database driver for postgres
)

var testQueries *Queries
var testDB *sql.DB

// *testing.M is the main entry point for testing in a package. It is used to run test cases in the package.
func TestMain(m *testing.M) {

	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("cannot load config:", err)

	}
	testDB, err = sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)
	os.Exit(m.Run())
}
