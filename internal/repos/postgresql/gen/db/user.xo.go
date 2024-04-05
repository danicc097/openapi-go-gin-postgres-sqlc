package db

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	models "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"

	"github.com/google/uuid"
)

// User represents a row from 'public.users'.
// Change properties via SQL column comments, joined with " && ":
//   - "properties":<p1>,<p2>,...
//     -- private: exclude a field from JSON.
//     -- not-required: make a schema field not required.
//     -- hidden: exclude field from OpenAPI generation.
//     -- refs-ignore: generate a field whose constraints are ignored by the referenced table,
//     i.e. no joins will be generated.
//     -- share-ref-constraints: for a FK column, it will generate the same M2O and M2M join fields the ref column has.
//   - "type":<pkg.type> to override the type annotation. An openapi schema named <type> must exist.
//   - "cardinality":<O2O|M2O|M2M> to generate/override joins explicitly. Only O2O is inferred.
//   - "tags":<tags> to append literal struct tag strings.
type User struct {
	UserID                   UserID        `json:"userID" db:"user_id" required:"true" nullable:"false"`                                      // user_id
	Username                 string        `json:"username" db:"username" required:"true" nullable:"false"`                                   // username
	Email                    string        `json:"email" db:"email" required:"true" nullable:"false"`                                         // email
	Age                      *int          `json:"age" db:"age"`                                                                              // age
	FirstName                *string       `json:"firstName" db:"first_name"`                                                                 // first_name
	LastName                 *string       `json:"lastName" db:"last_name"`                                                                   // last_name
	FullName                 *string       `json:"fullName" db:"full_name"`                                                                   // full_name
	ExternalID               string        `json:"-" db:"external_id" nullable:"false"`                                                       // external_id
	APIKeyID                 *UserAPIKeyID `json:"-" db:"api_key_id"`                                                                         // api_key_id
	Scopes                   models.Scopes `json:"scopes" db:"scopes" required:"true" nullable:"false" ref:"#/components/schemas/Scopes"`     // scopes
	RoleRank                 int           `json:"-" db:"role_rank" nullable:"false"`                                                         // role_rank
	HasPersonalNotifications bool          `json:"hasPersonalNotifications" db:"has_personal_notifications" required:"true" nullable:"false"` // has_personal_notifications
	HasGlobalNotifications   bool          `json:"hasGlobalNotifications" db:"has_global_notifications" required:"true" nullable:"false"`     // has_global_notifications
	CreatedAt                time.Time     `json:"createdAt" db:"created_at" required:"true" nullable:"false"`                                // created_at
	UpdatedAt                time.Time     `json:"updatedAt" db:"updated_at" required:"true" nullable:"false"`                                // updated_at
	DeletedAt                *time.Time    `json:"deletedAt" db:"deleted_at"`                                                                 // deleted_at

	ReceiverNotificationsJoin *[]Notification       `json:"-" db:"notifications_receiver" openapi-go:"ignore"`        // M2O users
	SenderNotificationsJoin   *[]Notification       `json:"-" db:"notifications_sender" openapi-go:"ignore"`          // M2O users
	TimeEntriesJoin           *[]TimeEntry          `json:"-" db:"time_entries" openapi-go:"ignore"`                  // M2O users
	UserNotificationsJoin     *[]UserNotification   `json:"-" db:"user_notifications" openapi-go:"ignore"`            // M2O users
	MemberProjectsJoin        *[]Project            `json:"-" db:"user_project_projects" openapi-go:"ignore"`         // M2M user_project
	MemberTeamsJoin           *[]Team               `json:"-" db:"user_team_teams" openapi-go:"ignore"`               // M2M user_team
	UserAPIKeyJoin            *UserAPIKey           `json:"-" db:"user_api_key_api_key_id" openapi-go:"ignore"`       // O2O user_api_keys (inferred)
	AssigneeWorkItemsJoin     *[]UserM2MWorkItemWIA `json:"-" db:"work_item_assignee_work_items" openapi-go:"ignore"` // M2M work_item_assignee
	WorkItemCommentsJoin      *[]WorkItemComment    `json:"-" db:"work_item_comments" openapi-go:"ignore"`            // M2O users
}

// UserCreateParams represents insert params for 'public.users'.
type UserCreateParams struct {
	Age                      *int          `json:"age"`                                                                       // age
	APIKeyID                 *UserAPIKeyID `json:"-"`                                                                         // api_key_id
	Email                    string        `json:"email" required:"true" nullable:"false"`                                    // email
	ExternalID               string        `json:"-" nullable:"false"`                                                        // external_id
	FirstName                *string       `json:"firstName"`                                                                 // first_name
	HasGlobalNotifications   bool          `json:"hasGlobalNotifications" required:"true" nullable:"false"`                   // has_global_notifications
	HasPersonalNotifications bool          `json:"hasPersonalNotifications" required:"true" nullable:"false"`                 // has_personal_notifications
	LastName                 *string       `json:"lastName"`                                                                  // last_name
	RoleRank                 int           `json:"-" nullable:"false"`                                                        // role_rank
	Scopes                   models.Scopes `json:"scopes" required:"true" nullable:"false" ref:"#/components/schemas/Scopes"` // scopes
	Username                 string        `json:"username" required:"true" nullable:"false"`                                 // username
}

// UserParams represents common params for both insert and update of 'public.users'.
type UserParams interface {
	GetAge() *int
	GetAPIKeyID() *UserAPIKeyID
	GetEmail() *string
	GetExternalID() *string
	GetFirstName() *string
	GetHasGlobalNotifications() *bool
	GetHasPersonalNotifications() *bool
	GetLastName() *string
	GetRoleRank() *int
	GetScopes() *models.Scopes
	GetUsername() *string
}

