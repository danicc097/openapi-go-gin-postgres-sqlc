package logging

import (
	"context"

	"golang.org/x/exp/slog"
)

type ctxKeyType struct{}

var ctxKey ctxKeyType

// FromContext takes a Logger from the context, if it was
// previously set by [ToContext]
//
// EXPERIMENTAL: Will change to log/slog import after we drop support for Go 1.20
func FromContext(ctx context.Context) (logger *slog.Logger, ok bool) {
	logger, ok = ctx.Value(ctxKey).(*slog.Logger)
	return logger, ok
}

// ToContext sets a Logger to the context.
//
// EXPERIMENTAL: Will change to log/slog import after we drop support for Go 1.20
func ToContext(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, ctxKey, logger)
}
