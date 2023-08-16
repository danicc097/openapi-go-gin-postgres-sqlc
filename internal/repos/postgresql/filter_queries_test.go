package postgresql_test

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type filterTestCase[T any] struct {
	name string
	// filter represents the slice of arguments to pass to repoMethod, excluding context and db connection.
	filter     any
	repoMethod reflect.Value
	callback   genericFilterCallbackFunc[T] // T needs to be a specific type in a single TDT so this is useless
}

// callback function that should test repo method output matches expectation.
type genericFilterCallbackFunc[T any] func(t *testing.T, result T)

// runGenericFilterTests tests repo filter methods for an entity.
func runGenericFilterTests[T any](t *testing.T, tc filterTestCase[T]) {
	t.Run(tc.name, func(t *testing.T) {
		t.Parallel()

		t.Run("rows_if_exists", func(t *testing.T) {
			t.Parallel()

			var foundEntity T
			var err error

			args := []reflect.Value{
				reflect.ValueOf(context.Background()),
				reflect.ValueOf(testPool),
			}

			filterargs, err := buildFilterArgs(reflect.ValueOf(tc.filter), false)
			require.NoError(t, err)

			result := tc.repoMethod.Call(append(args, filterargs...))

			if result[1].Interface() != nil {
				err = result[1].Interface().(error)
			} else {
				r := result[0].Interface()
				var ok bool
				foundEntity, ok = r.(T)
				require.True(t, ok, "mismatched entity type returned: got %T want %T", r, foundEntity)
			}
			require.NoError(t, err)

			tc.callback(t, foundEntity)
		})

		t.Run("no_rows_if_not_exists", func(t *testing.T) {
			t.Parallel()

			var err error

			args := []reflect.Value{
				reflect.ValueOf(context.Background()),
				reflect.ValueOf(testPool),
			}

			filterargs, err := buildFilterArgs(reflect.ValueOf(tc.filter), true)
			require.NoError(t, err)

			result := tc.repoMethod.Call(append(args, filterargs...))

			if result[1].Interface() != nil {
				var ok bool
				err, ok = result[1].Interface().(error)
				assert.True(t, ok)
			}

			if result[0].Kind() == reflect.Slice {
				require.Zero(t, result[0].Len())
				// in case of e.g. user.ByTeam it will fail if team not found, instead of returning empty slice
				// require.NoError(t, err) // not necessarily
				if err != nil {
					require.ErrorContains(t, err, errNoRows)
				}
			} else {
				require.ErrorContains(t, err, errNoRows)
			}
		})
	})
}

func buildFilterArgs(filter reflect.Value, zero bool) ([]reflect.Value, error) {
	args := []reflect.Value{}

	t := filter.Type()
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.String, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if !zero {
			args = append(args, filter)

			break
		}
		args = append(args, reflect.Zero(t))
	case reflect.Array:
		if !zero {
			args = append(args, filter)

			break
		}
		if t == reflect.TypeOf(uuid.UUID{}) {
			// FIXME: unsupported filter type: db.UserID
			// need to handle case reflect.Struct with field UUID
			args = append(args, reflect.ValueOf(uuid.Nil))
		}
	case reflect.Slice: // assume testing with multiple parameters
		for i := 0; i < filter.Len(); i++ {
			elem := filter.Index(i)
			elemArgs, err := buildFilterArgs(elem, zero)
			if err != nil {
				return nil, err
			}
			args = append(args, elemArgs...)
		}
	case reflect.Struct:
		if t.Kind() == reflect.Struct {
			if t.Field(0).Type == reflect.TypeOf(uuid.New()) {
				if !zero {
					args = append(args, filter)

					break
				}
				args = append(args, reflect.Zero(t)) // zero value of db.UserID, etc.
			}
		}
	case reflect.Interface: // handle `any`
		if !filter.IsNil() {
			value := reflect.ValueOf(filter.Interface())
			elemArgs, err := buildFilterArgs(value, zero)
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
