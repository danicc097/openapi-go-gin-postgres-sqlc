package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/format"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type WorkItemTag struct {
	logger *zap.SugaredLogger
	repos  *repos.Repos
}

// NewWorkItemTag returns a new WorkItemTag service.
func NewWorkItemTag(logger *zap.SugaredLogger, repos *repos.Repos) *WorkItemTag {
	return &WorkItemTag{
		logger: logger,
		repos:  repos,
	}
}

// ByID gets a work item tag by ID.
func (wit *WorkItemTag) ByID(ctx context.Context, d db.DBTX, id db.WorkItemTagID) (*db.WorkItemTag, error) {
	defer newOTelSpan().Build(ctx).End()

	witObj, err := wit.repos.WorkItemTag.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemTag.ByID: %w", err)
	}

	return witObj, nil
}

// ByName gets a work item tag by name.
func (wit *WorkItemTag) ByName(ctx context.Context, d db.DBTX, name string, projectID db.ProjectID) (*db.WorkItemTag, error) {
	defer newOTelSpan().Build(ctx).End()

	witObj, err := wit.repos.WorkItemTag.ByName(ctx, d, name, projectID)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemTag.ByName: %w", err)
	}

	return witObj, nil
}

// Create creates a new work item tag.
func (wit *WorkItemTag) Create(ctx context.Context, d db.DBTX, caller CtxUser, params *db.WorkItemTagCreateParams) (*db.WorkItemTag, error) {
	defer newOTelSpan().Build(ctx).End()

	// TODO: user set from authmiddleware should be typed user with predefined useful joins like teams and projects.
	// then pass over ass context. we could replace `caller CtxUser` with just its userID and fetch it every single time with its proper joins,
	// but why would we... caller is something used all over the place for every service method.
	// userInProject := false
	// for _, team := range caller.MemberTeamsJoin {
	// 	if team.ProjectID == params.ProjectID {
	// 		userInProject = true
	// 	}
	// }
	format.PrintJSON(caller)
	userProjects := make([]db.ProjectID, len(caller.Projects))
	for i, p := range caller.Projects {
		userProjects[i] = p.ProjectID
	}
	if slices.Contains(userProjects, params.ProjectID) {
		return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "user is not a member of project %q", internal.ProjectNameByID[params.ProjectID])
	}

	witObj, err := wit.repos.WorkItemTag.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemTag.Create: %w", err)
	}

	return witObj, nil
}

// Update updates an existing work item tag.
func (wit *WorkItemTag) Update(ctx context.Context, d db.DBTX, caller CtxUser, id db.WorkItemTagID, params *db.WorkItemTagUpdateParams) (*db.WorkItemTag, error) {
	defer newOTelSpan().Build(ctx).End()

	witObj, err := wit.repos.WorkItemTag.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("work item tag not found: %w", err)
	}

	userProjects := make([]db.ProjectID, len(caller.Projects))
	for i, p := range caller.Projects {
		userProjects[i] = p.ProjectID
	}
	if slices.Contains(userProjects, witObj.ProjectID) {
		return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "user is not a member of project %q", internal.ProjectNameByID[witObj.ProjectID])
	}

	witObj, err = wit.repos.WorkItemTag.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemTag.Update: %w", err)
	}

	return witObj, nil
}

// Delete deletes a work item tag by ID.
func (wit *WorkItemTag) Delete(ctx context.Context, d db.DBTX, caller CtxUser, id db.WorkItemTagID) (*db.WorkItemTag, error) {
	defer newOTelSpan().Build(ctx).End()

	witObj, err := wit.repos.WorkItemTag.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemTag.Delete: %w", err)
	}

	return witObj, nil
}
