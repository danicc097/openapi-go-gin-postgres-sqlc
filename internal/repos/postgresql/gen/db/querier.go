// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"
)

type Querier interface {
	// plpgsql-language-server:use-keyword-query-parameter
	GetUser(ctx context.Context, db DBTX, arg GetUserParams) (GetUserRow, error)
	// plpgsql-language-server:use-keyword-query-parameter
	GetUserNotifications(ctx context.Context, db DBTX, arg GetUserNotificationsParams) ([]GetUserNotificationsRow, error)
	IsTeamInProject(ctx context.Context, db DBTX, arg IsTeamInProjectParams) (bool, error)
	IsUserInProject(ctx context.Context, db DBTX, arg IsUserInProjectParams) (bool, error)
}

var _ Querier = (*Queries)(nil)
