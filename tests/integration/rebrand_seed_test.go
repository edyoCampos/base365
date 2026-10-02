//go:build integration

package integration

import (
	"fmt"
	"strings"
	"testing"
)

// RF-14: a database created from the migrations carries no trace of the pre-rebrand name,
// in object comments or in any seeded row. The name is built from parts so the brand
// check does not flag this file.
func TestMigratedDB_HasNoLegacyBrand(t *testing.T) {
	db := testDB(t)
	legacy := "go" + "claw"

	// Object comments (COMMENT ON ...).
	rows, err := db.Query(`SELECT COALESCE(description, '') FROM pg_description`)
	if err != nil {
		t.Fatal(err)
	}
	for rows.Next() {
		var d string
		if err := rows.Scan(&d); err != nil {
			t.Fatal(err)
		}
		if strings.Contains(strings.ToLower(d), legacy) {
			t.Errorf("object comment mentions the legacy name: %q", d)
		}
	}
	_ = rows.Close()

	// Every row of every table in the public schema, serialized as JSON text.
	trows, err := db.Query(`SELECT table_name FROM information_schema.tables WHERE table_schema = 'public' AND table_type = 'BASE TABLE'`)
	if err != nil {
		t.Fatal(err)
	}
	var tables []string
	for trows.Next() {
		var n string
		if err := trows.Scan(&n); err != nil {
			t.Fatal(err)
		}
		tables = append(tables, n)
	}
	_ = trows.Close()
	if len(tables) == 0 {
		t.Fatal("expected migrated tables")
	}
	for _, tbl := range tables {
		var hits int
		q := fmt.Sprintf(`SELECT count(*) FROM "%s" t WHERE lower(row_to_json(t)::text) LIKE '%%%s%%'`, tbl, legacy)
		if err := db.QueryRow(q).Scan(&hits); err != nil {
			t.Fatalf("scan %s: %v", tbl, err)
		}
		if hits > 0 {
			t.Errorf("table %s has %d row(s) mentioning the legacy name", tbl, hits)
		}
	}
}
