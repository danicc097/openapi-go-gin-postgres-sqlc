// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package models

import (
	"context"
)

type Querier interface {
	// plpgsql-language-server:use-keyword-query-parameter
	GetExtraSchemaNotifications(ctx context.Context, d DBTX, arg GetExtraSchemaNotificationsParams) ([]GetExtraSchemaNotificationsRow, error)
	// plpgsql-language-server:use-keyword-query-parameter
	GetUser(ctx context.Context, d DBTX, arg GetUserParams) (GetUserRow, error)
	// plpgsql-language-server:use-keyword-query-parameter
	GetUserNotifications(ctx context.Context, d DBTX, arg GetUserNotificationsParams) ([]GetUserNotificationsRow, error)
	IsTeamInProject(ctx context.Context, d DBTX, arg IsTeamInProjectParams) (bool, error)
}

var _ Querier = (*Queries)(nil)
