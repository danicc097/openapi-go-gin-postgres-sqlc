package services

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"go.uber.org/zap"
)

type CtxUser struct {
	// db user object with no joins
	*models.User
	Role     Role
	Teams    []models.Team
	Projects []models.Project
	APIKey   *models.UserAPIKey
}

// NewCtxUser returns a new CtxUser.
// Required joins: Teams, Projects.
func NewCtxUser(user *models.User) *CtxUser {
	authzsvc := NewAuthorization(zap.S())
	role, _ := authzsvc.RoleByRank(user.RoleRank)

	return &CtxUser{
		User:     user,
		Role:     role,
		Teams:    *user.MemberTeamsJoin,
		Projects: *user.MemberProjectsJoin,
	}
}
