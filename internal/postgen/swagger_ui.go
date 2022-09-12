package postgen

import (
	"bytes"
	"os"
	"path"
	"strconv"
	"text/template"
)

// SetupSwaggerUI sets url as the base path.
func SetupSwaggerUI(url string) error {
	buf := &bytes.Buffer{}
	staticDir := "internal/static"

	t, err := template.New("").Parse(`
window.onload = function () {
	//<editor-fold desc="Changeable Configuration Block">

	// the following lines will be replaced by docker/configurator, when it runs in a docker-container
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

	_, err = swaggerInit.Write(buf.Bytes())
	if err != nil {
		return err
	}

	os.Link("openapi.yaml", path.Join(staticDir, "swagger-ui/openapi.yaml"))

	return nil
}
