package services

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"go.uber.org/zap"
)

type Project struct {
	logger      *zap.Logger
	projectRepo repos.Project
	teamRepo    repos.Team
	authzsvc    *Authorization
}

// NewProject returns a new Project service.
func NewProject(logger *zap.Logger,
	projectRepo repos.Project,
	teamRepo repos.Team,
	authzsvc *Authorization,
) *Project {
	return &Project{
		logger:      logger,
		projectRepo: projectRepo,
		teamRepo:    teamRepo,
		authzsvc:    authzsvc,
	}
}

// addTeam, removeTeam, projectByID
