package models

import "golang.org/x/text/language"

// User implements oidc-server storage.User.
// It is used for development and testing purposes only.
// nolint: revive
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