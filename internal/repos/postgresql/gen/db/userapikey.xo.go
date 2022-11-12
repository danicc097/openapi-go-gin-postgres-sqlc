package db

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// UserAPIKey represents a row from 'public.user_api_keys'.
type UserAPIKey struct {
	UserID    uuid.UUID `json:"user_id" db:"user_id"`       // user_id
	APIKey    string    `json:"api_key" db:"api_key"`       // api_key
	ExpiresOn time.Time `json:"expires_on" db:"expires_on"` // expires_on

	// xo fields
	_exists, _deleted bool
}

type UserAPIKeySelectConfig struct {
	limit   string
	orderBy string
	joins   UserAPIKeyJoins
}

type UserAPIKeySelectConfigOption func(*UserAPIKeySelectConfig)

// UserAPIKeyWithLimit limits row selection.
func UserAPIKeyWithLimit(limit int) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type UserAPIKeyOrderBy = string

const (
	UserAPIKeyExpiresOnDescNullsFirst UserAPIKeyOrderBy = "expires_on DESC NULLS FIRST"
	UserAPIKeyExpiresOnDescNullsLast  UserAPIKeyOrderBy = "expires_on DESC NULLS LAST"
	UserAPIKeyExpiresOnAscNullsFirst  UserAPIKeyOrderBy = "expires_on ASC NULLS FIRST"
	UserAPIKeyExpiresOnAscNullsLast   UserAPIKeyOrderBy = "expires_on ASC NULLS LAST"
)

// UserAPIKeyWithOrderBy orders results by the given columns.
func UserAPIKeyWithOrderBy(rows ...UserAPIKeyOrderBy) UserAPIKeySelectConfigOption {
	return func(s *UserAPIKeySelectConfig) {
		s.orderBy = strings.Join(rows, ", ")
	}
}

type UserAPIKeyJoins struct{}

// UserAPIKeyWithJoin orders results by the given columns.
func UserAPIKeyWithJoin(joins UserAPIKeyJoins) UserAPIKeySelectConfigOption {
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
	// insert (manual)
	sqlstr := `INSERT INTO public.user_api_keys (` +
		`user_id, api_key, expires_on` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) `
	// run
	logf(sqlstr, uak.UserID, uak.APIKey, uak.ExpiresOn)
	if _, err := db.Exec(ctx, sqlstr, uak.UserID, uak.APIKey, uak.ExpiresOn); err != nil {
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
		`api_key = $1, expires_on = $2 ` +
		`WHERE user_id = $3 `
	// run
	logf(sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID)
	if _, err := db.Exec(ctx, sqlstr, uak.APIKey, uak.ExpiresOn, uak.UserID); err != nil {
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
		`user_id, api_key, expires_on` +
		`) VALUES (` +
		`$1, $2, $3` +
		`)` +
		` ON CONFLICT (user_id) DO ` +
		`UPDATE SET ` +
		`api_key = EXCLUDED.api_key, expires_on = EXCLUDED.expires_on  `
	// run
	logf(sqlstr, uak.UserID, uak.APIKey, uak.ExpiresOn)
	if _, err := db.Exec(ctx, sqlstr, uak.UserID, uak.APIKey, uak.ExpiresOn); err != nil {
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
		`WHERE user_id = $1 `
	// run
	logf(sqlstr, uak.UserID)
	if _, err := db.Exec(ctx, sqlstr, uak.UserID); err != nil {
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
	c := &UserAPIKeySelectConfig{
		joins: UserAPIKeyJoins{},
	}
	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_api_keys.user_id,
user_api_keys.api_key,
user_api_keys.expires_on ` +
		`FROM public.user_api_keys ` +
		`` +
		` WHERE api_key = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, apiKey)
	uak := UserAPIKey{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, apiKey).Scan(&uak.UserID, &uak.APIKey, &uak.ExpiresOn); err != nil {
		return nil, logerror(err)
	}
	return &uak, nil
}

// UserAPIKeyByUserID retrieves a row from 'public.user_api_keys' as a UserAPIKey.
//
// Generated from index 'user_api_keys_pkey'.
func UserAPIKeyByUserID(ctx context.Context, db DB, userID uuid.UUID, opts ...UserAPIKeySelectConfigOption) (*UserAPIKey, error) {
	c := &UserAPIKeySelectConfig{
		joins: UserAPIKeyJoins{},
	}
	for _, o := range opts {
		o(c)
	}

	// query
	sqlstr := `SELECT ` +
		`user_api_keys.user_id,
user_api_keys.api_key,
user_api_keys.expires_on ` +
		`FROM public.user_api_keys ` +
		`` +
		` WHERE user_id = $1 `
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, userID)
	uak := UserAPIKey{
		_exists: true,
	}

	if err := db.QueryRow(ctx, sqlstr, userID).Scan(&uak.UserID, &uak.APIKey, &uak.ExpiresOn); err != nil {
		return nil, logerror(err)
	}
	return &uak, nil
}

// User returns the User associated with the UserAPIKey's (UserID).
//
// Generated from foreign key 'user_api_keys_user_id_fkey'.
func (uak *UserAPIKey) User(ctx context.Context, db DB) (*User, error) {
	return UserByUserID(ctx, db, uak.UserID)
}
