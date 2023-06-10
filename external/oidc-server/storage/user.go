package storage

import (
	"crypto/rsa"

	"golang.org/x/text/language"
)

type User struct {
	ID                string
	Username          string
	Password          string
	FirstName         string
	LastName          string
	Email             string
	EmailVerified     bool
	Phone             string
	PhoneVerified     bool
	PreferredLanguage language.Tag
	IsAdmin           bool
}

type Service struct {
	keys map[string]*rsa.PublicKey
}

type UserStore interface {
	GetUserByID(string) *User
	GetUserByUsername(string) *User
	ExampleClientID() string
}

type userStore struct {
	users map[string]*User
}

func NewUserStore(issuer string) UserStore {
	return userStore{
		users: map[string]*User{
			"id1": {
				ID:                "id1",
				Username:          "admin",
				Password:          "admin",
				FirstName:         "Test",
				LastName:          "User",
				Email:             "admin@admin.com",
				EmailVerified:     true,
				Phone:             "",
				PhoneVerified:     false,
				PreferredLanguage: language.German,
			},
			"id2": {
				ID:                "id2",
				Username:          "user2",
				Password:          "user2",
				FirstName:         "Test 2",
				LastName:          "User 2",
				Email:             "user2@app.com",
				EmailVerified:     true,
				Phone:             "",
				PhoneVerified:     false,
				PreferredLanguage: language.German,
			},
			"id3": {
				ID:                "id3",
				Username:          "user3",
				Password:          "user3",
				FirstName:         "Test 3",
				LastName:          "User 3",
				Email:             "user2@app.com",
				EmailVerified:     true,
				Phone:             "",
				PhoneVerified:     false,
				PreferredLanguage: language.German,
			},
		},
	}
}

// ExampleClientID is only used in the example server
func (u userStore) ExampleClientID() string {
	return "service"
}

func (u userStore) GetUserByID(id string) *User {
	return u.users[id]
}

func (u userStore) GetUserByUsername(username string) *User {
	for _, user := range u.users {
		if user.Username == username {
			return user
		}
	}
	return nil
}
