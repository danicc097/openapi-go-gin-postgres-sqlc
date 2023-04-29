package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {
	var env, scopePolicyPath, rolePolicyPath string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&rolePolicyPath, "roles-path", "roles.json", "Roles policy JSON filepath")
	flag.StringVar(&scopePolicyPath, "scopes-path", "scopes.json", "Scopes policy JSON filepath")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	logger, _ := zap.NewDevelopment()
	pool, _, err := postgresql.New(logger)
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
	if err != nil {
		log.Fatalf("NewAuthorization: %s\n", err)
	}

	userSvc := services.NewUser(logger, userRepo, notifRepo, authzSvc)
	/*activitySvc :=*/ _ = services.NewActivity(logger, activityRepo)
	/*teamSvc :=*/ _ = services.NewTeam(logger, teamRepo)
	/*projectSvc :=*/ _ = services.NewProject(logger, projectRepo, teamRepo)
	/*teSvc :=*/ _ = services.NewTimeEntry(logger, teRepo)
	/*wiTypeSvc :=*/ _ = services.NewWorkItemType(logger, wiTypeRepo)
	/*wiTagSvc :=*/ _ = services.NewWorkItemTag(logger, wiTagRepo)

	ctx := context.Background()

	registerUsers(userSvc, ctx, pool, logger)
}

func registerUsers(usvc *services.User, ctx context.Context, pool *pgxpool.Pool, logger *zap.Logger) {
	for i := 0; i < 10; i++ {
		u, err := usvc.Register(ctx, pool, services.UserRegisterParams{
			Username:   "user_" + strconv.Itoa(i),
			FirstName:  pointers.New("Name " + strconv.Itoa(i)),
			Email:      "user_" + strconv.Itoa(i) + "@mail.com",
			ExternalID: "external_id_user_" + strconv.Itoa(i),
		})
		if err != nil {
			log.Fatalf("Could not register user: %s", err)
		}
		logger.Sugar().Info("Registered ", u.Username)
	}
	u, err := usvc.Register(ctx, pool, services.UserRegisterParams{
		Username:   "manager_1",
		FirstName:  pointers.New("MrManager"),
		Email:      "manager_1" + "@mail.com",
		ExternalID: "external_id_manager_1",
		Role:       models.RoleManager,
	})
	if err != nil {
		log.Fatalf("Could not register user: %s", err)
	}
	logger.Sugar().Info("Registered ", u.Username)
	u, err = usvc.Register(ctx, pool, services.UserRegisterParams{
		Username:   "superadmin_1",
		FirstName:  pointers.New("MrSuperadmin"),
		Email:      "superadmin_1" + "@mail.com",
		ExternalID: "external_id_superadmin_1",
		Role:       models.RoleSuperAdmin,
	})
	if err != nil {
		log.Fatalf("Could not register user: %s", err)
	}
	logger.Sugar().Info("Registered ", u.Username)
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}
