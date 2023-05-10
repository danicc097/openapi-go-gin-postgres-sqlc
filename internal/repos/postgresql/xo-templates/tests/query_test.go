package tests

import (
	"context"
	"fmt"
	"testing"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/xo-templates/tests/got"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestM2M(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	id := uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d")
	u, err := db.UserByUserID(ctx, testPool, id, db.WithUserJoin(db.UserJoins{Books: true}))
	fmt.Printf("u: %v\n", u)
	assert.NoError(t, err)
	assert.Len(t, u.BooksJoin, 1)
}
