package repository

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// Step 1. Init section
	os.Setenv("DBHOST","localhost")
	os.Setenv("DBUSER","postgres")
	os.Setenv("DBPASSWORD","mysecretpassword")
	os.Setenv("DBNAME","mgt")
	os.Setenv("DBPORT","5432")

	// Step 2: Run it
	code := m.Run()


	// Step 3. TearDown section


	// Step 4. Finalize
	os.Exit(code)
}
