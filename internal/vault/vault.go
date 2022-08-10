// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package vault

import (
	"os"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/envvar/vault"
)

// New instantiates the Vault client using configuration defined in environment variables.
func New() (*vault.Provider, error) {
	vaultPath := os.Getenv("VAULT_PATH")
	vaultToken := os.Getenv("VAULT_TOKEN")
	vaultAddress := os.Getenv("VAULT_ADDRESS")

	provider, err := vault.New(vaultToken, vaultAddress, vaultPath)
	if err != nil {
		return nil, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "vault.New ")
	}

	return provider, nil
}
