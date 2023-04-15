package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// User represents a row from 'public.users'.
// Include "property:private" in a SQL column comment to exclude a field from JSON.
type User struct {
	UserID                   uuid.UUID  `json:"userID" db:"user_id" required:"true"`                                      // user_id
	Username                 string     `json:"username" db:"username" required:"true"`                                   // username
	Email                    string     `json:"email" db:"email" required:"true"`                                         // email
	FirstName                *string    `json:"firstName" db:"first_name" required:"true"`                                // first_name
	LastName                 *string    `json:"lastName" db:"last_name" required:"true"`                                  // last_name
	FullName                 *string    `json:"fullName" db:"full_name" required:"true"`                                  // full_name
	ExternalID               string     `json:"-" db:"external_id"`                                                       // external_id
	APIKeyID                 *int       `json:"-" db:"api_key_id"`                                                        // api_key_id
	Scopes                   []string   `json:"-" db:"scopes"`                                                            // scopes
	RoleRank                 int16      `json:"-" db:"role_rank"`                                                         // role_rank
	HasPersonalNotifications bool       `json:"hasPersonalNotifications" db:"has_personal_notifications" required:"true"` // has_personal_notifications
	HasGlobalNotifications   bool       `json:"hasGlobalNotifications" db:"has_global_notifications" required:"true"`     // has_global_notifications
	CreatedAt                time.Time  `json:"createdAt" db:"created_at" required:"true"`                                // created_at
	UpdatedAt                time.Time  `json:"-" db:"updated_at"`                                                        // updated_at
	DeletedAt                *time.Time `json:"deletedAt" db:"deleted_at" required:"true"`                                // deleted_at

	TimeEntries *[]TimeEntry `json:"timeEntries" db:"time_entries"` // O2M
	UserAPIKey  *UserAPIKey  `json:"userAPIKey" db:"user_api_key"`  // O2O
	Teams       *[]Team      `json:"teams" db:"teams"`              // M2M
	WorkItems   *[]WorkItem  `json:"workItems" db:"work_items"`     // M2M
	// xo fields
	_exists, _deleted bool
}

// UserCreateParams represents insert params for 'public.users'
type UserCreateParams struct {
	Username                 string   `json:"username"`                 // username
	Email                    string   `json:"email"`                    // email
	FirstName                *string  `json:"firstName"`                // first_name
	LastName                 *string  `json:"lastName"`                 // last_name
	ExternalID               string   `json:"-"`                        // external_id
	APIKeyID                 *int     `json:"-"`                        // api_key_id
	Scopes                   []string `json:"-"`                        // scopes
	RoleRank                 int16    `json:"-"`                        // role_rank
	HasPersonalNotifications bool     `json:"hasPersonalNotifications"` // has_personal_notifications
	HasGlobalNotifications   bool     `json:"hasGlobalNotifications"`   // has_global_notifications
}

// UserUpdateParams represents update params for 'public.users'
type UserUpdateParams struct {
	Username                 *string   `json:"username"`                 // username
	Email                    *string   `json:"email"`                    // email
	FirstName                **string  `json:"firstName"`                // first_name
	LastName                 **string  `json:"lastName"`                 // last_name
	ExternalID               *string   `json:"-"`                        // external_id
	APIKeyID                 **int     `json:"-"`                        // api_key_id
	Scopes                   *[]string `json:"-"`                        // scopes
	RoleRank                 *int16    `json:"-"`                        // role_rank
	HasPersonalNotifications *bool     `json:"hasPersonalNotifications"` // has_personal_notifications
	HasGlobalNotifications   *bool     `json:"hasGlobalNotifications"`   // has_global_notifications
}

type UserSelectConfig struct {
	limit     string
	orderBy   string
	joins     UserJoins
	deletedAt string
}
type UserSelectConfigOption func(*UserSelectConfig)

// WithUserLimit limits row selection.
func WithUserLimit(limit int) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

// WithDeletedUserOnly limits result to records marked as deleted.
func WithDeletedUserOnly() UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.deletedAt = " not null "
	}
}

type UserOrderBy = string

