package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
)

// UserAPIKey represents a row from 'public.user_api_keys'.
// Change properties via SQL column comments, joined with ",":
//   - "property:private" to exclude a field from JSON.
//   - "type:<pkg.type>" to override the type annotation.
//   - "cardinality:O2O|M2O|M2M" to generate joins (not executed by default).
type UserAPIKey struct {
	UserAPIKeyID int       `json:"-" db:"user_api_key_id"`                    // user_api_key_id
	APIKey       string    `json:"apiKey" db:"api_key" required:"true"`       // api_key
	ExpiresOn    time.Time `json:"expiresOn" db:"expires_on" required:"true"` // expires_on
	UserID       uuid.UUID `json:"userID" db:"user_id" required:"true"`       // user_id

	UserJoin *User `json:"-" db:"user_user_api_key_id" openapi-go:"ignore"` // O2O (inferred)

}

// UserAPIKeyCreateParams represents insert params for 'public.user_api_keys'.
type UserAPIKeyCreateParams struct {
	APIKey    string    `json:"apiKey" required:"true"`    // api_key
	ExpiresOn time.Time `json:"expiresOn" required:"true"` // expires_on
	UserID    uuid.UUID `json:"userID" required:"true"`    // user_id
}

// CreateUserAPIKey creates a new UserAPIKey in the database with the given params.
func CreateUserAPIKey(ctx context.Context, db DB, params *UserAPIKeyCreateParams) (*UserAPIKey, error) {
	uak := &UserAPIKey{
		APIKey:    params.APIKey,
		ExpiresOn: params.ExpiresOn,
		UserID:    params.UserID,
	}

	return uak.Insert(ctx, db)
}

// UserAPIKeyUpdateParams represents update params for 'public.user_api_keys'
type UserAPIKeyUpdateParams struct {
	APIKey    *string    `json:"apiKey" required:"true"`    // api_key
	ExpiresOn *time.Time `json:"expiresOn" required:"true"` // expires_on
	UserID    *uuid.UUID `json:"userID" required:"true"`    // user_id
}

// SetUpdateParams updates public.user_api_keys struct fields with the specified params.
func (uak *UserAPIKey) SetUpdateParams(params *UserAPIKeyUpdateParams) {
	if params.APIKey != nil {
		uak.APIKey = *params.APIKey
	}
	if params.ExpiresOn != nil {
		uak.ExpiresOn = *params.ExpiresOn
	}
	if params.UserID != nil {
		uak.UserID = *params.UserID
	}
}

type UserAPIKeySelectConfig struct {
	limit   string
	orderBy string
	joins   UserAPIKeyJoins
}
type UserAPIKeySelectConfigOption func(*UserAPIKeySelectConfig)

// WithUserAPIKeyLimit limits row selection.
func WithUserAPIKeyLimit(limit int) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
		if limit > 0 {
			s.limit = fmt.Sprintf(" limit %d ", limit)
		}
	}
}

type UserAPIKeyOrderBy = string

const (
	UserAPIKeyExpiresOnDescNullsFirst UserAPIKeyOrderBy = " expires_on DESC NULLS FIRST "
	UserAPIKeyExpiresOnDescNullsLast  UserAPIKeyOrderBy = " expires_on DESC NULLS LAST "
	UserAPIKeyExpiresOnAscNullsFirst  UserAPIKeyOrderBy = " expires_on ASC NULLS FIRST "
	UserAPIKeyExpiresOnAscNullsLast   UserAPIKeyOrderBy = " expires_on ASC NULLS LAST "
)

// WithUserAPIKeyOrderBy orders results by the given columns.
func WithUserAPIKeyOrderBy(rows ...UserAPIKeyOrderBy) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
		if len(rows) > 0 {
			s.orderBy = " order by "
			s.orderBy += strings.Join(rows, ", ")
		}
	}
}

type UserAPIKeyJoins struct {
	User bool
}

// WithUserAPIKeyJoin joins with the given tables.
func WithUserAPIKeyJoin(joins UserAPIKeyJoins) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
		s.joins = UserAPIKeyJoins{
			User: s.joins.User || joins.User,
		}
	}
}

