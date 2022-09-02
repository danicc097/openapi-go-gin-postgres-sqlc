package rest

import (
	"regexp"

	validator "github.com/go-playground/validator/v10"
)

const (
	alphanumSpace = "^[a-zA-Z0-9 ]+$"
)

var (
	alphanumSpaceRegex = regexp.MustCompile(alphanumSpace)
)

var Alphanumspace validator.Func = func(fl validator.FieldLevel) bool {
	return alphanumSpaceRegex.MatchString(fl.Field().String())
}
