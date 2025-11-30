package handler

import (
	"net/http"
	"strings"
)

// SwaggerHandler serves the Swagger UI and OpenAPI spec
type SwaggerHandler struct {
	specPath string
}

// NewSwaggerHandler creates a new swagger handler
func NewSwaggerHandler(specPath string) *SwaggerHandler {
	return &SwaggerHandler{specPath: specPath}
}

// ServeSpec serves the OpenAPI YAML specification
func (h *SwaggerHandler) ServeSpec(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, h.specPath)
}

// ServeUI serves the Swagger UI HTML page
func (h *SwaggerHandler) ServeUI(w http.ResponseWriter, r *http.Request) {
	// Redirect /swagger to /swagger/
	if r.URL.Path == "/swagger" {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(swaggerUIHTML))
}

// swaggerUIHTML is the HTML template for Swagger UI
var swaggerUIHTML = strings.TrimSpace(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Copy Paste Service - API Documentation</title>
    <link rel="stylesheet" type="text/css" href="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui.css">
    <style>
        html { box-sizing: border-box; overflow: -moz-scrollbars-vertical; overflow-y: scroll; }
        *, *:before, *:after { box-sizing: inherit; }
        body { margin: 0; background: #fafafa; }
        .topbar { display: none; }
    </style>
</head>
<body>
    <div id="swagger-ui"></div>
    <script src="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui-bundle.js"></script>
    <script src="https://unpkg.com/swagger-ui-dist@5.9.0/swagger-ui-standalone-preset.js"></script>
    <script>
        window.onload = function() {
            const ui = SwaggerUIBundle({
                url: "/swagger/openapi.yaml",
                dom_id: '#swagger-ui',
                deepLinking: true,
                presets: [
                    SwaggerUIBundle.presets.apis,
                    SwaggerUIStandalonePreset
                ],
                plugins: [
                    SwaggerUIBundle.plugins.DownloadUrl
                ],
                layout: "StandaloneLayout"
            });
            window.ui = ui;
        };
    </script>
</body>
</html>
`)

