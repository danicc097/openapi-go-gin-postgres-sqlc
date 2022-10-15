package rest

import (
	"regexp"

	"github.com/gin-gonic/gin/binding"
	validator "github.com/go-playground/validator/v10"
)

const (
	alphanumSpace = "^[a-zA-Z0-9 ]+$"
)

var alphanumSpaceRegex = regexp.MustCompile(alphanumSpace)

var Alphanumspace validator.Func = func(fl validator.FieldLevel) bool {
	return alphanumSpaceRegex.MatchString(fl.Field().String())
}

func registerValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("alphanumspace", Alphanumspace)
	}
}
