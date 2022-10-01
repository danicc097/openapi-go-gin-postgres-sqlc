package postgen

type Conf struct {
	// CurrentHandlersDir is the directory with edited generated files.
	CurrentHandlersDir Dir
	// GenHandlersDir is the directory with raw generated files for a given spec.
	GenHandlersDir Dir
	// OutHandlersDir is the directory to store merged files,
	// which may differ from CurrentHandlersDir.
	OutHandlersDir Dir
	// OutServicesDir is the directory to store new default services.
	OutServicesDir Dir
}