func (p UserCreateParams) GetAge() *int {
	return p.Age
}

func (p UserUpdateParams) GetAge() *int {
	if p.Age != nil {
		return *p.Age
	}
	return nil
}

func (p UserCreateParams) GetAPIKeyID() *UserAPIKeyID {
	return p.APIKeyID
}

func (p UserUpdateParams) GetAPIKeyID() *UserAPIKeyID {
	if p.APIKeyID != nil {
		return *p.APIKeyID
	}
	return nil
}

func (p UserCreateParams) GetEmail() *string {
	x := p.Email
	return &x
}

func (p UserUpdateParams) GetEmail() *string {
	return p.Email
}

func (p UserCreateParams) GetExternalID() *string {
	x := p.ExternalID
	return &x
}

func (p UserUpdateParams) GetExternalID() *string {
	return p.ExternalID
}

func (p UserCreateParams) GetFirstName() *string {
	return p.FirstName
}

func (p UserUpdateParams) GetFirstName() *string {
	if p.FirstName != nil {
		return *p.FirstName
	}
	return nil
}

func (p UserCreateParams) GetHasGlobalNotifications() *bool {
	x := p.HasGlobalNotifications
	return &x
}

func (p UserUpdateParams) GetHasGlobalNotifications() *bool {
	return p.HasGlobalNotifications
}

func (p UserCreateParams) GetHasPersonalNotifications() *bool {
	x := p.HasPersonalNotifications
	return &x
}

func (p UserUpdateParams) GetHasPersonalNotifications() *bool {
	return p.HasPersonalNotifications
}

func (p UserCreateParams) GetLastName() *string {
	return p.LastName
}

func (p UserUpdateParams) GetLastName() *string {
	if p.LastName != nil {
		return *p.LastName
	}
	return nil
}

func (p UserCreateParams) GetRoleRank() *int {
	x := p.RoleRank
	return &x
}

func (p UserUpdateParams) GetRoleRank() *int {
	return p.RoleRank
}

func (p UserCreateParams) GetScopes() *models.Scopes {
	x := p.Scopes
	return &x
}

func (p UserUpdateParams) GetScopes() *models.Scopes {
	return p.Scopes
}

func (p UserCreateParams) GetUsername() *string {
	x := p.Username
	return &x
}

func (p UserUpdateParams) GetUsername() *string {
	return p.Username
}

type UserID struct {
	uuid.UUID
}

func NewUserID(id uuid.UUID) UserID {
	return UserID{
		UUID: id,
	}
}

// CreateUser creates a new User in the database with the given params.
func CreateUser(ctx context.Context, db DB, params *UserCreateParams) (*User, error) {
	u := &User{
		Age:                      params.Age,
		APIKeyID:                 params.APIKeyID,
		Email:                    params.Email,
		ExternalID:               params.ExternalID,
		FirstName:                params.FirstName,
		HasGlobalNotifications:   params.HasGlobalNotifications,
		HasPersonalNotifications: params.HasPersonalNotifications,
		LastName:                 params.LastName,
		RoleRank:                 params.RoleRank,
		Scopes:                   params.Scopes,
		Username:                 params.Username,
	}

	return u.Insert(ctx, db)
}

type UserSelectConfig struct {
	limit   string
	orderBy map[string]models.Direction
	joins   UserJoins
	filters map[string][]any
	having  map[string][]any

	deletedAt string
}
type UserSelectConfigOption func(*UserSelectConfig)

// WithUserLimit limits row selection.
func WithUserLimit(limit int) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

// WithDeletedUserOnly limits result to records marked as deleted.
func WithDeletedUserOnly() UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.deletedAt = " not null "
	}
}

type UserOrderBy string

const (
	UserCreatedAtDescNullsFirst UserOrderBy = " created_at DESC NULLS FIRST "
	UserCreatedAtDescNullsLast  UserOrderBy = " created_at DESC NULLS LAST "
	UserCreatedAtAscNullsFirst  UserOrderBy = " created_at ASC NULLS FIRST "
	UserCreatedAtAscNullsLast   UserOrderBy = " created_at ASC NULLS LAST "
	UserDeletedAtDescNullsFirst UserOrderBy = " deleted_at DESC NULLS FIRST "
	UserDeletedAtDescNullsLast  UserOrderBy = " deleted_at DESC NULLS LAST "
	UserDeletedAtAscNullsFirst  UserOrderBy = " deleted_at ASC NULLS FIRST "
	UserDeletedAtAscNullsLast   UserOrderBy = " deleted_at ASC NULLS LAST "
	UserUpdatedAtDescNullsFirst UserOrderBy = " updated_at DESC NULLS FIRST "
	UserUpdatedAtDescNullsLast  UserOrderBy = " updated_at DESC NULLS LAST "
	UserUpdatedAtAscNullsFirst  UserOrderBy = " updated_at ASC NULLS FIRST "
	UserUpdatedAtAscNullsLast   UserOrderBy = " updated_at ASC NULLS LAST "
)

// WithUserOrderBy accumulates orders results by the given columns.
// A nil entry removes the existing column sort, if any.
func WithUserOrderBy(rows map[string]*models.Direction) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		te := EntityFilters[TableEntityUser]
		for dbcol, dir := range rows {
			if _, ok := te[dbcol]; !ok {
				continue
			}
			if dir == nil {
				delete(s.orderBy, dbcol)
				continue
			}
			s.orderBy[dbcol] = *dir
		}
	}
}

