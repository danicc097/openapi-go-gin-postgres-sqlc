package postgen

type Conf struct {
	// CurrentHandlersDir is the directory with edited generated files.
	CurrentHandlersDir string
	// GenHandlersDir is the directory with raw generated files for a given spec.
	GenHandlersDir string
	// OutHandlersDir is the directory to store merged files,
	// which may differ from CurrentHandlersDir.
	OutHandlersDir string
	// OutServicesDir is the directory to store new default services.
	OutServicesDir string
}
