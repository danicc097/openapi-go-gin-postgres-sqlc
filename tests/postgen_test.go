package tests_test

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/postgen"
	"github.com/danicc097/openapi-go-gin-postgres-sqlc/tests"
	"github.com/google/go-cmp/cmp"
)

func setupTests(t *testing.T) {
	os.Setenv("IS_TESTING", "1")

	cmd := exec.Command(
		"../bin/build",
		"generate-tests-api",
	)
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Logf("combined out:\n%s\n", string(out))
		t.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func TestHandlerPostProcessing(t *testing.T) {
	t.Parallel()

	setupTests(t)

	const baseDir = "testdata/postgen/openapi_generator"

	cases := []struct {
		Name string
		Dir  string
	}{
		{
			"Merging",
			"merge_changes",
		},
		{
			"NameClashing",
			"name_clashing",
		},
	}

	for _, test := range cases {
		test := test
		t.Run(test.Name, func(t *testing.T) {
			t.Parallel()

			conf := &postgen.Conf{
				CurrentHandlersDir: path.Join(baseDir, test.Dir, "internal/handlers"),
				GenHandlersDir:     path.Join(baseDir, test.Dir, "internal/gen"),
				OutHandlersDir:     path.Join(baseDir, test.Dir, "got"),
				OutServicesDir:     path.Join(baseDir, test.Dir, "internal/services"),
			}

			err := os.RemoveAll(conf.OutHandlersDir)
			if err != nil {
				t.Fatal(err)
			}
			var stderr bytes.Buffer
			og := postgen.NewOpenapiGenerator(conf, &stderr)

			s := tests.GetStderr(t, path.Join(baseDir, test.Dir, "want"))

			err = og.Generate()
			if err != nil && s != "" {
				// check stderr.txt is exactly as output
				if diff := cmp.Diff(s, stderr.String()); diff != "" {
					t.Fatalf("stderr differed (-want +got):\n%s", diff)
				}

				return
			} else if err != nil {
				t.Fatalf("err: %s\nstderr: %s\n", err, &stderr)
			}

			pconf := &printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}
			ff, _ := filepath.Glob(path.Join(baseDir, test.Dir, "want", "/*"))
			for _, f := range ff {
				basename := path.Base(f)
				wp := path.Join(baseDir, test.Dir, "want", basename)
				gp := path.Join(conf.OutHandlersDir, basename)
				wantBlob, _ := os.ReadFile(wp)
				gotBlob, _ := os.ReadFile(gp)
				want := &bytes.Buffer{}
				got := &bytes.Buffer{}

				err := printContent(t, string(wantBlob), pconf, want)
				if err != nil {
					t.Fatal(err)
				}
				err = printContent(t, string(gotBlob), pconf, got)
				if err != nil {
					t.Fatal(err)
				}
				if diff := cmp.Diff(want.String(), got.String()); diff != "" {
					t.Errorf("strings differed (-want +got):\n%s", diff)
				}
			}
		})
	}
}

// printContent normalizes source code and prints to a dest.
func printContent(t *testing.T, content string, pconf *printer.Config, dest io.Writer) error {
	t.Helper()
	fset := token.NewFileSet()

	r, err := parser.ParseFile(fset, "", content, parser.ParseComments)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	if err := pconf.Fprint(buf, fset, r); err != nil {
		return err
	}

	out, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	dest.Write(out)

	return nil
}
