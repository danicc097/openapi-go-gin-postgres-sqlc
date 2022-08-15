package services

import "go.uber.org/zap"

// TODO only use services, not repos. No need to turn redis, etc. into repo for now.
// TODO service interfaces

type Default struct {
	Logger *zap.Logger
}
type Docs struct {
	Logger *zap.Logger
}
type Fake struct {
	Logger *zap.Logger
}
type Pet struct {
	Logger *zap.Logger
}
type Store struct {
	Logger *zap.Logger
}
type User struct {
	Logger *zap.Logger
}
