package postgresql_test

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWorkItemComment_Update(t *testing.T) {
	t.Parallel()

	workitemcomment := newRandomWorkItemComment(t, testPool, models.ProjectNameDemo)

	type args struct {
		id     models.WorkItemCommentID
		params models.WorkItemCommentUpdateParams
	}
	type params struct {
		name        string
		args        args
		want        *models.WorkItemComment
		errContains string
	}

	tests := []params{
		{
			name: "updated",
			args: args{
				id:     workitemcomment.WorkItemCommentID,
				params: models.WorkItemCommentUpdateParams{
					// TODO: set fields to update as in crud-api-tests.go.tmpl.bash
				},
			},
			want: func() *models.WorkItemComment {
				u := *workitemcomment
				// TODO: set updated fields to expected values as in crud-api-tests.go.tmpl.bash

				return &u
			}(),
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			r := postgresql.NewWorkItemComment()
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
				require.ErrorContains(t, err, tc.errContains)

				return
			}

			// NOTE: ignore unwanted fields
			// got.UpdatedAt = want.UpdatedAt

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestWorkItemComment_Delete(t *testing.T) {
	t.Parallel()

	workitemcomment := newRandomWorkItemComment(t, testPool, models.ProjectNameDemo)

	type args struct {
		id models.WorkItemCommentID
	}
	type params struct {
		name        string
		args        args
		errContains string
	}

	tests := []params{
		{
			name: "deleted work item comment not found",
			args: args{
				id: workitemcomment.WorkItemCommentID,
			},
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			workItemCommentRepo := postgresql.NewWorkItemComment()
			_, err := workItemCommentRepo.Delete(context.Background(), testPool, tc.args.id)
			require.NoError(t, err)

			_, err = workItemCommentRepo.ByID(context.Background(), testPool, tc.args.id)
			require.ErrorContains(t, err, errNoRows)
			/* row was deleted
			workitemcomment, err = workItemCommentRepo.ByID(context.Background(), testPool, tc.args.id, db.WithDeletedWorkItemCommentOnly())
			require.NoError(t, err)
			assert.Equal(t, workitemcomment.WorkItemCommentID, tc.args.id)
			*/
		})
	}
}

func TestWorkItemComment_ByIndexedQueries(t *testing.T) {
	t.Parallel()

	workitemcomment := newRandomWorkItemComment(t, testPool, models.ProjectNameDemo)
	logger := testutil.NewLogger(t)

	workItemCommentRepo := reposwrappers.NewWorkItemCommentWithRetry(postgresql.NewWorkItemComment(), logger, 10, 65*time.Millisecond)

	uniqueCallback := func(t *testing.T, res *models.WorkItemComment) {
		assert.Equal(t, res.WorkItemCommentID, workitemcomment.WorkItemCommentID)
	}

	uniqueTestCases := []filterTestCase[*models.WorkItemComment]{
		{
			name:       "id",
			filter:     workitemcomment.WorkItemCommentID,
			repoMethod: reflect.ValueOf(workItemCommentRepo.ByID),
			callback:   uniqueCallback,
		},
	}
	for _, tc := range uniqueTestCases {
		runGenericFilterTests(t, tc)
	}
}

func TestWorkItemComment_Create(t *testing.T) {
	t.Parallel()

	type want struct {
		// NOTE: include db-generated fields here to test equality as well
		models.WorkItemCommentCreateParams
	}

	type args struct {
		params models.WorkItemCommentCreateParams
	}

	t.Run("correct_workItemComment", func(t *testing.T) {
		t.Parallel()

		newRandomWorkItemComment(t, testPool, models.ProjectNameDemo)
	})
}
