package services

import (
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Services struct {
	EntityNotification *EntityNotification
	User               *User
	Project            *Project
	Team               *Team
	DemoWorkItem       *DemoWorkItem
	DemoTwoWorkItem    *DemoTwoWorkItem
	WorkItemTag        *WorkItemTag
	Authorization      *Authorization
	Authentication     *Authentication
	Notification       *Notification
	TimeEntry          *TimeEntry
	WorkItemType       *WorkItemType
}

func New(logger *zap.SugaredLogger, repos *repos.Repos, pool *pgxpool.Pool) *Services {
	usersvc := NewUser(logger, repos)
	teamsvc := NewTeam(logger, repos)
	projectsvc := NewProject(logger, repos)
	demoworkitemsvc := NewDemoWorkItem(logger, repos)
	demotwoworkitemsvc := NewDemoTwoWorkItem(logger, repos)
	workitemtagsvc := NewWorkItemTag(logger, repos)
	timeentrysvc := NewTimeEntry(logger, repos)
	workitemtypesvc := NewWorkItemType(logger, repos)
	notificationsvc := NewNotification(logger, repos)
	authnsvc := NewAuthentication(logger, repos, pool)
	authzsvc, err := NewAuthorization(authnsvc.logger)
	if err != nil {
		panic(fmt.Sprintf("NewAuthorization: %v", err))
	}
	entitynotificationsvc := NewEntityNotification(logger, repos)

	return &Services{
		EntityNotification: entitynotificationsvc,
		User:               usersvc,
		Team:               teamsvc,
		Project:            projectsvc,
		DemoWorkItem:       demoworkitemsvc,
		DemoTwoWorkItem:    demotwoworkitemsvc,
		WorkItemTag:        workitemtagsvc,
		Authorization:      authzsvc,
		Authentication:     authnsvc,
		Notification:       notificationsvc,
		TimeEntry:          timeentrysvc,
		WorkItemType:       workitemtypesvc,
	}
}
