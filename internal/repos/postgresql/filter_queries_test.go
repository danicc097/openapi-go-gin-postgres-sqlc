package postgresql_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

type filterTestCase struct {
	name       string
	filter     any
	repoMethod reflect.Value
	// callback   genericFilterCallbackFunc[T] // but we need different T for each test case...
}

type genericFilterCallbackFunc[T any] func(t *testing.T, foundEntity T)

// runGenericFilterTests tests db filter functions for an entity by running a callback function
// on the found entity that verifies filter output.
func runGenericFilterTests[T any](t *testing.T, tc filterTestCase, callback genericFilterCallbackFunc[T]) {
	t.Run(tc.name, func(t *testing.T) {
		t.Run("rows_if_exists", func(t *testing.T) {
			t.Parallel()

			var foundEntity T
			var err error

			args := []reflect.Value{
				reflect.ValueOf(context.Background()),
				reflect.ValueOf(testPool),
				reflect.ValueOf(tc.filter),
			}

			result := tc.repoMethod.Call(args)

			if result[1].Interface() != nil {
				err = result[1].Interface().(error)
			} else {
				r := result[0].Interface()
				var ok bool
				foundEntity, ok = r.(T)
				require.True(t, ok, "mismatched entity type returned: got %T want %T", r, foundEntity)
			}
			require.NoError(t, err)

			callback(t, foundEntity)
		})

		t.Run("no_rows_if_not_exists", func(t *testing.T) {
			t.Parallel()

			var err error

			args := []reflect.Value{
				reflect.ValueOf(context.Background()),
				reflect.ValueOf(testPool),
			}

			filterargs, err := buildFilterArgs(reflect.ValueOf(tc.filter))
			require.NoError(t, err)

			result := tc.repoMethod.Call(append(args, filterargs...))

			if result[1].Interface() != nil {
				err = result[1].Interface().(error)
			}

			require.ErrorContains(t, err, errNoRows)
		})
	})
}

func buildFilterArgs(filter reflect.Value) ([]reflect.Value, error) {
	args := []reflect.Value{}

	switch filter.Type().Kind() {
	case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		args = append(args, reflect.Zero(filter.Type()))
	case reflect.Array:
		if filter.Type() == reflect.TypeOf(uuid.UUID{}) {
			args = append(args, reflect.ValueOf(uuid.Nil))
		}
	case reflect.Slice: // assume testing with multiple parameters
		for i := 0; i < filter.Len(); i++ {
			elem := filter.Index(i)
			elemArgs, err := buildFilterArgs(elem)
			if err != nil {
				return nil, err
			}
			args = append(args, elemArgs...)
		}
	default:
		return nil, fmt.Errorf("unsupported filter type: %v", filter.Type())
	}

	return args, nil
}
