package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// EntityNotification represents the repository used for interacting with entity notification records.
type EntityNotification struct {
	q db.Querier
}

// NewEntityNotification instantiates the entity notification repository.
func NewEntityNotification() *EntityNotification {
	return &EntityNotification{
		q: NewQuerierWrapper(db.New()),
	}
}

var _ repos.EntityNotification = (*EntityNotification)(nil)

func (t *EntityNotification) Create(ctx context.Context, d db.DBTX, params *db.EntityNotificationCreateParams) (*db.EntityNotification, error) {
	entityNotification, err := db.CreateEntityNotification(ctx, d, params)
	if err != nil {
		return nil, fmt.Errorf("could not create entityNotification: %w", ParseDBErrorDetail(err))
	}

	return entityNotification, nil
}

func (t *EntityNotification) Update(ctx context.Context, d db.DBTX, id db.EntityNotificationID, params *db.EntityNotificationUpdateParams) (*db.EntityNotification, error) {
	entityNotification, err := t.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("could not get entity notification by id %w", ParseDBErrorDetail(err))
	}

	entityNotification.SetUpdateParams(params)

	entityNotification, err = entityNotification.Update(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not update entity notification: %w", ParseDBErrorDetail(err))
	}

	return entityNotification, err
}

func (t *EntityNotification) ByID(ctx context.Context, d db.DBTX, id db.EntityNotificationID, opts ...db.EntityNotificationSelectConfigOption) (*db.EntityNotification, error) {
	entityNotification, err := db.EntityNotificationByEntityNotificationID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get entity notification: %w", ParseDBErrorDetail(err))
	}

	return entityNotification, nil
}

func (t *EntityNotification) Delete(ctx context.Context, d db.DBTX, id db.EntityNotificationID) (*db.EntityNotification, error) {
	entityNotification := &db.EntityNotification{
		EntityNotificationID: id,
	}

	err := entityNotification.SoftDelete(ctx, d) // use SoftDelete if a deleted_at column exists.
	if err != nil {
		return nil, fmt.Errorf("could not delete entity notification: %w", ParseDBErrorDetail(err))
	}

	return entityNotification, err
}

func (t *EntityNotification) Restore(ctx context.Context, d db.DBTX, id db.EntityNotificationID) (*db.EntityNotification, error) {
	entityNotification := &db.EntityNotification{
		EntityNotificationID: id,
	}

	entityNotification, err := entityNotification.Restore(ctx, d)
	if err != nil {
		return nil, fmt.Errorf("could not restore entity notification: %w", ParseDBErrorDetail(err))
	}

	return entityNotification, err
}
