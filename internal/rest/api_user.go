package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/models"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/format"
	"github.com/gin-gonic/gin"
)

func (h *StrictHandlers) UpdateUser(c *gin.Context, request UpdateUserRequestObject) (UpdateUserResponseObject, error) {
	// span attribute not inheritable:
	// see https://github.com/open-telemetry/opentelemetry-collector-contrib/issues/14026
	caller, _ := GetUserCallerFromCtx(c)

	span := GetSpanFromCtx(c)
	span.AddEvent("update-user") // filterable with event="update-user"

	tx := GetTxFromCtx(c)

	user, err := h.svc.User.Update(c, tx, db.UserID{UUID: request.Id}, caller, request.Body)
	if err != nil {
		renderErrorResponse(c, "Could not update user", err)

		return nil, nil
	}

	role, ok := h.svc.Authorization.RoleByRank(user.RoleRank)
	if !ok {
		renderErrorResponse(c, fmt.Sprintf("Role with rank %d not found", user.RoleRank), nil)

		return nil, nil
	}

	res := User{User: user, Role: Role(role.Name)}

	return UpdateUser200JSONResponse(res), nil
}

func (h *StrictHandlers) DeleteUser(c *gin.Context, request DeleteUserRequestObject) (DeleteUserResponseObject, error) {
	tx := GetTxFromCtx(c)

	_, err := h.svc.User.Delete(c, tx, db.NewUserID(request.Id))
	if err != nil {
		renderErrorResponse(c, "Could not delete user", err)

		return nil, nil
	}

	return DeleteUser204Response{}, nil
}

func (h *StrictHandlers) GetCurrentUser(c *gin.Context, request GetCurrentUserRequestObject) (GetCurrentUserResponseObject, error) {
	caller, _ := GetUserCallerFromCtx(c)

	role, ok := h.svc.Authorization.RoleByRank(caller.RoleRank)
	if !ok {
		msg := fmt.Sprintf("role with rank %d not found", caller.RoleRank)
		renderErrorResponse(c, msg, errors.New(msg))

		return nil, nil
	}

	res := User{
		User:     caller.User,
		Role:     Role(role.Name),
		Teams:    &caller.Teams,
		Projects: &caller.Projects,
		APIKey:   caller.APIKey,
	}

	return GetCurrentUser200JSONResponse(res), nil
}

func (h *StrictHandlers) UpdateUserAuthorization(c *gin.Context, request UpdateUserAuthorizationRequestObject) (UpdateUserAuthorizationResponseObject, error) {
	caller, _ := GetUserCallerFromCtx(c)

	span := GetSpanFromCtx(c)
	span.AddEvent("update-user") // filterable with event="update-user"

	tx := GetTxFromCtx(c)

	if _, err := h.svc.User.UpdateUserAuthorization(c, tx, db.UserID{UUID: request.Id}, caller, request.Body); err != nil {
		renderErrorResponse(c, "Error updating user authorization", err)

		return nil, nil
	}

	return UpdateUserAuthorization204Response{}, nil
}

