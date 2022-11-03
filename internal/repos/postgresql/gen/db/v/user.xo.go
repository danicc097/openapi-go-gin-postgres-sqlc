package v

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/lib/pq"

	"github.com/google/uuid"
)

// User represents a row from 'v.users'.
type User struct {
	UserID     uuid.NullUUID   `json:"user_id"`     // user_id
	Username   sql.NullString  `json:"username"`    // username
	Email      sql.NullString  `json:"email"`       // email
	Scopes     pq.StringArray  `json:"scopes"`      // scopes
	FirstName  sql.NullString  `json:"first_name"`  // first_name
	LastName   sql.NullString  `json:"last_name"`   // last_name
	FullName   sql.NullString  `json:"full_name"`   // full_name
	ExternalID sql.NullString  `json:"external_id"` // external_id
	Role       db.NullUserRole `json:"role"`        // role
	CreatedAt  sql.NullTime    `json:"created_at"`  // created_at
	UpdatedAt  sql.NullTime    `json:"updated_at"`  // updated_at
	DeletedAt  sql.NullTime    `json:"deleted_at"`  // deleted_at
	Teams      pq.GenericArray `json:"teams"`       // teams
}

type UserSelectConfig struct {
	limit    string
	orderBy  string
	joinWith []UserJoinBy
}

type UserSelectConfigOption func(*UserSelectConfig)

// UserWithLimit limits row selection.
func UserWithLimit(limit int) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.limit = fmt.Sprintf(" limit %d ", limit)
	}
}

type UserOrderBy = string

const (
	UserCreatedAtDescNullsFirst UserOrderBy = "CreatedAt DescNullsFirst"
	UserCreatedAtDescNullsLast  UserOrderBy = "CreatedAt DescNullsLast"
	UserCreatedAtAscNullsFirst  UserOrderBy = "CreatedAt AscNullsFirst"
	UserCreatedAtAscNullsLast   UserOrderBy = "CreatedAt AscNullsLast"
	UserUpdatedAtDescNullsFirst UserOrderBy = "UpdatedAt DescNullsFirst"
	UserUpdatedAtDescNullsLast  UserOrderBy = "UpdatedAt DescNullsLast"
	UserUpdatedAtAscNullsFirst  UserOrderBy = "UpdatedAt AscNullsFirst"
	UserUpdatedAtAscNullsLast   UserOrderBy = "UpdatedAt AscNullsLast"
	UserDeletedAtDescNullsFirst UserOrderBy = "DeletedAt DescNullsFirst"
	UserDeletedAtDescNullsLast  UserOrderBy = "DeletedAt DescNullsLast"
	UserDeletedAtAscNullsFirst  UserOrderBy = "DeletedAt AscNullsFirst"
	UserDeletedAtAscNullsLast   UserOrderBy = "DeletedAt AscNullsLast"
)

// UserWithOrderBy orders results by the given columns.
func UserWithOrderBy(rows ...UserOrderBy) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.orderBy = strings.Join(rows, ", ")
	}
}

type UserJoinBy = string
