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
	"go.uber.org/zap"
)

// clear && go run cmd/initial-data/main.go -env .env.dev
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

	notifrepo := postgresql.NewNotification()
	urepo := postgresql.NewUser()

	authzsvc, err := services.NewAuthorization(logger, scopePolicyPath, rolePolicyPath)
	if err != nil {
		log.Fatalf("NewAuthorization: %s\n", err)
	}

	usvc := services.NewUser(logger, urepo, notifrepo, authzsvc)
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		u, err := usvc.Register(ctx, pool, services.UserRegisterParams{
			Username:   "user_" + strconv.Itoa(i),
			FirstName:  pointers.New("Name " + strconv.Itoa(i)),
			Email:      "user_" + strconv.Itoa(i) + "@mail.com",
			ExternalID: "external_id_user_" + strconv.Itoa(i),
			// TODO default role of User if not set
			Role: models.RoleUser,
			// TODO map of default Scopes for a given role
			Scopes: []models.Scope{models.ScopeUsersRead},
		})
		if err != nil {
			log.Fatalf("Could not register user: %s", err)
		}
		log.Default().Println("Registered", u.Username)
	}
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}
