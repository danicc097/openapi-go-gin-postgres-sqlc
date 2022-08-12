package main

import (
	"os"
	"testing"

	"github.com/dave/dst/decorator"
)

func TestHandlerPostProcessing(t *testing.T) {
	content, err := os.ReadFile("testdata/merge_changes/current/api_user.go")
	if err != nil {
		panic(err)
	}
	f, err := decorator.Parse(string(content))
	if err != nil {
		panic(err)
	}

	// list := f.Decls[0].(*dst.FuncDecl).Body.List
	// list[0], list[1] = list[1], list[0]

	if err := decorator.Print(f); err != nil {
		panic(err)
	}
}