type UserJoins struct {
	ReceiverNotifications bool `json:"receiverNotifications" required:"true" nullable:"false"` // M2O notifications
	SenderNotifications   bool `json:"senderNotifications" required:"true" nullable:"false"`   // M2O notifications
	TimeEntries           bool `json:"timeEntries" required:"true" nullable:"false"`           // M2O time_entries
	UserNotifications     bool `json:"userNotifications" required:"true" nullable:"false"`     // M2O user_notifications
	MemberProjects        bool `json:"memberProjects" required:"true" nullable:"false"`        // M2M user_project
	MemberTeams           bool `json:"memberTeams" required:"true" nullable:"false"`           // M2M user_team
	UserAPIKey            bool `json:"userAPIKey" required:"true" nullable:"false"`            // O2O user_api_keys
	AssigneeWorkItems     bool `json:"assigneeWorkItems" required:"true" nullable:"false"`     // M2M work_item_assignee
	WorkItemComments      bool `json:"workItemComments" required:"true" nullable:"false"`      // M2O work_item_comments
}

// WithUserJoin joins with the given tables.
func WithUserJoin(joins UserJoins) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.joins = UserJoins{
			ReceiverNotifications: s.joins.ReceiverNotifications || joins.ReceiverNotifications,
			SenderNotifications:   s.joins.SenderNotifications || joins.SenderNotifications,
			TimeEntries:           s.joins.TimeEntries || joins.TimeEntries,
			UserNotifications:     s.joins.UserNotifications || joins.UserNotifications,
			MemberProjects:        s.joins.MemberProjects || joins.MemberProjects,
			MemberTeams:           s.joins.MemberTeams || joins.MemberTeams,
			UserAPIKey:            s.joins.UserAPIKey || joins.UserAPIKey,
			AssigneeWorkItems:     s.joins.AssigneeWorkItems || joins.AssigneeWorkItems,
			WorkItemComments:      s.joins.WorkItemComments || joins.WorkItemComments,
		}
	}
}

// UserM2MWorkItemWIA represents a M2M join against "public.work_item_assignee"
type UserM2MWorkItemWIA struct {
	WorkItem WorkItem            `json:"workItem" db:"work_items" required:"true"`
	Role     models.WorkItemRole `json:"role" db:"role" required:"true" ref:"#/components/schemas/WorkItemRole" `
}

// WithUserFilters adds the given WHERE clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	filters := map[string][]any{
//		"NOT (col.name = any ($i))": {[]string{"excl_name_1", "excl_name_2"}},
//		`(col.created_at > $i OR
//		col.is_closed = $i)`: {time.Now().Add(-24 * time.Hour), true},
//	}
func WithUserFilters(filters map[string][]any) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.filters = filters
	}
}

// WithUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
// WithUserHavingClause adds the given HAVING clause conditions, which can be dynamically parameterized
// with $i to prevent SQL injection.
// Example:
//
//	// filter a given aggregate of assigned users to return results where at least one of them has id of userId.
//	// See xo_join_* alias used by the join db tag in the SelectSQL string.
//	filters := map[string][]any{
//	"$i = ANY(ARRAY_AGG(xo_join_assigned_users_join.user_id))": {userId},
//	}
func WithUserHavingClause(conditions map[string][]any) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.having = conditions
	}
}

const userTableReceiverNotificationsJoinSQL = `-- M2O join generated from "notifications_receiver_fkey"
left join (
  select
  receiver as notifications_user_id
    , row(notifications.*) as __notifications
  from
    notifications
  group by
	  notifications_user_id, notifications.notification_id
) as xo_join_notifications_receiver on xo_join_notifications_receiver.notifications_user_id = users.user_id
`

const userTableReceiverNotificationsSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_notifications_receiver.__notifications)) filter (where xo_join_notifications_receiver.notifications_user_id is not null), '{}') as notifications_receiver`

const userTableReceiverNotificationsGroupBySQL = `users.user_id`

const userTableSenderNotificationsJoinSQL = `-- M2O join generated from "notifications_sender_fkey"
left join (
  select
  sender as notifications_user_id
    , row(notifications.*) as __notifications
  from
    notifications
  group by
	  notifications_user_id, notifications.notification_id
) as xo_join_notifications_sender on xo_join_notifications_sender.notifications_user_id = users.user_id
`

const userTableSenderNotificationsSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_notifications_sender.__notifications)) filter (where xo_join_notifications_sender.notifications_user_id is not null), '{}') as notifications_sender`

const userTableSenderNotificationsGroupBySQL = `users.user_id`

const userTableTimeEntriesJoinSQL = `-- M2O join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , row(time_entries.*) as __time_entries
  from
    time_entries
  group by
	  time_entries_user_id, time_entries.time_entry_id
) as xo_join_time_entries on xo_join_time_entries.time_entries_user_id = users.user_id
`

const userTableTimeEntriesSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_time_entries.__time_entries)) filter (where xo_join_time_entries.time_entries_user_id is not null), '{}') as time_entries`

const userTableTimeEntriesGroupBySQL = `users.user_id`

const userTableUserNotificationsJoinSQL = `-- M2O join generated from "user_notifications_user_id_fkey"
left join (
  select
  user_id as user_notifications_user_id
    , row(user_notifications.*) as __user_notifications
  from
    user_notifications
  group by
	  user_notifications_user_id, user_notifications.user_notification_id
) as xo_join_user_notifications on xo_join_user_notifications.user_notifications_user_id = users.user_id
`

const userTableUserNotificationsSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_user_notifications.__user_notifications)) filter (where xo_join_user_notifications.user_notifications_user_id is not null), '{}') as user_notifications`

const userTableUserNotificationsGroupBySQL = `users.user_id`

const userTableMemberProjectsJoinSQL = `-- M2M join generated from "user_project_project_id_fkey"
left join (
	select
		user_project.member as user_project_member
		, projects.project_id as __projects_project_id
		, row(projects.*) as __projects
	from
		user_project
	join projects on projects.project_id = user_project.project_id
	group by
		user_project_member
		, projects.project_id
) as xo_join_user_project_projects on xo_join_user_project_projects.user_project_member = users.user_id
`

const userTableMemberProjectsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_user_project_projects.__projects
		)) filter (where xo_join_user_project_projects.__projects_project_id is not null), '{}') as user_project_projects`

const userTableMemberProjectsGroupBySQL = `users.user_id, users.user_id`

const userTableMemberTeamsJoinSQL = `-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_team.member as user_team_member
		, teams.team_id as __teams_team_id
		, row(teams.*) as __teams
	from
		user_team
	join teams on teams.team_id = user_team.team_id
	group by
		user_team_member
		, teams.team_id
) as xo_join_user_team_teams on xo_join_user_team_teams.user_team_member = users.user_id
`

const userTableMemberTeamsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_user_team_teams.__teams
		)) filter (where xo_join_user_team_teams.__teams_team_id is not null), '{}') as user_team_teams`

const userTableMemberTeamsGroupBySQL = `users.user_id, users.user_id`

const userTableUserAPIKeyJoinSQL = `-- O2O join generated from "users_api_key_id_fkey (inferred)"
left join user_api_keys as _users_api_key_id on _users_api_key_id.user_api_key_id = users.api_key_id
`

const userTableUserAPIKeySelectSQL = `(case when _users_api_key_id.user_api_key_id is not null then row(_users_api_key_id.*) end) as user_api_key_api_key_id`

const userTableUserAPIKeyGroupBySQL = `_users_api_key_id.user_api_key_id,
      _users_api_key_id.user_api_key_id,
	users.user_id`

const userTableAssigneeWorkItemsJoinSQL = `-- M2M join generated from "work_item_assignee_work_item_id_fkey"
left join (
	select
		work_item_assignee.assignee as work_item_assignee_assignee
		, work_item_assignee.role as role
		, work_items.work_item_id as __work_items_work_item_id
		, row(work_items.*) as __work_items
	from
		work_item_assignee
	join work_items on work_items.work_item_id = work_item_assignee.work_item_id
	group by
		work_item_assignee_assignee
		, work_items.work_item_id
		, role
) as xo_join_work_item_assignee_work_items on xo_join_work_item_assignee_work_items.work_item_assignee_assignee = users.user_id
`

const userTableAssigneeWorkItemsSelectSQL = `COALESCE(
		ARRAY_AGG( DISTINCT (
		xo_join_work_item_assignee_work_items.__work_items
		, xo_join_work_item_assignee_work_items.role
		)) filter (where xo_join_work_item_assignee_work_items.__work_items_work_item_id is not null), '{}') as work_item_assignee_work_items`

const userTableAssigneeWorkItemsGroupBySQL = `users.user_id, users.user_id`

const userTableWorkItemCommentsJoinSQL = `-- M2O join generated from "work_item_comments_user_id_fkey"
left join (
  select
  user_id as work_item_comments_user_id
    , row(work_item_comments.*) as __work_item_comments
  from
    work_item_comments
  group by
	  work_item_comments_user_id, work_item_comments.work_item_comment_id
) as xo_join_work_item_comments on xo_join_work_item_comments.work_item_comments_user_id = users.user_id
`

const userTableWorkItemCommentsSelectSQL = `COALESCE(ARRAY_AGG( DISTINCT (xo_join_work_item_comments.__work_item_comments)) filter (where xo_join_work_item_comments.work_item_comments_user_id is not null), '{}') as work_item_comments`

const userTableWorkItemCommentsGroupBySQL = `users.user_id`

// UserUpdateParams represents update params for 'public.users'.
type UserUpdateParams struct {
	Age                      **int          `json:"age"`                                                       // age
	APIKeyID                 **UserAPIKeyID `json:"-"`                                                         // api_key_id
	Email                    *string        `json:"email" nullable:"false"`                                    // email
	ExternalID               *string        `json:"-" nullable:"false"`                                        // external_id
	FirstName                **string       `json:"firstName"`                                                 // first_name
	HasGlobalNotifications   *bool          `json:"hasGlobalNotifications" nullable:"false"`                   // has_global_notifications
	HasPersonalNotifications *bool          `json:"hasPersonalNotifications" nullable:"false"`                 // has_personal_notifications
	LastName                 **string       `json:"lastName"`                                                  // last_name
	RoleRank                 *int           `json:"-" nullable:"false"`                                        // role_rank
	Scopes                   *models.Scopes `json:"scopes" nullable:"false" ref:"#/components/schemas/Scopes"` // scopes
	Username                 *string        `json:"username" nullable:"false"`                                 // username
}

