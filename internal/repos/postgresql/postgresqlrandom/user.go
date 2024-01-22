package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

// NOTE: FKs should always be passed explicitly.
func UserCreateParams() *db.UserCreateParams {
	return &db.UserCreateParams{
		Username:                 testutil.RandomNameIdentifier(1, "-") + testutil.RandomName(),
		Email:                    testutil.RandomEmail(),
		FirstName:                pointers.New(testutil.RandomFirstName()),
		LastName:                 pointers.New(testutil.RandomLastName()),
		ExternalID:               testutil.RandomString(10),
		Scopes:                   models.Scopes{"scope1", "scope2"},
		RoleRank:                 testutil.RandomInt(2, 4),
		APIKeyID:                 nil,
		HasGlobalNotifications:   false,
		HasPersonalNotifications: false,
	}
}
