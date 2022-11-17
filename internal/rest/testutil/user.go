package testutil

import (
	"context"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/pointers"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
)

type CreateUserParams struct {
	DeletedAt  time.Time
	Role       models.Role
	Scopes     []models.Scope
	WithToken  bool // if true, an access token is created and returned
	WithAPIKey bool // if true, an api key is created and returned
}

type CreateUserResult struct {
	User   *db.User
	APIKey string
	Token  string
}

// CreateUser creates a new random user with the given configuration.
func (ff *FixtureFactory) CreateUser(ctx context.Context, params CreateUserParams) (*CreateUserResult, error) {
	// TODO any value that has a unique constraint in db must be generated via randomXXX().
	// the only parameters accepted are high level, at the `rest` layer only.
	// IMPORTANT: functions in this package only make use of SERVICES.
	// do not use any repository or db layer components
	// services have absolutely all the logic we need for fixtures. dont want any magic or leaking.
	scopes := make([]string, len(params.Scopes))
	for i, s := range params.Scopes {
		scopes[i] = string(s)
	}
	r, err := ff.authzsvc.RoleByName(string(params.Role))
	if err != nil {
		return nil, fmt.Errorf("authzsvc.RoleByName: %v", err)
	}
	// TODO usvc.CreateUser with createUser params instead and then authn.createtoken or createapikey, usvc.deleteuser, etc. - no repo logic here
	u := &db.User{
		Username:  testutil.RandomNameIdentifier(1, "-") + testutil.RandomName(),
		Email:     testutil.RandomEmail(),
		FirstName: pointers.String(testutil.RandomFirstName()),
		LastName:  pointers.String(testutil.RandomLastName()),
		Scopes:    scopes,
		RoleRank:  r.Rank,
		DeletedAt: &params.DeletedAt,
	}

	ff.usvc.Register(ctx, ff.pool, u)

	var apiKey, accessToken string

	if params.WithAPIKey {
		apiKey = ff.authnsvc.CreateAPIKeyForUser(ctx, u) // TODO will save row in user_api_keys and also u.update() to save api_key_id
	}
	if params.WithToken {
		accessToken = ff.authnsvc.CreateAccessTokenForUser(ctx, u) // TODO simply returns a jwt
	}

	return &CreateUserResult{
		User:   u,
		APIKey: apiKey,
		Token:  accessToken,
	}, nil
}