// SetUpdateParams updates public.users struct fields with the specified params.
func (u *User) SetUpdateParams(params *UserUpdateParams) {
	if params.Age != nil {
		u.Age = *params.Age
	}
	if params.APIKeyID != nil {
		u.APIKeyID = *params.APIKeyID
	}
	if params.Email != nil {
		u.Email = *params.Email
	}
	if params.ExternalID != nil {
		u.ExternalID = *params.ExternalID
	}
	if params.FirstName != nil {
		u.FirstName = *params.FirstName
	}
	if params.HasGlobalNotifications != nil {
		u.HasGlobalNotifications = *params.HasGlobalNotifications
	}
	if params.HasPersonalNotifications != nil {
		u.HasPersonalNotifications = *params.HasPersonalNotifications
	}
	if params.LastName != nil {
		u.LastName = *params.LastName
	}
	if params.RoleRank != nil {
		u.RoleRank = *params.RoleRank
	}
	if params.Scopes != nil {
		u.Scopes = *params.Scopes
	}
	if params.Username != nil {
		u.Username = *params.Username
	}
}

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db DB) (*User, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.users (
	age, api_key_id, deleted_at, email, external_id, first_name, has_global_notifications, has_personal_notifications, last_name, role_rank, scopes, username
	) VALUES (
	$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12
	) RETURNING * `
	// run
	logf(sqlstr, u.Age, u.APIKeyID, u.DeletedAt, u.Email, u.ExternalID, u.FirstName, u.HasGlobalNotifications, u.HasPersonalNotifications, u.LastName, u.RoleRank, u.Scopes, u.Username)

	rows, err := db.Query(ctx, sqlstr, u.Age, u.APIKeyID, u.DeletedAt, u.Email, u.ExternalID, u.FirstName, u.HasGlobalNotifications, u.HasPersonalNotifications, u.LastName, u.RoleRank, u.Scopes, u.Username)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Insert/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	newu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Insert/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	*u = newu

	return u, nil
}

// Update updates a User in the database.
func (u *User) Update(ctx context.Context, db DB) (*User, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.users SET 
	age = $1, api_key_id = $2, deleted_at = $3, email = $4, external_id = $5, first_name = $6, has_global_notifications = $7, has_personal_notifications = $8, last_name = $9, role_rank = $10, scopes = $11, username = $12 
	WHERE user_id = $13 
	RETURNING * `
	// run
	logf(sqlstr, u.Age, u.APIKeyID, u.DeletedAt, u.Email, u.ExternalID, u.FirstName, u.HasGlobalNotifications, u.HasPersonalNotifications, u.LastName, u.RoleRank, u.Scopes, u.Username, u.UserID)

	rows, err := db.Query(ctx, sqlstr, u.Age, u.APIKeyID, u.DeletedAt, u.Email, u.ExternalID, u.FirstName, u.HasGlobalNotifications, u.HasPersonalNotifications, u.LastName, u.RoleRank, u.Scopes, u.Username, u.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Update/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	newu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Update/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}
	*u = newu

	return u, nil
}

// Upsert upserts a User in the database.
// Requires appropriate PK(s) to be set beforehand.
func (u *User) Upsert(ctx context.Context, db DB, params *UserCreateParams) (*User, error) {
	var err error

	u.Age = params.Age
	u.APIKeyID = params.APIKeyID
	u.Email = params.Email
	u.ExternalID = params.ExternalID
	u.FirstName = params.FirstName
	u.HasGlobalNotifications = params.HasGlobalNotifications
	u.HasPersonalNotifications = params.HasPersonalNotifications
	u.LastName = params.LastName
	u.RoleRank = params.RoleRank
	u.Scopes = params.Scopes
	u.Username = params.Username

	u, err = u.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", &XoError{Entity: "User", Err: err})
			}
			u, err = u.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", &XoError{Entity: "User", Err: err})
			}
		}
	}

	return u, err
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.users 
	WHERE user_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, u.UserID); err != nil {
		return logerror(err)
	}
	return nil
}

