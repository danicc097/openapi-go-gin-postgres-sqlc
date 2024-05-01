package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
)

// Project represents the repository used for interacting with Project records.
type Project struct {
	q models.Querier
}

// NewProject instantiates the Project repository.
func NewProject() *Project {
	return &Project{
		q: NewQuerierWrapper(models.New()),
	}
}

var _ repos.Project = (*Project)(nil)

func (u *Project) IsTeamInProject(ctx context.Context, d models.DBTX, arg models.IsTeamInProjectParams) (bool, error) {
	r, err := u.q.IsTeamInProject(ctx, d, arg)
	if err != nil {
		return false, fmt.Errorf("q.IsTeamInProject: %w", ParseDBErrorDetail(err))
	}

	return r, nil
}

func (u *Project) ByID(ctx context.Context, d models.DBTX, id models.ProjectID, opts ...models.ProjectSelectConfigOption) (*models.Project, error) {
	project, err := models.ProjectByProjectID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get project: %w", ParseDBErrorDetail(err))
	}

	return project, nil
}

func (u *Project) ByName(ctx context.Context, d models.DBTX, name models.ProjectName, opts ...models.ProjectSelectConfigOption) (*models.Project, error) {
	project, err := models.ProjectByName(ctx, d, name, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get project: %w", ParseDBErrorDetail(err))
	}

	return project, nil
}

func (u *Project) UpdateBoardConfig(ctx context.Context, d models.DBTX, projectID models.ProjectID, paths []string, obj any) error {
	sqlstr := `
	UPDATE public.projects
	SET board_config = jsonb_set_deep(board_config, $1, $2)
	WHERE project_id = $3`

	if _, err := d.Exec(ctx, sqlstr, paths, obj, projectID); err != nil {
		return fmt.Errorf("could not update project board config: %w", ParseDBErrorDetail(err))
	}

	return nil
}
