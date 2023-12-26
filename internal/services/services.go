package services

import (
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Services struct {
	User            *User
	Project         *Project
	Team            *Team
	DemoWorkItem    *DemoWorkItem
	DemoTwoWorkItem *DemoTwoWorkItem
	WorkItemTag     *WorkItemTag
	Authorization   *Authorization
	Authentication  *Authentication
	Notification    *Notification
	TimeEntry       *TimeEntry
	WorkItemType    *WorkItemType
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

	// imagine cyclic dep between these 2, in which case this solves nothing.
	// workitemtagsvc.usersvc may call a function that depends
	// on workitemtagsvc.usersvc.workitemtagsvc being set.
	// FIXME: the only way is to have another service
	// with the logic, which gets really ugly really quickly,
	// dep injection (pass services within services as function params only when cyclic dep found, meanwhile pass
	// to constructor)
	usersvc.workitemtagsvc = workitemtagsvc
	workitemtagsvc.usersvc = usersvc

	return &Services{
		User:            usersvc,
		Team:            teamsvc,
		Project:         projectsvc,
		DemoWorkItem:    demoworkitemsvc,
		DemoTwoWorkItem: demotwoworkitemsvc,
		WorkItemTag:     workitemtagsvc,
		Authorization:   authzsvc,
		Authentication:  authnsvc,
		Notification:    notificationsvc,
		TimeEntry:       timeentrysvc,
		WorkItemType:    workitemtypesvc,
	}
}
