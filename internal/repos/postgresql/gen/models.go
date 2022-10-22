// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleUser    Role = "user"
	RoleManager Role = "manager"
	RoleAdmin   Role = "admin"
)

func (e *Role) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Role(s)
	case string:
		*e = Role(s)
	default:
		return fmt.Errorf("unsupported scan type for Role: %T", src)
	}
	return nil
}

type NullRole struct {
	Role  Role
	Valid bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullRole) Scan(value interface{}) error {
	if value == nil {
		ns.Role, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Role.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullRole) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Role, nil
}

func (e Role) Valid() bool {
	switch e {
	case RoleUser,
		RoleManager,
		RoleAdmin:
		return true
	}
	return false
}

func AllRoleValues() []Role {
	return []Role{
		RoleUser,
		RoleManager,
		RoleAdmin,
	}
}

type ApiKeys struct {
	ApiKeyID int32     `db:"api_key_id" json:"api_key_id"`
	ApiKey   string    `db:"api_key" json:"api_key"`
	UserID   uuid.UUID `db:"user_id" json:"user_id"`
}

type Movies struct {
	MovieID  int32  `db:"movie_id" json:"movie_id"`
	Title    string `db:"title" json:"title"`
	Year     int32  `db:"year" json:"year"`
	Synopsis string `db:"synopsis" json:"synopsis"`
}

type Users struct {
	UserID      uuid.UUID      `db:"user_id" json:"user_id"`
	Username    string         `db:"username" json:"username"`
	Email       string         `db:"email" json:"email"`
	FirstName   sql.NullString `db:"first_name" json:"first_name"`
	LastName    sql.NullString `db:"last_name" json:"last_name"`
	FullName    sql.NullString `db:"full_name" json:"full_name"`
	ExternalID  string         `db:"external_id" json:"external_id"`
	Role        Role           `db:"role" json:"role"`
	IsSuperuser bool           `db:"is_superuser" json:"is_superuser"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at" json:"deleted_at"`
}
