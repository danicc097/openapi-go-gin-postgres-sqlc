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

type DemoWorkItem struct {
	logger     *zap.SugaredLogger
	demowiRepo repos.DemoWorkItem
	wiRepo     repos.WorkItem
}

type Member struct {
	UserID uuid.UUID           `json:"userID" required:"true"`
	Role   models.WorkItemRole `json:"role" required:"true"`
}

type DemoWorkItemCreateParams struct {
	repos.DemoWorkItemCreateParams
	TagIDs  []int    `json:"tagIDs" required:"true"`
	Members []Member `json:"members" required:"true"`
}

// NewDemoWorkItem returns a new DemoWorkItem service.
func NewDemoWorkItem(logger *zap.SugaredLogger, demowiRepo repos.DemoWorkItem, wiRepo repos.WorkItem) *DemoWorkItem {
	return &DemoWorkItem{
		logger:     logger,
		demowiRepo: demowiRepo,
		wiRepo:     wiRepo,
	}
}

// ByID gets a work item by ID.
func (a *DemoWorkItem) ByID(ctx context.Context, d db.DBTX, id int64) (*db.WorkItem, error) {
	defer newOTELSpan(ctx, "DemoWorkItem.ByID").End()

	wi, err := a.demowiRepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.ByID: %w", err)
	}

	return wi, nil
}

// Create creates a new work item.
func (a *DemoWorkItem) Create(ctx context.Context, d db.DBTX, params DemoWorkItemCreateParams) (*db.WorkItem, error) {
	defer newOTELSpan(ctx, "DemoWorkItem.Create").End()

	demoWi, err := a.demowiRepo.Create(ctx, d, params.DemoWorkItemCreateParams)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.Create: %w", err)
	}

	for _, id := range params.TagIDs {
		err := a.AssignTag(ctx, d, &db.WorkItemWorkItemTagCreateParams{
			WorkItemTagID: id,
			WorkItemID:    demoWi.WorkItemID,
		})
		var ierr *internal.Error
		if err != nil {
			if errors.As(err, &ierr); ierr.Code() != internal.ErrorCodeAlreadyExists {
				return nil, fmt.Errorf("db.CreateWorkItemWorkItemTag: %w", err)
			}
		}
	}

	for _, m := range params.Members {
		err := a.AssignMember(ctx, d, &db.WorkItemMemberCreateParams{
			Member:     m.UserID,
			WorkItemID: demoWi.WorkItemID,
			Role:       m.Role,
		})
		var ierr *internal.Error
		if err != nil {
			if errors.As(err, &ierr); ierr.Code() != internal.ErrorCodeAlreadyExists {
				return nil, fmt.Errorf("db.CreateWorkItemWorkItemMember: %w", err)
			}
		}
	}

	opts := db.WithWorkItemJoin(db.WorkItemJoins{DemoWorkItem: true, Members: true, WorkItemTags: true})
	wi, err := a.demowiRepo.ByID(ctx, d, demoWi.WorkItemID, opts)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.ByID: %w", err)
	}
	// TODO now query workitem.xo.go joining with tags, members, demoWorkItem, etc... and return that instead.
	// all work_item_***_project.go services must return db.WorkItem which has specific project joins, which
	// allows us to share the same generic logic and extend it based on current project, etc.

	return wi, nil
}

// Update updates an existing work item.
func (a *DemoWorkItem) Update(ctx context.Context, d db.DBTX, id int64, params repos.DemoWorkItemUpdateParams) (*db.WorkItem, error) {
	defer newOTELSpan(ctx, "DemoWorkItem.Update").End()

	wi, err := a.demowiRepo.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.Update: %w", err)
	}

	return wi, nil
}

// Delete deletes a work item by ID.
func (a *DemoWorkItem) Delete(ctx context.Context, d db.DBTX, id int64) (*db.WorkItem, error) {
	defer newOTELSpan(ctx, "DemoWorkItem.Delete").End()

	wi, err := a.demowiRepo.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.Delete: %w", err)
	}

	return wi, nil
}

func (a *DemoWorkItem) AssignTag(ctx context.Context, d db.DBTX, params *db.WorkItemWorkItemTagCreateParams) error {
	_, err := db.CreateWorkItemWorkItemTag(ctx, d, params)

	return err
}

func (a *DemoWorkItem) RemoveTag(ctx context.Context, d db.DBTX, tagID int, workItemID int64) error {
	wiwit := &db.WorkItemWorkItemTag{
		WorkItemTagID: tagID,
		WorkItemID:    workItemID,
	}

	return wiwit.Delete(ctx, d)
}

func (a *DemoWorkItem) AssignMember(ctx context.Context, d db.DBTX, params *db.WorkItemMemberCreateParams) error {
	_, err := db.CreateWorkItemMember(ctx, d, params)

	return err
}

func (a *DemoWorkItem) RemoveMember(ctx context.Context, d db.DBTX, memberID uuid.UUID, workItemID int64) error {
	wiwit := &db.WorkItemMember{
		Member:     memberID,
		WorkItemID: workItemID,
	}

	return wiwit.Delete(ctx, d)
}

// repo has Update only, then service has Close() (Update with closed=True), Move() (Update with kanban step change), ...)
// params for dedicated workItem require workItemID (FK-as-PK)
// TBD if useful: ByTag, ByType (for closed workitem searches. open ones simply return everything and filter in client)

func (a *DemoWorkItem) ListDeleted(ctx context.Context, d db.DBTX, teamID int) ([]db.WorkItem, error) {
	// WorkItemsByTeamID with deleted opt, orderby createdAt
	return []db.WorkItem{}, errors.New("not implemented")
}

func (a *DemoWorkItem) List(ctx context.Context, d db.DBTX, teamID int) ([]db.WorkItem, error) {
	// WorkItemsByTeamID with orderby createdAt
	return []db.WorkItem{}, errors.New("not implemented")
}

func (a *DemoWorkItem) Restore(ctx context.Context, d db.DBTX, id int64) (*db.WorkItem, error) {
	return a.demowiRepo.Restore(ctx, d, id)
}
