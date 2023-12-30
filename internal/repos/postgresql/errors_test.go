package postgresql_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserFriendlyPgErrors(t *testing.T) {
	t.Parallel()

	witRepo := reposwrappers.NewWorkItemTagWithRetry(postgresql.NewWorkItemTag(), 10, 65*time.Millisecond)

	type want struct {
		db.WorkItemTagCreateParams
	}

	type args struct {
		params db.WorkItemTagCreateParams
	}

	t.Run("unique and foreign key violations show user-friendly errors", func(t *testing.T) {
		t.Parallel()

		ucp := postgresqltestutil.RandomWorkItemTagCreateParams(t, internal.ProjectIDByName[models.ProjectDemo])

		want := want{
			WorkItemTagCreateParams: *ucp,
		}

		args := args{
			params: *ucp,
		}

		got, err := witRepo.Create(context.Background(), testPool, &args.params)
		require.NoError(t, err)

		assert.Equal(t, want.Name, got.Name)
		assert.Equal(t, want.Description, got.Description)
		assert.Equal(t, want.Color, got.Color)
		assert.Equal(t, want.ProjectID, got.ProjectID)

		_, err = witRepo.Create(context.Background(), testPool, &args.params)
		require.Error(t, err)
		require.Error(t, err)

		assert.ErrorContains(t, err, fmt.Sprintf("combination of name=%s and projectID=%d already exists", want.Name, want.ProjectID))

		args.params.ProjectID = -999
		_, err = witRepo.Create(context.Background(), testPool, &args.params)
		require.Error(t, err)
		require.Error(t, err)

		assert.ErrorContains(t, err, fmt.Sprintf("projectID \"%d\" is invalid", args.params.ProjectID))
	})
}
