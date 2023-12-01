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
				}),
		}
	case CreateWorkitem:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case CreateWorkitemComment:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case CreateWorkitemTag:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case DeleteUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
					RequiredScopes: models.Scopes{
						models.Scope("users:delete")},
				}),
		}
	case DeleteWorkitem:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case Events:
		return []gin.HandlerFunc{}
	case GetCurrentUser:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetPaginatedNotifications:
		return []gin.HandlerFunc{}
	case GetProject:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetProjectBoard:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetProjectConfig:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
				}),
		}
	case GetProjectWorkitems:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	case GetWorkitem:
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
	case UpdateProjectConfig:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
			h.authmw.EnsureAuthorized(
				AuthRestriction{
					MinimumRole: models.Role("admin"),
				}),
		}
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
					RequiredScopes: models.Scopes{
						models.Scope("scopes:write")},
				}),
		}
	case UpdateWorkitem:
		return []gin.HandlerFunc{
			h.authmw.EnsureAuthenticated(),
		}
	default:
		return []gin.HandlerFunc{}
	}
}
