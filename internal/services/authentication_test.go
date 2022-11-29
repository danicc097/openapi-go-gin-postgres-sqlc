package services_test

import (
	"testing"
)

func TestAuthentication(t *testing.T) {
	t.Parallel()

	/**
	 * TODO mock auth server. Test:
	- user created if not exists with given sub.
	- error already exists
	- updates user if userinfo changed (email, username at the very least)
	*/
}
