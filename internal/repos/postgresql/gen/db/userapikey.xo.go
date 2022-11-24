package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
)

// UserAPIKey represents a row from 'public.user_api_keys'.
type UserAPIKey struct {
	UserAPIKeyID int       `json:"userAPIKeyID" db:"user_api_key_id"` // user_api_key_id
	APIKey       string    `json:"apiKey" db:"api_key"`               // api_key
	ExpiresOn    time.Time `json:"expiresOn" db:"expires_on"`         // expires_on
	UserID       uuid.UUID `json:"userID" db:"user_id"`               // user_id

	// Usedr *User `json:"uñser" db:"userd"` // O2O
	User *User `json:"user" db:"user"` // O2O
	// xo fields
	_exists, _deleted bool
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
		s.limit = fmt.Sprintf(" limit %d ", limit)
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
		if len(rows) == 0 {
			s.orderBy = ""
			return
		}
		s.orderBy = " order by "
		s.orderBy += strings.Join(rows, ", ")
	}
}

type UserAPIKeyJoins struct {
	User bool
}

// WithUserAPIKeyJoin orders results by the given columns.
func WithUserAPIKeyJoin(joins UserAPIKeyJoins) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
		s.joins = joins
	}
}

// Exists returns true when the UserAPIKey exists in the database.
func (uak *UserAPIKey) Exists() bool {
	return uak._exists
}

// Deleted returns true when the UserAPIKey has been marked for deletion from
// the database.
func (uak *UserAPIKey) Deleted() bool {
	return uak._deleted
}

// Insert inserts the UserAPIKey to the database.
func (uak *UserAPIKey) Insert(ctx context.Context, db DB) error {
	switch {
	case uak._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case uak._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	sqlstr := `INSERT INTO public.user_api_keys (` +
		`api_key, expires_on, user_id` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING user_api_key_id `
	// run
	logf(sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID)
	if err := db.QueryRow(ctx, sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID).Scan(&uak.UserAPIKeyID); err != nil {
		return logerror(err)
	}
	// set exists
	uak._exists = true
	return nil
}

// Update updates a UserAPIKey in the database.
func (uak *UserAPIKey) Update(ctx context.Context, db DB) error {
	switch {
	case !uak._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case uak._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	sqlstr := `UPDATE public.user_api_keys SET ` +
		`api_key = $1, expires_on = $2, user_id = $3 ` +
		`WHERE user_api_key_id = $4 ` +
		`RETURNING user_api_key_id `
	// run
	logf(sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID, uak.UserAPIKeyID)
	if err := db.QueryRow(ctx, sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID, uak.UserAPIKeyID).Scan(&uak.UserAPIKeyID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the UserAPIKey to the database.
func (uak *UserAPIKey) Save(ctx context.Context, db DB) error {
	if uak.Exists() {
		return uak.Update(ctx, db)
	}
	return uak.Insert(ctx, db)
}

// Upsert performs an upsert for UserAPIKey.
func (uak *UserAPIKey) Upsert(ctx context.Context, db DB) error {
	switch {
	case uak._deleted: // deleted
		return logerror(&ErrUpsertFailed{ErrMarkedForDeletion})
	}
	// upsert
	sqlstr := `INSERT INTO public.user_api_keys (` +
		`user_api_key_id, api_key, expires_on, user_id` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`)` +
		` ON CONFLICT (user_api_key_id) DO ` +
		`UPDATE SET ` +
		`api_key = EXCLUDED.api_key, expires_on = EXCLUDED.expires_on, user_id = EXCLUDED.user_id  `
	// run
	logf(sqlstr, uak.UserAPIKeyID, uak.APIKey, uak.ExpiresOn, uak.UserID)
	if _, err := db.Exec(ctx, sqlstr, uak.UserAPIKeyID, uak.APIKey, uak.ExpiresOn, uak.UserID); err != nil {
		return logerror(err)
	}
	// set exists
	uak._exists = true
	return nil
}

// Delete deletes the UserAPIKey from the database.
func (uak *UserAPIKey) Delete(ctx context.Context, db DB) error {
	switch {
	case !uak._exists: // doesn't exist
		return nil
	case uak._deleted: // deleted
		return nil
	}
	// delete with single primary key
	sqlstr := `DELETE FROM public.user_api_keys ` +
		`WHERE user_api_key_id = $1 `
	// run
	logf(sqlstr, uak.UserAPIKeyID)
	if _, err := db.Exec(ctx, sqlstr, uak.UserAPIKeyID); err != nil {
		return logerror(err)
	}
	// set deleted
	uak._deleted = true
	return nil
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
(case when $1::boolean = true then users.* end) as user ` + // TODO maybe needs :: type annotation like with jsonb else wont work. (when join=False everything works)
		`FROM public.user_api_keys ` +
		`-- O2O join generated from "user_api_keys_user_id_fkey"
left join users on users.user_id = user_api_keys.user_id` +
		` WHERE user_api_keys.api_key = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, apiKey)
	// uak := UserAPIKey{
	// 	_exists: true,
	// }
	// .Scan(&uak.UserAPIKeyID, &uak.APIKey, &uak.ExpiresOn, &uak.UserID, &uak.User)
	rows, err := db.Query(ctx, sqlstr, true, apiKey)
	if err != nil {
		return nil, errors.Wrap(err, "Query")
	}
	// uakk, err := pgx.RowToStructByName[UserAPIKey](rows)
	// if err != nil {
	// 	return nil, errors.Wrap(err, "pgx.RowToStructByNam")
	// }
	// fmt.Printf("uakk: %v\n", uakk)
	// slice, err := pgx.CollectRows(rows, pgx.RowToStructByName[UserAPIKey]) // for single return
	slice, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[UserAPIKey]) // for array_agg
	// _, err := pgx.RowToStructByName[UserAPIKey](row)
	if err != nil {
		return nil, errors.Wrap(err, "CollectRows")
	}

	fmt.Printf("slice: %v\n", slice)

	return &slice, nil
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
(case when $1::boolean = true then users.* end) as user ` +
		`FROM public.user_api_keys ` +
		`-- O2O join generated from "user_api_keys_user_id_fkey"
left join users on users.user_id = user_api_keys.user_id` +
		` WHERE user_api_keys.user_api_key_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, userAPIKeyID)
	uak := UserAPIKey{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.User, userAPIKeyID).Scan(&uak.UserAPIKeyID, &uak.APIKey, &uak.ExpiresOn, &uak.UserID, &uak.User); err != nil {
		return nil, logerror(err)
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
(case when $1::boolean = true then users.* end) as user ` +
		`FROM public.user_api_keys ` +
		`-- O2O join generated from "user_api_keys_user_id_fkey"
left join users on users.user_id = user_api_keys.user_id` +
		` WHERE user_api_keys.user_id = $2 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, userID)
	uak := UserAPIKey{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, c.joins.User, userID).Scan(&uak.UserAPIKeyID, &uak.APIKey, &uak.ExpiresOn, &uak.UserID, &uak.User); err != nil {
		return nil, logerror(err)
	}
	return &uak, nil
}

// FKUser returns the User associated with the UserAPIKey's (UserID).
//
// Generated from foreign key 'user_api_keys_user_id_fkey'.
func (uak *UserAPIKey) FKUser(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, uak.UserID)
}
