package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	models1 "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"go.uber.org/zap"
)

type WorkItemComment struct {
	logger   *zap.SugaredLogger
	repos    *repos.Repos
	authzsvc *Authorization
}

// NewWorkItemComment returns a new work item comment service.
func NewWorkItemComment(logger *zap.SugaredLogger, repos *repos.Repos) *WorkItemComment {
	authzsvc := NewAuthorization(logger)

	return &WorkItemComment{
		logger:   logger,
		repos:    repos,
		authzsvc: authzsvc,
	}
}

// ByID gets a work item comment by ID.
func (t *WorkItemComment) ByID(ctx context.Context, d models1.DBTX, id models1.WorkItemCommentID) (*models1.WorkItemComment, error) {
	defer newOTelSpan().Build(ctx).End()

	workItemComment, err := t.repos.WorkItemComment.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemComment.ByID: %w", err)
	}

	return workItemComment, nil
}

// Create creates a new work item comment.
func (t *WorkItemComment) Create(ctx context.Context, d models1.DBTX, params *models1.WorkItemCommentCreateParams) (*models1.WorkItemComment, error) {
	defer newOTelSpan().Build(ctx).End()

	workItemComment, err := t.repos.WorkItemComment.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemComment.Create: %w", err)
	}

	return workItemComment, nil
}

// Update updates an existing work item comment.
func (t *WorkItemComment) Update(ctx context.Context, d models1.DBTX, caller CtxUser, id models1.WorkItemCommentID, params *models1.WorkItemCommentUpdateParams) (*models1.WorkItemComment, error) {
	defer newOTelSpan().Build(ctx).End()

	workItemComment, err := t.repos.WorkItemComment.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemComment.Update: %w", err)
	}

	return workItemComment, nil
}

// Delete deletes an existing work item comment.
func (t *WorkItemComment) Delete(ctx context.Context, d models1.DBTX, caller CtxUser, id models1.WorkItemCommentID) (*models1.WorkItemComment, error) {
	defer newOTelSpan().Build(ctx).End()

	workItemComment, err := t.repos.WorkItemComment.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemComment.ByID: %w", err)
	}

	err = t.authzsvc.HasRequiredRole(caller.Role, models.RoleAdmin)
	isAdmin := err == nil

	if workItemComment.UserID != caller.UserID && !isAdmin {
		return nil, internal.NewErrorf(models.ErrorCodeUnauthorized, "cannot delete another user's comment")
	}

	workItemComment, err = t.repos.WorkItemComment.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.WorkItemComment.Delete: %w", err)
	}

	return workItemComment, nil
}
