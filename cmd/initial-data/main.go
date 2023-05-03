package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
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
	projectRepo := postgresql.NewProject()
	teRepo := postgresql.NewTimeEntry()
	wiTypeRepo := postgresql.NewWorkItemType()
	wiTagRepo := postgresql.NewWorkItemTag()

	authzSvc, err := services.NewAuthorization(logger, scopePolicyPath, rolePolicyPath)
	handleError(err)

	userSvc := services.NewUser(logger, userRepo, notifRepo, authzSvc)
	authnSvc := services.NewAuthentication(logger, userSvc, pool)
	/*activitySvc :=*/ _ = services.NewActivity(logger, activityRepo)
	teamSvc := services.NewTeam(logger, teamRepo)
	/*projectSvc :=*/ _ = services.NewProject(logger, projectRepo, teamRepo)
	/*teSvc :=*/ _ = services.NewTimeEntry(logger, teRepo)
	/*wiTypeSvc :=*/ _ = services.NewWorkItemType(logger, wiTypeRepo)
	/*wiTagSvc :=*/ _ = services.NewWorkItemTag(logger, wiTagRepo)

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

	t1, err := teamSvc.Create(ctx, pool, &db.TeamCreateParams{
		ProjectID:   internal.ProjectIDByName[models.ProjectDemo],
		Name:        "Team 1",
		Description: "Team 1 description",
	})
	handleError(err)
	fmt.Printf("t1: %v\n", t1)
	// TODO usvc.AssignTeam(userIDs[0], t1.TeamID)
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
