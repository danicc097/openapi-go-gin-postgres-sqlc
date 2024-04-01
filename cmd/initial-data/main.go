package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var (
	pool *pgxpool.Pool
	l, _ = zap.NewDevelopment()
)

const (
	day   = 24 * time.Hour
	week  = 7 * day
	month = 30 * day
)

func main() {
	var err error
	var env, scopePolicyPath, rolePolicyPath string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&rolePolicyPath, "roles-path", "roles.json", "Roles policy JSON filepath")
	flag.StringVar(&scopePolicyPath, "scopes-path", "scopes.json", "Scopes policy JSON filepath")
	flag.Parse()

	logger := l.Sugar()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}
	pool, _, err = postgresql.New(logger)
	if err != nil {
		log.Fatalf("postgresql.New: %s\n", err)
	}

	repositories := services.CreateRepos()

	// TODO: services.Create(logger, repositories, pool)
	svc := services.New(logger, repositories, pool)

	ctx := context.Background()

	/**
	 *
	 * USERS
	 *
	 **/

	var users []*db.User

	cfg := internal.Config

	// register superAdmin, which is used for internal calls that require a (super)admin caller.
	// e.g. first user registration via auth callback requires an existing admin,
	// which wouldn't be possible without a registered admin beforehand.
	superAdmin, err := svc.User.Register(ctx, pool, services.UserRegisterParams{
		Username:   "superadmin",
		Email:      cfg.SuperAdmin.DefaultEmail,
		ExternalID: "", // will be updated on login
		Role:       models.RoleSuperAdmin,
	})
	handleError(err)
	_, err = svc.Authentication.CreateAPIKeyForUser(ctx, superAdmin)
	handleError(err)

	superAdmin, err = svc.User.ByID(ctx, pool, superAdmin.UserID) // get joins
	handleError(err)

	superAdminCaller := services.CtxUser{
		User:     superAdmin,
		Teams:    *superAdmin.MemberTeamsJoin,
		Projects: *superAdmin.MemberProjectsJoin,
	}

	//
	//
	// PROD guard
	//
	//
	if env == "prod" {
		// prod specific init, if any, and exit early
		fmt.Printf("TODO: create superAdmin only")
		os.Exit(0)
	}

	// TODO: use users which will exist in auth server. that way we can test out these users as well.
	// no need to do it for local.json. as for e2e, we dont want any initial data apart from the superadmin at all
	// so that it mimics real usage from an empty project.
	authServerUsersPath := "cmd/oidc-server/data/users/base.json"
	usersBlob, err := os.ReadFile(authServerUsersPath)
	handleError(err)
	var uu map[string]*models.AuthServerUser
	err = json.Unmarshal(usersBlob, &uu)
	handleError(err)

	logger.Info("Creating users...")
	for i := 0; i < 50; i++ {
		u, err := svc.User.Register(ctx, pool, services.UserRegisterParams{
			Username:   "user_" + strconv.Itoa(i),
			FirstName:  pointers.New(testutil.RandomFirstName()),
			LastName:   pointers.New(testutil.RandomLastName()),
			Email:      "user_" + strconv.Itoa(i) + "@mail.com",
			ExternalID: "external_id_user_" + strconv.Itoa(i),
			// Scopes: []models.Scope{models.}, // TODO:
		})
		handleError(err)
		_, err = svc.Authentication.CreateAPIKeyForUser(ctx, u)
		handleError(err)

		u, err = svc.User.ByID(ctx, pool, u.UserID) // get joins
		handleError(err)

		users = append(users, u)
	}
	u, err := svc.User.Register(ctx, pool, services.UserRegisterParams{
		Username:   "manager_1",
		FirstName:  pointers.New("MrManager"),
		Email:      "manager_1" + "@mail.com",
		ExternalID: "external_id_manager_1",
		Role:       models.RoleManager,
	})
	handleError(err)
	_, err = svc.Authentication.CreateAPIKeyForUser(ctx, u)
	handleError(err)
	users = append(users, u)

	/**
	 *
	 * TEAMS
	 *
	 **/
	logger.Info("Creating teams...")

	teamDemo, err := svc.Team.Create(ctx, pool, &db.TeamCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Team 1",
		Description: "Team 1 description",
	})
	handleError(err, teamDemo)
	teamDemo2, err := svc.Team.Create(ctx, pool, &db.TeamCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemoTwo],
		Name:        "Team 2-1",
		Description: "Team 2-1 description",
	})
	handleError(err, teamDemo2)
	team2Demo2, err := svc.Team.Create(ctx, pool, &db.TeamCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemoTwo],
		Name:        "Team 2-2",
		Description: "Team 2-2 description",
	})
	handleError(err, team2Demo2)

	for i, u := range users {
		users[i], err = svc.User.AssignTeam(ctx, pool, u.UserID, teamDemo.TeamID)
		handleError(err)
		users[i], err = svc.User.AssignTeam(ctx, pool, u.UserID, team2Demo2.TeamID)
		handleError(err)
	}
	// format.PrintJSONByTag(users, "db")

	superAdmin, err = svc.User.AssignTeam(ctx, pool, superAdmin.UserID, teamDemo.TeamID)
	handleError(err)
	superAdmin, err = svc.User.AssignTeam(ctx, pool, superAdmin.UserID, teamDemo2.TeamID)
	handleError(err)

	superAdminCaller = *services.NewCtxUser(superAdmin)

	/**
	 *
	 * ACTIVITIES
	 *
	 **/
	logger.Info("Creating activities...")

	activity1, err := svc.Activity.Create(ctx, pool, &db.ActivityCreateParams{
		ProjectID:    internal.ProjectIDByName[models.ProjectDemo],
		Name:         "Activity 1",
		Description:  "Activity 1 description",
		IsProductive: true,
	})
	handleError(err, activity1)
	activity2, err := svc.Activity.Create(ctx, pool, &db.ActivityCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Activity 2",
		Description: "Activity 2 description",
	})
	handleError(err, activity2)
	activity3, err := svc.Activity.Create(ctx, pool, &db.ActivityCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Activity 3",
		Description: "Activity 3 description",
	})
	handleError(err, activity3)

	/**
	 *
	 * WORK ITEM TAGS
	 *
	 **/
	logger.Info("Creating workitem tags...")
	wiTag1, err := svc.WorkItemTag.Create(ctx, pool, superAdminCaller, &db.WorkItemTagCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Tag 1",
		Description: "Tag 1 description",
		Color:       "#be6cc4",
	})
	handleError(err, wiTag1)

	wiTag2, err := svc.WorkItemTag.Create(ctx, pool, superAdminCaller, &db.WorkItemTagCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Tag 2",
		Description: "Tag 2 description",
		Color:       "#29b8db",
	})
	handleError(err, wiTag2)

	wiTagDemo2_1, err := svc.WorkItemTag.Create(ctx, pool, superAdminCaller, &db.WorkItemTagCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemoTwo],
		Name:        "Tag 1",
		Description: "Tag 1 description",
		Color:       "#be6cc4",
	})
	handleError(err, wiTagDemo2_1)

	/**
	 *
	 * DEMO WORK ITEMS
	 *
	 **/
	logger.Info("Creating workitems...")

	demoWorkItems := []*db.WorkItem{}
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 2000)
	for i := 1; i <= 1000; i++ {
		semaphore <- struct{}{} // acquire
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			demowi, err := svc.DemoWorkItem.Create(ctx, pool, superAdminCaller, services.DemoWorkItemCreateParams{
				DemoWorkItemCreateParams: repos.DemoWorkItemCreateParams{
					Base: db.WorkItemCreateParams{
						TeamID:         teamDemo.TeamID,
						Title:          fmt.Sprintf("A new work item (%d)", i),
						Description:    fmt.Sprintf("Description for a new work item (%d)", i),
						WorkItemTypeID: internal.DemoWorkItemTypesIDByName[models.DemoWorkItemTypesType1],
						// TODO if not passed then query where step order = 0 for a given project and use that
						// steporder could also be generated just like idByName and viceversa
						KanbanStepID: internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsReceived],
						TargetDate:   time.Now().Add(time.Duration(i) * day),
						Metadata:     map[string]any{"key": true},
					},
					DemoProject: db.DemoWorkItemCreateParams{
						LastMessageAt: time.Now().Add(time.Duration(-i) * day),
					},
				},
				WorkItemCreateParams: services.WorkItemCreateParams{
					TagIDs: []db.WorkItemTagID{wiTag1.WorkItemTagID, wiTag2.WorkItemTagID},
					Members: []services.Member{
						{UserID: users[0].UserID, Role: models.WorkItemRolePreparer},
						{UserID: users[1].UserID, Role: models.WorkItemRoleReviewer},
					},
				},
			})
			handleError(err)
			demoWorkItems = append(demoWorkItems, demowi)

			<-semaphore // release
		}(i)
	}

	wg.Wait()

	svc.DemoWorkItem.Update(ctx, pool, superAdminCaller, demoWorkItems[0].WorkItemID, repos.DemoWorkItemUpdateParams{
		Base: &db.WorkItemUpdateParams{
			KanbanStepID: pointers.New(internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsUnderReview]),
		},
	})

	/**
	 *
	 * DEMO TWO WORK ITEMS
	 *
	 **/
	demoTwoWorkItems := []*db.WorkItem{}
	for i := 1; i <= 20; i++ {
		semaphore <- struct{}{} // acquire
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			demoTwowi, err := svc.DemoTwoWorkItem.Create(ctx, pool, superAdminCaller, services.DemoTwoWorkItemCreateParams{
				DemoTwoWorkItemCreateParams: repos.DemoTwoWorkItemCreateParams{
					Base: db.WorkItemCreateParams{
						TeamID:         teamDemo.TeamID,
						Title:          fmt.Sprintf("A new work item (%d)", i),
						Description:    fmt.Sprintf("Description for a new work item (%d)", i),
						WorkItemTypeID: internal.DemoTwoWorkItemTypesIDByName[models.DemoTwoWorkItemTypesType1],
						// TODO if not passed then query where step order = 0 for a given project and use that
						// steporder could also be generated just like idByName and viceversa
						KanbanStepID: internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsReceived],
						TargetDate:   time.Now().Add(time.Duration(i) * day),
						Metadata:     map[string]any{"key": true},
					},
					DemoTwoProject: db.DemoTwoWorkItemCreateParams{
						CustomDateForProject2: pointers.New(time.Now().Add(time.Duration(i) * day)),
					},
				},
				WorkItemCreateParams: services.WorkItemCreateParams{
					TagIDs: []db.WorkItemTagID{wiTag1.WorkItemTagID, wiTag2.WorkItemTagID},
					Members: []services.Member{
						{UserID: users[0].UserID, Role: models.WorkItemRolePreparer},
						{UserID: users[1].UserID, Role: models.WorkItemRoleReviewer},
					},
				},
			})
			handleError(err)
			demoTwoWorkItems = append(demoTwoWorkItems, demoTwowi)

			<-semaphore // release
		}(i)
	}

	wg.Wait()

	/**
	 *
	 * TIME ENTRIES
	 *
	 **/
	logger.Info("Creating time entries...")

	ucaller := services.CtxUser{
		User:     users[0],
		Teams:    *users[0].MemberTeamsJoin,
		Projects: *users[0].MemberProjectsJoin,
	}
	te1, err := svc.TimeEntry.Create(ctx, pool, ucaller, &db.TimeEntryCreateParams{
		WorkItemID:      &demoWorkItems[0].WorkItemID,
		ActivityID:      activity1.ActivityID,
		UserID:          users[0].UserID,
		Comment:         "Doing really important stuff as part of a work item",
		Start:           time.Now(),
		DurationMinutes: pointers.New(56),
	})
	handleError(err, te1)

	te2, err := svc.TimeEntry.Create(ctx, pool, ucaller, &db.TimeEntryCreateParams{
		ActivityID:      activity2.ActivityID,
		UserID:          users[0].UserID,
		TeamID:          &teamDemo.TeamID,
		Comment:         "Doing really important stuff for the team",
		Start:           time.Now(),
		DurationMinutes: pointers.New(26),
	})
	handleError(err, te2)

	for _, u := range users {
		_, err := svc.TimeEntry.Create(ctx, pool, services.CtxUser{User: u, Teams: *u.MemberTeamsJoin}, &db.TimeEntryCreateParams{
			ActivityID: activity2.ActivityID,
			UserID:     u.UserID,
			TeamID:     &teamDemo.TeamID,
			Comment:    "Generic comment (ongoing activity)",
			Start:      time.Now().Add(time.Duration(rand.Intn(120)) * time.Hour),
		})
		handleError(err)
	}

	/**
	 *
	 * NOTIFICATIONS
	 *
	 **/

	for _, u := range users {
		_, err := svc.Notification.CreateNotification(ctx, pool, &services.NotificationCreateParams{
			NotificationCreateParams: db.NotificationCreateParams{
				Body:             "Notification for " + u.Email,
				Labels:           []string{"label 1", "label 2"},
				Link:             pointers.New("https://somelink"),
				Title:            "Important title",
				Sender:           superAdmin.UserID,
				Receiver:         &u.UserID,
				NotificationType: db.NotificationTypePersonal,
			},
		})
		handleError(err)
	}

	_, err = svc.Notification.CreateNotification(ctx, pool, &services.NotificationCreateParams{
		NotificationCreateParams: db.NotificationCreateParams{
			Body:             "Global notification for all users",
			Labels:           []string{"label 4"},
			Link:             pointers.New("https://somelink"),
			Title:            "Important title",
			Sender:           superAdmin.UserID,
			NotificationType: db.NotificationTypeGlobal,
		},
		ReceiverRole: pointers.New(models.RoleUser),
	})
	handleError(err)

	testUser := users[10]
	fmt.Printf("testUser.UserID: %v\n", testUser.UserID)
	err = svc.WorkItem.AssignUsers(ctx, pool, demoWorkItems[0].WorkItemID, []services.Member{{Role: models.WorkItemRolePreparer, UserID: testUser.UserID}})
	handleError(err)
	// TODO: tests later with paginated from cache.<project_name>
	// paginated queries have sortable id. for first query include previous results (-1 or -1 second)
	// and then use returned cursor.
	// wis, err := db.WorkItemPaginatedByWorkItemID(ctx, pool, demoWorkItems[0].WorkItemID-1, models.DirectionAsc, db.WithWorkItemHavingClause(map[string][]any{
	// 	// adding inside where clause yields `aggregate functions are not allowed in WHERE, since it makes no sense.
	// 	//  see https://www.postgresql.org/docs/current/tutorial-agg.html
	// 	"$i = ANY(ARRAY_AGG(xo_join_work_item_assigned_user_assigned_users.__users_user_id))": {testUser.UserID},
	// }), db.WithWorkItemJoin(db.WorkItemJoins{Assignees: true, DemoWorkItem: true}))
	// handleError(err)
	// fmt.Printf("wis len: %v - First workitem found:\n", len(wis))
	// format.PrintJSONByTag(wis[0], "db")
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}

func handleError(err error, info ...any) {
	if err != nil {
		l.Sugar().Fatalf("initial-data error: %s || info: %v", err, info)
	}
}
