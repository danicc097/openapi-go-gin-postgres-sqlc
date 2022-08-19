package crud

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"time"
)

// User represents a row from 'public.users'.
type User struct {
	UserID      int64          `json:"user_id"`      // user_id
	Username    string         `json:"username"`     // username
	Email       string         `json:"email"`        // email
	FirstName   sql.NullString `json:"first_name"`   // first_name
	LastName    sql.NullString `json:"last_name"`    // last_name
	Role        Role           `json:"role"`         // role
	IsVerified  bool           `json:"is_verified"`  // is_verified
	Salt        string         `json:"salt"`         // salt
	Password    string         `json:"password"`     // password
	IsActive    bool           `json:"is_active"`    // is_active
	IsSuperuser bool           `json:"is_superuser"` // is_superuser
	CreatedAt   time.Time      `json:"created_at"`   // created_at
	UpdatedAt   time.Time      `json:"updated_at"`   // updated_at
	// xo fields
	_exists, _deleted bool
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
func (u *User) Insert(ctx context.Context, db DB) error {
	switch {
	case u._exists: // already exists
		return logerror(&ErrInsertFailed{ErrAlreadyExists})
	case u._deleted: // deleted
		return logerror(&ErrInsertFailed{ErrMarkedForDeletion})
	}
	// insert (primary key generated and returned by database)
	const sqlstr = `INSERT INTO public.users (` +
		`username, email, first_name, last_name, role, is_verified, salt, password, is_active, is_superuser, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12` +
		`) RETURNING user_id`
	// run
	logf(sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.Role, u.IsVerified, u.Salt, u.Password, u.IsActive, u.IsSuperuser, u.CreatedAt, u.UpdatedAt)
	if err := db.QueryRowContext(ctx, sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.Role, u.IsVerified, u.Salt, u.Password, u.IsActive, u.IsSuperuser, u.CreatedAt, u.UpdatedAt).Scan(&u.UserID); err != nil {
		return logerror(err)
	}
	// set exists
	u._exists = true
	return nil
}

// Update updates a User in the database.
func (u *User) Update(ctx context.Context, db DB) error {
	switch {
	case !u._exists: // doesn't exist
		return logerror(&ErrUpdateFailed{ErrDoesNotExist})
	case u._deleted: // deleted
		return logerror(&ErrUpdateFailed{ErrMarkedForDeletion})
	}
	// update with composite primary key
	const sqlstr = `UPDATE public.users SET ` +
		`username = $1, email = $2, first_name = $3, last_name = $4, role = $5, is_verified = $6, salt = $7, password = $8, is_active = $9, is_superuser = $10, created_at = $11, updated_at = $12 ` +
		`WHERE user_id = $13`
	// run
	logf(sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.Role, u.IsVerified, u.Salt, u.Password, u.IsActive, u.IsSuperuser, u.CreatedAt, u.UpdatedAt, u.UserID)
	if _, err := db.ExecContext(ctx, sqlstr, u.Username, u.Email, u.FirstName, u.LastName, u.Role, u.IsVerified, u.Salt, u.Password, u.IsActive, u.IsSuperuser, u.CreatedAt, u.UpdatedAt, u.UserID); err != nil {
		return logerror(err)
	}
	return nil
}

// Save saves the User to the database.
func (u *User) Save(ctx context.Context, db DB) error {
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
	const sqlstr = `INSERT INTO public.users (` +
		`user_id, username, email, first_name, last_name, role, is_verified, salt, password, is_active, is_superuser, created_at, updated_at` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13` +
		`)` +
		` ON CONFLICT (user_id) DO ` +
		`UPDATE SET ` +
		`username = EXCLUDED.username, email = EXCLUDED.email, first_name = EXCLUDED.first_name, last_name = EXCLUDED.last_name, role = EXCLUDED.role, is_verified = EXCLUDED.is_verified, salt = EXCLUDED.salt, password = EXCLUDED.password, is_active = EXCLUDED.is_active, is_superuser = EXCLUDED.is_superuser, created_at = EXCLUDED.created_at, updated_at = EXCLUDED.updated_at `
	// run
	logf(sqlstr, u.UserID, u.Username, u.Email, u.FirstName, u.LastName, u.Role, u.IsVerified, u.Salt, u.Password, u.IsActive, u.IsSuperuser, u.CreatedAt, u.UpdatedAt)
	if _, err := db.ExecContext(ctx, sqlstr, u.UserID, u.Username, u.Email, u.FirstName, u.LastName, u.Role, u.IsVerified, u.Salt, u.Password, u.IsActive, u.IsSuperuser, u.CreatedAt, u.UpdatedAt); err != nil {
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
	const sqlstr = `DELETE FROM public.users ` +
		`WHERE user_id = $1`
	// run
	logf(sqlstr, u.UserID)
	if _, err := db.ExecContext(ctx, sqlstr, u.UserID); err != nil {
		return logerror(err)
	}
	// set deleted
	u._deleted = true
	return nil
}

// UserByEmail retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_email_key'.
func UserByEmail(ctx context.Context, db DB, email string) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`user_id, username, email, first_name, last_name, role, is_verified, salt, password, is_active, is_superuser, created_at, updated_at ` +
		`FROM public.users ` +
		`WHERE email = $1`
	// run
	logf(sqlstr, email)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, email).Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.Role, &u.IsVerified, &u.Salt, &u.Password, &u.IsActive, &u.IsSuperuser, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}

// UserByUserID retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_pkey'.
func UserByUserID(ctx context.Context, db DB, userID int64) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`user_id, username, email, first_name, last_name, role, is_verified, salt, password, is_active, is_superuser, created_at, updated_at ` +
		`FROM public.users ` +
		`WHERE user_id = $1`
	// run
	logf(sqlstr, userID)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, userID).Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.Role, &u.IsVerified, &u.Salt, &u.Password, &u.IsActive, &u.IsSuperuser, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}

// UserByUsername retrieves a row from 'public.users' as a User.
//
// Generated from index 'users_username_key'.
func UserByUsername(ctx context.Context, db DB, username string) (*User, error) {
	// query
	const sqlstr = `SELECT ` +
		`user_id, username, email, first_name, last_name, role, is_verified, salt, password, is_active, is_superuser, created_at, updated_at ` +
		`FROM public.users ` +
		`WHERE username = $1`
	// run
	logf(sqlstr, username)
	u := User{
		_exists: true,
	}
	if err := db.QueryRowContext(ctx, sqlstr, username).Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.Role, &u.IsVerified, &u.Salt, &u.Password, &u.IsActive, &u.IsSuperuser, &u.CreatedAt, &u.UpdatedAt); err != nil {
		return nil, logerror(err)
	}
	return &u, nil
}
