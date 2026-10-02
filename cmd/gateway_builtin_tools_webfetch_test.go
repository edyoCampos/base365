package cmd

import (
	"encoding/json"
	"strings"
	"testing"
)

// RN-07 / RF-18: a fresh install ships the external extractor disabled and without any endpoint.
func TestWebFetchSeed_ExternalExtractorDisabledByDefault(t *testing.T) {
	var raw json.RawMessage
	for _, def := range builtinToolSeedData() {
		if def.Name == "web_fetch" {
			raw = def.Settings
		}
	}
	if raw == nil {
		t.Fatal("web_fetch must be seeded")
	}
	if string(raw) != defaultWebFetchSettings {
		t.Fatalf("seed and backfill must share the same default settings, got %s", raw)
	}
	var cfg struct {
		Extractors []struct {
			Name    string `json:"name"`
			Enabled bool   `json:"enabled"`
			BaseURL string `json:"base_url"`
		} `json:"extractors"`
	}
	if err := json.Unmarshal(raw, &cfg); err != nil {
		t.Fatal(err)
	}
	for _, e := range cfg.Extractors {
		switch e.Name {
		case "defuddle":
			if e.Enabled || e.BaseURL != "" {
				t.Errorf("defuddle must be disabled without endpoint by default, got %+v", e)
			}
		case "html-to-markdown":
			if !e.Enabled {
				t.Error("in-process extractor must stay enabled")
			}
		}
	}
	if strings.Contains(string(raw), "http") {
		t.Errorf("default settings must not contain any endpoint: %s", raw)
	}
}
