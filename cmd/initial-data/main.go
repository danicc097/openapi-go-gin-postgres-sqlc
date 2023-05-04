package main

import (
	"context"
	"flag"
	"fmt"
	"log"
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
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var (
	logger, _ = zap.NewDevelopment()
	pool      *pgxpool.Pool
)

func main() {
	var err error
	var env, scopePolicyPath, rolePolicyPath string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&rolePolicyPath, "roles-path", "roles.json", "Roles policy JSON filepath")
	flag.StringVar(&scopePolicyPath, "scopes-path", "scopes.json", "Scopes policy JSON filepath")
	flag.Parse()

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
	wiTagRepo := postgresql.NewWorkItemTag()

	authzSvc, err := services.NewAuthorization(logger, scopePolicyPath, rolePolicyPath)
	handleError(err)

	userSvc := services.NewUser(logger, userRepo, notifRepo, authzSvc)
	authnSvc := services.NewAuthentication(logger, userSvc, pool)
	activitySvc := services.NewActivity(logger, activityRepo)
	teamSvc := services.NewTeam(logger, teamRepo)
	teSvc := services.NewTimeEntry(logger, teRepo)
	demoWiSvc := services.NewDemoWorkItem(logger, demoWiRepo)
	wiTagSvc := services.NewWorkItemTag(logger, wiTagRepo)

	ctx := context.Background()

	if env == "prod" {
		// prod specific init
		os.Exit(0)
	}

	/**
	 *
	 * USERS
	 *
	 **/

	var userIDs []uuid.UUID

	logger.Sugar().Info("Registering users...")
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

		logger.Sugar().Info("Registered ", u.Username)
		userIDs = append(userIDs, u.UserID)
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
	logger.Sugar().Info("Registered ", u.Username)
	userIDs = append(userIDs, u.UserID)

	u, err = userSvc.Register(ctx, pool, services.UserRegisterParams{
		Username:   "superadmin_1",
		FirstName:  pointers.New("MrSuperadmin"),
		Email:      "superadmin_1" + "@mail.com",
		ExternalID: "external_id_superadmin_1",
		Role:       models.RoleSuperAdmin,
	})
	handleError(err)
	_, err = authnSvc.CreateAPIKeyForUser(ctx, u)
	handleError(err)
	logger.Sugar().Info("Registered ", u.Username)
	userIDs = append(userIDs, u.UserID)

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

	for _, id := range userIDs {
		err = userSvc.AssignTeam(ctx, pool, id, team1.TeamID)
		handleError(err)
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
	logger.Sugar().Info("Created activity ", activity1.Name)
	activity2, err := activitySvc.Create(ctx, pool, &db.ActivityCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Activity 2",
		Description: "Activity 2 description",
	})
	handleError(err)
	logger.Sugar().Info("Created activity ", activity2.Name)
	activity3, err := activitySvc.Create(ctx, pool, &db.ActivityCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Activity 3",
		Description: "Activity 3 description",
	})
	handleError(err)
	logger.Sugar().Info("Created activity ", activity3.Name)

	/**
	 *
	 * WORK ITEM TAGS
	 *
	 **/

	wiTag1, err := wiTagSvc.Create(ctx, pool, &db.WorkItemTagCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Tag 1",
		Description: "Tag 1 description",
		Color:       "#be6cc4",
	})
	handleError(err)
	logger.Sugar().Info("Created tag ", wiTag1.Name)

	wiTag2, err := wiTagSvc.Create(ctx, pool, &db.WorkItemTagCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Tag 2",
		Description: "Tag 2 description",
		Color:       "#29b8db",
	})
	handleError(err)
	logger.Sugar().Info("Created tag ", wiTag2.Name)

	/**
	 *
	 * WORK ITEMS
	 *
	 **/

	demowi1, err := demoWiSvc.Create(ctx, pool, services.DemoWorkItemCreateParams{
		DemoWorkItemCreateParams: repos.DemoWorkItemCreateParams{
			Base: db.WorkItemCreateParams{
				TeamID:         team1.TeamID,
				Title:          "A new work item",
				Description:    "Description for a new work item",
				WorkItemTypeID: internal.DemoWorkItemTypesIDByName[models.DemoWorkItemTypesType1],
				// TODO if not passed then query where step order = 0 for a given project and use that
				// steporder could also be generated just like idByName and viceversa
				KanbanStepID: internal.DemoKanbanStepsIDByName[models.DemoKanbanStepsReceived],
				TargetDate:   time.Now().Add(1 * time.Hour),
				Metadata:     []byte(`{}`),
			},
			DemoProject: db.DemoWorkItemCreateParams{
				LastMessageAt: time.Now().Add(-30 * 24 * time.Hour),
			},
		},
		TagIDs: []int{wiTag1.WorkItemTagID, wiTag2.WorkItemTagID},
		Members: []services.Member{
			{UserID: userIDs[0], Role: models.WorkItemRolePreparer},
			{UserID: userIDs[1], Role: models.WorkItemRoleReviewer},
		},
	})
	handleError(err)
	logger.Sugar().Info("Created work item with title: ", demowi1.Title)

	/**
	 *
	 * TIME ENTRIES
	 *
	 **/

	timeEntry1, err := teSvc.Create(ctx, pool, &db.TimeEntryCreateParams{
		WorkItemID: &demowi1.WorkItemID,
		ActivityID: activity1.ActivityID,
		UserID:     userIDs[0],
		Comment:    "Doing really important stuff as part of a work item",
		Start:      time.Now(),
	})
	handleError(err)
	logger.Sugar().Info("Created time entry: ", timeEntry1.Comment)

	timeEntry2, err := teSvc.Create(ctx, pool, &db.TimeEntryCreateParams{
		ActivityID: activity2.ActivityID,
		UserID:     userIDs[0],
		TeamID:     &team1.TeamID,
		Comment:    "Doing really important stuff for the team",
		Start:      time.Now(),
	})
	handleError(err)
	logger.Sugar().Info("Created time entry: ", timeEntry2.Comment)
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}

func handleError(err error) {
	if err != nil {
		logger.Sugar().Fatalf("error: %s", err)
	}
}
