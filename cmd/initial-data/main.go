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
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
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

	notifRepo := postgresql.NewNotification()
	userRepo := postgresql.NewUser()
	activityRepo := postgresql.NewActivity()
	teamRepo := postgresql.NewTeam()
	teRepo := postgresql.NewTimeEntry()
	demoWiRepo := postgresql.NewDemoWorkItem()
	wiRepo := postgresql.NewWorkItem()
	wiTagRepo := postgresql.NewWorkItemTag()

	authzSvc, err := services.NewAuthorization(logger, scopePolicyPath, rolePolicyPath)
	handleError(err)

	userSvc := services.NewUser(logger, userRepo, notifRepo, authzSvc)
	authnSvc := services.NewAuthentication(logger, userSvc, pool)
	activitySvc := services.NewActivity(logger, activityRepo)
	teamSvc := services.NewTeam(logger, teamRepo)
	teSvc := services.NewTimeEntry(logger, teRepo, wiRepo)
	wiSvc := services.NewWorkItem(logger, demoWiRepo, userRepo)
	demoWiSvc := services.NewDemoWorkItem(logger, demoWiRepo, wiRepo, userRepo, wiSvc)
	wiTagSvc := services.NewWorkItemTag(logger, wiTagRepo)

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
	superAdmin, err := userSvc.Register(ctx, pool, services.UserRegisterParams{
		Username:   "superadmin",
		Email:      cfg.SuperAdmin.DefaultEmail,
		ExternalID: "", // will be updated on login
		Role:       models.RoleSuperAdmin,
	})
	handleError(err)
	_, err = authnSvc.CreateAPIKeyForUser(ctx, superAdmin)
	handleError(err)
	logger.Info("Registered ", superAdmin.Username)
	users = append(users, superAdmin)

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

	logger.Info("Registering users...")
	for i := 0; i < 10; i++ {
		u, err := userSvc.Register(ctx, pool, services.UserRegisterParams{
			Username:   "user_" + strconv.Itoa(i),
			FirstName:  pointers.New("Name " + strconv.Itoa(i)),
			Email:      "user_" + strconv.Itoa(i) + "@mail.com",
			ExternalID: "external_id_user_" + strconv.Itoa(i),
		})
		handleError(err)
		_, err = authnSvc.CreateAPIKeyForUser(ctx, u)
		handleError(err)

		logger.Info("Registered ", u.Username)
		users = append(users, u)
	}
	u, err := userSvc.Register(ctx, pool, services.UserRegisterParams{
		Username:   "manager_1",
		FirstName:  pointers.New("MrManager"),
		Email:      "manager_1" + "@mail.com",
		ExternalID: "external_id_manager_1",
		Role:       models.RoleManager,
	})
	handleError(err)
	_, err = authnSvc.CreateAPIKeyForUser(ctx, u)
	handleError(err)
	logger.Info("Registered ", u.Username)
	users = append(users, u)

	/**
	 *
	 * TEAMS
	 *
	 **/

	team1, err := teamSvc.Create(ctx, pool, &db.TeamCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Team 1",
		Description: "Team 1 description",
	})
	handleError(err)

	for _, u := range users {
		err = userSvc.AssignTeam(ctx, pool, u.UserID, team1.TeamID)
		handleError(err)
		// save up some extra calls
		u.MemberTeamsJoin = &[]db.Team{*team1}
	}

	/**
	 *
	 * ACTIVITIES
	 *
	 **/

	activity1, err := activitySvc.Create(ctx, pool, &db.ActivityCreateParams{
		ProjectID:    internal.ProjectIDByName[models.ProjectDemo],
		Name:         "Activity 1",
		Description:  "Activity 1 description",
		IsProductive: true,
	})
	handleError(err)
	logger.Info("Created activity ", activity1.Name)
	activity2, err := activitySvc.Create(ctx, pool, &db.ActivityCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Activity 2",
		Description: "Activity 2 description",
	})
	handleError(err)
	logger.Info("Created activity ", activity2.Name)
	activity3, err := activitySvc.Create(ctx, pool, &db.ActivityCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Activity 3",
		Description: "Activity 3 description",
	})
	handleError(err)
	logger.Info("Created activity ", activity3.Name)

	/**
	 *
	 * WORK ITEM TAGS
	 *
	 **/

	wiTag1, err := wiTagSvc.Create(ctx, pool, superAdmin, &db.WorkItemTagCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Tag 1",
		Description: "Tag 1 description",
		Color:       "#be6cc4",
	})
	handleError(err)
	logger.Info("Created tag ", wiTag1.Name)

	wiTag2, err := wiTagSvc.Create(ctx, pool, superAdmin, &db.WorkItemTagCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Tag 2",
		Description: "Tag 2 description",
		Color:       "#29b8db",
	})
	handleError(err)
	logger.Info("Created tag ", wiTag2.Name)

	/**
	 *
	 * WORK ITEMS
	 *
	 **/

	demoWorkItems := []*db.WorkItem{}
	for i := 1; i <= 20; i++ {
		demowi, err := demoWiSvc.Create(ctx, pool, services.DemoWorkItemCreateParams{
			DemoWorkItemCreateParams: repos.DemoWorkItemCreateParams{
				Base: db.WorkItemCreateParams{
					TeamID:         team1.TeamID,
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
			TagIDs: []int{wiTag1.WorkItemTagID, wiTag2.WorkItemTagID},
			Members: []services.Member{
				{UserID: users[0].UserID, Role: models.WorkItemRolePreparer},
				{UserID: users[1].UserID, Role: models.WorkItemRoleReviewer},
			},
		})
		handleError(err)
		logger.Info("Created work item with title: ", demowi.Title)
		fmt.Printf("demowi.WorkItemAssignedUsersJoin: %+v\n", demowi.WorkItemAssignedUsersJoin)
		demoWorkItems = append(demoWorkItems, demowi)
	}

	demoWiSvc.Update(ctx, pool, demoWorkItems[0].WorkItemID, repos.DemoWorkItemUpdateParams{
		Base: &db.WorkItemUpdateParams{
			KanbanStepID: pointers.New(internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsUnderReview]),
		},
	})

	/**
	 *
	 * TIME ENTRIES
	 *
	 **/

	_, err = teSvc.Create(ctx, pool, users[0], &db.TimeEntryCreateParams{
		WorkItemID:      &demoWorkItems[0].WorkItemID,
		ActivityID:      activity1.ActivityID,
		UserID:          users[0].UserID,
		Comment:         "Doing really important stuff as part of a work item",
		Start:           time.Now(),
		DurationMinutes: pointers.New(56),
	})
	handleError(err)

	_, err = teSvc.Create(ctx, pool, users[0], &db.TimeEntryCreateParams{
		ActivityID:      activity2.ActivityID,
		UserID:          users[0].UserID,
		TeamID:          &team1.TeamID,
		Comment:         "Doing really important stuff for the team",
		Start:           time.Now(),
		DurationMinutes: pointers.New(26),
	})
	handleError(err)

	for _, u := range users {
		_, err := teSvc.Create(ctx, pool, u, &db.TimeEntryCreateParams{
			ActivityID: activity2.ActivityID,
			UserID:     u.UserID,
			TeamID:     &team1.TeamID,
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
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}

func handleError(err error) {
	if err != nil {
		l.Sugar().Fatalf("error: %s", err)
	}
}
