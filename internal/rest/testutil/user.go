package testutil

func createUser() {
	// TODO any value that has a  unique constraint in db must be generated
	// via randomXXX().
	// the only parameters createUser accepts are high level, at the `rest` layer only.
	// functions in this package only make use of SERVICES.
	// it also accepts
	// e.g. roles -> roles from deepmap gen server, not from services or anywhere else.
	// when everything is create we return the user as well as any external data associated to it
	// after creation
}
