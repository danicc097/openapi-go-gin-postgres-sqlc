package internal_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/models"
	"github.com/stretchr/testify/assert"
)

func TestErrorCause(t *testing.T) {
	t.Parallel()

	var ierr *internal.Error

	err := internal.NewErrorf(models.ErrorCodeUnknown, "root")
	err = internal.WrapErrorf(err, models.ErrorCodeInvalidArgument, "wrapped 1")
	err = internal.WrapErrorf(err, models.ErrorCodeInvalidArgument, "wrapped 2")
	errors.As(err, &ierr)
	assert.Equal(t, "root", ierr.Cause().Error())

	err = internal.NewErrorf(models.ErrorCodeUnknown, "root")
	errors.As(err, &ierr)
	assert.Equal(t, "root", ierr.Cause().Error())

	err = internal.NewErrorf(models.ErrorCodeUnknown, "root")
	err = fmt.Errorf("wrapper not an internal.Error %s", err.Error())
	errors.As(err, &ierr)
	assert.Equal(t, "root", ierr.Cause().Error())

	err = fmt.Errorf("not an internal.Error")
	err = internal.WrapErrorf(err, models.ErrorCodeUnknown, "root")
	errors.As(err, &ierr)
	assert.Equal(t, "not an internal.Error", ierr.Cause().Error())
}

func TestErrorWithLoc(t *testing.T) {
	t.Parallel()

	var ierr *internal.Error

	err := internal.NewErrorWithLocf(models.ErrorCodeUnknown, []string{"nested", "0"}, "root")

	err = internal.WrapErrorWithLocf(err, models.ErrorCodeInvalidArgument, []string{}, "wrapped 1")
	errors.As(err, &ierr)
	assert.Equal(t, []string{"nested", "0"}, ierr.Loc())

	err = internal.WrapErrorWithLocf(err, models.ErrorCodeNotFound, []string{"parent"}, "wrapped 2")
	errors.As(err, &ierr)
	assert.Equal(t, []string{"parent", "nested", "0"}, ierr.Loc())
	assert.True(t, ierr.Code() == models.ErrorCodeNotFound)

	err = internal.WrapErrorWithLocf(errors.New("some error"), models.ErrorCodeInvalidArgument, []string{"loc"}, "wrapped 1")
	errors.As(err, &ierr)
	assert.Equal(t, []string{"loc"}, ierr.Loc())
}
