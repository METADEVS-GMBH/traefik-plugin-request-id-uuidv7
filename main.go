package request_id

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const defaultHeader = "X-Request-ID"
const defaultEnabled = true

type Config struct {
	HeaderName string `json:"headerName,omitempty"`
	Enabled    bool   `json:"enabled,omitempty"`
}

func CreateConfig() *Config {
	return &Config{
		HeaderName: defaultHeader,
		Enabled:    defaultEnabled,
	}
}

func New(ctx context.Context, next http.Handler, config *Config, _ string) (http.Handler, error) {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if config.Enabled {
			value := uuid.Must(uuid.NewV7()).String()
			request.Header.Set(config.HeaderName, value)
			writer.Header().Set(config.HeaderName, value)
		}
		next.ServeHTTP(writer, request)
	}), nil
}
