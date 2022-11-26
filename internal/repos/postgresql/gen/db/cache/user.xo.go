package cache

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

// UserPublic represents fields that may be exposed from 'cache.users'
// and embedded in other response models.
// Include "property:private" in a SQL column comment to exclude a field.
// Joins may be explicitly added in the Response struct.
type UserPublic struct {
	UserID     *uuid.UUID `json:"userID"`     // user_id
	Username   *string    `json:"username"`   // username
	Email      *string    `json:"email"`      // email
	FirstName  *string    `json:"firstName"`  // first_name
	LastName   *string    `json:"lastName"`   // last_name
	FullName   *string    `json:"fullName"`   // full_name
	ExternalID *string    `json:"externalID"` // external_id
	APIKeyID   *int       `json:"apiKeyID"`   // api_key_id
	Scopes     []string   `json:"scopes"`     // scopes
	RoleRank   *int16     `json:"roleRank"`   // role_rank
	CreatedAt  *time.Time `json:"createdAt"`  // created_at
	UpdatedAt  *time.Time `json:"updatedAt"`  // updated_at
	DeletedAt  *time.Time `json:"deletedAt"`  // deleted_at
	Teams      []any      `json:"teams"`      // teams
}

// User represents a row from 'cache.users'.
type User struct {
	UserID     *uuid.UUID `json:"user_id" db:"user_id" openapi-json:"userID"`             // user_id
	Username   *string    `json:"username" db:"username" openapi-json:"username"`         // username
	Email      *string    `json:"email" db:"email" openapi-json:"email"`                  // email
	FirstName  *string    `json:"first_name" db:"first_name" openapi-json:"firstName"`    // first_name
	LastName   *string    `json:"last_name" db:"last_name" openapi-json:"lastName"`       // last_name
	FullName   *string    `json:"full_name" db:"full_name" openapi-json:"fullName"`       // full_name
	ExternalID *string    `json:"external_id" db:"external_id" openapi-json:"externalID"` // external_id
	APIKeyID   *int       `json:"api_key_id" db:"api_key_id" openapi-json:"apiKeyID"`     // api_key_id
	Scopes     []string   `json:"scopes" db:"scopes" openapi-json:"scopes"`               // scopes
	RoleRank   *int16     `json:"role_rank" db:"role_rank" openapi-json:"roleRank"`       // role_rank
	CreatedAt  *time.Time `json:"created_at" db:"created_at" openapi-json:"createdAt"`    // created_at
	UpdatedAt  *time.Time `json:"updated_at" db:"updated_at" openapi-json:"updatedAt"`    // updated_at
	DeletedAt  *time.Time `json:"deleted_at" db:"deleted_at" openapi-json:"deletedAt"`    // deleted_at
	Teams      []any      `json:"teams" db:"teams" openapi-json:"teams"`                  // teams
}

func (x *User) ToPublic() UserPublic {
	return UserPublic{
		UserID:     x.UserID,
		Username:   x.Username,
		Email:      x.Email,
		FirstName:  x.FirstName,
		LastName:   x.LastName,
		FullName:   x.FullName,
		ExternalID: x.ExternalID,
		APIKeyID:   x.APIKeyID,
		Scopes:     x.Scopes,
		RoleRank:   x.RoleRank,
		CreatedAt:  x.CreatedAt,
		UpdatedAt:  x.UpdatedAt,
		DeletedAt:  x.DeletedAt,
		Teams:      x.Teams,
	}
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

type UserJoins struct{}

// WithUserJoin orders results by the given columns.
func WithUserJoin(joins UserJoins) UserSelectConfigOption {
	return func(s *UserSelectConfig) {
		s.joins = joins
	}
}

// UsersByExternalID retrieves a row from 'cache.users' as a User.
//
// Generated from index 'users_external_id_idx'.
func UsersByExternalID(ctx context.Context, db DB, externalID *string, opts ...UserSelectConfigOption) ([]*User, error) {
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
users.created_at,
users.updated_at,
users.deleted_at,
users.teams `+
		`FROM cache.users `+
		``+
		` WHERE users.external_id = $1  AND users.deleted_at is %s `, c.deletedAt)
	sqlstr += c.orderBy
	sqlstr += c.limit

	// run
	logf(sqlstr, externalID)
	rows, err := db.Query(ctx, sqlstr, externalID)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// process
	var res []*User
	for rows.Next() {
		u := User{}
		// scan
		if err := rows.Scan(&u.UserID, &u.Username, &u.Email, &u.FirstName, &u.LastName, &u.FullName, &u.ExternalID, &u.APIKeyID, &u.Scopes, &u.RoleRank, &u.CreatedAt, &u.UpdatedAt, &u.DeletedAt, &u.Teams); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &u)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