// Insert inserts the UserAPIKey to the database.
func (uak *UserAPIKey) Insert(ctx context.Context, db DB) (*UserAPIKey, error) {
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.user_api_keys (` +
		`api_key, expires_on, user_id` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING * `
	// run
	logf(sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID)

	rows, err := db.Query(ctx, sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Insert/db.Query: %w", err))
	}
	newuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Insert/pgx.CollectOneRow: %w", err))
	}

	*uak = newuak

	return uak, nil
}

// Update updates a UserAPIKey in the database.
func (uak *UserAPIKey) Update(ctx context.Context, db DB) (*UserAPIKey, error) {
	// update with composite primary key
	sqlstr := `UPDATE public.user_api_keys SET ` +
		`api_key = $1, expires_on = $2, user_id = $3 ` +
		`WHERE user_api_key_id = $4 ` +
		`RETURNING * `
	// run
	logf(sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID, uak.UserAPIKeyID)

	rows, err := db.Query(ctx, sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID, uak.UserAPIKeyID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Update/db.Query: %w", err))
	}
	newuak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Update/pgx.CollectOneRow: %w", err))
	}
	*uak = newuak

	return uak, nil
}

// Upsert upserts a UserAPIKey in the database.
// Requires appropiate PK(s) to be set beforehand.
func (uak *UserAPIKey) Upsert(ctx context.Context, db DB, params *UserAPIKeyCreateParams) (*UserAPIKey, error) {
	var err error

	uak.APIKey = params.APIKey
	uak.ExpiresOn = params.ExpiresOn
	uak.UserID = params.UserID

	uak, err = uak.Insert(ctx, db)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code != pgerrcode.UniqueViolation {
				return nil, fmt.Errorf("UpsertUser/Insert: %w", err)
			}
			uak, err = uak.Update(ctx, db)
			if err != nil {
				return nil, fmt.Errorf("UpsertUser/Update: %w", err)
			}
		}
	}

	return uak, err
}

// Delete deletes the UserAPIKey from the database.
func (uak *UserAPIKey) Delete(ctx context.Context, db DB) error {
	// delete with single primary key
	sqlstr := `DELETE FROM public.user_api_keys ` +
		`WHERE user_api_key_id = $1 `
	// run
	if _, err := db.Exec(ctx, sqlstr, uak.UserAPIKeyID); err != nil {
		return logerror(err)
	}
	return nil
}

// UserAPIKeyPaginatedByUserAPIKeyID returns a cursor-paginated list of UserAPIKey.
func UserAPIKeyPaginatedByUserAPIKeyID(ctx context.Context, db DB, userAPIKeyID int, opts ...UserAPIKeySelectConfigOption) ([]UserAPIKey, error) {
	c := &UserAPIKeySelectConfig{joins: UserAPIKeyJoins{}}

	for _, o := range opts {
		o(c)
	}

	sqlstr := `SELECT ` +
		`user_api_keys.user_api_key_id,
user_api_keys.api_key,
user_api_keys.expires_on,
user_api_keys.user_id,
(case when $1::boolean = true and user_api_key_ids.api_key_id is not null then row(user_api_key_ids.*) end) as user_user_api_key_id ` +
		`FROM public.user_api_keys ` +
		`-- O2O join generated from "users_api_key_id_fkey(O2O inferred)"
left join users as user_api_key_ids on user_api_key_ids.api_key_id = user_api_keys.user_api_key_id` +
		` WHERE user_api_keys.user_api_key_id > $2 GROUP BY user_api_key_ids.api_key_id, user_api_key_ids.user_id, user_api_keys.user_api_key_id `
	sqlstr += c.limit

	// run

	rows, err := db.Query(ctx, sqlstr, userAPIKeyID)
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Paginated/db.Query: %w", err))
	}
	res, err := pgx.CollectRows(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("UserAPIKey/Paginated/pgx.CollectRows: %w", err))
	}
	return res, nil
}

