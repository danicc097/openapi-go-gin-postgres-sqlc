package main

import (
	"context"
	"flag"

	oidc_server "github.com/danicc097/oidc-server/v3"
	"github.com/danicc097/oidc-server/v3/storage"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/zitadel/oidc/v2/pkg/oidc"
)

const (
	// CustomScope is an example for how to use custom scopes in this library
	// (in this scenario, when requested, it will return a custom claim).
	CustomScope = "custom_scope"
	AuthScope   = "auth"
	// CustomClaim is an example for how to return custom claims with this library.
	CustomClaim = "custom_claim"
	AuthClaim   = "auth"
)

// customClaim demonstrates how to return custom claims based on provided information.
func customClaim(clientID string) map[string]interface{} {
	return map[string]interface{}{
		"client": clientID,
		"other":  "stuff",
	}
}

func getPrivateClaimsFromScopesFunc(ctx context.Context, userID, clientID string, scopes []string) (claims map[string]interface{}, err error) {
	for _, scope := range scopes {
		switch scope {
		case CustomScope:
			claims = storage.AppendClaim(claims, CustomClaim, customClaim(clientID))
		}
	}

	return claims, nil
}

func setUserInfoFunc(user *models.AuthServerUser, userInfo *oidc.UserInfo, scope, clientID string) {
	switch scope {
	case oidc.ScopeOpenID:
		userInfo.Subject = user.ID()
	case oidc.ScopeEmail:
		userInfo.Email = user.Email
		userInfo.EmailVerified = oidc.Bool(user.EmailVerified)
	case oidc.ScopeProfile:
		userInfo.PreferredUsername = user.Username()
		userInfo.Name = user.FirstName + " " + user.LastName
		userInfo.FamilyName = user.LastName
		userInfo.GivenName = user.FirstName
		userInfo.Locale = oidc.NewLocale(user.PreferredLanguage)
	case oidc.ScopePhone:
		userInfo.PhoneNumber = user.Phone
		userInfo.PhoneNumberVerified = user.PhoneVerified
	case AuthScope:
		userInfo.AppendClaims(AuthClaim, map[string]interface{}{
			"is_admin": user.IsAdmin(),
		})
	case CustomScope:
		userInfo.AppendClaims(CustomClaim, customClaim(clientID))
	}
}

// locally: DATA_DIR="cmd/oidc-server/data" ISSUER="https://localhost:10001" go run cmd/oidc-server/main.go --cert-file certificates/localhost.pem --key-file certificates/localhost-key.pem
func main() {
	var env, certFile, keyFile, pathPrefix string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&pathPrefix, "path-prefix", "", "Domain path prefix. Example: /oidc")
	flag.StringVar(&certFile, "cert-file", "", "TLS certificate filepath")
	flag.StringVar(&keyFile, "key-file", "", "TLS certificate key filepath")

	flag.Parse()

	config := oidc_server.Config[models.AuthServerUser]{
		SetUserInfoFunc:                setUserInfoFunc,
		GetPrivateClaimsFromScopesFunc: getPrivateClaimsFromScopesFunc,
		PathPrefix:                     pathPrefix,
	}

	if certFile != "" && keyFile != "" {
		config.TLS = &struct {
			CertFile string
			KeyFile  string
		}{
			CertFile: certFile,
			KeyFile:  keyFile,
		}
	}

	oidc_server.Run(config)
}
