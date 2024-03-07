package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
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

	if err := wit.validateCreateParams(d, caller, params); err != nil {
		return nil, err
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

	if err := wit.validateUpdateParams(d, caller, id, params); err != nil {
		return nil, err
	}

	witObj, err := wit.repos.WorkItemTag.Update(ctx, d, id, params)
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

func (wit *WorkItemTag) validateCreateParams(d db.DBTX, caller CtxUser, params *db.WorkItemTagCreateParams) error {
	if err := wit.validateBaseParams(validateModeCreate, d, caller, nil, params); err != nil {
		return err
	}

	return nil
}

func (wit *WorkItemTag) validateUpdateParams(d db.DBTX, caller CtxUser, id db.WorkItemTagID, params *db.WorkItemTagUpdateParams) error {
	if err := wit.validateBaseParams(validateModeUpdate, d, caller, &id, params); err != nil {
		return err
	}

	return nil
}

func (wit *WorkItemTag) validateBaseParams(mode validateMode, d db.DBTX, caller CtxUser, id *db.WorkItemTagID, params db.WorkItemTagParams) error {
	var projectID db.ProjectID

	switch {
	case params.GetProjectID() != nil:
		projectID = *params.GetProjectID()
	case id != nil: // update may change a tags project, so default to current tag's project last
		witObj, err := wit.repos.WorkItemTag.ByID(context.Background(), d, *id)
		if err != nil {
			return fmt.Errorf("work item tag not found: %w", err)
		}
		projectID = witObj.ProjectID
	default:
		return internal.NewErrorf(models.ErrorCodeInvalidArgument, "missing project parameter")
	}

	userProjects := make([]db.ProjectID, len(caller.Projects))
	for i, p := range caller.Projects {
		userProjects[i] = p.ProjectID
	}
	if !slices.Contains(userProjects, projectID) {
		return internal.NewErrorf(models.ErrorCodeUnauthorized, "user is not a member of project %q", internal.ProjectNameByID[projectID])
	}

	return nil
}
