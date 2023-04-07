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
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"go.uber.org/zap"

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
	pool, _, err := postgresql.New(logger)
	if err != nil {
		log.Fatalf("postgresql.New: %s\n", err)
	}
	type Team struct {
		TeamID int    `json:"teamID" db:"team_id"`
		Name   string `json:"team" db:"team"`
	}
	type User struct {
		UserID int     `json:"userID" db:"user_id"`
		Name   string  `json:"name" db:"name"`
		Teams  []*Team `json:"teams" db:"teams"`
	}
	rows, _ := pool.Query(context.Background(), `
WITH user_team AS (
	SELECT 1 AS user_id, 1 AS team_id
	UNION ALL
	SELECT 1 AS user_id, 2 AS team_id
	UNION ALL
	SELECT 999 AS user_id, 1 AS team_id
	UNION ALL
	SELECT 999 AS user_id, 2 AS team_id
), users AS (
	SELECT 1 AS user_id, 'John Doe' AS name
	UNION ALL
	SELECT 999 AS user_id, '999' AS name
),teams AS (
	SELECT 1 AS team_id, 'team 1' AS name
	UNION ALL
	SELECT 2 AS team_id, 'team 2' AS name
)
SELECT users.user_id
,joined_teams.teams as teams
FROM users
left join (
	select
		user_team.user_id as teams_user_id
		, array_agg(teams.*) as teams
		from user_team
    join teams using (team_id)
    group by teams_user_id
  ) as joined_teams on joined_teams.teams_user_id = users.user_id
	`)
	users, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
	fmt.Printf("err: %v\n", err)
	js, _ := json.Marshal(users[0])
	fmt.Printf("user: %+v\n", string(js))
	// {"userID":1,"name":"","teams":[{"teamID":1,"team":"team 1"},{"teamID":2,"team":"team 2"}]}
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}
