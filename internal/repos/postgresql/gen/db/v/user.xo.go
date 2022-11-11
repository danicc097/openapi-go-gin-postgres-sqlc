package v

// Code generated by xo. DO NOT EDIT.

import (
	"fmt"
	"strings"

	"gopkg.in/guregu/null.v4"

	"github.com/google/uuid"
)

// User represents a row from 'v.users'.
type User struct {
	UserID     uuid.NullUUID `json:"user_id" db:"user_id"`         // user_id
	Username   null.String   `json:"username" db:"username"`       // username
	Email      null.String   `json:"email" db:"email"`             // email
	FirstName  null.String   `json:"first_name" db:"first_name"`   // first_name
	LastName   null.String   `json:"last_name" db:"last_name"`     // last_name
	FullName   null.String   `json:"full_name" db:"full_name"`     // full_name
	ExternalID null.String   `json:"external_id" db:"external_id"` // external_id
	Scopes     []string      `json:"scopes" db:"scopes"`           // scopes
	RoleRank   null.Int      `json:"role_rank" db:"role_rank"`     // role_rank
	CreatedAt  null.Time     `json:"created_at" db:"created_at"`   // created_at
	UpdatedAt  null.Time     `json:"updated_at" db:"updated_at"`   // updated_at
	DeletedAt  null.Time     `json:"deleted_at" db:"deleted_at"`   // deleted_at
	Teams      []any         `json:"teams" db:"teams"`             // teams
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
	UserCreatedAtDescNullsFirst UserOrderBy = "created_at DESC NULLS FIRST"
	UserCreatedAtDescNullsLast  UserOrderBy = "created_at DESC NULLS LAST"
	UserCreatedAtAscNullsFirst  UserOrderBy = "created_at ASC NULLS FIRST"
	UserCreatedAtAscNullsLast   UserOrderBy = "created_at ASC NULLS LAST"
	UserUpdatedAtDescNullsFirst UserOrderBy = "updated_at DESC NULLS FIRST"
	UserUpdatedAtDescNullsLast  UserOrderBy = "updated_at DESC NULLS LAST"
	UserUpdatedAtAscNullsFirst  UserOrderBy = "updated_at ASC NULLS FIRST"
	UserUpdatedAtAscNullsLast   UserOrderBy = "updated_at ASC NULLS LAST"
	UserDeletedAtDescNullsFirst UserOrderBy = "deleted_at DESC NULLS FIRST"
	UserDeletedAtDescNullsLast  UserOrderBy = "deleted_at DESC NULLS LAST"
	UserDeletedAtAscNullsFirst  UserOrderBy = "deleted_at ASC NULLS FIRST"
	UserDeletedAtAscNullsLast   UserOrderBy = "deleted_at ASC NULLS LAST"
)

// UserWithOrderBy orders results by the given columns.
func UserWithOrderBy(rows ...UserOrderBy) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.orderBy = strings.Join(rows, ", ")
	}
}

type UserJoinBy = string
