//go:build sqlite || sqliteonly

package cmd

import (
	"context"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/edyoCampos/base365/internal/store/sqlitestore"
)

// RF-18 on a real (SQLite) store: a fresh database gets the external extractor disabled.
func TestWebFetchSeed_FreshSQLiteDatabase(t *testing.T) {
	db, err := sqlitestore.OpenDB(filepath.Join(t.TempDir(), "seed.db"))
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = db.Close() })
	if err := sqlitestore.EnsureSchema(db); err != nil {
		t.Fatal(err)
	}
	bts := sqlitestore.NewSQLiteBuiltinToolStore(db)
	ctx := context.Background()
	seedBuiltinTools(ctx, bts)

	settings, err := bts.GetSettings(ctx, "web_fetch")
	if err != nil {
		t.Fatal(err)
	}
	var cfg struct {
		Extractors []struct {
			Name    string `json:"name"`
			Enabled bool   `json:"enabled"`
		} `json:"extractors"`
	}
	if err := json.Unmarshal(settings, &cfg); err != nil {
		t.Fatalf("settings %s: %v", settings, err)
	}
	for _, e := range cfg.Extractors {
		if e.Name == "defuddle" && e.Enabled {
			t.Error("defuddle must be disabled on a fresh database")
		}
	}
	if strings.Contains(strings.ToLower(string(settings)), "go"+"claw") {
		t.Errorf("seeded settings mention the legacy name: %s", settings)
	}
}
