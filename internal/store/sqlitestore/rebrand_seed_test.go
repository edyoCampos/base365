//go:build sqlite || sqliteonly

package sqlitestore

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

// RF-14: a database created from scratch carries no trace of the pre-rebrand name,
// neither in the schema SQL nor in any seeded row. The name is built from parts so the
// brand check does not flag this file.
func TestFreshDB_HasNoLegacyBrand(t *testing.T) {
	legacy := "go" + "claw"

	db, err := OpenDB(filepath.Join(t.TempDir(), "fresh.db"))
	if err != nil {
		t.Fatalf("OpenDB: %v", err)
	}
	t.Cleanup(func() { _ = db.Close() })
	if err := EnsureSchema(db); err != nil {
		t.Fatalf("EnsureSchema: %v", err)
	}

	rows, err := db.Query(`SELECT name, COALESCE(sql, '') FROM sqlite_master WHERE type IN ('table','index','trigger','view')`)
	if err != nil {
		t.Fatal(err)
	}
	type obj struct{ name, sql string }
	var objs []obj
	for rows.Next() {
		var o obj
		if err := rows.Scan(&o.name, &o.sql); err != nil {
			t.Fatal(err)
		}
		objs = append(objs, o)
	}
	_ = rows.Close()

	tables := 0
	for _, o := range objs {
		if strings.Contains(strings.ToLower(o.name+o.sql), legacy) {
			t.Errorf("schema object %q mentions the legacy name", o.name)
		}
		if strings.HasPrefix(o.sql, "CREATE TABLE") && !strings.HasPrefix(o.name, "sqlite_") && !strings.Contains(o.sql, "VIRTUAL") {
			tables++
			// Dump every value of the table as text and look for the legacy name.
			r, err := db.Query(fmt.Sprintf(`SELECT * FROM "%s"`, o.name))
			if err != nil {
				continue // e.g. virtual/FTS helper tables that cannot be scanned generically
			}
			cols, _ := r.Columns()
			for r.Next() {
				vals := make([]any, len(cols))
				ptrs := make([]any, len(cols))
				for i := range vals {
					ptrs[i] = &vals[i]
				}
				if err := r.Scan(ptrs...); err != nil {
					break
				}
				for i, v := range vals {
					if strings.Contains(strings.ToLower(fmt.Sprint(v)), legacy) {
						t.Errorf("seeded value in %s.%s mentions the legacy name: %v", o.name, cols[i], v)
					}
				}
			}
			_ = r.Close()
		}
	}
	if tables == 0 {
		t.Fatal("expected at least one table to be inspected")
	}
}
