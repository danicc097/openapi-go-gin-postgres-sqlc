package services

import (
	"context"
	"errors"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

type DemoTwoWorkItem struct {
	logger        *zap.SugaredLogger
	demotwowiRepo repos.DemoTwoWorkItem
	wiRepo        repos.WorkItem
}

type DemoTwoWorkItemCreateParams struct {
	repos.DemoTwoWorkItemCreateParams
	TagIDs  []int    `json:"tagIDs"  required:"true"`
	Members []Member `json:"members" required:"true"`
}

// NewDemoTwoWorkItem returns a new DemoTwoWorkItem service.
func NewDemoTwoWorkItem(logger *zap.SugaredLogger, demowiRepo repos.DemoTwoWorkItem, wiRepo repos.WorkItem) *DemoTwoWorkItem {
	return &DemoTwoWorkItem{
		logger:        logger,
		demotwowiRepo: demowiRepo,
		wiRepo:        wiRepo,
	}
}

// ByID gets a work item by ID.
func (w *DemoTwoWorkItem) ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error) {
	defer newOTELSpan(ctx, "DemoTwoWorkItem.ByID").End()

	wi, err := w.demotwowiRepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.ByID: %w", err)
	}

	return wi, nil
}

// Create creates a new work item.
func (w *DemoTwoWorkItem) Create(ctx context.Context, d db.DBTX, params DemoTwoWorkItemCreateParams) (*db.WorkItem, error) {
	defer newOTELSpan(ctx, "DemoTwoWorkItem.Create").End()

	demoWi, err := w.demotwowiRepo.Create(ctx, d, params.DemoTwoWorkItemCreateParams)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.Create: %w", err)
	}

	for _, id := range params.TagIDs {
		err := w.AssignTag(ctx, d, &db.WorkItemWorkItemTagCreateParams{
			WorkItemTagID: id,
			WorkItemID:    demoWi.WorkItemID,
		})
		var ierr *internal.Error
		if err != nil {
			if errors.As(err, &ierr); ierr.Code() != models.ErrorCodeAlreadyExists {
				return nil, fmt.Errorf("db.CreateWorkItemWorkItemTag: %w", err)
			}
		}
	}

	for _, m := range params.Members {
		err := w.AssignMember(ctx, d, &db.WorkItemAssignedUserCreateParams{
			AssignedUser: m.UserID,
			WorkItemID:   demoWi.WorkItemID,
			Role:         m.Role,
		})
		var ierr *internal.Error
		if err != nil {
			if errors.As(err, &ierr); ierr.Code() != models.ErrorCodeAlreadyExists {
				return nil, fmt.Errorf("a.AssignMember: %w", err)
			}
		}
	}

	// TODO rest response with non pointer required joins as usual, so that it is always up to date
	// (else tests - with response validation - will fail)
	// response validation could be disabled in prod for better availability in place of strictness
	opts := db.WithWorkItemJoin(db.WorkItemJoins{DemoTwoWorkItem: true, AssignedUsers: true, WorkItemTags: true})
	wi, err := w.demotwowiRepo.ByID(ctx, d, demoWi.WorkItemID, opts)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.ByID: %w", err)
	}

	return wi, nil
}

// Update updates an existing work item.
func (w *DemoTwoWorkItem) Update(ctx context.Context, d db.DBTX, id int, params repos.DemoTwoWorkItemUpdateParams) (*db.WorkItem, error) {
	defer newOTELSpan(ctx, "DemoTwoWorkItem.Update").End()

	wi, err := w.demotwowiRepo.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.Update: %w", err)
	}

	return wi, nil
}

// Delete deletes a work item by ID.
func (w *DemoTwoWorkItem) Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error) {
	defer newOTELSpan(ctx, "DemoTwoWorkItem.Delete").End()

	wi, err := w.demotwowiRepo.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.Delete: %w", err)
	}

	return wi, nil
}

func (w *DemoTwoWorkItem) AssignTag(ctx context.Context, d db.DBTX, params *db.WorkItemWorkItemTagCreateParams) error {
	_, err := db.CreateWorkItemWorkItemTag(ctx, d, params)

	return err
}

func (w *DemoTwoWorkItem) RemoveTag(ctx context.Context, d db.DBTX, tagID int, workItemID int) error {
	wiwit := &db.WorkItemWorkItemTag{
		WorkItemTagID: tagID,
		WorkItemID:    workItemID,
	}

	return wiwit.Delete(ctx, d)
}

func (w *DemoTwoWorkItem) AssignMember(ctx context.Context, d db.DBTX, params *db.WorkItemAssignedUserCreateParams) error {
	_, err := db.CreateWorkItemAssignedUser(ctx, d, params)

	return err
}

func (w *DemoTwoWorkItem) RemoveMember(ctx context.Context, d db.DBTX, memberID uuid.UUID, workItemID int) error {
	wim := &db.WorkItemAssignedUser{
		AssignedUser: memberID,
		WorkItemID:   workItemID,
	}

	return wim.Delete(ctx, d)
}

// repo has Update only, then service has Close() (Update with closed=True), Move() (Update with kanban step change), ...)
// params for dedicated workItem require workItemID (FK-as-PK)
// TBD if useful: ByTag, ByType (for closed workitem searches. open ones simply return everything and filter in client)

func (w *DemoTwoWorkItem) ListDeleted(ctx context.Context, d db.DBTX, teamID int) ([]db.WorkItem, error) {
	// WorkItemsByTeamID with deleted opt, orderby createdAt
	return []db.WorkItem{}, errors.New("not implemented")
}

func (w *DemoTwoWorkItem) List(ctx context.Context, d db.DBTX, teamID int) ([]db.WorkItem, error) {
	// WorkItemsByTeamID with orderby createdAt
	return []db.WorkItem{}, errors.New("not implemented")
}

func (w *DemoTwoWorkItem) Restore(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error) {
	return w.demotwowiRepo.Restore(ctx, d, id)
}
