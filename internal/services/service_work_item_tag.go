package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
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
func (wit *WorkItemTag) Create(ctx context.Context, d db.DBTX, caller *db.User, params *db.WorkItemTagCreateParams) (*db.WorkItemTag, error) {
	defer newOTelSpan().Build(ctx).End()

	// TODO: we should use GetUserInProject and not rely on db.User joins.
	// but we want the ctx user set from auth mw
	// to be as light as possible.
	// userInProject := false
	// for _, team := range *caller.MemberTeamsJoin {
	// 	if team.ProjectID == params.ProjectID {
	// 		userInProject = true
	// 	}
	// }

	userInProject, err := wit.repos.User.IsUserInProject(ctx, d, db.IsUserInProjectParams{
		UserID:    caller.UserID.UUID,
		ProjectID: int32(params.ProjectID),
	})
	if err != nil {
		return nil, fmt.Errorf("repos.User.IsUserInProject: %w", err)
	}
	fmt.Printf("userInProject: %v\n", userInProject)

	witObj, err := wit.repos.WorkItemTag.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemTag.Create: %w", err)
	}

	return witObj, nil
}

// Update updates an existing work item tag.
func (wit *WorkItemTag) Update(ctx context.Context, d db.DBTX, caller *db.User, id db.WorkItemTagID, params *db.WorkItemTagUpdateParams) (*db.WorkItemTag, error) {
	defer newOTelSpan().Build(ctx).End()

	witObj, err := wit.repos.WorkItemTag.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemTag.Update: %w", err)
	}

	return witObj, nil
}

// Delete deletes a work item tag by ID.
func (wit *WorkItemTag) Delete(ctx context.Context, d db.DBTX, caller *db.User, id db.WorkItemTagID) (*db.WorkItemTag, error) {
	defer newOTelSpan().Build(ctx).End()

	witObj, err := wit.repos.WorkItemTag.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemTag.Delete: %w", err)
	}

	return witObj, nil
}