// SoftDelete soft deletes the User from the database via 'deleted_at'.
func (u *User) SoftDelete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `UPDATE public.users 
	SET deleted_at = NOW() 
	WHERE user_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, u.UserID); err != nil {
		return logerror(err)
	}
	// set deleted
	u.DeletedAt = newPointer(time.Now())

	return nil
}

// Restore restores a soft deleted User from the database.
func (u *User) Restore(ctx context.Context, db DB) (*User, error) {
	u.DeletedAt = nil
	newu, err := u.Update(ctx, db)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Restore/pgx.CollectRows: %w", &XoError{Entity: "User", Err: err}))
	}
	return newu, nil
}

// UserPaginatedByCreatedAt returns a cursor-paginated list of User.
func UserPaginatedByCreatedAt(ctx context.Context, db DB, createdAt time.Time, direction models.Direction, opts ...UserSelectConfigOption) ([]User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.ReceiverNotifications {
		selectClauses = append(selectClauses, userTableReceiverNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableReceiverNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableReceiverNotificationsGroupBySQL)
	}

	if c.joins.SenderNotifications {
		selectClauses = append(selectClauses, userTableSenderNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableSenderNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableSenderNotificationsGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, userTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, userTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, userTableTimeEntriesGroupBySQL)
	}

	if c.joins.UserNotifications {
		selectClauses = append(selectClauses, userTableUserNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableUserNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserNotificationsGroupBySQL)
	}

	if c.joins.MemberProjects {
		selectClauses = append(selectClauses, userTableMemberProjectsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberProjectsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberProjectsGroupBySQL)
	}

	if c.joins.MemberTeams {
		selectClauses = append(selectClauses, userTableMemberTeamsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberTeamsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberTeamsGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.AssigneeWorkItems {
		selectClauses = append(selectClauses, userTableAssigneeWorkItemsSelectSQL)
		joinClauses = append(joinClauses, userTableAssigneeWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, userTableAssigneeWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, userTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemCommentsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	operator := "<"
	if direction == models.DirectionAsc {
		operator = ">"
	}

	sqlstr := fmt.Sprintf(`SELECT 
	users.age,
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.email,
	users.external_id,
	users.first_name,
	users.full_name,
	users.has_global_notifications,
	users.has_personal_notifications,
	users.last_name,
	users.role_rank,
	users.scopes,
	users.updated_at,
	users.user_id,
	users.username %s 
	 FROM public.users %s 
	 WHERE users.created_at %s $1
	 %s   AND users.deleted_at is %s  %s 
  %s 
  ORDER BY 
		created_at %s `, selects, joins, operator, filters, c.deletedAt, groupbys, havingClause, direction)
	sqlstr += c.limit
	sqlstr = "/* UserPaginatedByCreatedAt */\n" + sqlstr

	// run

	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Paginated/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Paginated/pgx.CollectRows: %w", &XoError{Entity: "User", Err: err}))
	}
	return res, nil
}

// UserByCreatedAt retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_created_at_key'.
func UserByCreatedAt(ctx context.Context, db DB, createdAt time.Time, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.ReceiverNotifications {
		selectClauses = append(selectClauses, userTableReceiverNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableReceiverNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableReceiverNotificationsGroupBySQL)
	}

	if c.joins.SenderNotifications {
		selectClauses = append(selectClauses, userTableSenderNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableSenderNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableSenderNotificationsGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, userTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, userTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, userTableTimeEntriesGroupBySQL)
	}

	if c.joins.UserNotifications {
		selectClauses = append(selectClauses, userTableUserNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableUserNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserNotificationsGroupBySQL)
	}

	if c.joins.MemberProjects {
		selectClauses = append(selectClauses, userTableMemberProjectsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberProjectsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberProjectsGroupBySQL)
	}

	if c.joins.MemberTeams {
		selectClauses = append(selectClauses, userTableMemberTeamsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberTeamsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberTeamsGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.AssigneeWorkItems {
		selectClauses = append(selectClauses, userTableAssigneeWorkItemsSelectSQL)
		joinClauses = append(joinClauses, userTableAssigneeWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, userTableAssigneeWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, userTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemCommentsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	users.age,
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.email,
	users.external_id,
	users.first_name,
	users.full_name,
	users.has_global_notifications,
	users.has_personal_notifications,
	users.last_name,
	users.role_rank,
	users.scopes,
	users.updated_at,
	users.user_id,
	users.username %s 
	 FROM public.users %s 
	 WHERE users.created_at = $1
	 %s   AND users.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	orderBy := ""
	if len(c.orderBy) > 0 {
		orderBy += " order by "
	}
	i := 0
	orderBys := make([]string, len(c.orderBy))
	for dbcol, dir := range c.orderBy {
		orderBys[i] = dbcol + " " + string(dir)
		i++
	}
	orderBy += " " + strings.Join(orderBys, ", ") + " "
	sqlstr += orderBy
	sqlstr += c.limit
	sqlstr = "/* UserByCreatedAt */\n" + sqlstr

	// run
	// logf(sqlstr, createdAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{createdAt}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByCreatedAt/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &u, nil
}

// UsersByDeletedAt_WhereDeletedAtIsNotNull retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_deleted_at_idx'.
func UsersByDeletedAt_WhereDeletedAtIsNotNull(ctx context.Context, db DB, deletedAt *time.Time, opts ...UserSelectConfigOption) ([]User, error) {
	c := &UserSelectConfig{deletedAt: " not null ", joins: UserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.ReceiverNotifications {
		selectClauses = append(selectClauses, userTableReceiverNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableReceiverNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableReceiverNotificationsGroupBySQL)
	}

	if c.joins.SenderNotifications {
		selectClauses = append(selectClauses, userTableSenderNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableSenderNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableSenderNotificationsGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, userTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, userTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, userTableTimeEntriesGroupBySQL)
	}

	if c.joins.UserNotifications {
		selectClauses = append(selectClauses, userTableUserNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableUserNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserNotificationsGroupBySQL)
	}

	if c.joins.MemberProjects {
		selectClauses = append(selectClauses, userTableMemberProjectsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberProjectsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberProjectsGroupBySQL)
	}

	if c.joins.MemberTeams {
		selectClauses = append(selectClauses, userTableMemberTeamsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberTeamsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberTeamsGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.AssigneeWorkItems {
		selectClauses = append(selectClauses, userTableAssigneeWorkItemsSelectSQL)
		joinClauses = append(joinClauses, userTableAssigneeWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, userTableAssigneeWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, userTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemCommentsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	users.age,
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.email,
	users.external_id,
	users.first_name,
	users.full_name,
	users.has_global_notifications,
	users.has_personal_notifications,
	users.last_name,
	users.role_rank,
	users.scopes,
	users.updated_at,
	users.user_id,
	users.username %s 
	 FROM public.users %s 
	 WHERE users.deleted_at = $1 AND (deleted_at IS NOT NULL)
	 %s   AND users.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UsersByDeletedAt_WhereDeletedAtIsNotNull */\n" + sqlstr

	// run
	// logf(sqlstr, deletedAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{deletedAt}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/UsersByDeletedAt/Query: %w", &XoError{Entity: "User", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/UsersByDeletedAt/pgx.CollectRows: %w", &XoError{Entity: "User", Err: err}))
	}
	return res, nil
}

// UserByEmail retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_email_key'.
func UserByEmail(ctx context.Context, db DB, email string, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.ReceiverNotifications {
		selectClauses = append(selectClauses, userTableReceiverNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableReceiverNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableReceiverNotificationsGroupBySQL)
	}

	if c.joins.SenderNotifications {
		selectClauses = append(selectClauses, userTableSenderNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableSenderNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableSenderNotificationsGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, userTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, userTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, userTableTimeEntriesGroupBySQL)
	}

	if c.joins.UserNotifications {
		selectClauses = append(selectClauses, userTableUserNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableUserNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserNotificationsGroupBySQL)
	}

	if c.joins.MemberProjects {
		selectClauses = append(selectClauses, userTableMemberProjectsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberProjectsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberProjectsGroupBySQL)
	}

	if c.joins.MemberTeams {
		selectClauses = append(selectClauses, userTableMemberTeamsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberTeamsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberTeamsGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.AssigneeWorkItems {
		selectClauses = append(selectClauses, userTableAssigneeWorkItemsSelectSQL)
		joinClauses = append(joinClauses, userTableAssigneeWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, userTableAssigneeWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, userTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemCommentsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	users.age,
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.email,
	users.external_id,
	users.first_name,
	users.full_name,
	users.has_global_notifications,
	users.has_personal_notifications,
	users.last_name,
	users.role_rank,
	users.scopes,
	users.updated_at,
	users.user_id,
	users.username %s 
	 FROM public.users %s 
	 WHERE users.email = $1
	 %s   AND users.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserByEmail */\n" + sqlstr

	// run
	// logf(sqlstr, email)
	rows, err := db.Query(ctx, sqlstr, append([]any{email}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByEmail/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByEmail/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &u, nil
}

// UserByExternalID retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_external_id_key'.
func UserByExternalID(ctx context.Context, db DB, externalID string, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.ReceiverNotifications {
		selectClauses = append(selectClauses, userTableReceiverNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableReceiverNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableReceiverNotificationsGroupBySQL)
	}

	if c.joins.SenderNotifications {
		selectClauses = append(selectClauses, userTableSenderNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableSenderNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableSenderNotificationsGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, userTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, userTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, userTableTimeEntriesGroupBySQL)
	}

	if c.joins.UserNotifications {
		selectClauses = append(selectClauses, userTableUserNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableUserNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserNotificationsGroupBySQL)
	}

	if c.joins.MemberProjects {
		selectClauses = append(selectClauses, userTableMemberProjectsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberProjectsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberProjectsGroupBySQL)
	}

	if c.joins.MemberTeams {
		selectClauses = append(selectClauses, userTableMemberTeamsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberTeamsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberTeamsGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.AssigneeWorkItems {
		selectClauses = append(selectClauses, userTableAssigneeWorkItemsSelectSQL)
		joinClauses = append(joinClauses, userTableAssigneeWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, userTableAssigneeWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, userTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemCommentsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	users.age,
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.email,
	users.external_id,
	users.first_name,
	users.full_name,
	users.has_global_notifications,
	users.has_personal_notifications,
	users.last_name,
	users.role_rank,
	users.scopes,
	users.updated_at,
	users.user_id,
	users.username %s 
	 FROM public.users %s 
	 WHERE users.external_id = $1
	 %s   AND users.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserByExternalID */\n" + sqlstr

	// run
	// logf(sqlstr, externalID)
	rows, err := db.Query(ctx, sqlstr, append([]any{externalID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByExternalID/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByExternalID/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &u, nil
}

// UserByUserID retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_pkey'.
func UserByUserID(ctx context.Context, db DB, userID UserID, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.ReceiverNotifications {
		selectClauses = append(selectClauses, userTableReceiverNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableReceiverNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableReceiverNotificationsGroupBySQL)
	}

	if c.joins.SenderNotifications {
		selectClauses = append(selectClauses, userTableSenderNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableSenderNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableSenderNotificationsGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, userTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, userTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, userTableTimeEntriesGroupBySQL)
	}

	if c.joins.UserNotifications {
		selectClauses = append(selectClauses, userTableUserNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableUserNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserNotificationsGroupBySQL)
	}

	if c.joins.MemberProjects {
		selectClauses = append(selectClauses, userTableMemberProjectsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberProjectsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberProjectsGroupBySQL)
	}

	if c.joins.MemberTeams {
		selectClauses = append(selectClauses, userTableMemberTeamsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberTeamsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberTeamsGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.AssigneeWorkItems {
		selectClauses = append(selectClauses, userTableAssigneeWorkItemsSelectSQL)
		joinClauses = append(joinClauses, userTableAssigneeWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, userTableAssigneeWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, userTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemCommentsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	users.age,
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.email,
	users.external_id,
	users.first_name,
	users.full_name,
	users.has_global_notifications,
	users.has_personal_notifications,
	users.last_name,
	users.role_rank,
	users.scopes,
	users.updated_at,
	users.user_id,
	users.username %s 
	 FROM public.users %s 
	 WHERE users.user_id = $1
	 %s   AND users.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserByUserID */\n" + sqlstr

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, append([]any{userID}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &u, nil
}

// UsersByUpdatedAt retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_updated_at_idx'.
func UsersByUpdatedAt(ctx context.Context, db DB, updatedAt time.Time, opts ...UserSelectConfigOption) ([]User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.ReceiverNotifications {
		selectClauses = append(selectClauses, userTableReceiverNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableReceiverNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableReceiverNotificationsGroupBySQL)
	}

	if c.joins.SenderNotifications {
		selectClauses = append(selectClauses, userTableSenderNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableSenderNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableSenderNotificationsGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, userTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, userTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, userTableTimeEntriesGroupBySQL)
	}

	if c.joins.UserNotifications {
		selectClauses = append(selectClauses, userTableUserNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableUserNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserNotificationsGroupBySQL)
	}

	if c.joins.MemberProjects {
		selectClauses = append(selectClauses, userTableMemberProjectsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberProjectsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberProjectsGroupBySQL)
	}

	if c.joins.MemberTeams {
		selectClauses = append(selectClauses, userTableMemberTeamsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberTeamsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberTeamsGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.AssigneeWorkItems {
		selectClauses = append(selectClauses, userTableAssigneeWorkItemsSelectSQL)
		joinClauses = append(joinClauses, userTableAssigneeWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, userTableAssigneeWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, userTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemCommentsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	users.age,
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.email,
	users.external_id,
	users.first_name,
	users.full_name,
	users.has_global_notifications,
	users.has_personal_notifications,
	users.last_name,
	users.role_rank,
	users.scopes,
	users.updated_at,
	users.user_id,
	users.username %s 
	 FROM public.users %s 
	 WHERE users.updated_at = $1
	 %s   AND users.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UsersByUpdatedAt */\n" + sqlstr

	// run
	// logf(sqlstr, updatedAt)
	rows, err := db.Query(ctx, sqlstr, append([]any{updatedAt}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/UsersByUpdatedAt/Query: %w", &XoError{Entity: "User", Err: err}))
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/UsersByUpdatedAt/pgx.CollectRows: %w", &XoError{Entity: "User", Err: err}))
	}
	return res, nil
}

// UserByUsername retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_username_key'.
func UserByUsername(ctx context.Context, db DB, username string, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}, filters: make(map[string][]any), having: make(map[string][]any)}

	for _, o := range opts {
		o(c)
	}

	paramStart := 1
	nth := func() string {
		paramStart++
		return strconv.Itoa(paramStart)
	}

	var filterClauses []string
	var filterParams []any
	for filterTmpl, params := range c.filters {
		filter := filterTmpl
		for strings.Contains(filter, "$i") {
			filter = strings.Replace(filter, "$i", "$"+nth(), 1)
		}
		filterClauses = append(filterClauses, filter)
		filterParams = append(filterParams, params...)
	}

	filters := ""
	if len(filterClauses) > 0 {
		filters = " AND " + strings.Join(filterClauses, " AND ") + " "
	}

	var havingClauses []string
	var havingParams []any
	for havingTmpl, params := range c.having {
		having := havingTmpl
		for strings.Contains(having, "$i") {
			having = strings.Replace(having, "$i", "$"+nth(), 1)
		}
		havingClauses = append(havingClauses, having)
		havingParams = append(havingParams, params...)
	}

	havingClause := "" // must be empty if no actual clause passed, else it errors out
	if len(havingClauses) > 0 {
		havingClause = " HAVING " + strings.Join(havingClauses, " AND ") + " "
	}

	var selectClauses []string
	var joinClauses []string
	var groupByClauses []string

	if c.joins.ReceiverNotifications {
		selectClauses = append(selectClauses, userTableReceiverNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableReceiverNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableReceiverNotificationsGroupBySQL)
	}

	if c.joins.SenderNotifications {
		selectClauses = append(selectClauses, userTableSenderNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableSenderNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableSenderNotificationsGroupBySQL)
	}

	if c.joins.TimeEntries {
		selectClauses = append(selectClauses, userTableTimeEntriesSelectSQL)
		joinClauses = append(joinClauses, userTableTimeEntriesJoinSQL)
		groupByClauses = append(groupByClauses, userTableTimeEntriesGroupBySQL)
	}

	if c.joins.UserNotifications {
		selectClauses = append(selectClauses, userTableUserNotificationsSelectSQL)
		joinClauses = append(joinClauses, userTableUserNotificationsJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserNotificationsGroupBySQL)
	}

	if c.joins.MemberProjects {
		selectClauses = append(selectClauses, userTableMemberProjectsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberProjectsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberProjectsGroupBySQL)
	}

	if c.joins.MemberTeams {
		selectClauses = append(selectClauses, userTableMemberTeamsSelectSQL)
		joinClauses = append(joinClauses, userTableMemberTeamsJoinSQL)
		groupByClauses = append(groupByClauses, userTableMemberTeamsGroupBySQL)
	}

	if c.joins.UserAPIKey {
		selectClauses = append(selectClauses, userTableUserAPIKeySelectSQL)
		joinClauses = append(joinClauses, userTableUserAPIKeyJoinSQL)
		groupByClauses = append(groupByClauses, userTableUserAPIKeyGroupBySQL)
	}

	if c.joins.AssigneeWorkItems {
		selectClauses = append(selectClauses, userTableAssigneeWorkItemsSelectSQL)
		joinClauses = append(joinClauses, userTableAssigneeWorkItemsJoinSQL)
		groupByClauses = append(groupByClauses, userTableAssigneeWorkItemsGroupBySQL)
	}

	if c.joins.WorkItemComments {
		selectClauses = append(selectClauses, userTableWorkItemCommentsSelectSQL)
		joinClauses = append(joinClauses, userTableWorkItemCommentsJoinSQL)
		groupByClauses = append(groupByClauses, userTableWorkItemCommentsGroupBySQL)
	}

	selects := ""
	if len(selectClauses) > 0 {
		selects = ", " + strings.Join(selectClauses, " ,\n ") + " "
	}
	joins := strings.Join(joinClauses, " \n ") + " "
	groupbys := ""
	if len(groupByClauses) > 0 {
		groupbys = "GROUP BY " + strings.Join(groupByClauses, " ,\n ") + " "
	}

	sqlstr := fmt.Sprintf(`SELECT 
	users.age,
	users.api_key_id,
	users.created_at,
	users.deleted_at,
	users.email,
	users.external_id,
	users.first_name,
	users.full_name,
	users.has_global_notifications,
	users.has_personal_notifications,
	users.last_name,
	users.role_rank,
	users.scopes,
	users.updated_at,
	users.user_id,
	users.username %s 
	 FROM public.users %s 
	 WHERE users.username = $1
	 %s   AND users.deleted_at is %s  %s 
  %s 
`, selects, joins, filters, c.deletedAt, groupbys, havingClause)
	sqlstr += c.orderBy
	sqlstr += c.limit
	sqlstr = "/* UserByUsername */\n" + sqlstr

	// run
	// logf(sqlstr, username)
	rows, err := db.Query(ctx, sqlstr, append([]any{username}, append(filterParams, havingParams...)...)...)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUsername/db.Query: %w", &XoError{Entity: "User", Err: err}))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUsername/pgx.CollectOneRow: %w", &XoError{Entity: "User", Err: err}))
	}

	return &u, nil
}

// FKUserAPIKey_APIKeyID returns the UserAPIKey associated with the User's (APIKeyID).
//
// Generated from foreign key 'users_api_key_id_fkey'.
func (u *User) FKUserAPIKey_APIKeyID(ctx context.Context, db DB) (*UserAPIKey, error) {
	return UserAPIKeyByUserAPIKeyID(ctx, db, *u.APIKeyID)
}