// UserAPIKeyByAPIKey retrieves a row from 'public.user_api_keys' as a UserAPIKey.
//
// Generated from index 'user_api_keys_api_key_key'.
func UserAPIKeyByAPIKey(ctx context.Context, db DB, apiKey string, opts ...UserAPIKeySelectConfigOption) (*UserAPIKey, error) {
	c := &UserAPIKeySelectConfig{joins: UserAPIKeyJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_api_keys.user_api_key_id,
user_api_keys.api_key,
user_api_keys.expires_on,
user_api_keys.user_id,
(case when $1::boolean = true and user_api_key_ids.api_key_id is not null then row(user_api_key_ids.*) end) as user_user_api_key_id ` +
		`FROM public.user_api_keys ` +
		`-- O2O join generated from "users_api_key_id_fkey(O2O inferred)"
left join users as user_api_key_ids on user_api_key_ids.api_key_id = user_api_keys.user_api_key_id` +
		` WHERE user_api_keys.api_key = $2 GROUP BY user_api_key_ids.api_key_id, user_api_key_ids.user_id, user_api_keys.user_api_key_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, apiKey)
	rows, err := db.Query(ctx, sqlstr, c.joins.User, apiKey)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByAPIKey/db.Query: %w", err))
	}
	uak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByAPIKey/pgx.CollectOneRow: %w", err))
	}

	return &uak, nil
}

// UserAPIKeyByUserAPIKeyID retrieves a row from 'public.user_api_keys' as a UserAPIKey.
//
// Generated from index 'user_api_keys_pkey'.
func UserAPIKeyByUserAPIKeyID(ctx context.Context, db DB, userAPIKeyID int, opts ...UserAPIKeySelectConfigOption) (*UserAPIKey, error) {
	c := &UserAPIKeySelectConfig{joins: UserAPIKeyJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_api_keys.user_api_key_id,
user_api_keys.api_key,
user_api_keys.expires_on,
user_api_keys.user_id,
(case when $1::boolean = true and user_api_key_ids.api_key_id is not null then row(user_api_key_ids.*) end) as user_user_api_key_id ` +
		`FROM public.user_api_keys ` +
		`-- O2O join generated from "users_api_key_id_fkey(O2O inferred)"
left join users as user_api_key_ids on user_api_key_ids.api_key_id = user_api_keys.user_api_key_id` +
		` WHERE user_api_keys.user_api_key_id = $2 GROUP BY user_api_key_ids.api_key_id, user_api_key_ids.user_id, user_api_keys.user_api_key_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userAPIKeyID)
	rows, err := db.Query(ctx, sqlstr, c.joins.User, userAPIKeyID)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserAPIKeyID/db.Query: %w", err))
	}
	uak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserAPIKeyID/pgx.CollectOneRow: %w", err))
	}

	return &uak, nil
}

// UserAPIKeyByUserID retrieves a row from 'public.user_api_keys' as a UserAPIKey.
//
// Generated from index 'user_api_keys_user_id_key'.
func UserAPIKeyByUserID(ctx context.Context, db DB, userID uuid.UUID, opts ...UserAPIKeySelectConfigOption) (*UserAPIKey, error) {
	c := &UserAPIKeySelectConfig{joins: UserAPIKeyJoins{}}

	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_api_keys.user_api_key_id,
user_api_keys.api_key,
user_api_keys.expires_on,
user_api_keys.user_id,
(case when $1::boolean = true and user_api_key_ids.api_key_id is not null then row(user_api_key_ids.*) end) as user_user_api_key_id ` +
		`FROM public.user_api_keys ` +
		`-- O2O join generated from "users_api_key_id_fkey(O2O inferred)"
left join users as user_api_key_ids on user_api_key_ids.api_key_id = user_api_keys.user_api_key_id` +
		` WHERE user_api_keys.user_id = $2 GROUP BY user_api_key_ids.api_key_id, user_api_key_ids.user_id, user_api_keys.user_api_key_id `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	// logf(sqlstr, userID)
	rows, err := db.Query(ctx, sqlstr, c.joins.User, userID)
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserID/db.Query: %w", err))
	}
	uak, err := pgx.CollectOneRow(rows, pgx.RowToStructByNameLax[UserAPIKey])
	if err != nil {
		return nil, logerror(fmt.Errorf("user_api_keys/UserAPIKeyByUserID/pgx.CollectOneRow: %w", err))
	}

	return &uak, nil
}

// FKUser_UserID returns the User associated with the UserAPIKey's (UserID).
//
// Generated from foreign key 'user_api_keys_user_id_fkey'.
func (uak *UserAPIKey) FKUser_UserID(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, uak.UserID)
}
