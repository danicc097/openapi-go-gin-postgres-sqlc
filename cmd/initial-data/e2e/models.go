package e2e

import "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"

/*
 NOTE: tygo just looks the ast from the current package.
 All spec models available in tstype as `models` namespace
*/

type User struct {
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	FirstName *string       `json:"firstName"`
	LastName  *string       `json:"lastName"`
	Scopes    models.Scopes `json:"scopes"    tstype:"models.Scopes"`
	Role      models.Role   `json:"role"      tstype:"models.Role"`
}
