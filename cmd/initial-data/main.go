package main

import (
	"fmt"
	"os"
)

// clear && go run cmd/initial-data/main.go -env .env.dev
func main() {
}

func errAndExit(out []byte, err error) {
	fmt.Fprintf(os.Stderr, "combined out:\n%s\n", string(out))
	fmt.Fprintf(os.Stderr, "cmd.Run() failed with %s\n", err)
	os.Exit(1)
}
