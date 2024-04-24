package e2e

import "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"

/*
 NOTE: tygo just looks the ast from the current package.
 All spec models available in tstype as `models` namespace
*/

// User ids are uuids, therefore we can use any unique column for
// e2e identifiers, which is also easier to reason about.
type User struct {
	Username  string        `json:"username"`
	Email     string        `json:"email"`
	FirstName *string       `json:"firstName"`
	LastName  *string       `json:"lastName"`
	Scopes    models.Scopes `json:"scopes"    tstype:"models.Scopes"`
	Role      models.Role   `json:"role"      tstype:"models.Role"`
}

// TODO: should include ids for the rest of entities with serial ids,
// given that e2e data will not be created concurrently.
type Team struct {
	Name        string             `json:"name"`
	ProjectName models.ProjectName `json:"projectName" tstype:"models.Project"`
}
