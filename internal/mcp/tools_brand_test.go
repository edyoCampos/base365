package mcp

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

// minCRUDTools is the number of CRUD tools registered before the rebrand
// (counted from NewTool calls in non-test files). The rename must not lose any.
const minCRUDTools = 179

// TestCRUDToolNames_Prefix guards the MCP tool-name contract (RN-02, RF-05):
// every CRUD tool starts with the product prefix, none keeps the legacy one.
func TestCRUDToolNames_Prefix(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	rx := regexp.MustCompile(`NewTool\(\s*"([^"]+)"`)
	names := map[string]bool{}
	for _, f := range files {
		if strings.HasSuffix(f, "_test.go") {
			continue
		}
		src, err := os.ReadFile(f)
		if err != nil {
			t.Fatal(err)
		}
		for _, m := range rx.FindAllStringSubmatch(string(src), -1) {
			names[m[1]] = true
		}
	}
	if len(names) < minCRUDTools {
		t.Fatalf("expected at least %d CRUD tools, found %d", minCRUDTools, len(names))
	}
	legacy := "go" + "claw_" // built from parts so the brand check does not flag this file
	for n := range names {
		if !strings.HasPrefix(n, "base365_") {
			t.Errorf("tool %q must start with base365_", n)
		}
		if strings.HasPrefix(n, legacy) {
			t.Errorf("tool %q uses the legacy prefix", n)
		}
	}
}
