package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgtype"
	"github.com/lib/pq"
	"gopkg.in/guregu/null.v4"
)

// Task represents a row from 'public.tasks'.
type Task struct {
	TaskID             int64        `json:"task_id" db:"task_id"`                           // task_id
	TaskTypeID         int          `json:"task_type_id" db:"task_type_id"`                 // task_type_id
	WorkItemID         int64        `json:"work_item_id" db:"work_item_id"`                 // work_item_id
	Title              string       `json:"title" db:"title"`                               // title
	Metadata           pgtype.JSONB `json:"metadata" db:"metadata"`                         // metadata
	TargetDate         time.Time    `json:"target_date" db:"target_date"`                   // target_date
	TargetDateTimezone string       `json:"target_date_timezone" db:"target_date_timezone"` // target_date_timezone
	CreatedAt          time.Time    `json:"created_at" db:"created_at"`                     // created_at
	UpdatedAt          time.Time    `json:"updated_at" db:"updated_at"`                     // updated_at
	DeletedAt          null.Time    `json:"deleted_at" db:"deleted_at"`                     // deleted_at
	// xo fields
	_exists, _deleted bool
}

// Team represents a row from 'public.teams'.
type Team struct {
	TeamID      int          `json:"team_id" db:"team_id"`         // team_id
	ProjectID   int          `json:"project_id" db:"project_id"`   // project_id
	Name        string       `json:"name" db:"name"`               // name
	Description string       `json:"description" db:"description"` // description
	Metadata    pgtype.JSONB `json:"metadata" db:"metadata"`       // metadata
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`   // created_at
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`   // updated_at
	// xo fields
	_exists, _deleted bool
}

// UserAPIKey represents a row from 'public.user_api_keys'.
type UserAPIKey struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id"`       // user_id
	APIKey    string    `json:"api_key" db:"api_key"`       // api_key
	ExpiresOn time.Time `json:"expires_on" db:"expires_on"` // expires_on
	// xo fields
	_exists, _deleted bool
}

// TimeEntry represents a row from 'public.time_entries'.
type TimeEntry struct {
	TimeEntryID     int64     `json:"time_entry_id" db:"time_entry_id"`       // time_entry_id
	TaskID          null.Int  `json:"task_id" db:"task_id"`                   // task_id
	ActivityID      int       `json:"activity_id" db:"activity_id"`           // activity_id
	TeamID          null.Int  `json:"team_id" db:"team_id"`                   // team_id
	UserID          uuid.UUID `json:"user_id" db:"user_id"`                   // user_id
	Comment         string    `json:"comment" db:"comment"`                   // comment
	Start           time.Time `json:"start" db:"start"`                       // start
	DurationMinutes null.Int  `json:"duration_minutes" db:"duration_minutes"` // duration_minutes
	// xo fields
	_exists, _deleted bool
}

type User struct {
	UserID     uuid.UUID      `json:"user_id"`     // user_id
	Username   string         `json:"username"`    // username
	Email      string         `json:"email"`       // email
	Scopes     pq.StringArray `json:"scopes"`      // scopes
	FirstName  null.String    `json:"first_name"`  // first_name
	LastName   null.String    `json:"last_name"`   // last_name
	FullName   null.String    `json:"full_name"`   // full_name
	ExternalID null.String    `json:"external_id"` // external_id
	Role       db.UserRole    `json:"role"`        // role
	CreatedAt  time.Time      `json:"created_at"`  // created_at
	UpdatedAt  time.Time      `json:"updated_at"`  // updated_at
	DeletedAt  null.Time      `json:"deleted_at"`  // deleted_at

	// omitempty to get explicit [] or no field at all later
	Tasks       *[]*Task      `json:"tasks,omitempty"`
	Teams       *[]*Team      `json:"teams,omitempty"`
	UserApiKey  *UserAPIKey   `json:"user_api_key,omitempty"`
	TimeEntries *[]*TimeEntry `json:"time_entries,omitempty"`

	// xo fields
	_exists, _deleted bool
}

const query = `
		select
	  (case when $1::boolean = true then joined_tasks.tasks end)::jsonb as tasks
	  , (case when $2::boolean = true then joined_teams.teams end)::jsonb as teams
	  , (case when $3::boolean = true then row_to_json(user_api_keys.*) end)::jsonb as user_api_key
	  , (case when $4::boolean = true then joined_time_entries.time_entries end)::jsonb as time_entries
	  , users.user_id
	  , users.username
	  , users.role
	from
	  users
	left join (
	  select
	    member as tasks_user_id
	    , json_agg(tasks.*) as tasks
	  from
	    task_member uo
	    join tasks using (task_id)
	  where
	    member in (
	      select
	        member
	      from
	        task_member
	      where
	        task_id = any (
	          select
	            task_id
	          from
	            tasks))
	      group by
	        member) joined_tasks on joined_tasks.tasks_user_id = users.user_id
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
	joinTasks := true
	joinTeams := true
	joinUserApiKeys := true
	joinTimeEntries := false

	fmt.Printf(`
joinTasks:= %t
joinTeams:= %t
joinUserApiKeys:= %t
joinTimeEntries:= %t
--------------------------
`, joinTasks, joinTeams, joinUserApiKeys, joinTimeEntries)

	// .Query --> Rows --- .QueryRow -> Row
	rows, err := pool.Query(context.Background(), query, joinTasks, joinTeams, joinUserApiKeys, joinTimeEntries)
	if err != nil {
		log.Fatalf("pool.Query: %s\n", err)
	}
	defer rows.Close()

	// https://stackoverflow.com/questions/63785376/inserting-empty-string-or-null-into-postgres-as-null-using-jackc-pgx
	// https://rodrigo.red/blog/go-lang-not-so-simple/
	users := make([]User, 0)
	for rows.Next() {
		var tasks []*Task
		var teams []*Team
		var userApiKey *UserAPIKey
		var timeEntries []*TimeEntry
		var u User
		// https://github.com/jackc/pgx/issues/180 cast as jsonb
		err := rows.Scan(&tasks, &teams, &userApiKey, &timeEntries, &u.UserID, &u.Username, &u.Role) // etc.
		if err != nil {
			log.Fatalf("rows.Scan: %s\n", err)
		}

		if u.UserID == uuid.Nil {
			fmt.Println("no row was found")
			return
		}

		fmt.Printf("tasks: %v\n", tasks)
		fmt.Printf("teams: %v\n", teams)
		fmt.Printf("timeEntries: %v\n", timeEntries)

		u.UserApiKey = userApiKey

		if len(tasks) > 0 {
			fmt.Println("len(tasks) > 0")
			u.Tasks = &tasks
		}
		if len(teams) > 0 {
			u.Teams = &teams
		}
		if len(timeEntries) > 0 {
			u.TimeEntries = &timeEntries
		}

		// TODO surely there are better ways to do this
		if joinTasks && u.Tasks == nil {
			fmt.Println("joinTasks && u.Tasks == nil")
			u.Tasks = &[]*Task{}
		}
		if joinTeams && u.Teams == nil {
			fmt.Println("joinTeams && u.Teams == nil")
			u.Teams = &[]*Team{}
		}
		if joinTimeEntries && u.TimeEntries == nil {
			fmt.Println("joinTimeEntries && u.TimeEntries == nil")
			u.TimeEntries = &[]*TimeEntry{}
		}
		users = append(users, u)
	}

	PrintJSON(users)
}

func PrintJSON(obj interface{}) {
	bytes, _ := json.MarshalIndent(obj, "  ", "  ")
	fmt.Println(string(bytes))
}
