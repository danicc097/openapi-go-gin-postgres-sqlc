package postgresqlrandom

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/pointers"
)

// NOTE: FKs should always be passed explicitly.
func UserCreateParams() *models.UserCreateParams {
	return &models.UserCreateParams{
		Username:                 testutil.RandomNameIdentifier(1, "-") + testutil.RandomName(),
		Email:                    testutil.RandomEmail(),
		FirstName:                pointers.New(testutil.RandomFirstName()),
		LastName:                 pointers.New(testutil.RandomLastName()),
		ExternalID:               testutil.RandomString(10),
		Scopes:                   models.Scopes{"scope1", "scope2"},
		RoleRank:                 testutil.RandomInt(2, 4),
		Age:                      pointers.New(testutil.RandomInt(20, 60)),
		APIKeyID:                 nil,
		HasGlobalNotifications:   false,
		HasPersonalNotifications: false,
	}
}
