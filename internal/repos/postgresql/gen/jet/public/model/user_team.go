//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
)

type UserTeam struct {
	TeamID int32     `sql:"primary_key" db:"team_id"`
	Member uuid.UUID `sql:"primary_key" db:"member"`
}
