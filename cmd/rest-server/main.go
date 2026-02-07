package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal"
	server "github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/format/colors"
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
		err = errors.New("unsupported platform")
	}
	if err != nil {
		log.Default().Printf("Couldn't launch local browser: %s", err)
	}
}

func main() {
	var env, specPath string

	flag.StringVar(&env, "env", "", "Environment Variables filename")
	flag.StringVar(&specPath, "spec-path", "openapi.yaml", "OpenAPI specification filepath")
	flag.Parse()

	var errs []string

	if env == "" {
		errs = append(errs, "    - env is required but unset")
	}

	if len(errs) > 0 {
		log.Fatalf("%s", "error: \n"+strings.Join(errs, "\n"))
	}

	// go openBrowser(url)

	// dummy values for dashboard
	cpuTemp := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	})
	prometheus.MustRegister(cpuTemp)
	cpuTemp.Set(65.3)

	errC, err := server.Run(env, specPath)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	fmt.Println("\n" + colors.G + colors.Bold +
		"Visit the docs: \n\t" +
		colors.G + internal.BuildAPIURL("docs") + "\n\t" +
		colors.G + internal.BuildAPIURL("docs-redoc") + " " +
		colors.Off + "\n")

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}
