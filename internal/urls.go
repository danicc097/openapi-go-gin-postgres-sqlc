package internal

import (
	"fmt"
	"net/url"
	"os"
)

// BuildAPIURL returns a fully-qualified URL with the given path elements
// accounting for reverse proxy configuration.
func BuildAPIURL(subpaths ...string) string {
	elems := []string{os.Getenv("API_PREFIX"), os.Getenv("API_VERSION")}
	elems = append(elems, subpaths...)

	path, err := url.JoinPath(
		"",
		elems...,
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	var host string

	switch os.Getenv("APP_ENV") {
	case "prod":
		host = os.Getenv("DOMAIN")
	default:
		host = os.Getenv("DOMAIN") + ":" + os.Getenv("API_PORT")
	}

	dsn := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}

	url := dsn.String()

	return url
}
