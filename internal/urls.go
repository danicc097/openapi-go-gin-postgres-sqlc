package internal

import (
	"fmt"
	"net/url"
	"os"
)

// BuildAPIURL returns a fully-qualified URL with the given path elements
// accounting for reverse proxy configuration.
func BuildAPIURL(subpaths ...string) string {
	elems := []string{config.APIPrefix, config.APIVersion}
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

	switch config.AppEnv {
	case "prod":
		host = config.Domain
	default:
		host = config.Domain + ":" + config.APIPort
	}

	dsn := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}

	url := dsn.String()

	return url
}
