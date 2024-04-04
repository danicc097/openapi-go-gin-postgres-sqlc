//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Users struct {
	UserID                   uuid.UUID  `sql:"primary_key" db:"user_id"`
	Username                 string     `db:"username"`
	Email                    string     `db:"email"`
	Age                      *int32     `db:"age"`
	FirstName                *string    `db:"first_name"`
	LastName                 *string    `db:"last_name"`
	FullName                 *string    `db:"full_name"`
	ExternalID               string     `db:"external_id"`
	APIKeyID                 *int32     `db:"api_key_id"`
	Scopes                   string     `db:"scopes"`
	RoleRank                 int16      `db:"role_rank"`
	HasPersonalNotifications bool       `db:"has_personal_notifications"`
	HasGlobalNotifications   bool       `db:"has_global_notifications"`
	CreatedAt                time.Time  `db:"created_at"`
	UpdatedAt                time.Time  `db:"updated_at"`
	DeletedAt                *time.Time `db:"deleted_at"`
}
