package postgresql_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqlrandom"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

func TestUserFriendlyPgErrors(t *testing.T) {
	t.Parallel()
	logger := testutil.NewLogger(t)
	witRepo := reposwrappers.NewWorkItemTagWithRetry(postgresql.NewWorkItemTag(), logger, 10, 65*time.Millisecond)

	type want struct {
		models.WorkItemTagCreateParams
	}

	type args struct {
		params models.WorkItemTagCreateParams
	}

	t.Run("unique and foreign key violations show user-friendly errors", func(t *testing.T) {
		t.Parallel()

		ucp := postgresqlrandom.WorkItemTagCreateParams(internal.ProjectIDByName[models.ProjectNameDemo])

		want := want{
			WorkItemTagCreateParams: *ucp,
		}

		args := args{
			params: *ucp,
		}

		got, err := witRepo.Create(t.Context(), testPool, &args.params)
		require.NoError(t, err)

		assert.Equal(t, want.Name, got.Name)
		assert.Equal(t, want.Description, got.Description)
		assert.Equal(t, want.Color, got.Color)
		assert.Equal(t, want.ProjectID, got.ProjectID)

		_, err = witRepo.Create(t.Context(), testPool, &args.params)
		require.Error(t, err)

		require.ErrorContains(t, err, fmt.Sprintf("combination of name=%s and projectID=%d already exists", want.Name, want.ProjectID))

		args.params.ProjectID = -999
		_, err = witRepo.Create(t.Context(), testPool, &args.params)
		require.Error(t, err)

		require.ErrorContains(t, err, fmt.Sprintf("projectID \"%d\" is invalid", args.params.ProjectID))
	})
}
