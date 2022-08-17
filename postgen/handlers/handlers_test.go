package main

import (
	"bytes"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"testing"

	"github.com/go-playground/assert/v2"
)

// stderr returns the contents of stderr.txt in dir.
func stderr(t *testing.T, dir string) string {
	t.Helper()
	path := filepath.Join(dir, "stderr.txt")

	if _, err := os.Stat(path); !os.IsNotExist(err) {
		blob, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}

		return string(blob)
	}

	return ""
}

func setupTests() {
	os.Setenv("IS_TESTING", "1")

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
	// TODO 2 seconds to use openapi-generator, can we cache it based on
	// templates dir content and spec?
	// setupTests()

	cases := []struct {
		Name string
		Dir  string
	}{
		{
			"Merging",
			"merge_changes",
		},
		// {
		// 	"Name clashing",
		// 	"name_clashing",
		// },
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			// TODO assert exit != 0 and want in stderr for name_clashing
			// if s := stderr(t, test.Dir); s != "" {
			// 	want := s
			// 	t.Logf("want stderr: %s", want)
			// }

			var (
				baseDir = "testdata"
				conf    = Conf{
					CurrentHandlersDir: path.Join(baseDir, string(test.Dir), "internal/handlers"),
					GenHandlersDir:     path.Join(baseDir, string(test.Dir), "internal/gen"),
					OutHandlersDir:     path.Join(baseDir, string(test.Dir), "got"),
					OutServicesDir:     path.Join(baseDir, string(test.Dir), "internal/services"),
				}
			)

			cb := getCommonBasenames(conf)
			handlers := analyzeHandlers(conf, cb)

			generateMergedFiles(handlers, conf)

			pconf := &printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}
			ff, _ := filepath.Glob(conf.OutHandlersDir + "/*")
			for _, f := range ff {
				basename := path.Base(f)
				wp := path.Join(baseDir, string(test.Dir), "want", basename)
				gp := path.Join(conf.OutHandlersDir, basename)
				wantBlob, _ := os.ReadFile(wp)
				gotBlob, _ := os.ReadFile(gp)

				want := &bytes.Buffer{}
				got := &bytes.Buffer{}
				t.Logf("file: %s\n", string(basename))
				t.Logf("want file: %s\n", wp)
				t.Logf("got file: %s\n", gp)
				printContent(t, string(wantBlob), pconf, want)
				printContent(t, string(gotBlob), pconf, got)

				assert.Equal(t, want.String(), got.String())
			}
			// var got []string
			// walk(test.Input, func(input string) {
			// 	got = append(got, input)
			// })

			// if !reflect.DeepEqual(got, test.ExpectedCalls) {
			// 	t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			// }
		})
	}
}

// printContent normalizes source code and prints to a dest.
func printContent(t *testing.T, content string, pconf *printer.Config, dest io.Writer) {
	t.Helper()
	fset := token.NewFileSet()

	r, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		panic(err)
	}

	buf := &bytes.Buffer{}
	if err := pconf.Fprint(buf, fset, r); err != nil {
		panic(err)
	}

	dest.Write(buf.Bytes())
}
