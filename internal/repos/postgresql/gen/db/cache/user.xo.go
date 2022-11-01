package cache

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/lib/pq"

	"github.com/google/uuid"
)

// User represents a row from 'cache.users'.
type User struct {
	UserID     uuid.NullUUID   `json:"user_id"`     // user_id
	Username   sql.NullString  `json:"username"`    // username
	Email      sql.NullString  `json:"email"`       // email
	FirstName  sql.NullString  `json:"first_name"`  // first_name
	LastName   sql.NullString  `json:"last_name"`   // last_name
	FullName   sql.NullString  `json:"full_name"`   // full_name
	ExternalID sql.NullString  `json:"external_id"` // external_id
	Role       db.NullUserRole `json:"role"`        // role
	CreatedAt  sql.NullTime    `json:"created_at"`  // created_at
	UpdatedAt  sql.NullTime    `json:"updated_at"`  // updated_at
	DeletedAt  sql.NullTime    `json:"deleted_at"`  // deleted_at
	Projects   pq.GenericArray `json:"projects"`    // projects
}

// TODO only create if exists
// GetMostRecentUser returns n most recent rows from 'users',
// ordered by "created_at" in descending order.
func GetMostRecentUser(ctx context.Context, db DB, n int) ([]*User, error) {
	// list
	const sqlstr = `SELECT ` +
		`user_id, username, email, first_name, last_name, full_name, external_id, role, created_at, updated_at, deleted_at, projects ` +
		`FROM cache.users ` +
		`ORDER BY created_at DESC LIMIT $1`
	// run
	logf(sqlstr, n)

	rows, err := db.Query(ctx, sqlstr, n)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()

	// load results
	var res []*User
	for rows.Next() {
		u := User{}
		// scan
		if err := rows.Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.FullName, &u.ExternalID, &u.Role, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.Projects); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
