package tests

import (
	"context"
	"fmt"
	"testing"

	db "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/xo-templates/tests/got"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

/**
 * TODO: test extensively
 */

func TestM2M(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	u, err := db.UserByUserID(ctx, testPool, uuid.MustParse("8bfb8359-28e0-4039-9259-3c98ada7300d"), db.WithUserJoin(db.UserJoins{Books: true}))
	fmt.Printf("u: %v\n", u)
	assert.NoError(t, err)
	assert.Len(t, *u.BooksJoin, 1)

	u, err = db.UserByUserID(ctx, testPool, uuid.MustParse("78b8db3e-9900-4ca2-9875-fd1eb59acf71"), db.WithUserJoin(db.UserJoins{Books: true}))
	fmt.Printf("u: %v\n", u)
	assert.NoError(t, err)
	assert.Len(t, *u.BooksJoin, 2)
}
