package postgresqltestutil

import (
	"context"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func NewRandomTimeEntry(t *testing.T, pool *pgxpool.Pool, activityID int, userID uuid.UUID, workItemID *int64, teamID *int) (*db.TimeEntry, error) {
	t.Helper()

	teRepo := postgresql.NewTimeEntry()

	ucp := RandomTimeEntryCreateParams(t, activityID, userID, workItemID, teamID)

	te, err := teRepo.Create(context.Background(), pool, ucp)
	if err != nil {
		return nil, err
	}

	return te, nil
}

func RandomTimeEntryCreateParams(t *testing.T, activityID int, userID uuid.UUID, workItemID *int64, teamID *int) db.TimeEntryCreateParams {
	t.Helper()

	return db.TimeEntryCreateParams{
		WorkItemID:      workItemID,
		ActivityID:      activityID,
		TeamID:          teamID,
		UserID:          userID,
		Comment:         testutil.RandomString(20),
		Start:           testutil.RandomDate(),
		DurationMinutes: pointers.New(int(testutil.RandomInt64(10, 400))),
	}
}
