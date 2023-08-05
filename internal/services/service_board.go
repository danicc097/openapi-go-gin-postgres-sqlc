package services

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"go.uber.org/zap"
)

type Board struct {
	logger           *zap.SugaredLogger
	projectRepo      repos.Project
	teamRepo         repos.Team
	workItemTagRepo  repos.WorkItemTag
	workItemTypeRepo repos.WorkItemType
	authzsvc         *Authorization
}

// NewBoard returns a new Board service.
func NewBoard(logger *zap.SugaredLogger, projectRepo repos.Project,
	teamRepo repos.Team,
	workItemTagRepo repos.WorkItemTag,
	workItemTypeRepo repos.WorkItemType,
	authzsvc *Authorization,
) *Board {
	return &Board{
		logger:           logger,
		projectRepo:      projectRepo,
		teamRepo:         teamRepo,
		workItemTagRepo:  workItemTagRepo,
		workItemTypeRepo: workItemTypeRepo,
		authzsvc:         authzsvc,
	}
}

// initialize(teams,), update()
