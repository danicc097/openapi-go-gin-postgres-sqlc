package models

import (
	"github.com/danicc097/oidc-server/v3/storage"
	"golang.org/x/text/language"
)

// AuthServerUser implements oidc-server storage.User.
// It is used for development and testing purposes only.
// nolint: revive
// Still cannot access common interface fields:
//
//	https://go101.org/generics/888-the-status-quo-of-go-custom-generics.html
type AuthServerUser struct {
	ID_               string       `json:"id"` // need exported for unmarshalling
	Username_         string       `json:"username"`
	Password_         string       `json:"password"`
	FirstName         string       `json:"firstName"`
	LastName          string       `json:"lastName"`
	Email             string       `json:"email"`
	EmailVerified     bool         `json:"emailVerified"`
	Phone             string       `json:"phone"`
	PhoneVerified     bool         `json:"phoneVerified"`
	PreferredLanguage language.Tag `json:"preferredLanguage"`
	IsAdmin_          bool         `json:"isAdmin"`
}

func (u AuthServerUser) ID() string {
	return u.ID_
}

func (u AuthServerUser) Username() string {
	return u.Username_
}

func (u AuthServerUser) IsAdmin() bool {
	return u.IsAdmin_
}

func (u AuthServerUser) Password() string {
	return u.Password_
}

var _ storage.User = (*AuthServerUser)(nil)
