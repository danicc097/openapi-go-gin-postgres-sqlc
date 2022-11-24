package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/pkg/errors"
)

// AppConfig contains app settings which are read from a config file.
type AppConfig struct {
	Test              string    `json:"test"`
	TestOptional      *string   `json:"testOptional"`
	TestArray         []string  `json:"testArray"`
	TestArrayOptional *[]string `json:"testArrayOptional"`
}

// Returns the directory of the file this function lives in.
func getFileRuntimeDirectory() string {
	_, b, _, _ := runtime.Caller(0)

	return path.Join(path.Dir(b))
}

var localConfigPath = filepath.Join(getFileRuntimeDirectory(), "config/%s.json")

// GenerateConfigTemplate creates a template.json config file for reference.
func GenerateConfigTemplate() error {
	cfg, err := json.MarshalIndent(&AppConfig{}, "", "  ")
	if err != nil {
		return errors.Wrap(err, "could not marshal template config json")
	}
	if err := os.WriteFile(fmt.Sprintf(localConfigPath, "template"), cfg, 0o777); err != nil {
		return errors.Wrap(err, "could not save template config json")
	}

	return nil
}