const (
	UserCreatedAtDescNullsFirst UserOrderBy = " created_at DESC NULLS FIRST "
	UserCreatedAtDescNullsLast  UserOrderBy = " created_at DESC NULLS LAST "
	UserCreatedAtAscNullsFirst  UserOrderBy = " created_at ASC NULLS FIRST "
	UserCreatedAtAscNullsLast   UserOrderBy = " created_at ASC NULLS LAST "
	UserUpdatedAtDescNullsFirst UserOrderBy = " updated_at DESC NULLS FIRST "
	UserUpdatedAtDescNullsLast  UserOrderBy = " updated_at DESC NULLS LAST "
	UserUpdatedAtAscNullsFirst  UserOrderBy = " updated_at ASC NULLS FIRST "
	UserUpdatedAtAscNullsLast   UserOrderBy = " updated_at ASC NULLS LAST "
	UserDeletedAtDescNullsFirst UserOrderBy = " deleted_at DESC NULLS FIRST "
	UserDeletedAtDescNullsLast  UserOrderBy = " deleted_at DESC NULLS LAST "
	UserDeletedAtAscNullsFirst  UserOrderBy = " deleted_at ASC NULLS FIRST "
	UserDeletedAtAscNullsLast   UserOrderBy = " deleted_at ASC NULLS LAST "
)

// WithUserOrderBy orders results by the given columns.
func WithUserOrderBy(rows ...UserOrderBy) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type UserJoins struct {
	TimeEntries bool
	UserAPIKey  bool
	Teams       bool
	WorkItems   bool
}

// WithUserJoin joins with the given tables.
func WithUserJoin(joins UserJoins) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the User exists in the database.
func (u *User) Exists() bool {
	return u._exists
}

// Deleted returns true when the User has been marked for deletion from
// the database.
func (u *User) Deleted() bool {
	return u._deleted
}

// Insert inserts the User to the database.
func (u *User) Insert(ctx context.Context, db DB) (*User, error) {
	switch {
	case u._exists: // already exists
		return nil, logerror(&ErrInsertFailed{ErrAlreadyExists})
	case u._deleted: // deleted
		return nil, logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.users (` +
		`username, email, first_name, last_name, external_id, api_key_id, scopes, role_rank, has_personal_notifications, has_global_notifications, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11` +
		`) RETURNING * `
	// run
	logf(sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.HasPersonalNotifications, u.HasGlobalNotifications, u.DeletedAt)

	rows, err := db.Query(ctx, sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.HasPersonalNotifications, u.HasGlobalNotifications, u.DeletedAt)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Insert/db.Query: %w", err))
	}
	newu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Insert/pgx.CollectOneRow: %w", err))
	}
	newu._exists = true
	*u = newu

	return u, nil
}

// Update updates a User in the database.
func (u *User) Update(ctx context.Context, db DB) (*User, error) {
	switch {
	case !u._exists: // doesn't exist
		return nil, logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case u._deleted: // deleted
		return nil, logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.users SET ` +
		`username = $1, email = $2, first_name = $3, last_name = $4, external_id = $5, api_key_id = $6, scopes = $7, role_rank = $8, has_personal_notifications = $9, has_global_notifications = $10, deleted_at = $11 ` +
		`WHERE user_id = $12 ` +
		`RETURNING * `
	// run
	logf(sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.HasPersonalNotifications, u.HasGlobalNotifications, u.CreatedAt, u.UpdatedAt, u.DeletedAt, u.UserID)

	rows, err := db.Query(ctx, sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.HasPersonalNotifications, u.HasGlobalNotifications, u.DeletedAt, u.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Update/db.Query: %w", err))
	}
	newu, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("User/Update/pgx.CollectOneRow: %w", err))
	}
	newu._exists = true
	*u = newu

	return u, nil
}

// Save saves the User to the database.
func (u *User) Save(ctx context.Context, db DB) (*User, error) {
	if u.Exists() {
		return u.Update(ctx, db)
	}
	return u.Insert(ctx, db)
}

