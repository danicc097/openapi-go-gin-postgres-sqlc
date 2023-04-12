package internal

import (
	"fmt"
	"net/url"
	"os"
)

// BuildAPIURL returns a fully-qualified URL with the given path elements
// accounting for reverse proxy configuration.
func BuildAPIURL(subpaths ...string) string {
	elems := []string{Config.APIPrefix, Config.APIVersion}
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

	switch Config.AppEnv {
	case "prod":
		host = Config.Domain
	default:
		host = Config.Domain + ":" + Config.APIPort
	}

	dsn := url.URL{
		Scheme: "https",
		Host:   host,
		Path:   path,
	}

	url := dsn.String()

	return url
}
