package services

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Authentication struct {
	Logger *zap.Logger
	Pool   *pgxpool.Pool
}

type AuthenticationService interface {
	// TODO Authentication delegated to auth server.
	// will use inmemory tokens for predefined users for simplicity.
	
}
