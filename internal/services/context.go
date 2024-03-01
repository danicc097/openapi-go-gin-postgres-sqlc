package services

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"
)

type CtxUser struct {
	// db user object with no joins
	*db.User
	Role     Role
	Teams    []db.Team
	Projects []db.Project
	APIKey   *db.UserAPIKey
}

// NewCtxUser returns a new CtxUser.
// Required joins: Teams, Projects.
func NewCtxUser(user *db.User) *CtxUser {
	authzsvc := NewAuthorization(zap.S())
	role, _ := authzsvc.RoleByRank(user.RoleRank)

	return &CtxUser{
		User:     user,
		Role:     role,
		Teams:    *user.MemberTeamsJoin,
		Projects: *user.MemberProjectsJoin,
	}
}
