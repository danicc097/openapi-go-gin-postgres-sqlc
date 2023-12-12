package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// Project represents the repository used for interacting with Project records.
type Project struct {
	q db.Querier
}

// NewProject instantiates the Project repository.
func NewProject() *Project {
	return &Project{
		q: NewQuerierWrapper(db.New()),
	}
}

var _ repos.Project = (*Project)(nil)

func (u *Project) IsTeamInProject(ctx context.Context, db db.DBTX, arg db.IsTeamInProjectParams) (bool, error) {
	r, err := u.q.IsTeamInProject(ctx, db, arg)
	if err != nil {
		return false, fmt.Errorf("q.IsTeamInProject: %w", parseDBErrorDetail(err))
	}

	return r, nil
}

func (u *Project) ByID(ctx context.Context, d db.DBTX, id db.ProjectID, opts ...db.ProjectSelectConfigOption) (*db.Project, error) {
	project, err := db.ProjectByProjectID(ctx, d, id, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get project: %w", parseDBErrorDetail(err))
	}

	return project, nil
}

func (u *Project) ByName(ctx context.Context, d db.DBTX, name models.Project, opts ...db.ProjectSelectConfigOption) (*db.Project, error) {
	project, err := db.ProjectByName(ctx, d, name, opts...)
	if err != nil {
		return nil, fmt.Errorf("could not get project: %w", parseDBErrorDetail(err))
	}

	return project, nil
}

func (u *Project) UpdateBoardConfig(ctx context.Context, d db.DBTX, projectID db.ProjectID, paths []string, obj any) error {
	sqlstr := `
	UPDATE public.projects
	SET board_config = jsonb_set_deep(board_config, $1, $2)
	WHERE project_id = $3`

	if _, err := d.Exec(ctx, sqlstr, paths, obj, projectID); err != nil {
		return fmt.Errorf("could not update project board config: %w", parseDBErrorDetail(err))
	}

	return nil
}
