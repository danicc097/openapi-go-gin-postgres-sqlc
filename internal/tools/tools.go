//go:build tools

package tools

import (
	_ "github.com/golang-migrate/migrate/v4/cmd/migrate"       // Database Migrations
	_ "github.com/golang-migrate/migrate/v4/database/postgres" // Database Migrations
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"    // Linter
	_ "github.com/kyleconroy/sqlc/cmd/sqlc"                    // Type-Safe SQL generator
	_ "github.com/lib/pq"                                      // PostgreSQL Database driver
)
