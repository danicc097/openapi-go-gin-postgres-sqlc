package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/format"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"go.uber.org/zap"

	// dot import so go code would resemble as much as native SQL
	// dot import is not mandatory
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/jet/public/model"
	. "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/jet/public/table"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"
)

// clear && go run cmd/cli/main.go -env .env.dev
func main() {
	var env string
	var init bool

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.BoolVar(&init, "init", false, "Run initial data script")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	format.PrintJSON(internal.Config())

	cmd := exec.Command(
		"bash", "-c",
		"source .envrc",
	)
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Print("1")
		errAndExit(out, err)
	}

	if init {
		cmd = exec.Command(
			"bash", "-c",
			"project db.initial-data",
		)
		cmd.Dir = "."
		if _, err := cmd.CombinedOutput(); err != nil {
			fmt.Print("2")
			// errAndExit(out, err) // exit code 1 for some reason
		}
	}

	logger, _ := zap.NewDevelopment()
	pool, sqlpool, err := postgresql.New(logger)
	if err != nil {
		log.Fatalf("postgresql.New: %s\n", err)
	}

	username := "user_1"
	// username := "doesntexist" // User should be nil
	// username := "superadmin"
	user, err := db.UserByUsername(context.Background(), pool, username,
		db.WithUserJoin(db.UserJoins{
			// TODO fix array_agg pgx collect and reenable
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

		UserNotifications.UserID.EQ(UUID(user.UserID)).
			AND(UserNotifications.UserID.EQ(UUID(user.UserID))),
	).ORDER_BY(
		Notifications.CreatedAt.DESC(),
	)
	query, args := getUserNotificationsByUserID.Sql()

	fmt.Printf("query: %v\n", query)
	fmt.Printf("args: %#v\n", args)

	type Res []struct {
		model.UserNotifications

		Notification model.Notifications
	}

	dest := &Res{}

	// won't be able to use same transaction and also need a sql.DB pool apart from pgxpool opened with postgresql.New
	// https://github.com/go-jet/jet/issues/59
	// this will break our repo and service (d db.DBTX) param
	err = getUserNotificationsByUserID.QueryContext(context.Background(), sqlpool, dest)
	if err != nil {
		fmt.Println(err.Error())
	}
	// fmt.Printf("dest: %#v\n", dest)
	// format.PrintJSON(dest)

	q := db.New()
	nn, err := q.GetUserNotifications(context.Background(), pool, db.GetUserNotificationsParams{UserID: user.UserID, Lim: pointers.New[int32](6), NotificationType: db.NotificationTypePersonal})
	if err != nil {
		log.Fatal(err.Error())
	}
	format.PrintJSON(nn)

	rows, _ := pool.Query(context.Background(), fmt.Sprintf(`SELECT user_api_keys.user_api_key_id,
	user_api_keys.api_key,
	user_api_keys.expires_on,
	user_api_keys.user_id,
	row_to_json(users.*) as user
	FROM public.user_api_keys
	left join users on users.user_id = user_api_keys.user_id
	WHERE user_api_keys.user_id = '%s'`, user.UserID)) // select api_key from user_api_keys limit 1;
	uaks, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[db.UserAPIKey])
	if err != nil {
		fmt.Printf("CollectRows error: %v", err)
		return
	}
	b, _ := json.Marshal(uaks[0])
	fmt.Printf("uaks[0]: %+v\n", string(b))
	//
	//
	//
	type AnotherTable struct{}
	type User struct {
		UserID int    `json:"userId" db:"user_id"`
		Name   string `json:"name" db:"name"`
	}
	type UserAPIKey struct {
		UserAPIKeyID int `json:"userApiKeyId" db:"user_api_key_id"`
		UserID       int `json:"userId" db:"user_id"`

		User         *User         `json:"user" db:"user"`
		AnotherTable *AnotherTable `json:"anotherTable" db:"another_table"`
	}
	rows, _ = pool.Query(context.Background(), `
	WITH user_api_keys AS (
		SELECT 1 AS user_id, 101 AS user_api_key_id, 'abc123' AS api_key
	), users AS (
		SELECT 1 AS user_id, 'John Doe' AS name
	)
	SELECT user_api_keys.user_api_key_id, user_api_keys.user_id, row(users.*) AS user
	FROM user_api_keys
	LEFT JOIN users ON users.user_id = user_api_keys.user_id
	WHERE user_api_keys.api_key = 'abc123';
	`)
	uaks_test, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserAPIKey])
	fmt.Printf("err: %v\n", err)
	bt, _ := json.Marshal(uaks_test[0])
	fmt.Printf("uaks_test[0]: %+v\n", string(bt))

	type Item struct {
		UserItemID int    `json:"userItemID" db:"user_item_id"`
		UserID     int    `json:"userID" db:"user_id"`
		Item       string `json:"item" db:"item"`
	}
	type CustomUser struct {
		UserID int     `json:"userID" db:"user_id"`
		Name   string  `json:"name" db:"name"`
		Items  []*Item `json:"items" db:"items"`
	}
	rows, _ = pool.Query(context.Background(), `
	WITH user_items AS (
		SELECT 1 AS user_id, 101 AS user_item_id, 'item 1' AS item
		UNION ALL
		SELECT 1 AS user_id, 102 AS user_item_id, 'item 2' AS item
	), users AS (
		SELECT 1 AS user_id, 'John Doe' AS name
	)
	SELECT users.user_id, array_agg(user_items.*) AS items
	FROM users
	LEFT JOIN user_items ON users.user_id = user_items.user_id
	GROUP BY users.user_id;
	`)
	userItems, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[CustomUser])
	fmt.Printf("err: %v\n", err)
	bt, _ = json.Marshal(userItems[0])
	fmt.Printf("userItems[0]: %+v\n", string(bt))
	// {"userID":1,"name":"","userItems":[{"userItemID":1,"userID":101,"item":"item 1"},{"userItemID":1,"userID":102,"item":"item 2"}]}
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}
