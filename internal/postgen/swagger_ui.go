package postgen

import (
	"bytes"
	"os"
	"path"
	"strconv"
	"text/template"
)

// SetupSwaggerUI sets url as the base path.
func SetupSwaggerUI(url string, specPath string) error {
	buf := &bytes.Buffer{}
	staticDir := "internal/static"

	t, err := template.New("").Parse(`
window.onload = function () {
	//<editor-fold desc="Changeable Configuration Block">

	window.ui = SwaggerUIBundle({
		url: {{.URL}},
		dom_id: "#swagger-ui",
		deepLinking: true,
		presets: [SwaggerUIBundle.presets.apis, SwaggerUIStandalonePreset],
		plugins: [SwaggerUIBundle.plugins.DownloadUrl],
		layout: "StandaloneLayout",
	});

	//</editor-fold>
};
	`)
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"URL": strconv.Quote(url),
	}

	if err := t.Execute(buf, params); err != nil {
		return err
	}

	swaggerInit, err := os.Create(path.Join(staticDir, "swagger-ui/swagger-initializer.js"))
	if err != nil {
		return err
	}

	if _, err := swaggerInit.Write(buf.Bytes()); err != nil {
		return err
	}

	bundleSpec := path.Join(staticDir, "swagger-ui/openapi.yaml")
	os.Remove(bundleSpec)
	if err := os.Link(specPath, bundleSpec); err != nil {
		return err
	}

	return nil
}
