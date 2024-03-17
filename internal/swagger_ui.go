package internal

import (
	"bytes"
	"fmt"
	"os"
	"path"
	"strconv"
	"text/template"
)

// SetupSwaggerUI sets url in the Swagger docs to the endpoint where specPath is served.
func SetupSwaggerUI(url string, specPath, swaggerUIDir string) error {
	buf := &bytes.Buffer{}

	// TODO: use indexes in array via plugin: https://github.com/swagger-api/swagger-ui/issues/6790#issuecomment-756513818
	// NOTE: cannot make plugins for responses in swagger ui, see https://github.com/swagger-api/swagger-ui/issues/5200
	// ie no SSE eventsource that accumulates messages instead of a useless endpoint

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
		onComplete: function () {
			const env = {{.Env}};
			let spec = ui.specSelectors.specJson().toJS();
			console.log(spec.servers)
			let servers = spec.servers.filter((item) => {
				return item["description"]?.toLowerCase() === env.toLowerCase();
			});
			spec.servers = servers;
			ui.specActions.updateJsonSpec(spec);
		},
	});

	//</editor-fold>
};
	`)
	if err != nil {
		return err
	}

	params := map[string]interface{}{
		"URL": strconv.Quote(url),
		"Env": strconv.Quote(string(Config.AppEnv)),
	}

	if err := t.Execute(buf, params); err != nil {
		return err
	}

	swaggerInitPath := path.Join(swaggerUIDir, "swagger-initializer.js")
	swaggerInit, err := os.Create(swaggerInitPath)
	if err != nil {
		return fmt.Errorf("could not create %s: %w", swaggerInitPath, err)
	}

	if _, err := swaggerInit.Write(buf.Bytes()); err != nil {
		return fmt.Errorf("could not write to %s: %w", swaggerInitPath, err)
	}

	return nil
}
