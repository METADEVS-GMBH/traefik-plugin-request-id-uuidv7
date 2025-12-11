// Package traefik_plugin_request_id_uuidv7 a Traefik plugin to add request ID to incoming HTTP requests.
//
//nolint:revive,stylecheck // Traefik plugin requires this package name
package traefik_plugin_request_id_uuidv7

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const (
	defaultHeader  = "X-Request-ID"
	defaultEnabled = true
)

// Config the plugin configuration.
type Config struct {
	HeaderName string `json:"headerName,omitempty"`
	Enabled    bool   `json:"enabled,omitempty"`
}

// CreateConfig creates the default plugin configuration.
func CreateConfig() *Config {
	return &Config{
		HeaderName: defaultHeader,
		Enabled:    defaultEnabled,
	}
}

// New created a new Request ID plugin.
func New(_ context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if config.Enabled {
			value := uuid.Must(uuid.NewV7()).String()
			request.Header.Set(config.HeaderName, value)
			writer.Header().Set(config.HeaderName, value)
		}
		next.ServeHTTP(writer, request)
	}), nil
}
