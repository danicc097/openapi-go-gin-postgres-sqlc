package postgresql

import (
	"context"
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
)

// Project represents the repository used for interacting with Project records.
type Project struct {
	q *db.Queries
}

// NewProject instantiates the Project repository.
func NewProject() *Project {
	return &Project{
		q: db.New(),
	}
}

var _ repos.Project = (*Project)(nil)

func (u *Project) ProjectByID(ctx context.Context, d db.DBTX, id int) (*db.Project, error) {
	return db.ProjectByProjectID(ctx, d, id)
}

func (u *Project) ProjectByName(ctx context.Context, d db.DBTX, name string) (*db.Project, error) {
	project, err := db.ProjectByName(ctx, d, name)
	if err != nil {
		return nil, fmt.Errorf("could not get project: %w", parseErrorDetail(err))
	}

	return project, nil
}
