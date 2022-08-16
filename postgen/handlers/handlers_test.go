package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"testing"
)

func setupTests() {
	os.Setenv("IS_TESTING", "1")

	cwd, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(cwd)

	cmd := exec.Command(
		"../../bin/build",
		"generate-tests-api",
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		fmt.Printf("combined out:\n%s\n", string(out))
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func TestHandlerPostProcessing(t *testing.T) {
	setupTests()
	// run build generate-api postgen/handlers/testdata/openapi.yaml postgen/handlers/testdata/gen
}
