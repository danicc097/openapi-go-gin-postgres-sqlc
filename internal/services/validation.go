package services

type validateMode int

const (
	validateModeCreate validateMode = iota
	validateModeUpdate
)
