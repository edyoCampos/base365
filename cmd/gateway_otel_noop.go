//go:build !otel

package cmd

import (
	"context"

	"github.com/edyoCampos/base365/internal/config"
	"github.com/edyoCampos/base365/internal/tracing"
)

// initOTelExporter is a no-op when built without the "otel" tag.
// Build with `go build -tags otel` to enable OpenTelemetry export.
func initOTelExporter(_ context.Context, _ *config.Config, _ *tracing.Collector) {
}
