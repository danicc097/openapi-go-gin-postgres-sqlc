package testutil

import (
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

func NewLogger(t *testing.T) *zap.SugaredLogger {
	logger := zaptest.NewLogger(t, zaptest.Level(zap.DebugLevel)).Sugar()

	return logger
}
