package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"go.uber.org/zap"

	// dot import so go code would resemble as much as native SQL
	// dot import is not mandatory

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/jet/public/model"
	. "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/jet/public/table"
	. "github.com/go-jet/jet/v2/postgres"
)

// clear && go run cmd/cli/main.go -env .env.dev
func main() {
	var env string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	appEnv := envvar.GetEnv("APP_ENV", "dev")
	if err := envvar.Load(path.Join(".env." + appEnv)); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	cmd := exec.Command(
		"bash", "-c",
		"source .envrc",
	)
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		errAndExit(out, err)
	}

	// cmd = exec.Command(
	// 	"bash", "-c",
	// 	"project db.initial-data",
	// )
	// cmd.Dir = "."
	// if out, err := cmd.CombinedOutput(); err != nil {
	// 	errAndExit(out, err)
	// }

	logger, _ := zap.NewDevelopment()
	conf := envvar.New()
	pool, err := postgresql.New(conf, logger)
	if err != nil {
		log.Fatalf("postgresql.New: %s\n", err)
	}

	username := "user_2"
	// username := "doesntexist" // User should be nil
	// username := "superadmin"
	user, err := db.UserByUsername(context.Background(), pool, username,
		db.WithUserJoin(db.UserJoins{
			TimeEntries: true,
			WorkItems:   true,
			Teams:       true,
		}),
		db.WithUserOrderBy(db.UserCreatedAtDescNullsLast))
	if err != nil {
		log.Fatalf("db.UserByUsername: %s\n", err)
	}
	format.PrintJSON(user)
	// test correct queries
	key := user.UserID.String() + "-key-hashed"
	uak, err := db.UserAPIKeyByAPIKey(context.Background(), pool, key, db.WithUserAPIKeyJoin(db.UserAPIKeyJoins{User: true}))
	if err != nil {
		log.Fatalf("UserAPIKeyByAPIKey: %v", err)
	}
	fmt.Printf("found user from its api key u: %v#\n", uak.User)

	getUserNotificationsByUserID := SELECT(
		UserNotifications.AllColumns,
		Notifications.AllColumns,
	).FROM(
		UserNotifications.
			INNER_JOIN(Notifications, Notifications.NotificationID.EQ(UserNotifications.NotificationID)),
	).WHERE(
		UserNotifications.UserID.EQ(UUID(user.UserID)),
	).ORDER_BY(
		UserNotifications.CreatedAt.DESC(),
	)
	query, args := getUserNotificationsByUserID.Sql()

	fmt.Println(query) // will print parameterized sql ($1, ...)
	fmt.Println(args)

	dbpool, err := sql.Open("pgx", pool.Config().ConnString())
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to test pool")
	}
	defer dbpool.Close()

	type Res []struct {
		model.UserNotifications

		Notification model.Notifications
	}

	dest := &Res{}

	// won't be able to use same transaction and also need a sql.DB pool apart from pgxpool opened with postgresql.New
	// https://github.com/go-jet/jet/issues/59
	// this will break our repo and service (d db.DBTX) param
	err = getUserNotificationsByUserID.QueryContext(context.Background(), dbpool, dest)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("dest: %#v\n", dest)
	format.PrintJSON(dest)
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}
