package rest

import (
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/services"
)

// TODO: cannot use body.Members (variable of type []models.ServicesMember) as []services.Member
// how should mapping be handled for equivalent structs from oapi-codegen, ideally automagically?
func restMembersToServices(mm []models.ServicesMember) []services.Member {
	members := make([]services.Member, 0, len(mm))
	for i, member := range mm {
		members[i] = services.Member{
			UserID: member.UserID,
			Role:   member.Role,
		}
	}
	return members
}
