// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package db

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
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

type Animals struct {
	AnimalID int64  `db:"animal_id" json:"animal_id"`
	Name     string `db:"name" json:"name"`
}

type PetTags struct {
	PetTagID int64  `db:"pet_tag_id" json:"pet_tag_id"`
	Name     string `db:"name" json:"name"`
}

type Pets struct {
	PetID    int64          `db:"pet_id" json:"pet_id"`
	Color    sql.NullString `db:"color" json:"color"`
	Metadata pgtype.JSONB   `db:"metadata" json:"metadata"`
}

type Users struct {
	UserID      int64          `db:"user_id" json:"user_id"`
	Username    string         `db:"username" json:"username"`
	Email       string         `db:"email" json:"email"`
	FirstName   sql.NullString `db:"first_name" json:"first_name"`
	LastName    sql.NullString `db:"last_name" json:"last_name"`
	Role        Role           `db:"role" json:"role"`
	IsVerified  bool           `db:"is_verified" json:"is_verified"`
	Salt        string         `db:"salt" json:"salt"`
	Password    string         `db:"password" json:"password"`
	IsActive    bool           `db:"is_active" json:"is_active"`
	IsSuperuser bool           `db:"is_superuser" json:"is_superuser"`
	CreatedAt   time.Time      `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" json:"updated_at"`
}