// Upsert performs an upsert for User.
func (u *User) Upsert(ctx context.Context, db DB) error {
	switch {
	case u._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.users (` +
		`user_id, username, email, first_name, last_name, full_name, external_id, api_key_id, scopes, role_rank, has_personal_notifications, has_global_notifications, deleted_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)` +
		` ON CONFLICT (user_id) DO ` +
		`UPDATE SET ` +
		`username = EXCLUDED.username, email = EXCLUDED.email, first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name, external_id = EXCLUDED.external_id, api_key_id = EXCLUDED.api_key_id, scopes = EXCLUDED.scopes, role_rank = EXCLUDED.role_rank, has_personal_notifications = EXCLUDED.has_personal_notifications, has_global_notifications = EXCLUDED.has_global_notifications, deleted_at = EXCLUDED.deleted_at ` +
		` RETURNING * `
	// run
	logf(sqlstr, u.UserID, u.Username, u.Email, u.FirstName, u.LastName, u.FullName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.HasPersonalNotifications, u.HasGlobalNotifications, u.DeletedAt)
	if _, err := db.Exec(ctx, sqlstr, u.UserID, u.Username, u.Email, u.FirstName, u.LastName, u.FullName, u.ExternalID, u.APIKeyID, u.Scopes, u.RoleRank, u.HasPersonalNotifications, u.HasGlobalNotifications, u.DeletedAt); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Delete deletes the User from the database.
func (u *User) Delete(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return nil
	case u._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.users ` +
		`WHERE user_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, u.UserID); err != nil {
		return logerror(err)
	}
	// set deleted
	u._deleted = true
	return nil
}

// SoftDelete soft deletes the User from the database via 'deleted_at'.
func (u *User) SoftDelete(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return nil
	case u._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `UPDATE public.users ` +
		`SET deleted_at = NOW() ` +
		`WHERE user_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, u.UserID); err != nil {
		return logerror(err)
	}
	// set deleted
	u._deleted = true
	u.DeletedAt = newPointer(time.Now())

	return nil
}

// Restore restores a soft deleted User from the database.
func (u *User) Restore(ctx context.Context, db DB) (*User, error) {
	u.DeletedAt = nil
	newu, err := u.Update(ctx, db)
	if err != nil {
		return nil, logerror(err)
	}
	return newu, nil
}

// UsersByCreatedAt retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_created_at_idx'.
func UsersByCreatedAt(ctx context.Context, db DB, createdAt time.Time, opts ...UserSelectConfigOption) ([]User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.has_personal_notifications,
users.has_global_notifications,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end) as time_entries,
(case when $2::boolean = true then row(user_api_keys.*) end) as user_api_key,
(case when $3::boolean = true then joined_teams.__teams end) as teams,
(case when $4::boolean = true then joined_work_items.__work_items end) as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- O2O join generated from "user_api_keys_user_id_fkey"
left join user_api_keys on user_api_keys.user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_team.user_id as user_team_user_id
		, array_agg(teams.*) as __teams
		from user_team
    join teams on teams.team_id = user_team.team_id
    group by user_team_user_id
  ) as joined_teams on joined_teams.user_team_user_id = users.user_id

-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		work_item_member.member as work_item_member_member
		, array_agg(work_items.*) as __work_items
		from work_item_member
    join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = users.user_id
`+
		` WHERE users.created_at = $5  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, createdAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.UserAPIKey, c.joins.Teams, c.joins.WorkItems, createdAt)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UsersByDeletedAt_WhereDeletedAtIsNotNull retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_deleted_at_idx'.
func UsersByDeletedAt_WhereDeletedAtIsNotNull(ctx context.Context, db DB, deletedAt *time.Time, opts ...UserSelectConfigOption) ([]User, error) {
	c := &UserSelectConfig{deletedAt: " not null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.has_personal_notifications,
users.has_global_notifications,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end) as time_entries,
(case when $2::boolean = true then row(user_api_keys.*) end) as user_api_key,
(case when $3::boolean = true then joined_teams.__teams end) as teams,
(case when $4::boolean = true then joined_work_items.__work_items end) as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- O2O join generated from "user_api_keys_user_id_fkey"
left join user_api_keys on user_api_keys.user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_team.user_id as user_team_user_id
		, array_agg(teams.*) as __teams
		from user_team
    join teams on teams.team_id = user_team.team_id
    group by user_team_user_id
  ) as joined_teams on joined_teams.user_team_user_id = users.user_id

