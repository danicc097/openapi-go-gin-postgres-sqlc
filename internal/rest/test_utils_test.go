package rest_test

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"reflect"
	"strings"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
)

func ReqWithAPIKey(apiKey string) RequestEditorFn {
	return func(ctx context.Context, req *http.Request) error {
		req.Header.Add("Content-Type", "application/json")
		req.Header.Add(rest.ApiKeyHeaderKey, apiKey)

		return nil
	}
}

type constructURLOptions struct {
	params any
}

// ConstructURLOption is the type for options that can be passed to ConstructInternalPath.
type ConstructURLOption func(*constructURLOptions)

// WithQueryParams specifies the struct containing the query parameters.
func WithQueryParams(params any) ConstructURLOption {
	return func(o *constructURLOptions) {
		o.params = params
	}
}

// ConstructInternalPath constructs a URL with encoded parameters based
// on the non-nil fields of the provided struct via the form tag.
// Required path prefixes are added automatically.
func ConstructInternalPath(subpath string, options ...ConstructURLOption) (string, error) {
	cleanSubpath := strings.TrimPrefix(strings.TrimPrefix(subpath, internal.Config.APIVersion), "/")
	u, err := url.Parse(internal.Config.APIVersion + "/" + cleanSubpath)
	if err != nil {
		return "", fmt.Errorf("could not parse URL: %w", err)
	}

	query := u.Query()

	// Process options
	opts := &constructURLOptions{}
	for _, opt := range options {
		opt(opts)
	}

	if opts.params == nil {
		return u.String(), nil
	}

	v := reflect.ValueOf(opts.params)
	t := reflect.TypeOf(opts.params)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return "", fmt.Errorf("params must be a struct")
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if value.Kind() == reflect.Ptr {
			if value.IsNil() {
				continue
			}

			value = value.Elem()
		}

		fieldName := field.Name
		formTag := field.Tag.Get("form")
		if formTag != "" {
			fieldName = formTag
		}

		fieldValue := fmt.Sprintf("%v", value.Interface())
		query.Add(fieldName, fieldValue)
	}

	u.RawQuery = query.Encode()

	return u.String(), nil
}

// MustConstructInternalPath constructs a URL with encoded parameters based on the non-nil fields of the provided struct.
// Required path prefixes are added automatically. It panics if an error occurs during URL construction.
func MustConstructInternalPath(subpath string, options ...ConstructURLOption) string {
	url, err := ConstructInternalPath(subpath, options...)
	if err != nil {
		panic(fmt.Errorf("could not construct URL: %w", err))
	}

	return url
}
