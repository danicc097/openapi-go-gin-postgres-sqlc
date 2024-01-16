/**
 * Previously defined in postgresqltestutil. However, these should just be used in repo layer tests (private) and use
 * postgresqltestutil create params for fixture factories in service and api layer so that actual service logic is used
 * for creation.
 */

package postgresql_test

import (
	"context"
	"testing"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/postgresqltestutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/stretchr/testify/require"
)

func newRandomActivity(t *testing.T, d db.DBTX, project models.Project) *db.Activity {
	t.Helper()

	activityRepo := postgresql.NewActivity()

	// shared between projects, will require one as params.
	ucp := postgresqltestutil.RandomActivityCreateParams(t, internal.ProjectIDByName[project])

	activity, err := activityRepo.Create(context.Background(), d, ucp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return activity
}

func newRandomDemoWorkItem(t *testing.T, d db.DBTX) *db.WorkItem {
	t.Helper()

	dpwiRepo := postgresql.NewDemoWorkItem()
	// project-specific workitem. for other randomized entities will accept models.Project
	team := newRandomTeam(t, d, internal.ProjectIDByName[models.ProjectDemo])

	kanbanStepID := internal.DemoKanbanStepsIDByName[testutil.RandomFrom(models.AllDemoKanbanStepsValues())]
	workItemTypeID := internal.DemoWorkItemTypesIDByName[testutil.RandomFrom(models.AllDemoWorkItemTypesValues())]
	cp := postgresqltestutil.RandomDemoWorkItemCreateParams(t, kanbanStepID, workItemTypeID, team.TeamID)
	dpwi, err := dpwiRepo.Create(context.Background(), d, cp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return dpwi
}

func newRandomWorkItemTag(t *testing.T, d db.DBTX, projectID db.ProjectID) *db.WorkItemTag {
	t.Helper()

	witRepo := postgresql.NewWorkItemTag()

	ucp := postgresqltestutil.RandomWorkItemTagCreateParams(t, projectID)

	wit, err := witRepo.Create(context.Background(), d, ucp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return wit
}

func newRandomTeam(t *testing.T, d db.DBTX, projectID db.ProjectID) *db.Team {
	t.Helper()

	teamRepo := reposwrappers.NewTeamWithRetry(postgresql.NewTeam(), testutil.NewLogger(t), 3, 200*time.Millisecond)

	ucp := postgresqltestutil.RandomTeamCreateParams(t, projectID)

	team, err := teamRepo.Create(context.Background(), d, ucp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return team
}

func newRandomDemoTwoWorkItem(t *testing.T, d db.DBTX) *db.WorkItem {
	t.Helper()

	dpwiRepo := postgresql.NewDemoTwoWorkItem()
	// project-specific workitem. for other randomized entities will accept models.Project
	team := newRandomTeam(t, d, internal.ProjectIDByName[models.ProjectDemoTwo])

	kanbanStepID := internal.DemoTwoKanbanStepsIDByName[testutil.RandomFrom(models.AllDemoTwoKanbanStepsValues())]
	workItemTypeID := internal.DemoTwoWorkItemTypesIDByName[testutil.RandomFrom(models.AllDemoTwoWorkItemTypesValues())]
	cp := postgresqltestutil.RandomDemoTwoWorkItemCreateParams(t, kanbanStepID, workItemTypeID, team.TeamID)
	dpwi, err := dpwiRepo.Create(context.Background(), d, cp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return dpwi
}

func newRandomUser(t *testing.T, d db.DBTX) *db.User {
	t.Helper()

	logger := testutil.NewLogger(t)

	userRepo := reposwrappers.NewUserWithRetry(postgresql.NewUser(), logger, 5, 65*time.Millisecond)

	ucp := postgresqltestutil.RandomUserCreateParams(t)

	user, err := userRepo.Create(context.Background(), d, ucp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return user
}

func newRandomTimeEntry(t *testing.T, d db.DBTX, activityID db.ActivityID, userID db.UserID, workItemID *db.WorkItemID, teamID *db.TeamID) *db.TimeEntry {
	t.Helper()

	teRepo := reposwrappers.NewTimeEntryWithRetry(postgresql.NewTimeEntry(), testutil.NewLogger(t), 5, 65*time.Millisecond)

	ucp := postgresqltestutil.RandomTimeEntryCreateParams(t, activityID, userID, workItemID, teamID)

	te, err := teRepo.Create(context.Background(), d, ucp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return te
}

func newRandomWorkItemComment(t *testing.T, d db.DBTX, project models.Project) *db.WorkItemComment {
	t.Helper()

	workItemCommentRepo := reposwrappers.NewWorkItemCommentWithRetry(postgresql.NewWorkItemComment(), testutil.NewLogger(t), 3, 200*time.Millisecond)

	var workItemID db.WorkItemID
	switch project {
	case models.ProjectDemo:
		workItemID = newRandomDemoWorkItem(t, d).WorkItemID
	case models.ProjectDemoTwo:
		workItemID = newRandomDemoTwoWorkItem(t, d).WorkItemID
	}

	user := newRandomUser(t, d)
	// these are repo test utils. don't care about logic concerning
	// "is user assigned to the same team as the workitem" or anything similar defined
	// at the service level, unless it's checked at the db level for some reason
	// If we need to test logic like that, use createParams.
	// services and api should use fixture factory instead so that it uses specific service logic for creation.
	// TODO: add project script checking newRandom* strings are not found outside repos

	cp := postgresqltestutil.RandomWorkItemCommentCreateParams(t, user.UserID, workItemID)

	workItemComment, err := workItemCommentRepo.Create(context.Background(), d, cp)
	require.NoError(t, err, "failed to create random entity") // IMPORTANT: must fail. If testing actual failures use random create params instead

	return workItemComment
}