func (h *StrictHandlers) GetPaginatedUsers(c *gin.Context, request GetPaginatedUsersRequestObject) (GetPaginatedUsersResponseObject, error) {
	users, err := h.svc.User.Paginated(c, h.pool, request.Params)
	if err != nil {
		renderErrorResponse(c, "Could not update user", err)

		return nil, nil
	}
	if request.Params.SearchQuery.Items != nil {
		for _, item := range *request.Params.SearchQuery.Items {
			v, _ := item.Filter.ValueByDiscriminator()
			switch t := v.(type) {
			case models.PaginationFilterPrimitive:
				fmt.Printf("t.Value: %+v\n", *t.Value)
			case models.PaginationFilterArray:
				fmt.Printf("t.Value (requires bind with openapi3.Schema): %+v\n", *t.Value)
			}
		}
	}

	// TODO: will save to handlers.Spec at server startup
	spec, err := GetSwagger()
	if err != nil {
		renderErrorResponse(c, "could not get openapi spec", err)

		return nil, nil
	}
	fp, ok := spec.Components.Schemas["PaginationFilter"]
	if !ok {
		renderErrorResponse(c, "PaginationFilter schema not found", err)

		return nil, nil
	}

	fmt.Printf("PaginationFilter schema: %+v\n", fp.Value)

	// qp := c.Request.URL.Query() // don't use these, have a base of map[string]interface{}
	// from json representation of request.Params instead which is more manageable
	// qpjson: {
	// ...,
	// "searchQuery":{
	// 	"items":{
	// 		"age":{
	// 			"filter":{"filterMode":"between","value":{"0":"123"}}}
	// 		},
	// 	},
	// }
	var qpItems map[string]interface{}
	qpjson, _ := json.Marshal(request.Params.SearchQuery.Items)
	fmt.Printf("qpjson: %v\n", string(qpjson))
	_ = json.Unmarshal(qpjson, &qpItems)
	// result := parseSchema(fp.Value, qp)
	// fmt.Printf("result: %+v\n", result)

	reconstructedMap, err := reconstructMapFromSchema(fp.Value, qpItems, "PaginationFilterArray")
	if err != nil {
		fmt.Println("Error:", err)
		return nil, nil
	}
	fmt.Printf("reconstructedMap: %v\n", reconstructedMap)

	// 1. pass schema to some fn, alongside SchemaName (to be found inside any|one|allof else err)
	// 2. this fn returns a new map[string]interface{} constructed based on types.
	// all it does is convert maps to arrays if required when type is "array" - rest.sliceMapToSlice
	// else it returns the same map

	format.PrintJSON(request.Params)
	// TODO: need a custom Unmarshal when we use a struct as query params
	// that takes care of converting url.Values map indexed by pos to array (like kin-openapi util)
	// need to generate a func (t PaginationFilter) MarshalJSON() ([]byte, error) {
	// that deals with it based dynamically based on a loaded openapi3.Schema and avoid generating all possible options at compile time, since the json we get from the query params uses indexes.
	// we can replicate the same logic in kin-openapi/openapi3filter/req_resp_decoder.go,
	// by checking the list of anyof,oneof,allof schemas and converting query param  maps recursively
	// to arrays based on the given schema name, where type is array.
	// IMPORTANT: or maybe just use POST body and be done with it... who cares,
	// we cannot even cache most pagination requests
	// -------
	// NOTE: to handle anyof, oneof, allof with arrays we must convert union json before
	// calling Unmarshal on the users side, cannot be done by the runtime package in UnmarshalDeepObject.
	// ie each As(.*) interface{} method, eg:
	// (t PaginationFilter) AsPaginationFilterPrimitive() (PaginationFilterPrimitive, error)
	// (t PaginationFilter) AsPaginationFilterArray() (PaginationFilterArray, error)
	// will have to modify t.union by calling some fn with parameter schemaName
	// "PaginationFilterPrimitive" or "PaginationFilterArray",
	// and pass the openapi3.Schema, so that this fn can find the given schema with name == schemaName
	// and build a new json object as map[string]interface{} recursively from the existing
	// one which uses map[0:map[key:true] 1:map[key:false]] and convert to
	// slice [map[key:true] map[key:false]]
	// the overhead should be minimal

	nextCursor := ""
	if len(users) > 0 {
		nextCursor = users[len(users)-1].CreatedAt.Format(time.RFC3339Nano)
	}
	items := make([]User, len(users))
	for i, u := range users {
		u := u
		role, _ := h.svc.Authorization.RoleByRank(u.RoleRank)
		items[i] = User{
			User:     &u,
			Role:     Role(role.Name),
			Teams:    u.MemberTeamsJoin,
			Projects: u.MemberProjectsJoin,
		}
	}
	res := PaginatedUsersResponse{
		Page: PaginationPage{
			NextCursor: nextCursor,
		},
		Items: items,
	}

	return GetPaginatedUsers200JSONResponse(res), nil
}
