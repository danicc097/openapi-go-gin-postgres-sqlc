package services

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type EntityNotification struct {
	logger *zap.SugaredLogger
	repos  *repos.Repos
}

// NewEntityNotification returns a new entity notification service.
func NewEntityNotification(logger *zap.SugaredLogger, repos *repos.Repos) *EntityNotification {
	return &EntityNotification{
		logger: logger,
		repos:  repos,
	}
}

// ByID gets a entity notification by ID.
func (t *EntityNotification) ByID(ctx context.Context, d db.DBTX, id db.EntityNotificationID) (*db.EntityNotification, error) {
	defer newOTelSpan().Build(ctx).End()

	entityNotification, err := t.repos.EntityNotification.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.EntityNotification.ByID: %w", err)
	}

	return entityNotification, nil
}

// Create creates a new entity notification.
func (t *EntityNotification) Create(ctx context.Context, d db.DBTX, params *db.EntityNotificationCreateParams) (*db.EntityNotification, error) {
	defer newOTelSpan().Build(ctx).End()

	entityNotification, err := t.repos.EntityNotification.Create(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("repos.EntityNotification.Create: %w", err)
	}

	return entityNotification, nil
}

// Update updates an existing entity notification.
func (t *EntityNotification) Update(ctx context.Context, d db.DBTX, id db.EntityNotificationID, params *db.EntityNotificationUpdateParams) (*db.EntityNotification, error) {
	defer newOTelSpan().Build(ctx).End()

	entityNotification, err := t.repos.EntityNotification.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("repos.EntityNotification.Update: %w", err)
	}

	return entityNotification, nil
}

// Delete deletes an existing entity notification.
func (t *EntityNotification) Delete(ctx context.Context, d db.DBTX, id db.EntityNotificationID) (*db.EntityNotification, error) {
	defer newOTelSpan().Build(ctx).End()

	entityNotification, err := t.repos.EntityNotification.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("repos.EntityNotification.Delete: %w", err)
	}

	return entityNotification, nil
}