-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		work_item_member.member as work_item_member_member
		, array_agg(work_items.*) as __work_items
		from work_item_member
    join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = users.user_id
`+
		` WHERE users.deleted_at = $5 AND (deleted_at IS NOT NULL)  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, deletedAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.UserAPIKey, c.joins.Teams, c.joins.WorkItems, deletedAt)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserByEmail retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_email_key'.
func UserByEmail(ctx context.Context, db DB, email string, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.has_personal_notifications,
users.has_global_notifications,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end) as time_entries,
(case when $2::boolean = true then row(user_api_keys.*) end) as user_api_key,
(case when $3::boolean = true then joined_teams.__teams end) as teams,
(case when $4::boolean = true then joined_work_items.__work_items end) as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- O2O join generated from "user_api_keys_user_id_fkey"
left join user_api_keys on user_api_keys.user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_team.user_id as user_team_user_id
		, array_agg(teams.*) as __teams
		from user_team
    join teams on teams.team_id = user_team.team_id
    group by user_team_user_id
  ) as joined_teams on joined_teams.user_team_user_id = users.user_id

-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		work_item_member.member as work_item_member_member
		, array_agg(work_items.*) as __work_items
		from work_item_member
    join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = users.user_id
`+
		` WHERE users.email = $5  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, email)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.UserAPIKey, c.joins.Teams, c.joins.WorkItems, email)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByEmail/db.Query: %w", err))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByEmail/pgx.CollectOneRow: %w", err))
	}
	u._exists = true
	return &u, nil
}

// UserByExternalID retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_external_id_key'.
func UserByExternalID(ctx context.Context, db DB, externalID string, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.has_personal_notifications,
users.has_global_notifications,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end) as time_entries,
(case when $2::boolean = true then row(user_api_keys.*) end) as user_api_key,
(case when $3::boolean = true then joined_teams.__teams end) as teams,
(case when $4::boolean = true then joined_work_items.__work_items end) as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- O2O join generated from "user_api_keys_user_id_fkey"
left join user_api_keys on user_api_keys.user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_team.user_id as user_team_user_id
		, array_agg(teams.*) as __teams
		from user_team
    join teams on teams.team_id = user_team.team_id
    group by user_team_user_id
  ) as joined_teams on joined_teams.user_team_user_id = users.user_id

-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		work_item_member.member as work_item_member_member
		, array_agg(work_items.*) as __work_items
		from work_item_member
    join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = users.user_id
`+
		` WHERE users.external_id = $5  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, externalID)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.UserAPIKey, c.joins.Teams, c.joins.WorkItems, externalID)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByExternalID/db.Query: %w", err))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByExternalID/pgx.CollectOneRow: %w", err))
	}
	u._exists = true
	return &u, nil
}

// UserByUserID retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_pkey'.
func UserByUserID(ctx context.Context, db DB, userID uuid.UUID, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.has_personal_notifications,
users.has_global_notifications,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end) as time_entries,
(case when $2::boolean = true then row(user_api_keys.*) end) as user_api_key,
(case when $3::boolean = true then joined_teams.__teams end) as teams,
(case when $4::boolean = true then joined_work_items.__work_items end) as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- O2O join generated from "user_api_keys_user_id_fkey"
left join user_api_keys on user_api_keys.user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_team.user_id as user_team_user_id
		, array_agg(teams.*) as __teams
		from user_team
    join teams on teams.team_id = user_team.team_id
    group by user_team_user_id
  ) as joined_teams on joined_teams.user_team_user_id = users.user_id

-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		work_item_member.member as work_item_member_member
		, array_agg(work_items.*) as __work_items
		from work_item_member
    join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = users.user_id
`+
		` WHERE users.user_id = $5  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.UserAPIKey, c.joins.Teams, c.joins.WorkItems, userID)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/db.Query: %w", err))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUserID/pgx.CollectOneRow: %w", err))
	}
	u._exists = true
	return &u, nil
}

