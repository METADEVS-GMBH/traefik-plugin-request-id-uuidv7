// Package traefik_request_id a Traefik plugin to add request ID to incoming HTTP requests.

package requestid

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const defaultHeader = "X-Request-ID"
const defaultEnabled = true

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
