package postgen

import (
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/tests"
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	tests.Setup()

	return m.Run()
}