// UsersByUpdatedAt retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_updated_at_idx'.
func UsersByUpdatedAt(ctx context.Context, db DB, updatedAt time.Time, opts ...UserSelectConfigOption) ([]User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.has_personal_notifications,
users.has_global_notifications,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end) as time_entries,
(case when $2::boolean = true then row(user_api_keys.*) end) as user_api_key,
(case when $3::boolean = true then joined_teams.__teams end) as teams,
(case when $4::boolean = true then joined_work_items.__work_items end) as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- O2O join generated from "user_api_keys_user_id_fkey"
left join user_api_keys on user_api_keys.user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_team.user_id as user_team_user_id
		, array_agg(teams.*) as __teams
		from user_team
    join teams on teams.team_id = user_team.team_id
    group by user_team_user_id
  ) as joined_teams on joined_teams.user_team_user_id = users.user_id

-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		work_item_member.member as work_item_member_member
		, array_agg(work_items.*) as __work_items
		from work_item_member
    join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = users.user_id
`+
		` WHERE users.updated_at = $5  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, updatedAt)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.UserAPIKey, c.joins.Teams, c.joins.WorkItems, updatedAt)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process

	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserByUsername retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_username_key'.
func UserByUsername(ctx context.Context, db DB, username string, opts ...UserSelectConfigOption) (*User, error) {
	c := &UserSelectConfig{deletedAt: " null ", joins: UserJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := fmt.Sprintf(`SELECT `+
		`users.user_id,
users.username,
users.email,
users.first_name,
users.last_name,
users.full_name,
users.external_id,
users.api_key_id,
users.scopes,
users.role_rank,
users.has_personal_notifications,
users.has_global_notifications,
users.created_at,
users.updated_at,
users.deleted_at,
(case when $1::boolean = true then joined_time_entries.time_entries end) as time_entries,
(case when $2::boolean = true then row(user_api_keys.*) end) as user_api_key,
(case when $3::boolean = true then joined_teams.__teams end) as teams,
(case when $4::boolean = true then joined_work_items.__work_items end) as work_items `+
		`FROM public.users `+
		`-- O2M join generated from "time_entries_user_id_fkey"
left join (
  select
  user_id as time_entries_user_id
    , array_agg(time_entries.*) as time_entries
  from
    time_entries
   group by
        user_id) joined_time_entries on joined_time_entries.time_entries_user_id = users.user_id
-- O2O join generated from "user_api_keys_user_id_fkey"
left join user_api_keys on user_api_keys.user_id = users.user_id
-- M2M join generated from "user_team_team_id_fkey"
left join (
	select
		user_team.user_id as user_team_user_id
		, array_agg(teams.*) as __teams
		from user_team
    join teams on teams.team_id = user_team.team_id
    group by user_team_user_id
  ) as joined_teams on joined_teams.user_team_user_id = users.user_id

-- M2M join generated from "work_item_member_work_item_id_fkey"
left join (
	select
		work_item_member.member as work_item_member_member
		, array_agg(work_items.*) as __work_items
		from work_item_member
    join work_items on work_items.work_item_id = work_item_member.work_item_id
    group by work_item_member_member
  ) as joined_work_items on joined_work_items.work_item_member_member = users.user_id
`+
		` WHERE users.username = $5  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, username)
	rows, err := db.Query(ctx, sqlstr, c.joins.TimeEntries, c.joins.UserAPIKey, c.joins.Teams, c.joins.WorkItems, username)
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUsername/db.Query: %w", err))
	}
	u, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[User])
	if err != nil {
		return nil, logerror(fmt.Errorf("users/UserByUsername/pgx.CollectOneRow: %w", err))
	}
	u._exists = true
	return &u, nil
}

// FKUserAPIKey_APIKeyID returns the UserAPIKey associated with the User's (APIKeyID).
//
// Generated from foreign key 'users_api_key_id_fkey'.
func (u *User) FKUserAPIKey_APIKeyID(ctx context.Context, db DB) (*UserAPIKey, error) {
	return UserAPIKeyByUserAPIKeyID(ctx, db, *u.APIKeyID)
}
