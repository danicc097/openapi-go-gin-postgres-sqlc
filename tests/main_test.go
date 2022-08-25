package tests_test

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("POSTGRES_DB", "postgres_test")
	os.Setenv("IS_TESTING", "1")
	// call flag.Parse() here if TestMain uses flags
	os.Exit(m.Run())
}
