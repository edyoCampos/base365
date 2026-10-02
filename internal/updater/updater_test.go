package updater

import "testing"

// RF-11: the auto-updater queries releases of the product repository, with the desktop tag prefix.
func TestUpdaterTargetsProductRepository(t *testing.T) {
	if githubRepo != "edyoCampos/base365" {
		t.Fatalf("githubRepo = %q, want edyoCampos/base365", githubRepo)
	}
	if tagPrefix != "lite-v" {
		t.Fatalf("tagPrefix = %q, want lite-v", tagPrefix)
	}
}
