package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"runtime"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/format"
	"go.uber.org/zap"

	"github.com/jackc/pgx/v5"
)

// Returns the directory of the file this function lives in.
func getFileRuntimeDirectory() string {
	_, b, _, _ := runtime.Caller(0)

	return path.Join(path.Dir(b))
}

// clear && go run cmd/cli/main.go -env .env.dev
func main() {
	var env string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	format.PrintJSON(internal.Config)

	cmd := exec.Command(
		"bash", "-c",
		"source .envrc",
	)
	cmd.Dir = "."
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Print("1")
		errAndExit(out, err)
	}

	logger, _ := zap.NewDevelopment()
	pool, _, err := postgresql.New(logger.Sugar())
	if err != nil {
		log.Fatalf("postgresql.New: %s\n", err)
	}
	type Team struct {
		TeamID int    `json:"teamID" db:"team_id"`
		Name   string `json:"team" db:"team"`
	}
	type User struct {
		UserID  int       `json:"userID" db:"user_id"`
		Name    string    `json:"name" db:"name"`
		Teams   *[]Team   `json:"teams" db:"teams"`
		Strings *[]string `json:"strings" db:"strings"`
	}

	// query, _ := os.ReadFile("cmd/pgx-tests/query.sql")
	query := `
	WITH user_team AS (
		SELECT 1 AS user_id, 1 AS team_id
		UNION ALL
		SELECT 1 AS user_id, 2 AS team_id
	), users AS (
		SELECT 1 AS user_id, 'John Doe' AS name
	),teams AS (
		SELECT 1 AS team_id, 'team 1' AS name
		UNION ALL
		SELECT 2 AS team_id, 'team 2' AS name
	)
	SELECT users.user_id
	,joined_teams.teams as teams
	, string_to_array('a b c', ' ')::text[] as strings
	FROM users
	left join (
		select
			user_team.user_id as teams_user_id
			, array_agg(teams.*) as teams
			from user_team
			join teams using (team_id)
			group by teams_user_id
		) as joined_teams on joined_teams.teams_user_id = users.user_id
	`
	rows, _ := pool.Query(context.Background(), query)
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
