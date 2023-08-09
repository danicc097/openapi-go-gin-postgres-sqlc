package services

import (
	"context"
	"errors"
	"fmt"
	"strconv"

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
	userRepo   repos.User
	wiSvc      *WorkItem
}

type Member struct {
	Role   models.WorkItemRole `json:"role"   ref:"#/components/schemas/WorkItemRole" required:"true"`
	UserID uuid.UUID           `json:"userID" required:"true"`
}

type DemoWorkItemCreateParams struct {
	repos.DemoWorkItemCreateParams
	TagIDs  []int    `json:"tagIDs"  nullable:"false" required:"true"`
	Members []Member `json:"members" nullable:"false" required:"true"`
}

// NewDemoWorkItem returns a new DemoWorkItem service.
func NewDemoWorkItem(logger *zap.SugaredLogger, demowiRepo repos.DemoWorkItem, wiRepo repos.WorkItem, userRepo repos.User, wiSvc *WorkItem) *DemoWorkItem {
	return &DemoWorkItem{
		logger:     logger,
		demowiRepo: demowiRepo,
		wiRepo:     wiRepo,
		userRepo:   userRepo,
		wiSvc:      wiSvc,
	}
}

// ByID gets a work item by ID.
func (w *DemoWorkItem) ByID(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error) {
	defer newOTelSpan(ctx, "").End()

	wi, err := w.demowiRepo.ByID(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.ByID: %w", err)
	}

	return wi, nil
}

// Create creates a new work item.
func (w *DemoWorkItem) Create(ctx context.Context, d db.DBTX, params DemoWorkItemCreateParams) (*db.WorkItem, error) {
	defer newOTelSpan(ctx, "").End()

	demoWi, err := w.demowiRepo.Create(ctx, d, params.DemoWorkItemCreateParams)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.Create: %w", err)
	}

	for i, id := range params.TagIDs {
		err := w.AssignTag(ctx, d, &db.WorkItemWorkItemTagCreateParams{
			WorkItemTagID: id,
			WorkItemID:    demoWi.WorkItemID,
		})
		var ierr *internal.Error
		if err != nil {
			if errors.As(err, &ierr) && ierr.Code() == models.ErrorCodeAlreadyExists {
				w.logger.Infof("skipping already assigned tag: %s\n", id)

				continue
			}

			// will be done in w.wiSvc.AssignWorkItemTags which returns loc []string{strconv.Itoa(i)}
			// and here we wrap it again in loc: tagIDs which is specific to Create only...
			return nil, internal.WrapErrorWithLocf(err, "", []string{"tagIDs", strconv.Itoa(i)}, "could not assign tag %d", id)
		}
	}

	err = w.wiSvc.AssignWorkItemMembers(ctx, d, demoWi, params.Members)
	if err != nil {
		return nil, internal.WrapErrorWithLocf(err, "", []string{"members"}, "could not assign members")
	}

	// TODO rest response with non pointer required joins as usual, so that it is always up to date
	// (else tests - with response validation - will fail)
	// response validation could be disabled in prod for better availability in place of strictness
	opts := db.WithWorkItemJoin(db.WorkItemJoins{DemoWorkItem: true, AssignedUsers: true, WorkItemTags: true})
	wi, err := w.demowiRepo.ByID(ctx, d, demoWi.WorkItemID, opts)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.ByID: %w", err)
	}

	return wi, nil
}

// Update updates an existing work item.
func (w *DemoWorkItem) Update(ctx context.Context, d db.DBTX, id int, params repos.DemoWorkItemUpdateParams) (*db.WorkItem, error) {
	defer newOTelSpan(ctx, "").End()

	wi, err := w.demowiRepo.Update(ctx, d, id, params)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.Update: %w", err)
	}

	return wi, nil
}

// Delete deletes a work item by ID.
func (w *DemoWorkItem) Delete(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error) {
	defer newOTelSpan(ctx, "").End()

	wi, err := w.wiRepo.Delete(ctx, d, id)
	if err != nil {
		return nil, fmt.Errorf("demowiRepo.Delete: %w", err)
	}

	return wi, nil
}

// TODO: same as assign/remove members.
func (w *DemoWorkItem) AssignTag(ctx context.Context, d db.DBTX, params *db.WorkItemWorkItemTagCreateParams) error {
	_, err := db.CreateWorkItemWorkItemTag(ctx, d, params)

	return err
}

// TODO: same as assign/remove members.
func (w *DemoWorkItem) RemoveTag(ctx context.Context, d db.DBTX, tagID int, workItemID int) error {
	wiwit := &db.WorkItemWorkItemTag{
		WorkItemTagID: tagID,
		WorkItemID:    workItemID,
	}

	return wiwit.Delete(ctx, d)
}

// TODO: remove in favor of assignmembers generic workitem function.
func (w *DemoWorkItem) AssignMember(ctx context.Context, d db.DBTX, params *db.WorkItemAssignedUserCreateParams) error {
	_, err := db.CreateWorkItemAssignedUser(ctx, d, params)

	return err
}

// TODO: remove in favor of removemembers generic workitem function.
func (w *DemoWorkItem) RemoveMember(ctx context.Context, d db.DBTX, memberID uuid.UUID, workItemID int) error {
	wim := &db.WorkItemAssignedUser{
		AssignedUser: memberID,
		WorkItemID:   workItemID,
	}

	return wim.Delete(ctx, d)
}

// repo has Update only, then service has Close() (Update with closed=True), Move() (Update with kanban step change), ...)
// params for dedicated workItem require workItemID (FK-as-PK)
// TBD if useful: ByTag, ByType (for closed workitem searches. open ones simply return everything and filter in client)

func (w *DemoWorkItem) ListDeleted(ctx context.Context, d db.DBTX, teamID int) ([]db.WorkItem, error) {
	// WorkItemsByTeamID with deleted opt, orderby createdAt
	return []db.WorkItem{}, errors.New("not implemented")
}

func (w *DemoWorkItem) List(ctx context.Context, d db.DBTX, teamID int) ([]db.WorkItem, error) {
	// WorkItemsByTeamID with orderby createdAt
	return []db.WorkItem{}, errors.New("not implemented")
}

func (w *DemoWorkItem) Restore(ctx context.Context, d db.DBTX, id int) (*db.WorkItem, error) {
	return w.wiRepo.Restore(ctx, d, id)
}
