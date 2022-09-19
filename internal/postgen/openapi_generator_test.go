package postgen

import (
	"bytes"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"os"
	"path"
	"path/filepath"
	"testing"

	"github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/testutil"
	"github.com/google/go-cmp/cmp"
)

func TestHandlerPostProcessing(t *testing.T) {
	t.Parallel()

	const baseDir = "testdata/openapi_generator"

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

	for _, tc := range cases {
		t.Run(tc.Name, func(t *testing.T) {
			conf := &Conf{
				CurrentHandlersDir: path.Join(baseDir, tc.Dir, "internal/rest"),
				GenHandlersDir:     path.Join(baseDir, tc.Dir, "internal/gen"),
				OutHandlersDir:     path.Join(baseDir, tc.Dir, "got"),
				OutServicesDir:     path.Join(baseDir, tc.Dir, "internal/services"),
			}

			err := os.RemoveAll(conf.OutHandlersDir)
			if err != nil {
				t.Fatal(err)
			}
			var stderr bytes.Buffer
			og := NewOpenapiGenerator(conf, &stderr, "")

			s := testutil.GetStderr(t, path.Join(baseDir, tc.Dir, "want"))

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
			ff, _ := filepath.Glob(path.Join(baseDir, tc.Dir, "want", "/*"))
			for _, f := range ff {
				basename := path.Base(f)
				wp := path.Join(baseDir, tc.Dir, "want", basename)
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
					t.Errorf("%s: strings differed (-want +got):\n%s", wp, diff)
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
