//go:build sqliteonly

package main

import (
	"os"
	"path/filepath"
	"testing"
)

// RF-06 / RF-10: the desktop keeps its secrets under ~/.base365 and never creates the
// pre-rebrand directory. The old name is built from parts so the brand check does not flag this file.
func TestSecretsDir_UsesProductDirectory(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)

	dir := secretsDir()
	want := filepath.Join(home, ".base365", "secrets")
	if dir != want {
		t.Fatalf("secretsDir() = %q, want %q", dir, want)
	}
	if _, err := os.Stat(want); err != nil {
		t.Fatalf("secrets dir must be created: %v", err)
	}
	if _, err := os.Stat(filepath.Join(home, ".go"+"claw")); !os.IsNotExist(err) {
		t.Fatalf("the legacy directory must not be created (err=%v)", err)
	}
	if serviceName != "base365-desktop" {
		t.Fatalf("keyring service name = %q, want base365-desktop", serviceName)
	}
}
