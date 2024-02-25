package services

import "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"

type CtxUser struct {
	// db user object with no joins
	*db.User
	Teams    []db.Team
	Projects []db.Project
	APIKey   *db.UserAPIKey
}

// NewCtxUser returns a new CtxUser.
// Required joins: Teams, Projects.
func NewCtxUser(user *db.User) *CtxUser {
	return &CtxUser{
		User:     user,
		Teams:    *user.MemberTeamsJoin,
		Projects: *user.ProjectsJoin,
	}
}
