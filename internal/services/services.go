package services

import (
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Services struct {
	User            *User
	DemoWorkItem    *DemoWorkItem
	DemoTwoWorkItem *DemoTwoWorkItem
	WorkItemTag     *WorkItemTag
	Authorization   *Authorization
	Authentication  *Authentication
}

func New(logger *zap.SugaredLogger, repos *repos.Repos, pool *pgxpool.Pool) *Services {
	usvc := NewUser(logger, repos)
	demoworkitemsvc := NewDemoWorkItem(logger, repos)
	demotwoworkitemsvc := NewDemoTwoWorkItem(logger, repos)
	workitemtagsvc := NewWorkItemTag(logger, repos)
	authnsvc := NewAuthentication(logger, repos, pool)
	authzsvc, err := NewAuthorization(authnsvc.logger)
	if err != nil {
		panic(fmt.Sprintf("NewAuthorization: %v", err))
	}

	return &Services{
		User:            usvc,
		DemoWorkItem:    demoworkitemsvc,
		DemoTwoWorkItem: demotwoworkitemsvc,
		WorkItemTag:     workitemtagsvc,
		Authorization:   authzsvc,
		Authentication:  authnsvc,
	}
}
