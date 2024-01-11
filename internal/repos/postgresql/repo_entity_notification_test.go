package postgresql_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEntityNotification_Update(t *testing.T) {
	t.Parallel()

	entitynotification := postgresqltestutil.NewRandomEntityNotification(t, testPool)

	type args struct {
		id     db.EntityNotificationID
		params db.EntityNotificationUpdateParams
	}
	type params struct {
		name        string
		args        args
		want        *db.EntityNotification
		errContains string
	}

	tests := []params{
		{
			name: "updated",
			args: args{
				id:     entitynotification.EntityNotificationID,
				params: db.EntityNotificationUpdateParams{
					// TODO: set fields to update as in crud-api-tests.go.tmpl.bash
				},
			},
			want: func() *db.EntityNotification {
				u := *entitynotification
				// TODO: set updated fields to expected values as in crud-api-tests.go.tmpl.bash

				return &u
			}(),
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := postgresql.NewEntityNotification()
			got, err := r.Update(context.Background(), testPool, tc.args.id, &tc.args.params)
			if err != nil && tc.errContains == "" {
				t.Errorf("unexpected error: %v", err)

				return
			}
			if tc.errContains != "" {
				if err == nil {
					t.Errorf("expected error but got nothing")

					return
				}
				assert.ErrorContains(t, err, tc.errContains)

				return
			}

			// NOTE: ignore unwanted fields
			// got.UpdatedAt = want.UpdatedAt

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestEntityNotification_SoftDelete(t *testing.T) {
	t.Parallel()

	entitynotification := postgresqltestutil.NewRandomEntityNotification(t, testPool)

	type args struct {
		id db.EntityNotificationID
	}
	type params struct {
		name        string
		args        args
		errContains string
	}

	tests := []params{
		{
			name: "deleted entity notification not found",
			args: args{
				id: entitynotification.EntityNotificationID,
			},
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			entityNotificationRepo := postgresql.NewEntityNotification()
			_, err := entityNotificationRepo.Delete(context.Background(), testPool, tc.args.id)
			require.NoError(t, err)

			_, err = entityNotificationRepo.ByID(context.Background(), testPool, tc.args.id)
			require.ErrorContains(t, err, errNoRows)

			entitynotification, err = entityNotificationRepo.ByID(context.Background(), testPool, tc.args.id, db.WithDeletedEntityNotificationOnly())
			require.NoError(t, err)
			assert.Equal(t, entitynotification.EntityNotificationID, tc.args.id)
		})
	}
}

func TestEntityNotification_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	entitynotification := postgresqltestutil.NewRandomEntityNotification(t, testPool)

	entityNotificationRepo := reposwrappers.NewEntityNotificationWithRetry(postgresql.NewEntityNotification(), 10, 65*time.Millisecond)

	uniqueCallback := func(t *testing.T, res *db.EntityNotification) {
		assert.Equal(t, res.EntityNotificationID, entitynotification.EntityNotificationID)
	}

	uniqueTestCases := []filterTestCase[*db.EntityNotification]{
		{
			name:       "id",
			filter:     entitynotification.EntityNotificationID,
			repoMethod: reflect.ValueOf(entityNotificationRepo.ByID),
			callback:   uniqueCallback,
		},
	}
	for _, tc := range uniqueTestCases {
		tc := tc
		runGenericFilterTests(t, tc)
	}
}

func TestEntityNotification_Create(t *testing.T) {
	t.Parallel()

	entityNotificationRepo := reposwrappers.NewEntityNotificationWithRetry(postgresql.NewEntityNotification(), 10, 65*time.Millisecond)

	type want struct {
		// NOTE: include db-generated fields here to test equality as well
		db.EntityNotificationCreateParams
	}

	type args struct {
		params db.EntityNotificationCreateParams
	}

	t.Run("correct_entityNotification", func(t *testing.T) {
		t.Parallel()

		entityNotificationCreateParams := postgresqltestutil.RandomEntityNotificationCreateParams(t)

		want := want{
			EntityNotificationCreateParams: *entityNotificationCreateParams,
		}

		args := args{
			params: *entityNotificationCreateParams,
		}

		got, err := entityNotificationRepo.Create(context.Background(), testPool, &args.params)
		require.NoError(t, err)

		assert.Equal(t, want.ID, got.ID)
		assert.Equal(t, want.Message, got.Message)
		assert.Equal(t, want.Topic, got.Topic)
	})

	// implement if needed
	t.Run("check constraint raises violation error", func(t *testing.T) {
		t.Skip("not implemented")
		t.Parallel()

		entityNotificationCreateParams := postgresqltestutil.RandomEntityNotificationCreateParams(t)
		// NOTE: update params to trigger check error

		args := args{
			params: *entityNotificationCreateParams,
		}

		_, err := entityNotificationRepo.Create(context.Background(), testPool, &args.params)
		require.Error(t, err)

		assert.ErrorContains(t, err, errViolatesCheckConstraint)
	})
}
