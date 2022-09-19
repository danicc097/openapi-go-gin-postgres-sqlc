package postgen

import (
	"os"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

func TestMain(m *testing.M) {
	os.Exit(testMain(m))
}

func testMain(m *testing.M) int {
	testutil.Setup()

	return m.Run()
}
