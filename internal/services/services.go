package services

import (
	"fmt"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

type Services struct {
	Activity        *Activity
	WorkItemComment *WorkItemComment
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
	WorkItem        *WorkItem
}

func New(logger *zap.SugaredLogger, repos *repos.Repos, pool *pgxpool.Pool) *Services {
	usersvc := NewUser(logger, repos)
	teamsvc := NewTeam(logger, repos)
	projectsvc := NewProject(logger, repos)
	activitysvc := NewActivity(logger, repos)
	demoworkitemsvc := NewDemoWorkItem(logger, repos)
	demotwoworkitemsvc := NewDemoTwoWorkItem(logger, repos)
	workitemsvc := NewWorkItem(logger, repos)
	workitemtagsvc := NewWorkItemTag(logger, repos)
	timeentrysvc := NewTimeEntry(logger, repos)
	workitemtypesvc := NewWorkItemType(logger, repos)
	notificationsvc := NewNotification(logger, repos)
	authnsvc := NewAuthentication(logger, repos, pool)
	authzsvc, err := NewAuthorization(logger)
	if err != nil {
		panic(fmt.Sprintf("NewAuthorization: %v", err))
	}

	// this would solve nothing. workitemtagsvc.usersvc may call a function that depends
	// on workitemtagsvc.usersvc.workitemtagsvc being set.
	// Instead create needed svcs within the function that needs them.
	// usersvc.workitemtagsvc = workitemtagsvc
	// workitemtagsvc.usersvc = usersvc
	// solution: create needed services in each service method, e.g. witSvc := NewWorkItemTag(u.logger, u.repos)
	// in our user service methods.
	workitemcommentsvc := NewWorkItemComment(logger, repos)

	// using struct for easier automated appending later on. forced fill via exhaustruct
	return &Services{
		WorkItemComment: workitemcommentsvc,
		Activity:        activitysvc,
		User:            usersvc,
		Team:            teamsvc,
		Project:         projectsvc,
		DemoWorkItem:    demoworkitemsvc,
		DemoTwoWorkItem: demotwoworkitemsvc,
		WorkItem:        workitemsvc,
		WorkItemTag:     workitemtagsvc,
		Authorization:   authzsvc,
		Authentication:  authnsvc,
		Notification:    notificationsvc,
		TimeEntry:       timeentrysvc,
		WorkItemType:    workitemtypesvc,
	}
}
