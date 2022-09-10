// Code adapted from:
// https://github.com/MarioCarrion/todo-api-microservice-example

package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"time"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/format"
	server "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
)

func openBrowser(url string) {
	var err error

	time.Sleep(time.Second * 2)

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Default().Printf("Couldn't launch local browser: %s", err)
	}
}

func main() {
	var env, address, specPath string

	flag.StringVar(&env, "env", ".env", "Environment Variables filename")
	flag.StringVar(&address, "address", ":8090", "HTTP Server Address")
	flag.StringVar(&specPath, "spec-path", "openapi.yaml", "OpenAPI specification filepath")
	flag.Parse()

	url := fmt.Sprintf("https://localhost%s/v2/docs", address)
	// go openBrowser(url)

	errC, err := server.Run(env, address, specPath)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}
	fmt.Printf("\n%sVisit the docs at %s%s\n\n", format.Green, url, format.Off)

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}

}
