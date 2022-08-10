# ``environment`` package

the alternative is to generate service structs by reparsing and editing the generated files again...
we already have api_{tag}.go files, a {tag}Service struct
that has logger, conn pool, etc. and implements all methods isn't too hard,
but would need rethinking how non-implemented handlers are created,
in this case we will have to
parse existing handlers/api_*.go, not just render a brand new template,
and 1. remove inexisting handlers, 2. dont touch the existing ones 3. append
the new unimplemented methods.
We will also need to regenerate, or completely rethink routes.go to initialize
everything based on each service and handler... perhaps best to keep all routes in each api_{tag}.go and concatenate and add to the main router as needed (would need to also parse cmd/**/main.go to ensure we are adding all possible route lists...)
Perhaps some globals here and there wont hurt as much.
