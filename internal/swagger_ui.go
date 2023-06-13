package internal

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
	swaggerUIDir := "internal/static/swagger-ui"

	t, err := template.New("").Parse(`
const detectEnvironment = () => {
	const host = window.location.host.toLowerCase();
	if (host.indexOf("localhost") > -1) {
		return "dev";
	}
	if (host.indexOf("e2e") > -1) {
		return "e2e";
	}
	if (host.indexOf("prod") > -1) {
		return "prod";
	}
};

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
			const env = detectEnvironment();
			let spec = ui.specSelectors.specJson().toJS();
			let servers = spec.servers.filter((item) => {
				return item.description.toLowerCase() === env.toLowerCase();
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
	}

	if err := t.Execute(buf, params); err != nil {
		return err
	}

	swaggerInit, err := os.Create(path.Join(swaggerUIDir, "swagger-initializer.js"))
	if err != nil {
		return err
	}

	if _, err := swaggerInit.Write(buf.Bytes()); err != nil {
		return err
	}

	// not needed, handler will use spec path from entrypoint args instead of reading the embed
	// bundleSpec := path.Join(swaggerUIDir, "openapi.yaml")
	// os.Remove(bundleSpec)
	// if err := os.Link(specPath, bundleSpec); err != nil {
	// 	return err
	// }

	return nil
}