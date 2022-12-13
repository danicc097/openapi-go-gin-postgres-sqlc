// Code generated by pregen. DO NOT EDIT.

package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) authMiddlewares(opID OperationID) []gin.HandlerFunc {
	switch opID {
	case AdminPing:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
					RequiredScopes: []models.Scope{
						models.Scope("test-scope")},
				}),
		}
	case DeleteUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
					RequiredScopes: []models.Scope{
						models.Scope("test-scope"),
						models.Scope("users:write")},
				}),
		}
	case Events:
		return []gin.HandlerFunc{}
	case GetCurrentUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetProject:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetProjectBoard:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetProjectWorkitems:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case InitializeProject:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
				}),
		}
	case MyProviderCallback:
		return []gin.HandlerFunc{}
	case MyProviderLogin:
		return []gin.HandlerFunc{}
	case OpenapiYamlGet:
		return []gin.HandlerFunc{}
	case Ping:
		return []gin.HandlerFunc{}
	case UpdateUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("user"),
				}),
		}
	case UpdateUserAuthorization:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("user"),
				}),
		}
	default:
		return []gin.HandlerFunc{}
	}
}
