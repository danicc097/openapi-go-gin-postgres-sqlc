package services

import (
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/reposwrappers"
)

// CreateRepos creates repositories for service usage.
func CreateRepos() *repos.Repos {
	activityrepo := reposwrappers.NewActivityWithTracing(
		reposwrappers.NewActivityWithTimeout(
			postgresql.NewActivity(),
			reposwrappers.ActivityWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	workitemrepo := reposwrappers.NewWorkItemWithTracing(
		reposwrappers.NewWorkItemWithTimeout(
			postgresql.NewWorkItem(),
			reposwrappers.WorkItemWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	demoworkitemrepo := reposwrappers.NewDemoWorkItemWithTracing(
		reposwrappers.NewDemoWorkItemWithTimeout(
			postgresql.NewDemoWorkItem(),
			reposwrappers.DemoWorkItemWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	demotwoworkitemrepo := reposwrappers.NewDemoTwoWorkItemWithTracing(
		reposwrappers.NewDemoTwoWorkItemWithTimeout(
			postgresql.NewDemoTwoWorkItem(),
			reposwrappers.DemoTwoWorkItemWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	workitemtagrepo := reposwrappers.NewWorkItemTagWithTracing(
		reposwrappers.NewWorkItemTagWithTimeout(
			postgresql.NewWorkItemTag(),
			reposwrappers.WorkItemTagWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	workitemcommentrepo := reposwrappers.NewWorkItemCommentWithTracing(
		reposwrappers.NewWorkItemCommentWithTimeout(
			postgresql.NewWorkItemComment(),
			reposwrappers.WorkItemCommentWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	workitemtyperepo := reposwrappers.NewWorkItemTypeWithTracing(
		reposwrappers.NewWorkItemTypeWithTimeout(
			postgresql.NewWorkItemType(),
			reposwrappers.WorkItemTypeWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	projectrepo := reposwrappers.NewProjectWithTracing(
		reposwrappers.NewProjectWithTimeout(
			postgresql.NewProject(),
			reposwrappers.ProjectWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	urepo := reposwrappers.NewUserWithTracing(
		reposwrappers.NewUserWithTimeout(
			postgresql.NewUser(),
			reposwrappers.UserWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	notifrepo := reposwrappers.NewNotificationWithTracing(
		reposwrappers.NewNotificationWithTimeout(
			postgresql.NewNotification(),
			reposwrappers.NotificationWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	ksrepo := reposwrappers.NewKanbanStepWithTracing(
		reposwrappers.NewKanbanStepWithTimeout(
			postgresql.NewKanbanStep(),
			reposwrappers.KanbanStepWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	teamrepo := reposwrappers.NewTeamWithTracing(
		reposwrappers.NewTeamWithTimeout(
			postgresql.NewTeam(),
			reposwrappers.TeamWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	timeentryrepo := reposwrappers.NewTimeEntryWithTracing(
		reposwrappers.NewTimeEntryWithTimeout(
			postgresql.NewTimeEntry(),
			reposwrappers.TimeEntryWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)
	entitynotificationrepo := reposwrappers.NewEntityNotificationWithTracing(
		reposwrappers.NewEntityNotificationWithTimeout(
			postgresql.NewEntityNotification(),
			reposwrappers.EntityNotificationWithTimeoutConfig{},
		),
		postgresql.OtelName,
		nil,
	)

	return &repos.Repos{
		EntityNotification: entitynotificationrepo,
		Activity:           activityrepo,
		DemoTwoWorkItem:    demotwoworkitemrepo,
		DemoWorkItem:       demoworkitemrepo,
		KanbanStep:         ksrepo,
		Notification:       notifrepo,
		Project:            projectrepo,
		Team:               teamrepo,
		TimeEntry:          timeentryrepo,
		User:               urepo,
		WorkItem:           workitemrepo,
		WorkItemComment:    workitemcommentrepo,
		WorkItemTag:        workitemtagrepo,
		WorkItemType:       workitemtyperepo,
	}
}

// CreateTestRepos creates repositories with convenient wrappers for testing.
func CreateTestRepos() *repos.Repos {
	repos := CreateRepos()

	repos.User = reposwrappers.NewUserWithRetry(repos.User, 5, 65*time.Millisecond) // created_at unique

	return repos
}
