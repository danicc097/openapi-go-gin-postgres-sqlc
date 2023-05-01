package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/format"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
	"github.com/google/uuid"
	"go.uber.org/zap"

	// dot import so go code would resemble as much as native SQL
	// dot import is not mandatory
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/jet/public/model"
	. "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/jet/public/table"
	. "github.com/go-jet/jet/v2/postgres"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Team struct {
	TeamID int    `json:"teamID" db:"team_id"`
	Name   string `json:"team" db:"team"`
}

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
		fmt.Println(cmd.Stdout)
	}

	logger, _ := zap.NewDevelopment()
	pool, sqlpool, err := postgresql.New(logger)
	if err != nil {
		log.Fatalf("postgresql.New: %s\n", err)
	}

	var rows pgx.Rows

	//
	//
	//
	// pgxArrayAggIssueWorkingQuery(pool)
	// pgxArrayAggIssueQuery(pool)

	username := "user_1"
	// username := "doesntexist" // User should be nil
	// username := "superadmin"
	user, err := db.UserByUsername(context.Background(), pool, username,
		db.WithUserJoin(db.UserJoins{
			TimeEntries: true,
			WorkItems:   true,
			Teams:       true,
			UserAPIKey:  true,
		}),
		db.WithUserOrderBy(db.UserCreatedAtDescNullsLast))
	if err != nil {
		log.Fatalf("db.UserByUsername: %s\n", err)
	}

	// for i := 0; i < 10; i++ {
	// 	fmt.Println("\n-.-.-.-.-.-.-.-.-.-.-.-.-.-.-.-\n")
	// 	_, err = db.UserByUsername(context.Background(), pool, username,
	// 		db.WithUserJoin(db.UserJoins{
	// 			TimeEntries: true,
	// 			WorkItems:   true,
	// 			Teams:       true,
	// 			UserAPIKey:  false,
	// 		}),
	// 		db.WithUserOrderBy(db.UserCreatedAtDescNullsLast))
	// 	if err != nil {
	// 		log.Fatalf("db.UserByUsername: %s\n", err)
	// 	}
	// }

	fmt.Printf("user: %+v\n", user)
	// test correct queries
	key := user.UserAPIKeyJoin.APIKey
	uak, err := db.UserAPIKeyByAPIKey(context.Background(), pool, key, db.WithUserAPIKeyJoin(db.UserAPIKeyJoins{User: true}))
	if err != nil {
		log.Fatalf("UserAPIKeyByAPIKey: %v", err)
	}
	fmt.Printf("found user from its api key u: %v#\n", uak.UserJoin)

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

	rows, _ = pool.Query(context.Background(), fmt.Sprintf(`SELECT user_api_keys.user_api_key_id,
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
	format.PrintJSON(uaks[0])
}

type Item struct {
	UserItemID int    `json:"userItemID" db:"user_item_id"`
	UserID     int    `json:"userID" db:"user_id"`
	Item       string `json:"item" db:"item"`
}

type Team1 struct {
	TeamID      int       `json:"teamID" db:"team_id"`
	ProjectID   int       `json:"projectID" db:"project_id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`

	Users *[]User1 `json:"users" db:"users"`
	// xo fields
	_exists, _deleted bool
}

type User1 struct {
	UserID   uuid.UUID `json:"userID" db:"user_id"`
	Username string    `json:"username" db:"username"`

	Teams *[]Team1 `json:"teams" db:"teams"`
	// xo fields
	_exists, _deleted bool
}

func OIDCheck(conn *pgxpool.Conn) {
	rows, err := conn.Query(context.Background(), "SELECT * FROM pg_class")
	if err != nil {
		log.Fatal(err)
	}

	fields := rows.FieldDescriptions()
	columnNames := make([]string, len(fields))
	for i, fd := range fields {
		columnNames[i] = fd.Name
	}

	for rows.Next() {
		rowValues := make([]any, len(fields))
		for i := range fields {
			rowValues[i] = new(any)
		}
		if err := rows.Scan(rowValues...); err != nil {
			log.Fatal(err)
		}

		for i, col := range rowValues {
			if columnNames[i] != "oid" {
				continue
			}
			if *col.(*any) != 26 {
				continue
			}
			fmt.Println("---------------")
			for i, col := range rowValues {
				fmt.Printf("%s: %v\n", columnNames[i], *col.(*any))
			}
			fmt.Println()
		}
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func pgxArrayAggIssueWorkingQuery(pool *pgxpool.Pool) {
	query := `
	WITH user_team AS (
		SELECT '19270107-1b9c-4f52-a578-7390d5b31513'::uuid AS user_id, 1 AS team_id
		UNION ALL
		SELECT '19270107-1b9c-4f52-a578-7390d5b31513'::uuid AS user_id, 2 AS team_id
	), users AS (
		SELECT '19270107-1b9c-4f52-a578-7390d5b31513'::uuid AS user_id, 'John Doe' AS name
	),teams AS (
		SELECT 1 AS team_id, 1 as project_id, 'team 1' AS name, 'This is team 1 from project 1' as description, now() AS created_at, now() AS updated_at
		UNION ALL
		SELECT 2 AS team_id, 2 as project_id, 'team 2' AS name, 'This is team 2 from project 1' as description, now() AS created_at, now() AS updated_at
	)
	SELECT users.user_id
	, joined_teams.__teams as teams
	FROM users
	left join (
		select
			user_team.user_id as user_team_user_id
			, array_agg(teams.*) as __teams
			from user_team
			join teams using (team_id)
			group by user_team_user_id
		) as joined_teams on joined_teams.user_team_user_id = users.user_id
	`

	rows, err := pool.Query(context.Background(), query)
	if err != nil {
		log.Fatalf("error pool.Query: %s\n", err)
	}

	users, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User1])
	if err != nil {
		log.Fatalf("error pgx.CollectRows: %s\n", err)
	}
	b, _ := json.Marshal(users[0])
	fmt.Printf("users[0]: %+v\n", string(b)) //{"userID":"19270107-1b9c-4f52-a578-7390d5b31513","username":"","teams":[{"teamID":1,"projectID":1,"name":"team 1","description":"This is team 1 from project 1","createdAt":"2023-04-16T08:51:29.108119Z","updatedAt":"2023-04-16T08:51:29.108119Z","users":null},{"teamID":2,"projectID":2,"name":"team 2","description":"This is team 2 from project 1","createdAt":"2023-04-16T08:51:29.108119Z","updatedAt":"2023-04-16T08:51:29.108119Z","users":null}]}
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}

func pgxArrayAggIssueQuery(pool *pgxpool.Pool) {
	conn, err := pool.Acquire(context.Background())
	if err != nil {
		log.Fatalf("error pool.Acquire: %s\n", err)
	}
	_, err = conn.Exec(context.Background(), `
create temporary table projects (
	project_id serial primary key
	, name text not null unique
);

create temporary table teams (
	team_id serial primary key
	, project_id int not null --limited to a project only
	, name text not null
	, description text not null
	, created_at timestamp with time zone default current_timestamp not null
	, updated_at timestamp with time zone default current_timestamp not null
	, foreign key (project_id) references projects (project_id) on delete cascade
	, unique (name , project_id)
);

create temporary table users (
  user_id uuid primary key
  , username text not null unique
);

create temporary table user_team (
  team_id int not null
  , user_id uuid not null
  , primary key (user_id , team_id)
  , foreign key (user_id) references users (user_id) on delete cascade
  , foreign key (team_id) references teams (team_id) on delete cascade
);

INSERT INTO users (user_id , username)
VALUES ('19270107-1b9c-4f52-a578-7390d5b31513' , 'user_1');

INSERT INTO projects ("name" , project_id)
VALUES ('project 1' , 1);

INSERT INTO teams ("name" , project_id , description)
VALUES ('team 1' , 1 , 'This is team 1 from project 1');
INSERT INTO teams ("name" , project_id , description)
VALUES ('team 2' , 1 , 'This is team 2 from project 1');

INSERT INTO user_team (team_id , user_id)
VALUES (1 , '19270107-1b9c-4f52-a578-7390d5b31513');
INSERT INTO user_team (team_id , user_id)
VALUES (2 , '19270107-1b9c-4f52-a578-7390d5b31513');
	`)
	if err != nil {
		log.Fatalf("error conn.Exec: %s\n", err)
	}

	query := `
	SELECT users.user_id
	, joined_teams.__teams as teams
	FROM users
	left join (
		select
			user_team.user_id as user_team_user_id
			, array_agg(teams.*) filter (where teams.* is not null) as __teams
			from user_team
			join teams on teams.team_id = user_team.team_id
			group by user_team_user_id
		) as joined_teams on joined_teams.user_team_user_id = users.user_id
;
	`
	/**
		 * user_id │ 19270107-1b9c-4f52-a578-7390d5b31513
	teams   │ {"(1,1,\"team 1\",\"This is team 1 from project 1\",\"2023-04-16 08:09:13.172744+00\",\"2023-04-16 08:09:13.172744+00\")","(2,1,\"team 2\",\"This is team 2 from project 1\",\"2023-04-16 08:09:13.173884+00\",\"2023-04-16 08:09:13.173884+00\")"}
	*/

	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		log.Fatalf("error conn.Query: %s\n", err)
	}
	users, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User1])
	if err != nil {
		OIDCheck(conn)
		log.Fatalf("error pgx.CollectRows: %s\n", err)
	}
	b, _ := json.Marshal(users[0])
	fmt.Printf("users[0]: %+v\n", string(b))
}
