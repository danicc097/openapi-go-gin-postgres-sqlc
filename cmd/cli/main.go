package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/google/uuid"
)

type User struct {
	db.User

	WorkItems   []*db.WorkItem  `json:"work_items,omitempty"`
	Teams       []*db.Team      `json:"teams,omitempty"`
	UserApiKey  *db.UserAPIKey  `json:"user_api_key,omitempty"`
	TimeEntries []*db.TimeEntry `json:"time_entries,omitempty"`

	// xo fields
	_exists, _deleted bool
}

const query = `
		select
	  (case when $1::boolean = true then joined_work_items.work_items end)::jsonb as work_items
	  , (case when $2::boolean = true then joined_teams.teams end)::jsonb as teams
	  , (case when $3::boolean = true then row_to_json(user_api_keys.*) end)::jsonb as user_api_key
	  , (case when $4::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries
	  , users.user_id
	  , users.username
	  , users.role_rank
	  , users.scopes
	from
	  users
	left join (
	  select
	    member as work_items_user_id
	    , json_agg(work_items.*) as work_items
	  from
	    work_item_member uo
	    join work_items using (work_item_id)
	  where
	    member in (
	      select
	        member
	      from
	        work_item_member
	      where
	        work_item_id = any (
	          select
	            work_item_id
	          from
	            work_items))
	      group by
	        member) joined_work_items on joined_work_items.work_items_user_id = users.user_id
	left join (
	  select
	    user_id as teams_user_id
	    , json_agg(teams.*) as teams
	  from
	    user_team uo
	    join teams using (team_id)
	  where
	    user_id in (
	      select
	        user_id
	      from
	        user_team
	      where
	        team_id = any (
	          select
	            team_id
	          from
	            teams))
	      group by
	        user_id) joined_teams on joined_teams.teams_user_id = users.user_id
	left join (
	  select
	  user_id
	    , json_agg(time_entries.*) as time_entries
	  from
	    time_entries
	   group by
	        user_id) joined_time_entries using (user_id)
	left join user_api_keys using (user_id)
	-- where username = $5
	;
		`

func main() {
	var env string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.Parse()

	if err := envvar.Load(env); err != nil {
		log.Fatalf("envvar.Load: %s\n", err)
	}

	conf := envvar.New()
	pool, err := postgresql.New(conf)
	if err != nil {
		log.Fatalf("postgresql.New: %s\n", err)
	}

	// username := "user_1"
	// username := "doesntexist" // User should be nil
	// username := "superadmin"
	joinWorkItems := true
	joinTeams := true
	joinUserApiKeys := true
	joinTimeEntries := false

	fmt.Printf(`
joinWorkItems:= %t
joinTeams:= %t
joinUserApiKeys:= %t
joinTimeEntries:= %t
--------------------------
`, joinWorkItems, joinTeams, joinUserApiKeys, joinTimeEntries)

	// .Query --> Rows --- .QueryRow -> Row
	rows, err := pool.Query(context.Background(), query, joinWorkItems, joinTeams, joinUserApiKeys, joinTimeEntries)
	if err != nil {
		log.Fatalf("pool.Query: %s\n", err)
	}
	defer rows.Close()

	// https://stackoverflow.com/questions/63785376/inserting-empty-string-or-null-into-postgres-as-null-using-jackc-pgx
	// https://rodrigo.red/blog/go-lang-not-so-simple/
	users := make([]User, 0)
	for rows.Next() {
		var u User
		// https://github.com/jackc/pgx/issues/180 cast as jsonb
		err := rows.Scan(&u.WorkItems, &u.Teams, &u.UserApiKey, &u.TimeEntries, &u.UserID, &u.Username, &u.RoleRank, &u.Scopes) // etc.
		if err != nil {
			log.Fatalf("rows.Scan: %s\n", err)
		}

		if u.UserID == uuid.Nil {
			fmt.Println("no row was found")
			return
		}

		// NOTE: Consumer, e.g. frontend will not care the slightest and
		// will simply check if the key exists (openapi fields WorkItems, teams, ... will be nullable)
		// TODO For internal backend use, we should probably have pointers to any Join field
		// in case we don't set the JoinXXX flag and explicitly set to nil if the flag is not set,
		// else we get empty array and mistake it for no values
		// when in reality we forgot to add the flag
		// but this won't work for o2o since zero value of struct is nil as well...

		users = append(users, u)
	}

	format.PrintJSON(users)
}
