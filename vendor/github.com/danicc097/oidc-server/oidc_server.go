/**
 * Package oidc_server is a modified version of the example server at https://github.com/zitadel/oidc/tree/main/example/server.
 */
package oidc_server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/danicc097/oidc-server/exampleop"
	"github.com/danicc097/oidc-server/storage"
)

// Config defines OIDC server configuration.
type Config[T storage.User] struct {
	// SetUserInfoFunc overrides population of userinfo based on scope.
	// Example:

	// 	const (
	// 		// CustomScope is an example for how to use custom scopes in this library
	// 		// (in this scenario, when requested, it will return a custom claim)
	// 		CustomScope = "custom_scope"
	// 		AuthScope   = "auth"
	// 		// CustomClaim is an example for how to return custom claims with this library
	// 		CustomClaim = "custom_claim"
	// 		AuthClaim   = "auth"
	// 	)
	//
	//	// customClaim demonstrates how to return custom claims based on provided information
	//	func customClaim(clientID string) map[string]interface{} {
	//		return map[string]interface{}{
	//			"client": clientID,
	//			"other":  "stuff",
	//		}
	//	}
	//
	// func SetUserInfoFunc(user *CustomUser, userInfo *oidc.UserInfo, scope string, clientID string) {
	// 	switch scope {
	// 	case oidc.ScopeOpenID:
	// 		userInfo.Subject = user.ID
	// 	case oidc.ScopeEmail:
	// 		userInfo.Email = user.Email
	// 		userInfo.EmailVerified = oidc.Bool(user.EmailVerified)
	// 	case oidc.ScopeProfile:
	// 		userInfo.PreferredUsername = user.Username
	// 		userInfo.Name = user.FirstName + " " + user.LastName
	// 		userInfo.FamilyName = user.LastName
	// 		userInfo.GivenName = user.FirstName
	// 		userInfo.Locale = oidc.NewLocale(user.PreferredLanguage)
	// 	case oidc.ScopePhone:
	// 		userInfo.PhoneNumber = user.Phone
	// 		userInfo.PhoneNumberVerified = user.PhoneVerified
	// 	case AuthScope:
	// 		userInfo.AppendClaims(AuthClaim, map[string]interface{}{
	// 			"is_admin": user.IsAdmin,
	// 		})
	// 	case CustomScope:
	// 		userInfo.AppendClaims(CustomClaim, customClaim(clientID))
	// 	}
	// }
	SetUserInfoFunc storage.SetUserInfoFunc[T]

	// GetPrivateClaimsFromScopesFunc will be called for the creation of a JWT access token to assert claims for custom scopes.
	// Example:
	// 	func getPrivateClaimsFromScopes(ctx context.Context, userID, clientID string, scopes []string) (claims map[string]interface{}, err error) {
	// 		for _, scope := range scopes {
	// 			switch scope {
	// 			case CustomScope:
	// 				claims = storage.AppendClaim(claims, CustomClaim, customClaim(clientID))
	// 			}
	// 		}
	// 		return claims, nil
	// 	}
	GetPrivateClaimsFromScopesFunc storage.GetPrivateClaimsFromScopesFunc

	// TLS runs the server with the given certificate.
	TLS *struct {
		CertFile string
		KeyFile  string
	}

	// PathPrefix represents domain subdirectories for the base URL, if any.
	PathPrefix string
}

// Runs starts the OIDC server.
func Run[T storage.User](config Config[T]) {
	redirectURIsPath := path.Join(os.Getenv("DATA_DIR"), "redirect_uris.txt")
	content, err := os.ReadFile(redirectURIsPath)
	if err != nil {
		panic(fmt.Errorf("could not read %s: %w", redirectURIsPath, err))
	}

	redirectURIs := strings.Split(string(content), "\n")

	log.Default().Printf("Redirect URIs: %s\n", redirectURIs)

	if config.PathPrefix != "" {
		log.Default().Printf("Using domain path prefix: %v\n", config.PathPrefix)
	}

	storage.RegisterClients(
		storage.NativeClient("native", config.PathPrefix, redirectURIs...),
		storage.WebClient("web", "secret", config.PathPrefix, redirectURIs...),
		storage.WebClient("api", "secret", config.PathPrefix, redirectURIs...),
	)

	ctx := context.Background()

	issuer := os.Getenv("ISSUER")
	port := "10001" // for internal network
	usersDataDir := path.Join(os.Getenv("DATA_DIR"), "users")

	us, err := storage.NewUserStore[T](issuer, usersDataDir)
	if err != nil {
		log.Fatal("could not create user store: ", err)
	}

	storage := storage.NewStorage(us, config.SetUserInfoFunc, config.GetPrivateClaimsFromScopesFunc)

	router := exampleop.SetupServer(issuer, storage, config.PathPrefix)

	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}
	log.Default().Printf("listening at: %s", server.Addr)
	if config.TLS == nil {
		err = server.ListenAndServe()
	} else {
		err = server.ListenAndServeTLS(config.TLS.CertFile, config.TLS.KeyFile)
	}
	if err != nil {
		log.Fatal(err)
	}
	<-ctx.Done()
}
